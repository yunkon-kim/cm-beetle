package recommendation

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRecommendationHooks(matcher func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair) recommendationHooks {
	return recommendationHooks{
		recommendVNet: func(ctx context.Context, csp, region string, srcInfra onpremmodel.OnpremInfra) ([]cloudmodel.VNetReq, error) {
			return []cloudmodel.VNetReq{
				{
					CidrBlock: "10.0.0.0/16",
					SubnetInfoList: []cloudmodel.SubnetReq{
						{IPv4_CIDR: "10.0.1.0/24"},
					},
				},
			}, nil
		},
		recommendSG: func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty) (cloudmodel.SecurityGroupReq, error) {
			rules := generateSecurityGroupRules(node.FirewallTable)
			rules = DeduplicateFirewallRules(rules)
			return cloudmodel.SecurityGroupReq{
				ConnectionName: csp + "-" + region,
				FirewallRules:  &rules,
			}, nil
		},
		matchCatalog:       matcher,
		preflightCheck:     nil,
		enforceZonalSubnet: nil,
	}
}

func mockGpuNlbCatalog(allowConsolidatedHeteroA100 bool) func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
	ubuntuImg := cloudmodel.ImageInfo{
		Id:             "aws+ap-northeast-2+ubuntu-2204",
		CspImageName:   "ubuntu-jammy-22.04-amd64",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "x86_64",
		OSDistribution: "ubuntu jammy 22.04",
	}

	return func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
		if !hasGpu(node) {
			return []CompatibleSpecImagePair{
				{
					Spec: cloudmodel.SpecInfo{
						Id:          "aws+ap-northeast-2+m6i.2xlarge",
						CspSpecName: "m6i.2xlarge",
						VCPU:        8,
						MemoryGiB:   32,
						CostPerHour: 0.384,
					},
					Image: ubuntuImg,
				},
			}
		}

		clusters := clusterGpuCards(node.GPUCards)
		primary := clusters[0]
		vendor := strings.ToLower(detectGpuVendor(primary.Vendor))
		model := detectGpuModel(primary.Model)

		if vendor == "amd" {
			return []CompatibleSpecImagePair{
				{
					Spec: cloudmodel.SpecInfo{
						Id:                  "aws+ap-northeast-2+g4ad.2xlarge",
						CspSpecName:         "g4ad.2xlarge",
						VCPU:                8,
						MemoryGiB:           32,
						AcceleratorCount:    1,
						AcceleratorMemoryGB: 8,
						AcceleratorModel:    "AMD Radeon Pro V520",
						CostPerHour:         0.379,
					},
					Image: ubuntuImg,
				},
			}
		}

		if model == "a100" {
			if !allowConsolidatedHeteroA100 {
				return nil
			}
			vram := float32(40)
			if primary.MemoryTotalGB >= 80 {
				vram = 80
			}
			return []CompatibleSpecImagePair{
				{
					Spec: cloudmodel.SpecInfo{
						Id:                  "aws+ap-northeast-2+p4d.24xlarge",
						CspSpecName:         "p4d.24xlarge",
						VCPU:                96,
						MemoryGiB:           1152,
						AcceleratorCount:    8,
						AcceleratorMemoryGB: vram,
						AcceleratorModel:    "NVIDIA A100",
						CostPerHour:         32.77,
					},
					Image: ubuntuImg,
				},
			}
		}

		if model == "l4" {
			return []CompatibleSpecImagePair{
				{
					Spec: cloudmodel.SpecInfo{
						Id:                  "aws+ap-northeast-2+g6.2xlarge",
						CspSpecName:         "g6.2xlarge",
						VCPU:                8,
						MemoryGiB:           32,
						AcceleratorCount:    1,
						AcceleratorMemoryGB: 24,
						AcceleratorModel:    "NVIDIA L4",
						CostPerHour:         0.978,
					},
					Image: ubuntuImg,
				},
			}
		}

		// Default NVIDIA T4 tier
		return []CompatibleSpecImagePair{
			{
				Spec: cloudmodel.SpecInfo{
					Id:                  "aws+ap-northeast-2+g4dn.2xlarge",
					CspSpecName:         "g4dn.2xlarge",
					VCPU:                8,
					MemoryGiB:           32,
					AcceleratorCount:    1,
					AcceleratorMemoryGB: 16,
					AcceleratorModel:    "NVIDIA T4",
					CostPerHour:         0.752,
				},
				Image: ubuntuImg,
			},
		}
	}
}

