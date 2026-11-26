// Package recommendation provides infrastructure and resource recommendation algorithms for CM-Beetle.
package recommendation

import (
	"cmp"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"slices"
	"sort"
	"strconv"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

// Sizing policy constants for GPU+NLB infrastructure recommendation.
const (
	SizingPolicyRightSizeUp = "rightSizeUp"
	SizingPolicyDividing    = "dividing"
	SizingPolicyAuto        = "auto"

	labelHardwareCase     = "cm-hardware-case"
	labelSizingPolicy     = "cm-sizing-policy"
	labelTopologyDecision = "cm-topology-decision"

	// Objective function weights for multi-objective adaptive policy evaluation:
	// J(T) = w_m * (WMR/100) - w_c * NormCost - w_e * SplitPenalty
	autoWeightMatchRate    = 0.50
	autoWeightCost         = 0.40
	autoWeightSplitPenalty = 0.10
	maxExhaustiveAutoTiers = 6
)

// ErrCrossVendorConsolidation indicates cross-vendor GPU consolidation cannot be fulfilled under rightSizeUp.
var ErrCrossVendorConsolidation = errors.New("cross-vendor GPU consolidation is physically impossible under rightSizeUp")

// HardwareCase represents the hardware accelerator categorization of an NLB backend pool.
type HardwareCase int

const (
	// CasePureCpu indicates all backend servers are CPU-only without accelerators.
	CasePureCpu HardwareCase = 0
	// CaseMixedCpuGpu indicates both CPU-only servers and GPU servers are present.
	CaseMixedCpuGpu HardwareCase = 1
	// CaseCrossVendorGpu indicates GPUs from multiple vendors are present.
	CaseCrossVendorGpu HardwareCase = 2
	// CaseHeteroGpuModel indicates same-vendor GPUs with differing models or VRAM.
	CaseHeteroGpuModel HardwareCase = 3
	// CaseHomoGpuModel indicates all servers share identical GPU model and VRAM.
	CaseHomoGpuModel HardwareCase = 4
)

// String returns a human-readable name for the HardwareCase.
func (c HardwareCase) String() string {
	switch c {
	case CasePureCpu:
		return "Case 0 (Pure CPU)"
	case CaseMixedCpuGpu:
		return "Case 1 (Mixed CPU+GPU)"
	case CaseCrossVendorGpu:
		return "Case 2 (Cross-Vendor GPU)"
	case CaseHeteroGpuModel:
		return "Case 3 (Heterogeneous GPU Models)"
	case CaseHomoGpuModel:
		return "Case 4 (Homogeneous GPU Models)"
	default:
		return fmt.Sprintf("Unknown Case (%d)", int(c))
	}
}

// ValidateSizingPolicy normalizes and validates a sizing policy without silent fallback.
func ValidateSizingPolicy(policy string) (string, error) {
	trimmed := strings.TrimSpace(policy)
	if trimmed == "" || strings.EqualFold(trimmed, SizingPolicyRightSizeUp) {
		return SizingPolicyRightSizeUp, nil
	}
	if strings.EqualFold(trimmed, SizingPolicyDividing) {
		return SizingPolicyDividing, nil
	}
	if strings.EqualFold(trimmed, SizingPolicyAuto) {
		return SizingPolicyAuto, nil
	}
	return "", fmt.Errorf(
		"unsupported sizingPolicy '%s'; must be '%s', '%s', or '%s'",
		trimmed, SizingPolicyRightSizeUp, SizingPolicyDividing, SizingPolicyAuto,
	)
}

// ClassifyBackendHardware evaluates the hardware accelerator profile of backend servers or LAUs.
func ClassifyBackendHardware(members []onpremmodel.NodeProperty) HardwareCase {
	gpuCount := 0
	cpuCount := 0
	vendors := make(map[string]bool)
	modelKeys := make(map[string]bool)

	for _, m := range members {
		if hasGpu(m) {
			gpuCount++
			for _, card := range m.GPUCards {
				v := strings.ToLower(detectGpuVendor(card.Vendor))
				if v == "" || v == strings.ToLower(GpuVendorOther) {
					v = strings.ToLower(detectGpuVendor(card.Model))
				}
				if v == "" || v == strings.ToLower(GpuVendorOther) {
					v = strings.ToLower(strings.TrimSpace(card.Vendor))
				}
				if v == "" {
					v = "unknown"
				}
				vendors[v] = true
				key := fmt.Sprintf("%s:%s:%.1f", v, strings.ToLower(strings.TrimSpace(card.Model)), card.MemoryTotalGB)
				modelKeys[key] = true
			}
		} else {
			cpuCount++
		}
	}

	if gpuCount == 0 {
		return CasePureCpu
	}
	// Check cross-vendor before CPU+GPU mix so NVIDIA+AMD+CPU never bypasses vendor isolation.
	if len(vendors) > 1 {
		return CaseCrossVendorGpu
	}
	if cpuCount > 0 && gpuCount > 0 {
		return CaseMixedCpuGpu
	}
	if len(modelKeys) > 1 {
		return CaseHeteroGpuModel
	}
	return CaseHomoGpuModel
}

// recommendationHooks encapsulates external catalog and provider dependencies for testability.
type recommendationHooks struct {
	recommendVNet      func(ctx context.Context, csp, region string, srcInfra onpremmodel.OnpremInfra) ([]cloudmodel.VNetReq, error)
	recommendSG        func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty) (cloudmodel.SecurityGroupReq, error)
	matchCatalog       func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair
	preflightCheck     func(ctx context.Context, specId, imageId, currentCspImageName, rootDiskType string) (CspProvisioningPrecheck, error)
	enforceZonalSubnet func(candidate *cloudmodel.RecommendedInfra)
}

func defaultRecommendationHooks() recommendationHooks {
	return recommendationHooks{
		recommendVNet:      RecommendVNet,
		recommendSG:        RecommendSecurityGroup,
		matchCatalog:       defaultCatalogMatcher,
		preflightCheck:     PreflightCheckCspProvisioning,
		enforceZonalSubnet: alignCandidateSubnetsWithSpecZones,
	}
}

