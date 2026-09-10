package recommendation

import (
	"testing"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasAnyGpu(t *testing.T) {
	tests := []struct {
		name     string
		infra    onpremmodel.OnpremInfra
		expected bool
	}{
		{
			name: "CPU-only infrastructure",
			infra: onpremmodel.OnpremInfra{
				Nodes: []onpremmodel.NodeProperty{
					{MachineId: "web-01"},
					{MachineId: "was-01"},
				},
			},
			expected: false,
		},
		{
			name: "Infrastructure with single GPU node",
			infra: onpremmodel.OnpremInfra{
				Nodes: []onpremmodel.NodeProperty{
					{MachineId: "web-01"},
					{
						MachineId: "gpu-node-01",
						GPUCards: []onpremmodel.GpuCardProperty{
							{DriverIndex: "0", Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16},
						},
					},
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, HasAnyGpu(tc.infra))
		})
	}
}

func TestDecomposeOnpremNodes(t *testing.T) {
	// Mixed infrastructure:
	// Node 1: CPU-only web server
	// Node 2: Homogeneous 2x A100 server
	// Node 3: Heterogeneous server (2x A100 + 1x T4 + 1x MI350)
	inputNodes := []onpremmodel.NodeProperty{
		{
			MachineId: "web-srv",
			CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 1},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		},
		{
			MachineId: "homo-gpu-srv",
			CPU:       onpremmodel.CpuProperty{Cpus: 16, Threads: 2},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
			GPUCards: []onpremmodel.GpuCardProperty{
				{DriverIndex: "0", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
				{DriverIndex: "1", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
			},
		},
		{
			MachineId: "hetero-gpu-srv",
			CPU:       onpremmodel.CpuProperty{Cpus: 32, Threads: 2},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 128},
			GPUCards: []onpremmodel.GpuCardProperty{
				{DriverIndex: "0", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
				{DriverIndex: "1", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
				{DriverIndex: "2", Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16},
				{DriverIndex: "card0", Vendor: "AMD", Model: "MI350", MemoryTotalGB: 288},
			},
		},
	}

	decomposed := DecomposeOnpremNodes(inputNodes)

	// Expect 1 (web) + 1 (homo) + 3 (hetero: A100, MI350, T4) = 5 nodes
	require.Len(t, decomposed, 5)

	// Node 0: web-srv (preserved 1:1)
	assert.Equal(t, "web-srv", decomposed[0].MachineId)
	assert.Empty(t, decomposed[0].GPUCards)

	// Node 1: homo-gpu-srv (preserved 1:1)
	assert.Equal(t, "homo-gpu-srv", decomposed[1].MachineId)
	assert.Len(t, decomposed[1].GPUCards, 2)

	// Node 2: hetero primary cluster (2x A100)
	assert.Equal(t, "hetero-gpu-srv-primary-a100", decomposed[2].MachineId)
	require.Len(t, decomposed[2].GPUCards, 2)
	assert.Equal(t, "NVIDIA", decomposed[2].GPUCards[0].Vendor)
	assert.Equal(t, "A100", decomposed[2].GPUCards[0].Model)

	// Node 3: hetero secondary cluster (1x MI350 - tie-broken by 288GB VRAM)
	assert.Equal(t, "hetero-gpu-srv-gpu-mi350", decomposed[3].MachineId)
	require.Len(t, decomposed[3].GPUCards, 1)
	assert.Equal(t, "AMD", decomposed[3].GPUCards[0].Vendor)
	assert.Equal(t, "MI350", decomposed[3].GPUCards[0].Model)

	// Node 4: hetero tertiary cluster (1x T4)
	assert.Equal(t, "hetero-gpu-srv-gpu-t4", decomposed[4].MachineId)
	require.Len(t, decomposed[4].GPUCards, 1)
	assert.Equal(t, "NVIDIA", decomposed[4].GPUCards[0].Vendor)
	assert.Equal(t, "T4", decomposed[4].GPUCards[0].Model)
}

func TestRecommendGpuInfraCandidates_Live(t *testing.T) {
	t.Skip("Live integration test requiring live Tumblebug server with >100ms response time; covered by E2E API verification")
}