func baseOsAndCpuNode(id, ip string, vcpu uint32, memGiB uint64) onpremmodel.NodeProperty {
	return onpremmodel.NodeProperty{
		Hostname:  id,
		MachineId: id,
		CPU:       onpremmodel.CpuProperty{Architecture: "x86_64", Cpus: vcpu, Threads: 1},
		Memory:    onpremmodel.MemoryProperty{TotalSize: memGiB},
		RootDisk:  onpremmodel.DiskProperty{Label: "boot", Type: "SSD", TotalSize: 50},
		Interfaces: []onpremmodel.NetworkInterfaceProperty{
			{Name: "eth0", IPv4CidrBlocks: []string{ip + "/24"}},
		},
		OS: onpremmodel.OsProperty{
			PrettyName: "Ubuntu 22.04",
			ID:         "ubuntu",
			VersionID:  "22.04",
		},
	}
}

func TestValidateSizingPolicy(t *testing.T) {
	p, err := ValidateSizingPolicy("")
	require.NoError(t, err)
	assert.Equal(t, SizingPolicyRightSizeUp, p)

	p, err = ValidateSizingPolicy("rightSizeUp")
	require.NoError(t, err)
	assert.Equal(t, SizingPolicyRightSizeUp, p)

	p, err = ValidateSizingPolicy("dividing")
	require.NoError(t, err)
	assert.Equal(t, SizingPolicyDividing, p)

	p, err = ValidateSizingPolicy("AUTO")
	require.NoError(t, err)
	assert.Equal(t, SizingPolicyAuto, p)

	_, err = ValidateSizingPolicy("hybrid")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported sizingPolicy")
}

func TestClassifyBackendHardware(t *testing.T) {
	tests := []struct {
		name     string
		members  []onpremmodel.NodeProperty
		expected HardwareCase
	}{
		{
			name: "Case 0: Pure CPU servers",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-cpu-1", CPU: onpremmodel.CpuProperty{Cpus: 8}},
				{MachineId: "srv-cpu-2", CPU: onpremmodel.CpuProperty{Cpus: 16}},
			},
			expected: CasePureCpu,
		},
		{
			name: "Case 1: Mixed CPU and GPU servers",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-cpu", CPU: onpremmodel.CpuProperty{Cpus: 8}},
				{
					MachineId: "srv-gpu",
					CPU:       onpremmodel.CpuProperty{Cpus: 16},
					GPUCards:  []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40}},
				},
			},
			expected: CaseMixedCpuGpu,
		},
		{
			name: "Case 2: Cross-vendor GPU servers (NVIDIA and AMD plus CPU)",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-cpu"},
				{MachineId: "srv-nvidia", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80}}},
				{MachineId: "srv-amd", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "AMD", Model: "MI250", MemoryTotalGB: 128}}},
			},
			expected: CaseCrossVendorGpu,
		},
		{
			name: "Case 3: Heterogeneous GPU models (same vendor, different VRAM)",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-a100-40", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40}}},
				{MachineId: "srv-a100-80", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80}}},
			},
			expected: CaseHeteroGpuModel,
		},
		{
			name: "Case 4: Homogeneous GPU models",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-t4-1", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16}}},
				{MachineId: "srv-t4-2", GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16}}},
			},
			expected: CaseHomoGpuModel,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ClassifyBackendHardware(tc.members))
		})
	}
}

func TestRecommendGpuNlb_AllCasesAndPolicies(t *testing.T) {
	hooks := testRecommendationHooks(mockGpuNlbCatalog(true))

	// Case 1: Mixed CPU (2 nodes) + GPU T4 (2 nodes)
	cpu1 := baseOsAndCpuNode("cpu-1", "10.0.0.1", 4, 16)
	cpu2 := baseOsAndCpuNode("cpu-2", "10.0.0.2", 4, 16)
	gpu1 := baseOsAndCpuNode("gpu-1", "10.0.0.3", 8, 32)
	gpu1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	gpu2 := baseOsAndCpuNode("gpu-2", "10.0.0.4", 8, 32)
	gpu2.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}

	mixedSrc := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{cpu1, cpu2, gpu1, gpu2},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name: "infer-pool",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "s1", IP: "10.0.0.1", Port: 8000},
						{Name: "s2", IP: "10.0.0.2", Port: 8000},
						{Name: "s3", IP: "10.0.0.3", Port: 8000},
						{Name: "s4", IP: "10.0.0.4", Port: 8000},
					},
				},
			},
		},
	}

	// rightSizeUp: 1 GPU NodeGroup (Size=4) + 1 NLB
	resRightSize, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", mixedSrc, SizingPolicyRightSizeUp, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resRightSize)
	require.Len(t, resRightSize[0].TargetInfra.NodeGroups, 1)
	require.Len(t, resRightSize[0].TargetNlbList, 1)
	assert.Equal(t, 4, resRightSize[0].TargetInfra.NodeGroups[0].NodeGroupSize)
	assert.Equal(t, resRightSize[0].TargetInfra.NodeGroups[0].Name, resRightSize[0].TargetNlbList[0].TargetGroup.NodeGroupId)

	// dividing: 2 NodeGroups (2 CPU + 2 GPU) + 2 NLBs
	resDividing, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", mixedSrc, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resDividing)
	require.Len(t, resDividing[0].TargetInfra.NodeGroups, 2)
	require.Len(t, resDividing[0].TargetNlbList, 2)
	costRightSize := CalculateCandidateHourlyCost(resRightSize[0])
	costDividing := CalculateCandidateHourlyCost(resDividing[0])
	assert.Less(t, costDividing.TotalHourlyCostUSD, costRightSize.TotalHourlyCostUSD)

	// auto: must select dividing here due to significant cost/match optimization
	resAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", mixedSrc, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resAuto)
	require.Len(t, resAuto[0].TargetInfra.NodeGroups, 2)
	require.Len(t, resAuto[0].TargetNlbList, 2)

	// Case 2: Cross-vendor GPU (NVIDIA + AMD)
	amd1 := baseOsAndCpuNode("amd-1", "10.0.0.5", 8, 32)
	amd1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "AMD", Model: "Radeon Pro V520", MemoryTotalGB: 8}}
	crossVendorSrc := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{gpu1, amd1},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 9000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name: "cross-pool",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "n1", IP: "10.0.0.3", Port: 9000},
						{Name: "a1", IP: "10.0.0.5", Port: 9000},
					},
				},
			},
		},
	}

	_, err = recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", crossVendorSrc, SizingPolicyRightSizeUp, 3, 85.0, hooks)
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrCrossVendorConsolidation))

	resCrossAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", crossVendorSrc, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	require.Len(t, resCrossAuto[0].TargetInfra.NodeGroups, 2)
	require.Len(t, resCrossAuto[0].TargetNlbList, 2)
}