func defaultCatalogMatcher(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair {
	specList, _, specErr := RecommendNodeSpecs(ctx, csp, region, node, limitSpecs)
	if specErr != nil {
		log.Warn().Err(specErr).Str("machineId", node.MachineId).Msg("failed to recommend specs")
	}
	imageList, imgErr := RecommendVmOsImages(ctx, csp, region, node, limitImages)
	if imgErr != nil {
		log.Warn().Err(imgErr).Str("machineId", node.MachineId).Msg("failed to recommend images")
	}
	if len(specList) == 0 || len(imageList) == 0 {
		return nil
	}
	pairs, pairErr := FindCompatibleVmSpecAndImagePairs(specList, imageList, csp)
	if pairErr != nil {
		log.Warn().Err(pairErr).Str("machineId", node.MachineId).Msg("failed to find compatible spec-image pairs")
		return nil
	}
	return pairs
}

// cachedCatalogMatcher memoizes catalog lookups within a single recommendation call by hardware/OS profile.
type cachedCatalogMatcher struct {
	limitSpecs  int
	limitImages int
	matcher     func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair
	cache       map[string][]CompatibleSpecImagePair
}

func newCachedCatalogMatcher(
	limitSpecs, limitImages int,
	matcher func(ctx context.Context, csp, region string, node onpremmodel.NodeProperty, limitSpecs, limitImages int) []CompatibleSpecImagePair,
) *cachedCatalogMatcher {
	return &cachedCatalogMatcher{
		limitSpecs:  limitSpecs,
		limitImages: limitImages,
		matcher:     matcher,
		cache:       make(map[string][]CompatibleSpecImagePair),
	}
}

func (m *cachedCatalogMatcher) lookup(ctx context.Context, csp, region string, node onpremmodel.NodeProperty) []CompatibleSpecImagePair {
	key := buildNodeProfileCacheKey(csp, region, node)
	if pairs, ok := m.cache[key]; ok {
		return pairs
	}
	pairs := m.matcher(ctx, csp, region, node, m.limitSpecs, m.limitImages)
	m.cache[key] = pairs
	return pairs
}

func buildNodeProfileCacheKey(csp, region string, node onpremmodel.NodeProperty) string {
	threads := node.CPU.Threads
	if threads == 0 {
		threads = 1
	}
	vcpu := node.CPU.Cpus * threads
	clusters := clusterGpuCards(node.GPUCards)
	gpuSig := "none"
	if len(clusters) > 0 {
		c := clusters[0]
		gpuSig = fmt.Sprintf("%s|%s|%.1f|%d", strings.ToLower(c.Vendor), strings.ToLower(c.Model), c.MemoryTotalGB, c.Count)
	}
	return fmt.Sprintf("%s|%s|%s|%d|%d|%s|%s|%s|%s",
		csp, region,
		strings.ToLower(node.CPU.Architecture),
		vcpu, node.Memory.TotalSize,
		strings.ToLower(node.OS.ID), strings.ToLower(node.OS.VersionID), strings.ToLower(node.OS.VersionCodename),
		gpuSig,
	)
}

// RecommendGpuNlbInfraCandidates recommends cloud infrastructure candidates using only static hardware and network metadata.
// It resolves conflicting GPU and NLB constraints via two-stage nested decomposition and multi-objective topology optimization.
func RecommendGpuNlbInfraCandidates(
	ctx context.Context,
	desiredCsp, desiredRegion string,
	srcInfra onpremmodel.OnpremInfra,
	sizingPolicy string,
	limit int,
	minMatchRate float64,
) ([]cloudmodel.RecommendedInfra, error) {
	return recommendGpuNlbInfraWithHooks(
		ctx, desiredCsp, desiredRegion, srcInfra, sizingPolicy, limit, minMatchRate, defaultRecommendationHooks(),
	)
}

func recommendGpuNlbInfraWithHooks(
	ctx context.Context,
	desiredCsp, desiredRegion string,
	srcInfra onpremmodel.OnpremInfra,
	sizingPolicy string,
	limit int,
	minMatchRate float64,
	hooks recommendationHooks,
) ([]cloudmodel.RecommendedInfra, error) {
	srcInfra = SanitizeSourceInfra(srcInfra)
	if len(srcInfra.Nodes) == 0 {
		return nil, fmt.Errorf("sourceInfra.nodes is empty")
	}

	policy, err := ValidateSizingPolicy(sizingPolicy)
	if err != nil {
		return nil, err
	}

	if limit <= 0 {
		limit = 3
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)
	limitSpecs := GetDefaultSpecsLimit()
	limitImages := GetDefaultImagesLimit()

	nodeByIP := buildNodeByIPIndex(srcInfra.Nodes)
	nodeByMachineId := buildNodeByMachineIdIndex(srcInfra.Nodes)

	// Stage 1: Correlate NLB backends with physical nodes and detect N:M overlaps.
	var resolvedNlbs []resolvedNlb
	var warnings []string
	nlbMembershipCount := make(map[string]int, len(srcInfra.Nodes))

	for _, nlb := range srcInfra.NLBs {
		backendName := nlb.Backend.Name
		var matchedMachineIds []string
		seenInBackend := make(map[string]bool)

		for _, srv := range nlb.Backend.Servers {
			node, ok := nodeByIP[srv.IP]
			if !ok {
				warnings = append(warnings, fmt.Sprintf("NLB '%s': server IP '%s' not found", backendName, srv.IP))
				continue
			}
			if !seenInBackend[node.MachineId] {
				matchedMachineIds = append(matchedMachineIds, node.MachineId)
				seenInBackend[node.MachineId] = true
			}
		}

		if len(matchedMachineIds) == 0 {
			warnings = append(warnings, fmt.Sprintf("NLB '%s': no matching nodes found", backendName))
			continue
		}

		backendPort, portWarning := resolvePort(nlb.Backend.Servers)
		if portWarning != "" {
			warnings = append(warnings, fmt.Sprintf("NLB '%s': %s", backendName, portWarning))
		}

		for _, machineId := range matchedMachineIds {
			nlbMembershipCount[machineId]++
		}

		if nlb.Backend.Balance != "" {
			warnings = append(warnings, fmt.Sprintf(
				"NLB '%s': load-balancing algorithm '%s' is mapped to CSP default algorithm",
				backendName, nlb.Backend.Balance,
			))
		}

		resolvedNlbs = append(resolvedNlbs, resolvedNlb{
			sourceNlb:        nlb,
			backendPort:      backendPort,
			memberMachineIds: matchedMachineIds,
		})
	}

	for mid, count := range nlbMembershipCount {
		if count > 1 {
			warnings = append(warnings, fmt.Sprintf(
				"Node '%s' belongs to %d NLB backends; port-scoped role isolation will create dedicated VMs per backend",
				mid, count,
			))
		}
	}

	// Stage 2: Build shared VNet and SSH key skeleton.
	connectionName := fmt.Sprintf("%s-%s", csp, region)
	skeleton := cloudmodel.RecommendedInfra{
		TargetCloud: cloudmodel.CloudProperty{Csp: csp, Region: region},
		TargetInfra: cloudmodel.InfraReq{
			Name:        "infra101",
			Description: fmt.Sprintf("Recommended infrastructure (%s policy)", policy),
			NodeGroups:  []cloudmodel.CreateNodeGroupReq{},
		},
	}

	recommendedVNetList, vnetErr := hooks.recommendVNet(ctx, csp, region, srcInfra)
	if vnetErr != nil || len(recommendedVNetList) == 0 {
		log.Warn().Err(vnetErr).Msg("failed to recommend vNet for GPU+NLB infra")
	} else {
		skeleton.TargetVNet = recommendedVNetList[0]
	}
	skeleton.TargetVNet.Name = "mig-vnet-01"
	skeleton.TargetVNet.Description = "a recommended vNet for migration"
	for i := range skeleton.TargetVNet.SubnetInfoList {
		skeleton.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		skeleton.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
	}

	firstSubnetId := ""
	if len(skeleton.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = skeleton.TargetVNet.SubnetInfoList[0].Name
	}

	skeleton.TargetSshKey = cloudmodel.SshKeyReq{
		Name:           "mig-sshkey-01",
		ConnectionName: connectionName,
		Description:    "SSH key pair for migration (Note: provided ONLY once, MUST be downloaded)",
	}

	// Stage 3: Two-stage nested decomposition (1:K LAU + N:1/N:M NLB) with constraint-aware auto policy optimization.
	catalogCache := newCachedCatalogMatcher(limitSpecs, limitImages, hooks.matchCatalog)

	var ngBlueprints []nodeGroupBlueprint
	var targetNlbList []cloudmodel.NlbReq
	var deduplicatedSgList []cloudmodel.SecurityGroupReq
	precomputedPairs := make(map[int][]CompatibleSpecImagePair)
	decisionNotes := make([]string, 0, len(resolvedNlbs))

	for _, rnlb := range resolvedNlbs {
		rawMembers := make([]onpremmodel.NodeProperty, 0, len(rnlb.memberMachineIds))
		for _, mid := range rnlb.memberMachineIds {
			if n, ok := nodeByMachineId[mid]; ok {
				rawMembers = append(rawMembers, n)
			}
		}

		// Decompose intra-node heterogeneous GPUs inside NLB backends into homogeneous LAUs first.
		lauMembers := DecomposeOnpremNodes(rawMembers)
		if len(lauMembers) > len(rawMembers) {
			warnings = append(warnings, fmt.Sprintf(
				"NLB '%s': decomposed %d physical member(s) into %d homogeneous accelerator units",
				rnlb.sourceNlb.Backend.Name, len(rawMembers), len(lauMembers),
			))
		}

		lauById := make(map[string]onpremmodel.NodeProperty, len(lauMembers))
		lauIds := make([]string, 0, len(lauMembers))
		for _, lau := range lauMembers {
			lauById[lau.MachineId] = lau
			lauIds = append(lauIds, lau.MachineId)
		}

		hwCase := ClassifyBackendHardware(lauMembers)
		log.Info().
			Str("backend", rnlb.sourceNlb.Backend.Name).
			Str("case", hwCase.String()).
			Str("policy", policy).
			Msg("Classified NLB backend hardware")

		bps, nlbs, sgs, cachedPairs, decision, bpErr := resolveBackendPolicyBlueprints(
			ctx, csp, region, connectionName, firstSubnetId, skeleton,
			rnlb, resolvedNlbs, nlbMembershipCount,
			lauMembers, lauIds, lauById, hwCase, policy, minMatchRate, hooks, catalogCache,
		)
		if bpErr != nil {
			return nil, bpErr
		}
		decisionNotes = append(decisionNotes, fmt.Sprintf("%s:%s", rnlb.sourceNlb.Backend.Name, decision))

		sgNameRemap := make(map[string]string, len(sgs))
		for _, sg := range sgs {
			origName := sg.Name
			exists, _, existingSg := containSg(deduplicatedSgList, sg)
			if !exists {
				deduplicatedSgList = append(deduplicatedSgList, sg)
				sgNameRemap[origName] = sg.Name
			} else {
				sgNameRemap[origName] = existingSg.Name
			}
		}

		baseIdx := len(ngBlueprints)
		for i := range bps {
			for sIdx, sgId := range bps[i].skeleton.SecurityGroupIds {
				if mapped, ok := sgNameRemap[sgId]; ok {
					bps[i].skeleton.SecurityGroupIds[sIdx] = mapped
				}
			}
			if i < len(cachedPairs) && cachedPairs[i] != nil {
				precomputedPairs[baseIdx+i] = cachedPairs[i]
			}
		}

		ngBlueprints = append(ngBlueprints, bps...)
		targetNlbList = append(targetNlbList, nlbs...)
	}

	// Stage 4: Standalone nodes (1 Server/LAU -> 1 NodeGroup, Size=1).
	var unrelatedNodes []onpremmodel.NodeProperty
	for _, node := range srcInfra.Nodes {
		if nlbMembershipCount[node.MachineId] == 0 {
			unrelatedNodes = append(unrelatedNodes, node)
		}
	}
	expandedUnrelated := DecomposeOnpremNodes(unrelatedNodes)

	shortIds := buildShortMachineIdMap(srcInfra.Nodes)
	usedNames := make(map[string]bool, len(ngBlueprints)+len(expandedUnrelated))
	for _, bp := range ngBlueprints {
		usedNames[bp.skeleton.Name] = true
	}

	for i, node := range expandedUnrelated {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		ngName, origMachineId := buildNodeGroupNameAndSrcMachineId(node, shortIds, usedNames)

		sg, sgErr := hooks.recommendSG(ctx, csp, region, node)
		if sgErr != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return nil, ctxErr
			}
			log.Warn().Err(sgErr).Str("machineId", node.MachineId).Msg("failed to recommend SG for standalone node")
		}
		exists, _, existingSg := containSg(deduplicatedSgList, sg)
		if !exists {
			sg.Name = fmt.Sprintf("mig-sg-%02d", len(deduplicatedSgList)+1)
			sg.ConnectionName = connectionName
			sg.Description = fmt.Sprintf("Recommended security group for %s", node.MachineId)
			sg.VNetId = skeleton.TargetVNet.Name
			deduplicatedSgList = append(deduplicatedSgList, sg)
		} else {
			sg = existingSg
		}

		minDisk := getCspMinRootDiskSizeGB(csp)
		if hasGpu(node) && minDisk < 100 {
			minDisk = 100
		}
		rootDiskSize := max(int(node.RootDisk.TotalSize), minDisk)

		ngBlueprints = append(ngBlueprints, nodeGroupBlueprint{
			representativeNode: node,
			isNlbRelated:       false,
			skeleton: cloudmodel.CreateNodeGroupReq{
				ConnectionName:   connectionName,
				Name:             ngName,
				VNetId:           skeleton.TargetVNet.Name,
				SubnetId:         firstSubnetId,
				SecurityGroupIds: []string{sg.Name},
				SshKeyId:         skeleton.TargetSshKey.Name,
				RootDiskType:     "",
				RootDiskSize:     rootDiskSize,
				NodeGroupSize:    1,
				Description:      fmt.Sprintf("Standalone VM %02d for %s", i+1, node.MachineId),
				Label: map[string]string{
					lineage.LabelSourceMachineIds: origMachineId,
					labelSizingPolicy:             policy,
				},
			},
		})
	}

	skeleton.TargetSecurityGroupList = deduplicatedSgList

	// Stage 5: Ensure catalog spec-image pairs are populated for every blueprint.
	pairsByGroup := make([][]CompatibleSpecImagePair, len(ngBlueprints))
	for ngIdx, bp := range ngBlueprints {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if pairs, ok := precomputedPairs[ngIdx]; ok {
			pairsByGroup[ngIdx] = pairs
			continue
		}
		pairsByGroup[ngIdx] = catalogCache.lookup(ctx, csp, region, bp.representativeNode)
	}

	// Stage 6: Apply CSP-specific NLB constraints and assemble Pareto candidates.
	targetNlbList = sanitizeNlbListByCsp(targetNlbList, csp)

	maxCandidates := 0
	for _, pairs := range pairsByGroup {
		if len(pairs) > maxCandidates {
			maxCandidates = len(pairs)
		}
	}
	if maxCandidates > limit {
		maxCandidates = limit
	}

	if maxCandidates == 0 {
		skeleton.Status = "nothing-to-recommend"
		skeleton.Description = "No compatible VM spec and OS image pairs found for infrastructure."
		skeleton.TargetNlbList = targetNlbList
		return []cloudmodel.RecommendedInfra{skeleton}, nil
	}

	var candidates []cloudmodel.RecommendedInfra
	for cIdx := 0; cIdx < maxCandidates; cIdx++ {
		candidateNodeGroups := make([]cloudmodel.CreateNodeGroupReq, len(ngBlueprints))
		for ngIdx := range ngBlueprints {
			candidateNodeGroups[ngIdx] = ngBlueprints[ngIdx].skeleton
		}

		var candidateSpecList []cloudmodel.SpecInfo
		var candidateImageList []cloudmodel.ImageInfo

		for ngIdx := range ngBlueprints {
			pairs := pairsByGroup[ngIdx]
			if len(pairs) == 0 {
				continue
			}
			pairIdx := cIdx
			if pairIdx >= len(pairs) {
				pairIdx = 0
			}
			selectedPair := pairs[pairIdx]
			selectedSpec := selectedPair.Spec
			selectedImage := selectedPair.Image

			if hooks.preflightCheck != nil {
				precheck, prefErr := hooks.preflightCheck(
					ctx, selectedSpec.Id, selectedImage.Id, selectedImage.CspImageName, "",
				)
				if prefErr == nil {
					if precheck.SuggestedSystemDisk != "" {
						candidateNodeGroups[ngIdx].RootDiskType = precheck.SuggestedSystemDisk
					}
					if precheck.ResolvedCspImageName != "" && precheck.ResolvedCspImageName != selectedImage.CspImageName {
						candidateNodeGroups[ngIdx].CspImageName = precheck.ResolvedCspImageName
					}
				}
			}

			matchRateVec := calculateMatchRateVector(csp, ngBlueprints[ngIdx].representativeNode, selectedSpec, selectedImage)
			candidateNodeGroups[ngIdx].SpecId = selectedSpec.Id
			candidateNodeGroups[ngIdx].ImageId = selectedImage.Id
			candidateNodeGroups[ngIdx].Description = fmt.Sprintf(
				"%s | %s",
				ngBlueprints[ngIdx].skeleton.Description,
				matchRateVec.FormatMatchRateSummary(),
			)

			specAdded := false
			for _, sp := range candidateSpecList {
				if sp.Id == selectedSpec.Id {
					specAdded = true
					break
				}
			}
			if !specAdded {
				candidateSpecList = append(candidateSpecList, selectedSpec)
			}

			imgAdded := false
			for _, img := range candidateImageList {
				if img.CspImageName == selectedImage.CspImageName {
					imgAdded = true
					break
				}
			}
			if !imgAdded {
				candidateImageList = append(candidateImageList, selectedImage)
			}
		}

		candidate := skeleton
		candidate.TargetInfra.NodeGroups = candidateNodeGroups
		candidate.TargetSpecList = candidateSpecList
		candidate.TargetOsImageList = candidateImageList
		candidate.TargetNlbList = targetNlbList
		if hooks.enforceZonalSubnet != nil {
			hooks.enforceZonalSubnet(&candidate)
		}
		assignDisjointZoneSubnets(&candidate)

		overallStatus, overallStatusDesc, matchSummary := calculateMemberWeightedMatchRateWithDetails(
			csp, candidateNodeGroups, ngBlueprints, candidateSpecList, candidateImageList, minMatchRate,
		)
		candidate.Status = overallStatus

		if isDuplicateInfraCandidate(candidates, candidate) {
			continue
		}

		costSummary := CalculateCandidateHourlyCost(candidate)
		decisionText := ""
		if len(decisionNotes) > 0 {
			decisionText = fmt.Sprintf(" | Decisions: %s", strings.Join(decisionNotes, ", "))
		}
		warningText := ""
		if len(warnings) > 0 {
			warningText = fmt.Sprintf(" | %d warning(s): %s", len(warnings), strings.Join(warnings, "; "))
		}

		candidate.Description = fmt.Sprintf(
			"Candidate #%d | %s | Policy=%s | %d NLB(s) | EstCost=$%.4f/hr | Match Rate: Min=%.1f%% Max=%.1f%% Avg=%.1f%% | %s%s%s",
			len(candidates)+1,
			overallStatus,
			policy,
			len(targetNlbList),
			costSummary.TotalHourlyCostUSD,
			matchSummary.MinMatchRate, matchSummary.MaxMatchRate, matchSummary.AvgMatchRate,
			overallStatusDesc,
			decisionText,
			warningText,
		)

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

// resolveBackendPolicyBlueprints selects and constructs blueprints for one NLB backend under the active sizing policy.
func resolveBackendPolicyBlueprints(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	lauMembers []onpremmodel.NodeProperty,
	lauIds []string,
	lauById map[string]onpremmodel.NodeProperty,
	hwCase HardwareCase,
	sizingPolicy string,
	minMatchRate float64,
	hooks recommendationHooks,
	catalogCache *cachedCatalogMatcher,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, [][]CompatibleSpecImagePair, string, error) {
	backendName := rnlb.sourceNlb.Backend.Name

	buildConsolidated := func(decision string) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq) {
		bp, nlbReq, sg := buildGpuNlbConsolidatedBlueprint(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, lauIds, lauById, backendName, hwCase, sizingPolicy, decision, hooks,
		)
		return []nodeGroupBlueprint{bp}, []cloudmodel.NlbReq{nlbReq}, []cloudmodel.SecurityGroupReq{sg}
	}

	switch hwCase {
	case CaseCrossVendorGpu:
		if sizingPolicy == SizingPolicyRightSizeUp {
			return nil, nil, nil, nil, "", fmt.Errorf(
				"%w: NLB backend '%s' mixes incompatible GPU vendors; use '%s' or '%s' policy",
				ErrCrossVendorConsolidation, backendName, SizingPolicyDividing, SizingPolicyAuto,
			)
		}
		if sizingPolicy == SizingPolicyAuto {
			tiers := buildOrderedHardwareTiers(lauMembers)
			if len(generateValidTierPartitions(tiers)) > 1 {
				return evaluateAutoBackendPartitions(
					ctx, csp, region, connectionName, subnetId, skeleton,
					rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, hwCase, minMatchRate, hooks, catalogCache,
				)
			}
		}
		decision := "dividing(cross-vendor)"
		if sizingPolicy == SizingPolicyAuto {
			decision = "auto->dividing(cross-vendor-infeasible)"
		}
		bps, nlbs, sgs := buildDividedGroupBlueprints(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, partitionKeyByVendor, hwCase, sizingPolicy, decision, hooks,
		)
		return bps, nlbs, sgs, nil, decision, nil

	case CaseMixedCpuGpu:
		switch sizingPolicy {
		case SizingPolicyRightSizeUp:
			bps, nlbs, sgs := buildConsolidated("rightSizeUp")
			return bps, nlbs, sgs, nil, "rightSizeUp", nil
		case SizingPolicyDividing:
			bps, nlbs, sgs := buildMixedCpuGpuDivided(
				ctx, csp, region, connectionName, subnetId, skeleton,
				rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, hwCase, sizingPolicy, "dividing", hooks,
			)
			return bps, nlbs, sgs, nil, "dividing", nil
		default:
			return evaluateAutoBackendPartitions(
				ctx, csp, region, connectionName, subnetId, skeleton,
				rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, hwCase, minMatchRate, hooks, catalogCache,
			)
		}

	case CaseHeteroGpuModel:
		switch sizingPolicy {
		case SizingPolicyRightSizeUp:
			bps, nlbs, sgs := buildConsolidated("rightSizeUp")
			return bps, nlbs, sgs, nil, "rightSizeUp", nil
		case SizingPolicyDividing:
			bps, nlbs, sgs := buildDividedGroupBlueprints(
				ctx, csp, region, connectionName, subnetId, skeleton,
				rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, partitionKeyByModel, hwCase, sizingPolicy, "dividing", hooks,
			)
			return bps, nlbs, sgs, nil, "dividing", nil
		default:
			return evaluateAutoBackendPartitions(
				ctx, csp, region, connectionName, subnetId, skeleton,
				rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, hwCase, minMatchRate, hooks, catalogCache,
			)
		}

	default:
		// CasePureCpu and CaseHomoGpuModel produce a single NodeGroup under all policies.
		decision := sizingPolicy
		if sizingPolicy == SizingPolicyAuto {
			decision = "auto->single(homogeneous)"
		}
		bps, nlbs, sgs := buildConsolidated(decision)
		return bps, nlbs, sgs, nil, decision, nil
	}
}