func TestRecommendGpuNlb_IntraNodeHeterogeneousGpuInNlb(t *testing.T) {
	hooks := testRecommendationHooks(mockGpuNlbCatalog(true))

	// Single physical server inside an NLB backend contains heterogeneous GPUs (1x A100 + 1x T4)
	heteroHost := baseOsAndCpuNode("host-hetero", "10.0.1.10", 16, 64)
	heteroHost.GPUCards = []onpremmodel.GpuCardProperty{
		{Vendor: "NVIDIA", Model: "NVIDIA A100", MemoryTotalGB: 40},
		{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16},
	}
	t4Host := baseOsAndCpuNode("host-t4", "10.0.1.11", 8, 32)
	t4Host.GPUCards = []onpremmodel.GpuCardProperty{
		{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16},
	}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{heteroHost, t4Host},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name: "ai-gw",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "h1", IP: "10.0.1.10", Port: 8000},
						{Name: "h2", IP: "10.0.1.11", Port: 8000},
					},
				},
			},
		},
	}

	res, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", src, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	// Two homogeneous groups after 1:K LAU + N:1 partition: 1x A100 group (Size=1) and 1x T4 group (Size=2)
	cand := res[0]
	require.Len(t, cand.TargetInfra.NodeGroups, 2)
	require.Len(t, cand.TargetNlbList, 2)

	sizesByName := map[string]int{}
	srcLineageByName := map[string]string{}
	for _, ng := range cand.TargetInfra.NodeGroups {
		sizesByName[ng.Name] = ng.NodeGroupSize
		srcLineageByName[ng.Name] = ng.Label[lineage.LabelSourceMachineIds]
	}
	assert.Equal(t, 1, sizesByName["ng-ai-gw-nvidia-a100-40gb"])
	assert.Equal(t, "host-hetero", srcLineageByName["ng-ai-gw-nvidia-a100-40gb"])
	assert.Equal(t, 2, sizesByName["ng-ai-gw-nvidia-t4-16gb"])
	assert.Equal(t, "host-hetero,host-t4", srcLineageByName["ng-ai-gw-nvidia-t4-16gb"])
}