type topologyEval struct {
	feasible   bool
	hourlyCost float64
	wmr        float64
	pairs      [][]CompatibleSpecImagePair
}

func evaluateBlueprintTopology(
	ctx context.Context,
	csp, region string,
	bps []nodeGroupBlueprint,
	minMatchRate float64,
	catalogCache *cachedCatalogMatcher,
) topologyEval {
	eval := topologyEval{
		feasible: len(bps) > 0,
		pairs:    make([][]CompatibleSpecImagePair, len(bps)),
	}
	if len(bps) == 0 {
		return eval
	}

	var weightedSum float64
	var weightCount int

	for i, bp := range bps {
		pairs := catalogCache.lookup(ctx, csp, region, bp.representativeNode)
		eval.pairs[i] = pairs
		if len(pairs) == 0 {
			eval.feasible = false
			continue
		}
		top := pairs[0]

		size := bp.skeleton.NodeGroupSize
		if size <= 0 {
			size = 1
		}
		if top.Spec.CostPerHour > 0 {
			eval.hourlyCost += float64(top.Spec.CostPerHour) * float64(size)
		}

		// Score every actual LAU member against the assigned target spec to expose over-provisioning.
		members := bp.memberNodes
		if len(members) == 0 {
			members = []onpremmodel.NodeProperty{bp.representativeNode}
		}
		for _, member := range members {
			vec := calculateMatchRateVector(csp, member, top.Spec, top.Image)
			dims := vec.ActiveDimensions()
			for _, d := range dims {
				weightedSum += d
			}
			weightCount += len(dims)
			// Penalize assigning a GPU instance to a CPU-only member.
			if !vec.HasGPU && (top.Spec.AcceleratorCount > 0 || top.Spec.AcceleratorMemoryGB > 0) {
				weightCount++
			}
		}
	}

	if eval.feasible && weightCount > 0 {
		eval.wmr = weightedSum / float64(weightCount)
	}
	return eval
}