func TestRecommendGpuNlb_MultiNlbOverlap_PortScopedFirewall(t *testing.T) {
	hooks := testRecommendationHooks(mockGpuNlbCatalog(true))

	webNode := baseOsAndCpuNode("web-1", "10.0.2.1", 4, 16)
	sharedNode := baseOsAndCpuNode("shared-gpu", "10.0.2.2", 8, 32)
	sharedNode.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	sharedNode.FirewallTable = []onpremmodel.FirewallRuleProperty{
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "80"},
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "8000"},
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "10.0.0.0/8", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "22"},
	}
	aiNode := baseOsAndCpuNode("ai-2", "10.0.2.3", 8, 32)
	aiNode.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{webNode, sharedNode, aiNode},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 80, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "web-pool",
					Servers: []onpremmodel.NlbServerProperty{{Name: "w1", IP: "10.0.2.1", Port: 80}, {Name: "sh", IP: "10.0.2.2", Port: 80}},
				},
			},
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "ai-pool",
					Servers: []onpremmodel.NlbServerProperty{{Name: "sh", IP: "10.0.2.2", Port: 8000}, {Name: "a2", IP: "10.0.2.3", Port: 8000}},
				},
			},
		},
	}

	res, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", src, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	cand := res[0]
	// web-pool splits into web-pool-cpu (1) and web-pool-gpu (1); ai-pool stays homogeneous ai-pool (2)
	require.Len(t, cand.TargetInfra.NodeGroups, 3)
	require.Len(t, cand.TargetNlbList, 3)

	sgById := map[string]cloudmodel.SecurityGroupReq{}
	for _, sg := range cand.TargetSecurityGroupList {
		sgById[sg.Name] = sg
	}

	// Verify ai-pool SG retains 8000 and 22, but filters foreign NLB port 80 from sharedNode
	for _, ng := range cand.TargetInfra.NodeGroups {
		if ng.Name == "ng-ai-pool" {
			sg := sgById[ng.SecurityGroupIds[0]]
			require.NotNil(t, sg.FirewallRules)
			ports := map[string]bool{}
			for _, r := range *sg.FirewallRules {
				ports[r.Ports] = true
			}
			assert.True(t, ports["8000"])
			assert.True(t, ports["22"])
			assert.False(t, ports["80"], "foreign NLB port 80 must be filtered from ai-pool SG")
		}
	}
}

func TestRecommendGpuNlb_AutoClosedLoop_PhantomSpecPruned(t *testing.T) {
	// Disable p4d (A100 80GB host) in catalog so consolidated A100+T4 becomes a phantom spec while dividing works
	hooks := testRecommendationHooks(mockGpuNlbCatalog(false))

	a100Host := baseOsAndCpuNode("srv-a100", "10.0.3.1", 16, 64)
	a100Host.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA A100", MemoryTotalGB: 80}}
	t4Host := baseOsAndCpuNode("srv-t4", "10.0.3.2", 8, 32)
	t4Host.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{a100Host, t4Host},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name: "phantom-pool",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "a100", IP: "10.0.3.1", Port: 8080},
						{Name: "t4", IP: "10.0.3.2", Port: 8080},
					},
				},
			},
		},
	}

	// rightSizeUp hits phantom spec -> returns nothing-to-recommend
	resRightSize, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", src, SizingPolicyRightSizeUp, 3, 85.0, hooks)
	require.NoError(t, err)
	assert.Equal(t, "nothing-to-recommend", resRightSize[0].Status)

	// auto detects infeasible consolidation and still fails closed unless BOTH are infeasible; verify if A100 80GB missing -> let's test CPU+GPU phantom spec case
	cpuNode := baseOsAndCpuNode("srv-cpu", "10.0.3.3", 16, 64)
	// When mixed CPU (huge host vCPU/RAM) + L4 creates consolidated spec that misses catalog, but CPU and L4 each match individually:
	customMatcher := func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
		if hasGpu(node) && node.CPU.Cpus*node.CPU.Threads >= 64 {
			// Consolidated (Max vCPU=64 + L4 GPU) does not exist in catalog (phantom spec)
			return nil
		}
		return mockGpuNlbCatalog(false)(ctx, csp, region, node, limitSpecs, limitImages)
	}
	hugeCpu := baseOsAndCpuNode("srv-huge-cpu", "10.0.3.10", 64, 256)
	l4Host := baseOsAndCpuNode("srv-l4", "10.0.3.11", 8, 32)
	l4Host.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24}}
	phantomSrc := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{hugeCpu, l4Host, cpuNode},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "mixed-phantom",
					Servers: []onpremmodel.NlbServerProperty{{Name: "hc", IP: "10.0.3.10", Port: 8080}, {Name: "gpu", IP: "10.0.3.11", Port: 8080}},
				},
			},
		},
	}

	resAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", phantomSrc, SizingPolicyAuto, 3, 85.0, testRecommendationHooks(customMatcher))
	require.NoError(t, err)
	require.NotEmpty(t, resAuto)
	assert.NotEqual(t, "nothing-to-recommend", resAuto[0].Status)
	// Expected NodeGroups: 2 divided (cpu, gpu) + 1 standalone = 3
	assert.Len(t, resAuto[0].TargetInfra.NodeGroups, 3)
	assert.Contains(t, resAuto[0].TargetInfra.NodeGroups[0].Label[labelTopologyDecision], "infeasible-spec-pruned")
}