type hardwareTier struct {
	key      string
	vendor   string
	isGpu    bool
	vramGB   float64
	gpuCount int
	lauIds   []string
}

type evaluatedTierPartition struct {
	blocks [][]hardwareTier
	bps    []nodeGroupBlueprint
	nlbs   []cloudmodel.NlbReq
	sgs    []cloudmodel.SecurityGroupReq
	eval   topologyEval
	score  float64
	order  int
}

func buildOrderedHardwareTiers(lauMembers []onpremmodel.NodeProperty) []hardwareTier {
	tierMap := make(map[string]*hardwareTier)
	for _, m := range lauMembers {
		key := partitionKeyByModel(m)
		t, ok := tierMap[key]
		if !ok {
			t = &hardwareTier{key: key, isGpu: hasGpu(m)}
			if t.isGpu {
				t.vendor = partitionKeyByVendor(m)
				if clusters := clusterGpuCards(m.GPUCards); len(clusters) > 0 {
					t.vramGB = float64(clusters[0].MemoryTotalGB)
					t.gpuCount = int(clusters[0].Count)
				}
			}
			tierMap[key] = t
		}
		t.lauIds = append(t.lauIds, m.MachineId)
	}

	tiers := make([]hardwareTier, 0, len(tierMap))
	for _, t := range tierMap {
		slices.Sort(t.lauIds)
		tiers = append(tiers, *t)
	}

	slices.SortFunc(tiers, func(a, b hardwareTier) int {
		if a.isGpu != b.isGpu {
			if !a.isGpu {
				return -1
			}
			return 1
		}
		if c := cmp.Compare(a.vendor, b.vendor); c != 0 {
			return c
		}
		if c := cmp.Compare(a.vramGB, b.vramGB); c != 0 {
			return c
		}
		if c := cmp.Compare(a.gpuCount, b.gpuCount); c != 0 {
			return c
		}
		return cmp.Compare(a.key, b.key)
	})
	return tiers
}

func isValidTierBlock(block []hardwareTier) bool {
	var seenVendor string
	for _, t := range block {
		if t.vendor == "" {
			continue
		}
		if seenVendor == "" {
			seenVendor = t.vendor
		} else if t.vendor != seenVendor {
			return false
		}
	}
	return true
}

func isValidTierPartition(blocks [][]hardwareTier) bool {
	for _, block := range blocks {
		if !isValidTierBlock(block) {
			return false
		}
	}
	return true
}

func tierPartitionSignature(blocks [][]hardwareTier) string {
	parts := make([]string, len(blocks))
	for i, block := range blocks {
		keys := make([]string, len(block))
		for j, t := range block {
			keys[j] = t.key
		}
		parts[i] = strings.Join(keys, "+")
	}
	return strings.Join(parts, "|")
}

func generateValidTierPartitions(tiers []hardwareTier) [][][]hardwareTier {
	if len(tiers) <= 1 {
		if len(tiers) == 0 {
			return nil
		}
		return [][][]hardwareTier{{tiers}}
	}

	var result [][][]hardwareTier
	seen := make(map[string]bool)
	appendIfValid := func(blocks [][]hardwareTier) {
		if len(blocks) == 0 || !isValidTierPartition(blocks) {
			return
		}
		sig := tierPartitionSignature(blocks)
		if !seen[sig] {
			seen[sig] = true
			result = append(result, blocks)
		}
	}

	if len(tiers) <= maxExhaustiveAutoTiers {
		numMasks := 1 << (len(tiers) - 1)
		for cutMask := 0; cutMask < numMasks; cutMask++ {
			var blocks [][]hardwareTier
			start := 0
			for i := 0; i < len(tiers)-1; i++ {
				if (cutMask>>i)&1 == 1 {
					blocks = append(blocks, tiers[start:i+1])
					start = i + 1
				}
			}
			blocks = append(blocks, tiers[start:])
			appendIfValid(blocks)
		}
		return result
	}

	// Bound large-tier search to canonical consolidated, per-vendor merged, and fully divided partitions.
	appendIfValid([][]hardwareTier{tiers})

	var cpuIsolated [][]hardwareTier
	i := 0
	for i < len(tiers) {
		if !tiers[i].isGpu {
			cpuIsolated = append(cpuIsolated, tiers[i:i+1])
			i++
			continue
		}
		end := i + 1
		for end < len(tiers) && tiers[end].vendor == tiers[i].vendor {
			end++
		}
		cpuIsolated = append(cpuIsolated, tiers[i:end])
		i = end
	}
	appendIfValid(cpuIsolated)

	allDivided := make([][]hardwareTier, len(tiers))
	for idx := range tiers {
		allDivided[idx] = tiers[idx : idx+1]
	}
	appendIfValid(allDivided)
	return result
}

func buildTierPartitionBlueprints(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	tiers []hardwareTier,
	blocks [][]hardwareTier,
	lauById map[string]onpremmodel.NodeProperty,
	hwCase HardwareCase,
	hooks recommendationHooks,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq) {
	backendName := rnlb.sourceNlb.Backend.Name
	gpuTierCount := 0
	for _, t := range tiers {
		if t.isGpu {
			gpuTierCount++
		}
	}

	vendorBlockCount := make(map[string]int)
	for _, block := range blocks {
		if !isValidTierBlock(block) {
			continue
		}
		for _, t := range block {
			if t.isGpu && t.vendor != "" {
				vendorBlockCount[t.vendor]++
				break
			}
		}
	}

	bps := make([]nodeGroupBlueprint, 0, len(blocks))
	nlbs := make([]cloudmodel.NlbReq, 0, len(blocks))
	sgs := make([]cloudmodel.SecurityGroupReq, 0, len(blocks))

	for blockIdx, block := range blocks {
		var blockLauIds []string
		var blockVendor string
		allGpu := true
		for _, t := range block {
			blockLauIds = append(blockLauIds, t.lauIds...)
			if !t.isGpu {
				allGpu = false
			} else if blockVendor == "" {
				blockVendor = t.vendor
			}
		}

		var suffix string
		switch {
		case len(blocks) == 1:
			suffix = backendName
		case len(block) == 1:
			if !block[0].isGpu {
				suffix = fmt.Sprintf("%s-cpu", backendName)
			} else if hwCase == CaseMixedCpuGpu && gpuTierCount == 1 {
				suffix = fmt.Sprintf("%s-gpu", backendName)
			} else {
				suffix = fmt.Sprintf("%s-%s", backendName, block[0].key)
			}
		case allGpu && blockVendor != "" && vendorBlockCount[blockVendor] == 1:
			suffix = fmt.Sprintf("%s-%s-gpu", backendName, blockVendor)
		default:
			suffix = fmt.Sprintf("%s-p%d", backendName, blockIdx+1)
		}

		bp, nlbReq, sg := buildGpuNlbConsolidatedBlueprint(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, blockLauIds, lauById, suffix, hwCase, SizingPolicyAuto, "auto", hooks,
		)
		bps = append(bps, bp)
		nlbs = append(nlbs, nlbReq)
		sgs = append(sgs, sg)
	}
	return bps, nlbs, sgs
}

func setBlueprintDecision(bps []nodeGroupBlueprint, decision string) {
	for i := range bps {
		if bps[i].skeleton.Label == nil {
			bps[i].skeleton.Label = make(map[string]string)
		}
		bps[i].skeleton.Label[labelTopologyDecision] = decision
	}
}

func betterTierPartition(a, b evaluatedTierPartition) bool {
	if math.Abs(a.score-b.score) > 1e-9 {
		return a.score > b.score
	}
	if len(a.bps) != len(b.bps) {
		return len(a.bps) < len(b.bps)
	}
	if math.Abs(a.eval.hourlyCost-b.eval.hourlyCost) > 1e-9 {
		return a.eval.hourlyCost < b.eval.hourlyCost
	}
	return a.order < b.order
}