func TestRecommendGpuNlb_ZeroSwInvarianceAndReproducibility(t *testing.T) {
	hooks := testRecommendationHooks(mockGpuNlbCatalog(true))

	cpu1 := baseOsAndCpuNode("srv-1", "10.0.4.1", 4, 16)
	gpu1 := baseOsAndCpuNode("srv-2", "10.0.4.2", 8, 32)
	gpu1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}

	swAbsent := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{cpu1, gpu1},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "vllm-pool",
					Servers: []onpremmodel.NlbServerProperty{{Name: "n1", IP: "10.0.4.1", Port: 8000}, {Name: "n2", IP: "10.0.4.2", Port: 8000}},
				},
			},
		},
	}

	// Add SW/Workload hints (ProcessName + Role) to a copy and verify output is byte-identical
	swInjected := swAbsent
	swInjected.Nodes = make([]onpremmodel.NodeProperty, len(swAbsent.Nodes))
	copy(swInjected.Nodes, swAbsent.Nodes)
	swInjected.Nodes[0].Role = "api-gateway"
	swInjected.Nodes[0].ListeningPorts = []onpremmodel.ListeningPortProperty{{Protocol: "tcp", Port: 8000, ProcessName: "nginx"}}
	swInjected.Nodes[1].Role = "llm-worker"
	swInjected.Nodes[1].ListeningPorts = []onpremmodel.ListeningPortProperty{{Protocol: "tcp", Port: 8000, ProcessName: "vllm"}}

	resAbsent, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", swAbsent, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	bytesAbsent, err := json.Marshal(resAbsent)
	require.NoError(t, err)

	resInjected, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", swInjected, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	bytesInjected, err := json.Marshal(resInjected)
	require.NoError(t, err)

	assert.Equal(t, string(bytesAbsent), string(bytesInjected), "recommendation output must be invariant to SW/workload hints")

	// Verify 100-run deterministic hash reproducibility
	baselineHash := sha256.Sum256(bytesAbsent)
	expectedHex := hex.EncodeToString(baselineHash[:])
	for i := 0; i < 100; i++ {
		resRun, runErr := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", swAbsent, SizingPolicyAuto, 3, 85.0, hooks)
		require.NoError(t, runErr)
		b, mErr := json.Marshal(resRun)
		require.NoError(t, mErr)
		h := sha256.Sum256(b)
		assert.Equal(t, expectedHex, hex.EncodeToString(h[:]))
	}
}

func TestRecommendGpuNlb_DisjointZoneMultiSubnetAllocation(t *testing.T) {
	ubuntuImg := cloudmodel.ImageInfo{
		Id:             "alibaba+ap-northeast-2+ubuntu-2204",
		CspImageName:   "ubuntu_22_04_x64.vhd",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "x86_64",
		OSDistribution: "ubuntu jammy 22.04",
	}
	disjointMatcher := func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
		if !hasGpu(node) {
			return []CompatibleSpecImagePair{{
				Spec: cloudmodel.SpecInfo{
					Id:               "alibaba+ap-northeast-2+cpu-zone-a",
					CspSpecName:      "ecs.u2i-c1m8.large",
					VCPU:             2,
					MemoryGiB:        16,
					CostPerHour:      0.11,
					AvailableZones:   []string{"ap-northeast-2a"},
					UnavailableZones: []string{"ap-northeast-2b"},
				},
				Image: ubuntuImg,
			}}
		}
		return []CompatibleSpecImagePair{{
			Spec: cloudmodel.SpecInfo{
				Id:                  "alibaba+ap-northeast-2+gpu-zone-b",
				CspSpecName:         "ecs.ebmgn8is.32xlarge",
				VCPU:                128,
				MemoryGiB:           1024,
				AcceleratorCount:    8,
				AcceleratorMemoryGB: 80,
				AcceleratorModel:    "NVIDIA H20",
				CostPerHour:         41.0,
				AvailableZones:      []string{"ap-northeast-2b"},
				UnavailableZones:    []string{"ap-northeast-2a"},
			},
			Image: ubuntuImg,
		}}
	}

	hooks := testRecommendationHooks(disjointMatcher)
	hooks.recommendVNet = func(ctx context.Context, csp, region string, srcInfra onpremmodel.OnpremInfra) ([]cloudmodel.VNetReq, error) {
		return []cloudmodel.VNetReq{{
			CidrBlock:      "10.0.0.0/22",
			SubnetInfoList: []cloudmodel.SubnetReq{{IPv4_CIDR: "10.0.1.0/24"}},
		}}, nil
	}
	hooks.enforceZonalSubnet = alignCandidateSubnetsWithSpecZones

	cpuNode := baseOsAndCpuNode("cpu-1", "10.0.1.11", 4, 16)
	gpuNode := baseOsAndCpuNode("gpu-1", "10.0.1.12", 8, 32)
	gpuNode.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{cpuNode, gpuNode},
		NLBs: []onpremmodel.NlbProperty{{
			Software: "haproxy",
			Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
			Backend: onpremmodel.NlbBackendProperty{
				Name:    "disjoint-pool",
				Servers: []onpremmodel.NlbServerProperty{{Name: "c1", IP: "10.0.1.11", Port: 8080}, {Name: "g1", IP: "10.0.1.12", Port: 8080}},
			},
		}},
	}

	res, err := recommendGpuNlbInfraWithHooks(context.Background(), "alibaba", "ap-northeast-2", src, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	cand := res[0]
	require.Len(t, cand.TargetVNet.SubnetInfoList, 2)
	assert.Equal(t, "ap-northeast-2a", cand.TargetVNet.SubnetInfoList[0].Zone)
	assert.Equal(t, "ap-northeast-2b", cand.TargetVNet.SubnetInfoList[1].Zone)
	assert.Equal(t, "mig-subnet-01", cand.TargetInfra.NodeGroups[0].SubnetId)
	assert.Equal(t, "mig-subnet-02", cand.TargetInfra.NodeGroups[1].SubnetId)
}