func evaluateAutoBackendPartitions(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	lauMembers []onpremmodel.NodeProperty,
	lauById map[string]onpremmodel.NodeProperty,
	hwCase HardwareCase,
	minMatchRate float64,
	hooks recommendationHooks,
	catalogCache *cachedCatalogMatcher,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, [][]CompatibleSpecImagePair, string, error) {
	tiers := buildOrderedHardwareTiers(lauMembers)
	partitions := generateValidTierPartitions(tiers)

	hasCons := false
	consFeasible := false
	var firstCand *evaluatedTierPartition
	feasibleCands := make([]evaluatedTierPartition, 0, len(partitions))

	for idx, blocks := range partitions {
		bps, nlbs, sgs := buildTierPartitionBlueprints(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, tiers, blocks, lauById, hwCase, hooks,
		)
		eval := evaluateBlueprintTopology(ctx, csp, region, bps, minMatchRate, catalogCache)
		cand := evaluatedTierPartition{
			blocks: blocks,
			bps:    bps,
			nlbs:   nlbs,
			sgs:    sgs,
			eval:   eval,
			order:  idx,
		}
		if firstCand == nil {
			copied := cand
			firstCand = &copied
		}
		if len(blocks) == 1 {
			hasCons = true
			consFeasible = eval.feasible
		}
		if eval.feasible {
			feasibleCands = append(feasibleCands, cand)
		}
	}

	if len(feasibleCands) == 0 {
		decision := "auto->rightSizeUp(no-feasible-catalog-match)"
		if hwCase == CaseCrossVendorGpu {
			decision = "auto->dividing(no-feasible-catalog-match)"
		}
		setBlueprintDecision(firstCand.bps, decision)
		return firstCand.bps, firstCand.nlbs, firstCand.sgs, firstCand.eval.pairs, decision, nil
	}

	if len(feasibleCands) == 1 {
		best := feasibleCands[0]
		var decision string
		switch {
		case hasCons && !consFeasible && len(best.blocks) == len(tiers):
			decision = "auto->dividing(infeasible-spec-pruned)"
		case hasCons && !consFeasible:
			decision = fmt.Sprintf("auto->partial(%d/%d-tiers,infeasible-spec-pruned)", len(best.blocks), len(tiers))
		case len(best.blocks) == 1:
			decision = "auto->rightSizeUp(divided-catalog-infeasible)"
		case hwCase == CaseCrossVendorGpu && len(best.blocks) == len(tiers):
			decision = "auto->dividing(cross-vendor-infeasible)"
		default:
			decision = "auto->dividing(partial-catalog-infeasible)"
		}
		setBlueprintDecision(best.bps, decision)
		return best.bps, best.nlbs, best.sgs, best.eval.pairs, decision, nil
	}

	maxCost := 0.0
	for _, cand := range feasibleCands {
		if cand.eval.hourlyCost > maxCost {
			maxCost = cand.eval.hourlyCost
		}
	}
	for i := range feasibleCands {
		c := &feasibleCands[i]
		normCost := 0.0
		if maxCost > 0 {
			normCost = c.eval.hourlyCost / maxCost
		}
		splitPenalty := float64(len(c.bps)-1) / float64(len(c.bps))
		c.score = autoWeightMatchRate*(c.eval.wmr/100.0) - autoWeightCost*normCost - autoWeightSplitPenalty*splitPenalty
	}

	bestIdx := 0
	for i := 1; i < len(feasibleCands); i++ {
		if betterTierPartition(feasibleCands[i], feasibleCands[bestIdx]) {
			bestIdx = i
		}
	}
	best := feasibleCands[bestIdx]

	runnerUpIdx := -1
	for i := range feasibleCands {
		if i == bestIdx {
			continue
		}
		if runnerUpIdx == -1 || betterTierPartition(feasibleCands[i], feasibleCands[runnerUpIdx]) {
			runnerUpIdx = i
		}
	}
	runnerUp := feasibleCands[runnerUpIdx]

	prunedSuffix := ""
	if hasCons && !consFeasible {
		prunedSuffix = ",infeasible-spec-pruned"
	}

	var decision string
	switch {
	case len(best.blocks) == 1:
		decision = fmt.Sprintf("auto->rightSizeUp(J=%.3f>=%.3f,cost=%.2fvs%.2f)", best.score, runnerUp.score, best.eval.hourlyCost, runnerUp.eval.hourlyCost)
	case len(best.blocks) == len(tiers):
		decision = fmt.Sprintf("auto->dividing(J=%.3f>%.3f,cost=%.2fvs%.2f%s)", best.score, runnerUp.score, best.eval.hourlyCost, runnerUp.eval.hourlyCost, prunedSuffix)
	default:
		decision = fmt.Sprintf("auto->partial(%d/%d-tiers,J=%.3f>%.3f,cost=%.2fvs%.2f%s)", len(best.blocks), len(tiers), best.score, runnerUp.score, best.eval.hourlyCost, runnerUp.eval.hourlyCost, prunedSuffix)
	}
	setBlueprintDecision(best.bps, decision)
	return best.bps, best.nlbs, best.sgs, best.eval.pairs, decision, nil
}

func buildMixedCpuGpuDivided(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	lauMembers []onpremmodel.NodeProperty,
	lauById map[string]onpremmodel.NodeProperty,
	hwCase HardwareCase,
	sizingPolicy, decision string,
	hooks recommendationHooks,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq) {
	// If multiple GPU models coexist alongside CPU, partition by distinct model key to maintain homogeneity.
	modelKeys := make(map[string]bool)
	for _, m := range lauMembers {
		if hasGpu(m) {
			modelKeys[partitionKeyByModel(m)] = true
		}
	}
	if len(modelKeys) > 1 {
		return buildDividedGroupBlueprints(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, partitionKeyByModel, hwCase, sizingPolicy, decision, hooks,
		)
	}
	return buildDividedGroupBlueprints(
		ctx, csp, region, connectionName, subnetId, skeleton,
		rnlb, allResolvedNlbs, nlbMembershipCount, lauMembers, lauById, partitionKeyCpuOrGpu, hwCase, sizingPolicy, decision, hooks,
	)
}

func partitionKeyCpuOrGpu(m onpremmodel.NodeProperty) string {
	if hasGpu(m) {
		return "gpu"
	}
	return "cpu"
}

func partitionKeyByVendor(m onpremmodel.NodeProperty) string {
	if len(m.GPUCards) == 0 {
		return "cpu"
	}
	vendor := detectGpuVendor(m.GPUCards[0].Vendor)
	if vendor == GpuVendorOther {
		vendor = detectGpuVendor(m.GPUCards[0].Model)
	}
	if vendor == "" || vendor == GpuVendorOther {
		vendor = strings.TrimSpace(m.GPUCards[0].Vendor)
	}
	sanitized := sanitizeName(vendor)
	if sanitized == "" {
		return "gpu"
	}
	return sanitized
}

func partitionKeyByModel(m onpremmodel.NodeProperty) string {
	if len(m.GPUCards) == 0 {
		return "cpu"
	}
	card := m.GPUCards[0]
	vendor := partitionKeyByVendor(m)
	model := detectGpuModel(card.Model)
	if model == "" {
		model = sanitizeName(card.Model)
	}
	if card.MemoryTotalGB > 0 {
		return sanitizeName(fmt.Sprintf("%s-%s-%.0fgb", vendor, model, card.MemoryTotalGB))
	}
	return sanitizeName(fmt.Sprintf("%s-%s", vendor, model))
}