func TestRecommendGpuNlb_MultiPortFirewallRuleProjection(t *testing.T) {
	hooks := testRecommendationHooks(mockGpuNlbCatalog(true))

	w1 := baseOsAndCpuNode("web-1", "10.0.5.10", 4, 16)
	w1.FirewallTable = []onpremmodel.FirewallRuleProperty{
		{Direction: "inbound", Protocol: "tcp", DstPorts: "80", SrcCIDR: "0.0.0.0/0", Action: "allow"},
	}

	shared := baseOsAndCpuNode("shared-gpu", "10.0.5.11", 8, 32)
	shared.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	shared.FirewallTable = []onpremmodel.FirewallRuleProperty{
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "22,80,8000"},
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "10.0.0.0/8", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "80, 8000"},
	}

	aiWorker := baseOsAndCpuNode("ai-1", "10.0.5.12", 8, 32)
	aiWorker.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	aiWorker.FirewallTable = []onpremmodel.FirewallRuleProperty{
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "8000"},
	}

	w1.FirewallTable = []onpremmodel.FirewallRuleProperty{
		{Action: "allow", Direction: "inbound", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*", DstCIDR: "0.0.0.0/0", DstPorts: "80"},
	}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{w1, shared, aiWorker},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 80, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "web-pool",
					Servers: []onpremmodel.NlbServerProperty{{Name: "w1", IP: "10.0.5.10", Port: 80}, {Name: "sg", IP: "10.0.5.11", Port: 80}},
				},
			},
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:    "ai-pool",
					Servers: []onpremmodel.NlbServerProperty{{Name: "sg", IP: "10.0.5.11", Port: 8000}, {Name: "a1", IP: "10.0.5.12", Port: 8000}},
				},
			},
		},
	}

	res, err := recommendGpuNlbInfraWithHooks(context.Background(), "aws", "ap-northeast-2", src, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	cand := res[0]
	sgByName := make(map[string]cloudmodel.SecurityGroupReq, len(cand.TargetSecurityGroupList))
	for _, sg := range cand.TargetSecurityGroupList {
		sgByName[sg.Name] = sg
	}

	for _, ng := range cand.TargetInfra.NodeGroups {
		rawRulePorts := map[string]bool{}
		portTokens := map[string]bool{}
		for _, sgID := range ng.SecurityGroupIds {
			if sg, ok := sgByName[sgID]; ok && sg.FirewallRules != nil {
				for _, rule := range *sg.FirewallRules {
					rawRulePorts[rule.Ports] = true
					for _, tok := range strings.Split(rule.Ports, ",") {
						if p := strings.TrimSpace(tok); p != "" {
							portTokens[p] = true
						}
					}
				}
			}
		}
		if strings.Contains(ng.Name, "web") {
			assert.True(t, portTokens["80"])
			assert.False(t, portTokens["8000"], "web pool SG must exclude foreign ai-pool port 8000")
			if strings.Contains(ng.Name, "gpu") {
				assert.True(t, portTokens["22"], "shared GPU SG in web pool should keep port 22 from multi-port rule")
			}
		}
		if strings.Contains(ng.Name, "ai") {
			assert.True(t, portTokens["8000"])
			assert.True(t, portTokens["22"], "ai pool SG should keep port 22 from multi-port rule")
			assert.False(t, portTokens["80"], "ai pool SG must exclude foreign web-pool port 80")
		}
	}
}