func buildDividedGroupBlueprints(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	lauMembers []onpremmodel.NodeProperty,
	lauById map[string]onpremmodel.NodeProperty,
	keyFn func(onpremmodel.NodeProperty) string,
	hwCase HardwareCase,
	sizingPolicy, decision string,
	hooks recommendationHooks,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq) {
	groups := make(map[string][]string)
	for _, m := range lauMembers {
		key := keyFn(m)
		groups[key] = append(groups[key], m.MachineId)
	}

	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	var bps []nodeGroupBlueprint
	var nlbs []cloudmodel.NlbReq
	var sgs []cloudmodel.SecurityGroupReq

	for _, k := range keys {
		suffix := fmt.Sprintf("%s-%s", rnlb.sourceNlb.Backend.Name, k)
		bp, nlbReq, sg := buildGpuNlbConsolidatedBlueprint(
			ctx, csp, region, connectionName, subnetId, skeleton,
			rnlb, allResolvedNlbs, nlbMembershipCount, groups[k], lauById, suffix, hwCase, sizingPolicy, decision, hooks,
		)
		bps = append(bps, bp)
		nlbs = append(nlbs, nlbReq)
		sgs = append(sgs, sg)
	}
	return bps, nlbs, sgs
}

func buildGpuNlbConsolidatedBlueprint(
	ctx context.Context,
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	lauIds []string,
	lauById map[string]onpremmodel.NodeProperty,
	groupSuffix string,
	hwCase HardwareCase,
	sizingPolicy, decision string,
	hooks recommendationHooks,
) (nodeGroupBlueprint, cloudmodel.NlbReq, cloudmodel.SecurityGroupReq) {
	ngName := "ng-" + sanitizeName(groupSuffix)
	repNode := synthesizeGroupRepresentativeNode(lauIds, lauById)

	// Collect unique physical source machine IDs for lineage traceability.
	physSeen := make(map[string]bool, len(lauIds))
	physIds := make([]string, 0, len(lauIds))
	for _, lid := range lauIds {
		origId, _ := splitDecomposedMachineId(lid)
		if !physSeen[origId] {
			physSeen[origId] = true
			physIds = append(physIds, origId)
		}
	}

	repWithFirewall := projectBackendFirewallRules(lauIds, lauById, rnlb, allResolvedNlbs, nlbMembershipCount, repNode)
	repWithFirewall = supplementFirewallIncomingFromRemote(repWithFirewall, rnlb, csp)
	repWithFirewall = supplementFirewallNlbToNode(repWithFirewall, rnlb, csp, skeleton)

	sg, err := hooks.recommendSG(ctx, csp, region, repWithFirewall)
	if err != nil {
		log.Warn().Err(err).Str("nodeGroup", ngName).Msg("failed to recommend SG for consolidated group")
	}
	sg.Name = fmt.Sprintf("mig-sg-%s", sanitizeName(groupSuffix))
	sg.ConnectionName = connectionName
	sg.Description = fmt.Sprintf("Security group for %s", ngName)
	sg.VNetId = skeleton.TargetVNet.Name

	minDisk := getCspMinRootDiskSizeGB(csp)
	if hasGpu(repNode) && minDisk < 100 {
		minDisk = 100
	}
	rootDiskSize := max(int(repNode.RootDisk.TotalSize), minDisk)

	members := make([]onpremmodel.NodeProperty, 0, len(lauIds))
	for _, lid := range lauIds {
		if m, ok := lauById[lid]; ok {
			members = append(members, m)
		}
	}

	bp := nodeGroupBlueprint{
		representativeNode: repNode,
		memberNodes:        members,
		isNlbRelated:       true,
		skeleton: cloudmodel.CreateNodeGroupReq{
			ConnectionName:   connectionName,
			Name:             ngName,
			VNetId:           skeleton.TargetVNet.Name,
			SubnetId:         subnetId,
			SecurityGroupIds: []string{sg.Name},
			SshKeyId:         skeleton.TargetSshKey.Name,
			RootDiskType:     "",
			RootDiskSize:     rootDiskSize,
			NodeGroupSize:    len(lauIds),
			Description:      fmt.Sprintf("Consolidated NodeGroup for NLB backend %s (%d nodes)", groupSuffix, len(lauIds)),
			Label: map[string]string{
				lineage.LabelSourceMachineIds: strings.Join(physIds, ","),
				lineage.LabelNlbBackend:       rnlb.sourceNlb.Backend.Name,
				labelHardwareCase:             fmt.Sprintf("%d", int(hwCase)),
				labelSizingPolicy:             sizingPolicy,
				labelTopologyDecision:         decision,
			},
		},
	}

	nlbReq := createManagedNlbReq(rnlb, ngName, groupSuffix)
	return bp, nlbReq, sg
}

// projectBackendFirewallRules merges member firewall rules while filtering foreign NLB ports on multi-NLB shared hosts.
func projectBackendFirewallRules(
	lauIds []string,
	lauById map[string]onpremmodel.NodeProperty,
	currentNlb resolvedNlb,
	allResolvedNlbs []resolvedNlb,
	nlbMembershipCount map[string]int,
	representativeNode onpremmodel.NodeProperty,
) onpremmodel.NodeProperty {
	currentPorts := make(map[string]bool)
	if currentNlb.sourceNlb.Listener.Port > 0 {
		currentPorts[strconv.Itoa(currentNlb.sourceNlb.Listener.Port)] = true
	}
	if currentNlb.backendPort > 0 {
		currentPorts[strconv.Itoa(currentNlb.backendPort)] = true
	}

	foreignPorts := make(map[string]bool)
	for _, other := range allResolvedNlbs {
		if other.sourceNlb.Backend.Name == currentNlb.sourceNlb.Backend.Name {
			continue
		}
		if other.sourceNlb.Listener.Port > 0 {
			p := strconv.Itoa(other.sourceNlb.Listener.Port)
			if !currentPorts[p] {
				foreignPorts[p] = true
			}
		}
		if other.backendPort > 0 {
			p := strconv.Itoa(other.backendPort)
			if !currentPorts[p] {
				foreignPorts[p] = true
			}
		}
	}

	merged := representativeNode
	ruleSet := make(map[string]onpremmodel.FirewallRuleProperty)
	orderedKeys := make([]string, 0)

	addRuleIfScoped := func(rule onpremmodel.FirewallRuleProperty, isOverlapping bool) {
		if isOverlapping && strings.EqualFold(rule.Direction, "inbound") && strings.EqualFold(rule.Action, "allow") {
			trimmedPorts := strings.TrimSpace(rule.DstPorts)
			if trimmedPorts != "" && trimmedPorts != "*" {
				tokens := strings.Split(trimmedPorts, ",")
				kept := make([]string, 0, len(tokens))
				for _, tok := range tokens {
					p := strings.TrimSpace(tok)
					if p == "" {
						continue
					}
					if !currentPorts[p] && foreignPorts[p] {
						continue
					}
					kept = append(kept, p)
				}
				if len(kept) == 0 {
					return
				}
				rule.DstPorts = strings.Join(kept, ",")
			}
		}
		key := firewallRuleKey(rule)
		if _, exists := ruleSet[key]; !exists {
			orderedKeys = append(orderedKeys, key)
			ruleSet[key] = rule
		}
	}

	for _, lid := range lauIds {
		node, ok := lauById[lid]
		if !ok {
			continue
		}
		origId, _ := splitDecomposedMachineId(node.MachineId)
		isOverlapping := nlbMembershipCount[origId] > 1
		for _, rule := range node.FirewallTable {
			addRuleIfScoped(rule, isOverlapping)
		}
	}

	rules := make([]onpremmodel.FirewallRuleProperty, 0, len(orderedKeys))
	for _, k := range orderedKeys {
		rules = append(rules, ruleSet[k])
	}
	merged.FirewallTable = rules
	return merged
}