func testPartialPartitionCatalog(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
	ubuntuImg := cloudmodel.ImageInfo{
		Id:             "ibm+au-syd+ubuntu-2204",
		CspImageName:   "ubuntu-jammy-22.04-amd64",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "x86_64",
		OSDistribution: "ubuntu jammy 22.04",
	}
	if !hasGpu(node) {
		return []CompatibleSpecImagePair{{
			Spec: cloudmodel.SpecInfo{
				Id:          "ibm+au-syd+bx2-4x16",
				CspSpecName: "bx2-4x16",
				VCPU:        4,
				MemoryGiB:   16,
				CostPerHour: 0.184,
			},
			Image: ubuntuImg,
		}}
	}
	clusters := clusterGpuCards(node.GPUCards)
	primary := clusters[0]
	vendor := strings.ToLower(detectGpuVendor(primary.Vendor))
	model := detectGpuModel(primary.Model)

	if vendor == "amd" {
		return []CompatibleSpecImagePair{{
			Spec: cloudmodel.SpecInfo{
				Id:                  "ibm+au-syd+v520-8x32",
				CspSpecName:         "v520-8x32",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 8,
				AcceleratorModel:    "AMD Radeon Pro V520",
				CostPerHour:         0.379,
			},
			Image: ubuntuImg,
		}}
	}
	if model == "a100" {
		return []CompatibleSpecImagePair{{
			Spec: cloudmodel.SpecInfo{
				Id:                  "ibm+au-syd+gx3-16x128-1a100",
				CspSpecName:         "gx3-16x128-1a100",
				VCPU:                16,
				MemoryGiB:           128,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 80,
				AcceleratorModel:    "NVIDIA A100",
				CostPerHour:         4.50,
			},
			Image: ubuntuImg,
		}}
	}
	if model == "l4" {
		return []CompatibleSpecImagePair{{
			Spec: cloudmodel.SpecInfo{
				Id:                  "ibm+au-syd+gx3-16x64-1l4",
				CspSpecName:         "gx3-16x64-1l4",
				VCPU:                16,
				MemoryGiB:           64,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 24,
				AcceleratorModel:    "NVIDIA L4",
				CostPerHour:         1.898,
			},
			Image: ubuntuImg,
		}}
	}
	return []CompatibleSpecImagePair{{
		Spec: cloudmodel.SpecInfo{
			Id:                  "ibm+au-syd+gx2-8x64-1t4",
			CspSpecName:         "gx2-8x64-1t4",
			VCPU:                8,
			MemoryGiB:           64,
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 16,
			AcceleratorModel:    "NVIDIA T4",
			CostPerHour:         3.148,
		},
		Image: ubuntuImg,
	}}
}

func TestRecommendGpuNlb_AutoPartialPartition_MixedCpuAndHeteroGpu(t *testing.T) {
	hooks := testRecommendationHooks(testPartialPartitionCatalog)

	c1 := baseOsAndCpuNode("s7-cpu-1", "10.0.7.11", 4, 16)
	c2 := baseOsAndCpuNode("s7-cpu-2", "10.0.7.12", 4, 16)
	t1 := baseOsAndCpuNode("s7-t4-1", "10.0.7.13", 8, 32)
	t1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	t2 := baseOsAndCpuNode("s7-t4-2", "10.0.7.14", 8, 32)
	t2.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	l1 := baseOsAndCpuNode("s7-l4-1", "10.0.7.15", 16, 64)
	l1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24}}
	mgmt := baseOsAndCpuNode("s7-mgmt-1", "10.0.7.21", 4, 16)

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{c1, c2, t1, t2, l1, mgmt},
		NLBs: []onpremmodel.NlbProperty{{
			Software: "haproxy",
			Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
			Backend: onpremmodel.NlbBackendProperty{
				Name: "s7-tier3-pool",
				Servers: []onpremmodel.NlbServerProperty{
					{Name: "c1", IP: "10.0.7.11", Port: 8000},
					{Name: "c2", IP: "10.0.7.12", Port: 8000},
					{Name: "t1", IP: "10.0.7.13", Port: 8000},
					{Name: "t2", IP: "10.0.7.14", Port: 8000},
					{Name: "l1", IP: "10.0.7.15", Port: 8000},
				},
			},
		}},
	}

	resRight, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyRightSizeUp, 3, 85.0, hooks)
	require.NoError(t, err)
	require.Len(t, resRight[0].TargetInfra.NodeGroups, 2)
	require.Len(t, resRight[0].TargetNlbList, 1)

	resDiv, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyDividing, 3, 85.0, hooks)
	require.NoError(t, err)
	require.Len(t, resDiv[0].TargetInfra.NodeGroups, 4)
	require.Len(t, resDiv[0].TargetNlbList, 3)

	resAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resAuto)
	// Expected NodeGroups: {CPU} + {T4, L4} + standalone = 3 NodeGroups and 2 NLBs
	require.Len(t, resAuto[0].TargetInfra.NodeGroups, 3)
	require.Len(t, resAuto[0].TargetNlbList, 2)

	costRight := CalculateCandidateHourlyCost(resRight[0]).TotalHourlyCostUSD
	costDiv := CalculateCandidateHourlyCost(resDiv[0]).TotalHourlyCostUSD
	costAuto := CalculateCandidateHourlyCost(resAuto[0]).TotalHourlyCostUSD
	assert.Less(t, costAuto, costRight)
	assert.Less(t, costAuto, costDiv)
	assert.Contains(t, resAuto[0].TargetInfra.NodeGroups[0].Label[labelTopologyDecision], "auto->partial(2/3-tiers")
}