func createManagedNlbReq(rnlb resolvedNlb, targetNodeGroupId, suffix string) cloudmodel.NlbReq {
	nlbType := "PUBLIC"
	if rnlb.sourceNlb.Listener.BindAddress != "" && rnlb.sourceNlb.Listener.BindAddress != "*" {
		nlbType = "INTERNAL"
	}

	listenerProtocol := normalizeProtocol(rnlb.sourceNlb.Listener.Protocol)
	targetGroupProtocol := normalizeProtocol(rnlb.sourceNlb.Backend.Protocol)

	hc := cloudmodel.NlbHealthCheckerReq{
		Interval:  defaultHealthCheckInterval,
		Threshold: defaultHealthCheckThreshold,
		Timeout:   defaultHealthCheckTimeout,
	}
	if rnlb.sourceNlb.HealthCheck.Enabled {
		if rnlb.sourceNlb.HealthCheck.Interval > 0 {
			hc.Interval = rnlb.sourceNlb.HealthCheck.Interval
		}
		if rnlb.sourceNlb.HealthCheck.Timeout > 0 {
			hc.Timeout = rnlb.sourceNlb.HealthCheck.Timeout
		}
		if rnlb.sourceNlb.HealthCheck.Threshold > 0 {
			hc.Threshold = rnlb.sourceNlb.HealthCheck.Threshold
		}
	}

	return cloudmodel.NlbReq{
		Description: fmt.Sprintf("Migrated from NLB backend: %s", suffix),
		Scope:       defaultNlbScope,
		Type:        nlbType,
		Listener: cloudmodel.NlbListenerReq{
			Protocol: listenerProtocol,
			Port:     fmt.Sprintf("%d", rnlb.sourceNlb.Listener.Port),
		},
		TargetGroup: cloudmodel.NlbTargetGroupReq{
			Protocol:    targetGroupProtocol,
			Port:        fmt.Sprintf("%d", rnlb.backendPort),
			NodeGroupId: targetNodeGroupId,
		},
		HealthChecker: hc,
	}
}

// assignDisjointZoneSubnets allocates per-zone subnets on zonal CSPs when NodeGroups lack a single common zone.
func assignDisjointZoneSubnets(candidate *cloudmodel.RecommendedInfra) {
	if candidate == nil || len(candidate.TargetVNet.SubnetInfoList) == 0 || len(candidate.TargetInfra.NodeGroups) == 0 {
		return
	}
	cspName := resolveCandidateCSP(*candidate)
	if !IsZonalSubnetCSP(cspName) {
		return
	}
	// If a common zone was already assigned onto subnet[0], nothing to split.
	if strings.TrimSpace(candidate.TargetVNet.SubnetInfoList[0].Zone) != "" {
		return
	}

	ngZones := make([]string, len(candidate.TargetInfra.NodeGroups))
	zoneSeen := make(map[string]bool)
	var distinctZones []string
	for i, ng := range candidate.TargetInfra.NodeGroups {
		spec, ok := findSpecForNodeGroup(ng.SpecId, candidate.TargetSpecList)
		if !ok {
			continue
		}
		z := pickNodeGroupAvailableZone(spec)
		if z == "" {
			continue
		}
		ngZones[i] = z
		if !zoneSeen[z] {
			zoneSeen[z] = true
			distinctZones = append(distinctZones, z)
		}
	}
	if len(distinctZones) == 0 {
		return
	}
	sort.Strings(distinctZones)

	subnets := make([]cloudmodel.SubnetReq, len(candidate.TargetVNet.SubnetInfoList))
	copy(subnets, candidate.TargetVNet.SubnetInfoList)

	zoneToSubnet := make(map[string]string, len(distinctZones))
	subnets[0].Zone = distinctZones[0]
	zoneToSubnet[distinctZones[0]] = subnets[0].Name

	usedCidrs := make([]string, 0, len(subnets)+len(distinctZones))
	for _, s := range subnets {
		if s.IPv4_CIDR != "" {
			usedCidrs = append(usedCidrs, s.IPv4_CIDR)
		}
	}

	for _, z := range distinctZones[1:] {
		existingMatch := ""
		for idx := range subnets {
			if strings.EqualFold(strings.TrimSpace(subnets[idx].Zone), z) {
				existingMatch = subnets[idx].Name
				break
			}
		}
		if existingMatch != "" {
			zoneToSubnet[z] = existingMatch
			continue
		}

		nextCidr, ok := allocateNextSubnetCIDR(candidate.TargetVNet.CidrBlock, usedCidrs)
		if !ok {
			log.Warn().Str("csp", cspName).Str("zone", z).Msg("unable to carve additional zone subnet inside VNet CIDR")
			continue
		}
		usedCidrs = append(usedCidrs, nextCidr)
		subName := fmt.Sprintf("mig-subnet-%02d", len(subnets)+1)
		subnets = append(subnets, cloudmodel.SubnetReq{
			Name:        subName,
			IPv4_CIDR:   nextCidr,
			Zone:        z,
			Description: fmt.Sprintf("Zone-aligned subnet for %s", z),
		})
		zoneToSubnet[z] = subName
	}

	candidate.TargetVNet.SubnetInfoList = subnets
	for i := range candidate.TargetInfra.NodeGroups {
		if subName, ok := zoneToSubnet[ngZones[i]]; ok && subName != "" {
			candidate.TargetInfra.NodeGroups[i].SubnetId = subName
		}
	}
}

func pickNodeGroupAvailableZone(spec cloudmodel.SpecInfo) string {
	avail := normalizeZoneList(spec.AvailableZones)
	if len(avail) == 0 {
		return ""
	}
	sort.Strings(avail)
	unavail := normalizeZoneList(spec.UnavailableZones)
	for _, z := range avail {
		if !slices.Contains(unavail, z) {
			return z
		}
	}
	return avail[0]
}

func allocateNextSubnetCIDR(vnetCidr string, usedCidrs []string) (string, bool) {
	_, vnetNet, err := net.ParseCIDR(strings.TrimSpace(vnetCidr))
	if err != nil || vnetNet == nil {
		return "", false
	}
	vnetOnes, bits := vnetNet.Mask.Size()
	if bits != 32 {
		return "", false
	}

	var parsedUsed []*net.IPNet
	prefSize := 24
	for _, u := range usedCidrs {
		if _, n, e := net.ParseCIDR(strings.TrimSpace(u)); e == nil && n != nil {
			parsedUsed = append(parsedUsed, n)
			if len(parsedUsed) == 1 {
				if ones, _ := n.Mask.Size(); ones > vnetOnes {
					prefSize = ones
				}
			}
		}
	}
	if prefSize <= vnetOnes {
		prefSize = vnetOnes + 2
		if prefSize > 28 {
			prefSize = 28
		}
	}

	vnetStart := binary.BigEndian.Uint32(vnetNet.IP.To4())
	vnetSize := uint32(1) << uint(32-vnetOnes)
	vnetEnd := vnetStart + vnetSize - 1

	for _, prefix := range []int{prefSize, 24, 25, 26, 27, 28} {
		if prefix <= vnetOnes || prefix > 28 {
			continue
		}
		blockSize := uint32(1) << uint(32-prefix)
		mask := net.CIDRMask(prefix, 32)
		// Start searching cleanly from second block if available, then wrap around.
		for offset := blockSize; offset < vnetSize; offset += blockSize {
			addr := vnetStart + offset
			if addr+blockSize-1 > vnetEnd {
				break
			}
			ip := make(net.IP, 4)
			binary.BigEndian.PutUint32(ip, addr)
			candNet := &net.IPNet{IP: ip, Mask: mask}
			if !cidrOverlapsAny(candNet, parsedUsed) {
				return candNet.String(), true
			}
		}
		for offset := uint32(0); offset < blockSize && offset < vnetSize; offset += blockSize {
			addr := vnetStart + offset
			ip := make(net.IP, 4)
			binary.BigEndian.PutUint32(ip, addr)
			candNet := &net.IPNet{IP: ip, Mask: mask}
			if !cidrOverlapsAny(candNet, parsedUsed) {
				return candNet.String(), true
			}
		}
	}
	return "", false
}

func cidrOverlapsAny(cand *net.IPNet, used []*net.IPNet) bool {
	for _, u := range used {
		if cand.Contains(u.IP) || u.Contains(cand.IP) {
			return true
		}
	}
	return false
}