func TestRecommendGpuNlb_AutoPartialPartition_ThreeSameVendorGpuTiers(t *testing.T) {
	hooks := testRecommendationHooks(testPartialPartitionCatalog)

	t1 := baseOsAndCpuNode("gpu-t4-1", "10.0.8.11", 8, 32)
	t1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	t2 := baseOsAndCpuNode("gpu-t4-2", "10.0.8.12", 8, 32)
	t2.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	l1 := baseOsAndCpuNode("gpu-l4-1", "10.0.8.13", 16, 64)
	l1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24}}
	a1 := baseOsAndCpuNode("gpu-a100-1", "10.0.8.14", 16, 128)
	a1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA A100", MemoryTotalGB: 80}}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{t1, t2, l1, a1},
		NLBs: []onpremmodel.NlbProperty{{
			Software: "haproxy",
			Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
			Backend: onpremmodel.NlbBackendProperty{
				Name: "gpu3-pool",
				Servers: []onpremmodel.NlbServerProperty{
					{Name: "t1", IP: "10.0.8.11", Port: 8000},
					{Name: "t2", IP: "10.0.8.12", Port: 8000},
					{Name: "l1", IP: "10.0.8.13", Port: 8000},
					{Name: "a1", IP: "10.0.8.14", Port: 8000},
				},
			},
		}},
	}

	resAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resAuto)
	// {T4, L4} merges into 1 GPU group (3 nodes) and {A100} stays isolated (1 node) = 2 NLBs, 2 NGs
	require.Len(t, resAuto[0].TargetInfra.NodeGroups, 2)
	require.Len(t, resAuto[0].TargetNlbList, 2)
	assert.Contains(t, resAuto[0].TargetInfra.NodeGroups[0].Label[labelTopologyDecision], "auto->partial(2/3-tiers")

	sizes := []int{resAuto[0].TargetInfra.NodeGroups[0].NodeGroupSize, resAuto[0].TargetInfra.NodeGroups[1].NodeGroupSize}
	assert.ElementsMatch(t, []int{3, 1}, sizes)
}

func TestRecommendGpuNlb_AutoPartialPartition_CrossVendorHardCut(t *testing.T) {
	hooks := testRecommendationHooks(testPartialPartitionCatalog)

	c1 := baseOsAndCpuNode("hyb-cpu-1", "10.0.9.11", 4, 16)
	t1 := baseOsAndCpuNode("hyb-t4-1", "10.0.9.12", 8, 32)
	t1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	t2 := baseOsAndCpuNode("hyb-t4-2", "10.0.9.13", 8, 32)
	t2.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}}
	l1 := baseOsAndCpuNode("hyb-l4-1", "10.0.9.14", 16, 64)
	l1.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24}}
	amd := baseOsAndCpuNode("hyb-amd-1", "10.0.9.15", 8, 32)
	amd.GPUCards = []onpremmodel.GpuCardProperty{{Vendor: "AMD", Model: "AMD Radeon Pro V520", MemoryTotalGB: 8}}

	src := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{c1, t1, t2, l1, amd},
		NLBs: []onpremmodel.NlbProperty{{
			Software: "haproxy",
			Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8080, Protocol: "tcp"},
			Backend: onpremmodel.NlbBackendProperty{
				Name: "hybrid4-pool",
				Servers: []onpremmodel.NlbServerProperty{
					{Name: "c1", IP: "10.0.9.11", Port: 8000},
					{Name: "t1", IP: "10.0.9.12", Port: 8000},
					{Name: "t2", IP: "10.0.9.13", Port: 8000},
					{Name: "l1", IP: "10.0.9.14", Port: 8000},
					{Name: "a1", IP: "10.0.9.15", Port: 8000},
				},
			},
		}},
	}

	_, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyRightSizeUp, 3, 85.0, hooks)
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrCrossVendorConsolidation))

	resAuto, err := recommendGpuNlbInfraWithHooks(context.Background(), "ibm", "au-syd", src, SizingPolicyAuto, 3, 85.0, hooks)
	require.NoError(t, err)
	require.NotEmpty(t, resAuto)
	// Hard cut preserves AMD isolation while merging {T4, L4}: {CPU}, {AMD V520}, {NVIDIA T4, NVIDIA L4}
	require.Len(t, resAuto[0].TargetInfra.NodeGroups, 3)
	require.Len(t, resAuto[0].TargetNlbList, 3)
	assert.Contains(t, resAuto[0].TargetInfra.NodeGroups[0].Label[labelTopologyDecision], "auto->partial(3/4-tiers")

	for _, ng := range resAuto[0].TargetInfra.NodeGroups {
		lineageIds := ng.Label[lineage.LabelSourceMachineIds]
		hasNvidia := strings.Contains(lineageIds, "hyb-t4") || strings.Contains(lineageIds, "hyb-l4")
		hasAmd := strings.Contains(lineageIds, "hyb-amd")
		assert.False(t, hasNvidia && hasAmd, "cross-vendor boundary must never merge NVIDIA and AMD in NodeGroup %s", ng.Name)
	}
}
