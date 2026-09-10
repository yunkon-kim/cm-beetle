# GPU Infrastructure and Node Spec Recommendation Pipeline Plan

> **Status**: Draft / Proposed  
> **Author**: Cloud-Barista CM-Beetle Development Team  
> **Target Package**: `pkg/core/recommendation`  
> **Target Dependency**: `CB-Tumblebug` (`/recommendSpec`, `/recommendSpecOptions`)  
> **Reference Issue**: [#477](https://github.com/cloud-barista/cm-beetle/issues/477)

---

## 1. Executive Summary & Problem Statement

When an on-premise physical server contains specialized hardware accelerators (e.g., NVIDIA A100, Tesla T4, AMD Instinct), migrating it to public clouds requires matching specialized GPU-accelerated cloud VM instances (e.g., AWS `p4d`/`g5`, Azure `ND`/`NC`, GCP `a2`/`g2`).

Cloud GPU instances differ fundamentally from general-purpose CPU instances:

1. **Hard Binary & Architectural Boundaries**: Workloads compiled against NVIDIA CUDA runtimes cannot execute on AMD ROCm or Intel OneAPI hardware without source/binary recompilation.
2. **Strict Physical Memory Ceiling**: GPU device memory (VRAM) cannot be transparently paged out to disk without crashing AI/ML training or inference pipelines. Under-provisioning VRAM causes fatal Out-Of-Memory (OOM) errors.
3. **Coupled Chassis Proportions**: General-purpose cloud VMs allow varying CPU and RAM within broad ranges (e.g. 1:2 to 1:8). In contrast, cloud GPU instances feature fixed host-to-accelerator ratios determined by physical CSP chassis (e.g. AWS `p4d.24xlarge` strictly pairs 8× A100 with 96 vCPUs and 1,152 GiB host RAM).
4. **Heterogeneous Physical Nodes vs. Homogeneous Cloud VMs**: Physical on-premise servers may install mixed GPU cards across PCIe slots (e.g. 2× NVIDIA A100 + 1× AMD MI350), whereas cloud providers strictly provision homogeneous GPU models per VM instance.

This plan details the architectural design for a clean, extensible, configuration-driven GPU recommendation engine adhering to Go Single Responsibility Principle (SRP) and Cloud-Barista project standards.

---

## 2. Architectural Analysis: Dedicated Pipelines vs. Inlined Extensions (Go SRP)

The introduction of GPU acceleration expands the problem space across two distinct tiers:
1. **Node Spec Tier**: Recommending an optimal cloud VM specification for a single workload.
2. **Infrastructure (MCI) Tier**: Recommending an entire multi-cloud infrastructure topology (VNet, Subnets, Security Groups, and Node Group sizing) for $N$ source physical servers.

Below, we evaluate both tiers through the lens of the **Go Single Responsibility Principle (SRP)** and **Open-Closed Principle (OCP)**.

### 2.1 Node-Spec Tier: `recommendGpuNodeSpec` vs. Inlined Branching in `RecommendNodeSpecs`

| Dimension                          | Inlined Branching (`if isGpu`)                                                                                                                                  | Dedicated Node Spec Pipeline (`resource-node-spec-gpu.go`)                                                               |
| :--------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------- |
| **Go Single Responsibility (SRP)** | **Violated**: A single function mixes CPU prime-number searching, CPU machine-type ratio classification (1:2 compute vs 1:8 memory), and GPU VRAM constraints. | **Adhered**: `recommendGpuNodeSpec()` and `recommendCpuNodeSpec()` encapsulate distinct workload models. Dispatcher only routes based on `hasGpu(node)`. |
| **Maintainability & Coupling**     | **High Coupling**: Modifications to GPU search heuristics or retry bounds risk introducing regressions into CPU sizing routines.                                 | **Zero Coupling**: GPU sizing policies and vendor lookup tables evolve independently without affecting CPU node migration. |
| **Testability**                    | Complex mocks required to isolate GPU execution branches inside CPU helper functions.                                                                            | Direct unit-testing of `recommendGpuNodeSpec` and `recommendCpuNodeSpec` with focused test fixtures.                      |

### 2.2 Infrastructure (MCI) Tier: Dedicated Pipeline (`infra-gpu.go`) vs. Extending `infra.go`

In `pkg/core/recommendation/infra.go` (1,273 lines), `RecommendVmInfraCandidates` operates on a strict **1:1 invariant**: $N$ source nodes produce exactly $N$ target NodeGroups and $N$ spec-image pairs.

When source servers contain heterogeneous GPUs (e.g. 2× A100 + 1× T4 + 1× MI350), this invariant is shattered: 1 physical server must be decomposed into $K$ target VM instances ($1:K$).

| Evaluation Dimension | Extending `infra.go` (In-place Expansion) | Dedicated GPU Infra Pipeline (`infra-gpu.go`) |
| :--- | :--- | :--- |
| **Go Single Responsibility (SRP)** | **Violated (Responsibility Creep)**:<br>`infra.go` would combine general CPU Pareto evaluation, 1:1 node mapping, $1:K$ heterogeneous GPU decomposition, multi-card topology quantization, and GPU inter-node networking into a massive monolithic file (>1,800 lines). | **Adhered (Clean Separation)**:<br>`infra.go` retains exclusive ownership of general CPU infra.<br>`infra-gpu.go` encapsulates the entire GPU topology decomposition ($1:K$), GPU subnet grouping, and GPU match rate scoring. |
| **Open-Closed Principle (OCP)** | **Poor**: Modifying the core loop of `RecommendVmInfraCandidates` introduces regression risks to established, production-validated CPU infrastructure migration workflows. | **Excellent**: Existing CPU infra flows remain untouched. GPU infrastructure logic evolves cleanly in isolation. |
| **Match Rate Evaluation Vector** | `MatchRateVector` in `infra.go` only scores `CPU`, `Memory`, `Image` (3 dimensions). Forcing GPU metrics (VRAM, Card Count, Vendor) into it distorts CPU matching scores. | Implements `GpuMatchRateVector` (CPU, Memory, Image, AcceleratorCount, VRAM, Architecture) without polluting the general CPU vector. |
| **Architectural Precedent** | Inconsistent with established project conventions. | **Directly follows the proven `infra-with-nlb.go` precedent**, where NLB-aware infrastructure recommendation was isolated into a dedicated parallel pipeline. |

### 2.3 Adopted Two-Tier Dedicated Architecture

We adopt **Dedicated Pipelines across both tiers**:

```
[RecommendVmInfraCandidates (Infra Dispatcher)]
            │
            ├─► hasAnyGpu(srcInfra) == true  ──► [infra-gpu.go: RecommendGpuVmInfraCandidates]
            │                                             │
            │                                             ├─► 1. Decompose Heterogeneous Nodes (1:K Topology)
            │                                             ├─► 2. Provision Shared Intra-Node VNet & Security Groups
            │                                             ├─► 3. Delegate to [recommendGpuNodeSpec] per Workload
            │                                             └─► 4. GPU Pareto-Frontier Candidate Assembly
            │
            └─► hasAnyGpu(srcInfra) == false ──► [infra.go: RecommendCpuVmInfraCandidates]
                                                          │
                                                          ├─► 1. 1:1 NodeGroup Sizing
                                                          ├─► 2. Standard CPU VNet & Security Groups
                                                          ├─► 3. Delegate to [recommendCpuNodeSpec]
                                                          └─► 4. CPU Pareto-Frontier Candidate Assembly
```

```
[RecommendNodeSpec (Node Dispatcher)]
            │
            ├─► hasGpu(node) == true  ──► [resource-node-spec-gpu.go: recommendGpuNodeSpec]
            │                                   │
            │                                   ├─► 1. Cluster Homogeneous GPU Cards
            │                                   ├─► 2. Evaluate Strategy (Rightsizing & Modernization)
            │                                   ├─► 3. Discrete VRAM Staircase & 2^k Topology Quantization
            │                                   ├─► 4. Targeted Query (TB /recommendSpec with Fallback)
            │                                   └─► 5. Multi-Dimensional Proximity Ranking
            │
            └─► hasGpu(node) == false ──► [resource-node-spec.go: recommendCpuNodeSpec]
                                                │
                                                ├─► 1. Derive Workload Machine Type (1:2 vs 1:8 ratio)
                                                ├─► 2. Prime Range Search (Continuous Step Iterations)
                                                ├─► 3. Targeted Query (TB /recommendSpec)
                                                └─► 4. CPU Host Proximity Ranking
```

---

## 3. Clean In-Code Constants and Vendor Normalization (Idiomatic Go)

To keep server configuration clean and avoid runtime file dependencies, we adopt the idiomatic Go pattern established in [`resource-node-spec.go`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/resource-node-spec.go#L440-L455) for CPU vendor resolution:

### 3.1 Strongly Typed Protocol Constants (`resource-node-spec-gpu.go`)

CB-Tumblebug protocol keys, canonical vendor names, and internal search defaults are declared as typed Go constants:

```go
// Tumblebug deployment plan metrics
const (
	MetricAcceleratorCount    = "acceleratorCount"
	MetricAcceleratorMemoryGB = "acceleratorMemoryGB"
	MetricAcceleratorType     = "acceleratorType"
	MetricAcceleratorModel    = "acceleratorModel"
	AcceleratorTypeGPU        = "gpu"

	// Internal search heuristics (kept in code to avoid over-configuration)
	DefaultRangeWeightMin = 1
	DefaultRangeWeightMax = 5
)

// Canonical GPU Vendor identities
const (
	GpuVendorNVIDIA = "NVIDIA"
	GpuVendorAMD    = "AMD"
	GpuVendorIntel  = "Intel"
	GpuVendorGoogle = "Google"
	GpuVendorOther  = "Other"
)
```

### 3.2 Vendor Normalization Alias Table (`gpuVendorAliases`)

A static slice maps model/chip keywords and abbreviations to canonical vendor identities. This matches the exact structure of `cpuVendorAliases` in [`resource-node-spec.go`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/resource-node-spec.go):

```go
// gpuVendorAliases maps normalized model/family tokens to canonical vendor identities.
var gpuVendorAliases = []struct {
	token  string
	vendor string
}{
	// NVIDIA
	{"nvidia", GpuVendorNVIDIA},
	{"tesla", GpuVendorNVIDIA},
	{"geforce", GpuVendorNVIDIA},
	{"a100", GpuVendorNVIDIA},
	{"h100", GpuVendorNVIDIA},
	{"h200", GpuVendorNVIDIA},
	{"b200", GpuVendorNVIDIA},
	{"b300", GpuVendorNVIDIA},
	{"gb200", GpuVendorNVIDIA},
	{"a10g", GpuVendorNVIDIA},
	{"a10", GpuVendorNVIDIA},
	{"a30", GpuVendorNVIDIA},
	{"a40", GpuVendorNVIDIA},
	{"t4", GpuVendorNVIDIA},
	{"t4g", GpuVendorNVIDIA},
	{"v100", GpuVendorNVIDIA},
	{"l4", GpuVendorNVIDIA},
	{"l20", GpuVendorNVIDIA},
	{"l40s", GpuVendorNVIDIA},
	{"rtx", GpuVendorNVIDIA},
	{"quadro", GpuVendorNVIDIA},

	// AMD
	{"amd", GpuVendorAMD},
	{"radeon", GpuVendorAMD},
	{"instinct", GpuVendorAMD},
	{"mi100", GpuVendorAMD},
	{"mi200", GpuVendorAMD},
	{"mi210", GpuVendorAMD},
	{"mi250", GpuVendorAMD},
	{"mi300", GpuVendorAMD},
	{"mi300x", GpuVendorAMD},
	{"mi350", GpuVendorAMD},
	{"v520", GpuVendorAMD},
	{"v620", GpuVendorAMD},
	{"v710", GpuVendorAMD},

	// Intel
	{"intel", GpuVendorIntel},
	{"gaudi", GpuVendorIntel},
	{"gaudi3", GpuVendorIntel},
	{"ponte vecchio", GpuVendorIntel},
	{"flex", GpuVendorIntel},
	{"max", GpuVendorIntel},

	// Google TPU
	{"tpu", GpuVendorGoogle},
	{"ct3", GpuVendorGoogle},
	{"ct5", GpuVendorGoogle},
	{"ct6", GpuVendorGoogle},
}

// detectGpuVendor normalizes arbitrary vendor or model strings to a canonical vendor.
func detectGpuVendor(raw string) string {
	lower := strings.ToLower(strings.TrimSpace(raw))
	for _, entry := range gpuVendorAliases {
		if strings.Contains(lower, entry.token) {
			return entry.vendor
		}
	}
	return GpuVendorOther
}
```

### 3.3 Design Advantages

1. **Zero Config Pollution**: `conf/config.yaml` remains focused strictly on server startup parameters (endpoints, credentials, log levels, DB paths, call pacer limits).
2. **Operational Reliability**: No external YAML file I/O, no container volume/path resolution dependencies, and no startup failure modes if a config file is missing.
3. **Architectural Consistency**: Adheres to the established Go design pattern already proven in CM-Beetle CPU recommendation (`cpuVendorAliases`).

---

## 4. Multi-Dimensional Proximity Ranking Algorithm

When CB-Tumblebug returns a candidate list of GPU VM specs, they are sorted using a multi-dimensional hierarchy ensuring functional correctness before economic optimization:

```
[Candidate VM Specs]
         │
         ▼
  ┌─────────────────────────────────────────────────────────────┐
  │ 1. Vendor Match                                            │
  │    (Canonical vendor match: NVIDIA == NVIDIA, AMD == AMD)   │
  │    *Absolute prerequisite: Prevents binary incompatibility  │
  └──────────────────────────────┬──────────────────────────────┘
                                 ▼ (Tie-break)
  ┌─────────────────────────────────────────────────────────────┐
  │ 2. Accelerator Count Proximity                              │
  │    (Exact count first; minimizes surplus GPU waste)         │
  └──────────────────────────────┬──────────────────────────────┘
                                 ▼ (Tie-break)
  ┌─────────────────────────────────────────────────────────────┐
  │ 3. VRAM (MemoryTotalGB) Proximity                           │
  │    (Closest VRAM >= Target VRAM; avoids OOM)                │
  └──────────────────────────────┬──────────────────────────────┘
                                 ▼ (Tie-break)
  ┌─────────────────────────────────────────────────────────────┐
  │ 4. Host Resource Distance (Manhattan L1)                    │
  │    |vCPU - target_vCPU| + |RAM - target_RAM|                │
  └──────────────────────────────┬──────────────────────────────┘
                                 ▼ (Tie-break)
  ┌─────────────────────────────────────────────────────────────┐
  │ 5. CostPerHour                                              │
  │    (Lowest cost per hour wins final tie-break)              │
  └─────────────────────────────────────────────────────────────┘
```

---

## 5. Handling Heterogeneous (Mixed) GPU Physical Nodes

### 5.1 Physical Reality & The Cloud Constraint

An on-premise physical chassis can host heterogeneous cards across PCIe slots:
- **PCIe Slot 1 & 2**: 2× NVIDIA A100-PCIE-40GB (Heavy LLM Training / PyTorch CUDA)
- **PCIe Slot 3**: 1× NVIDIA Tesla T4 (Light Inference / Video Transcode)
- **PCIe Slot 4**: 1× AMD Instinct MI350 (ROCm Development / 288GB VRAM)

Because no public CSP (AWS, Azure, GCP, Alibaba, NCP) provisions mixed accelerator models in a single VM—and cross-vendor instances (NVIDIA + AMD in one VM) do not exist—a single cloud VM cannot directly mirror this physical server.

### 5.2 Architectural Tradeoff: Multi-Node Decomposition vs. Single-Node Consolidation

| Strategy | Feasibility & Stability | Workload & Driver Impact | Decision |
| :--- | :--- | :--- | :---: |
| **Multi-Node Decomposition** *(이기종 분할 추천)* | **100% Deterministic**: Maps each homogeneous card cluster to a dedicated cloud VM instance. | **Zero Impact**: CUDA and ROCm run in their native environments without code modifications. | **ADOPTED** |
| **Single-Node Consolidation** *(동종 단일 VM 통합)* | **Infeasible**: Requires arbitrary heuristic guesses (which vendor to retain? how to re-architect ROCm code to CUDA? how to size host CPU/RAM?). | **Fatal**: Workloads compiled for the omitted vendor fail to execute at runtime. | **REJECTED** |

> **Architectural Decision**: Single-Node Consolidation is intentionally **rejected**. Automating consolidation requires subjective business decisions that only the workload owner can make. Attempting to force-consolidate heterogeneous GPUs produces high runtime failure risks and excessive cost over-provisioning.

### 5.3 Adopted Multi-Node Decomposition Workflow

When a source server has mixed GPU cards, CM-Beetle executes **Workload-Preserving Multi-Node Decomposition**:

1. **Card Clustering ([`clusterGpuCards`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/resource-node-spec-gpu.go#L28))**:
   Groups cards by `(Vendor, Model, MemoryTotalGB)` into distinct homogeneous clusters:
   - `Cluster[0]` (**Primary**): Highest card count or total capability (e.g., 2× NVIDIA A100, 40GB)
   - `Cluster[1]` (**Secondary**): Next largest cluster (e.g., 1× AMD Instinct MI350, 288GB)
   - `Cluster[2]` (**Tertiary**): Remaining auxiliary cluster (e.g., 1× NVIDIA Tesla T4, 16GB)

2. **Single Node Recommendation API (`RecommendNodeSpecs`)**:
   - The primary cluster drives the node's main VM spec recommendation.
   - For `len(clusters) > 1`, auxiliary recommendations for secondary/tertiary clusters are attached as supplementary recommendations (`SubClusterRecommendations`), informing the user of the exact cloud instances needed to fully absorb the physical hardware.

3. **Full Infrastructure Recommendation API (`RecommendVmInfraCandidates`)**:
   - Decomposes the single on-premise physical machine into distinct target cloud nodes within the recommended Multi-Cloud Infrastructure (MCI).
   - Configures them inside the same Virtual Network (VNet) so that services previously communicating via localhost can seamlessly communicate over private cloud IP addresses.

---

## 6. Specialized GPU Recommendation Algorithms (+ @ Engine Points)

Unlike CPU/RAM where memory and cores can be linearly scaled using continuous ranges (e.g., Prime Range Search), GPU cloud instances are constrained by physical hardware topologies, discrete VRAM sizes, and architectural generations. We incorporate four specialized algorithmic techniques into CM-Beetle:

### 6.1 Discrete VRAM Staircase Search (불연속 VRAM 계단식 매칭)
- **Problem**: Cloud GPUs do not offer arbitrary VRAM (e.g., 20 GB or 35 GB does not exist in standard CSP instances). They exist in strict physical tiers: `8 GB -> 16 GB -> 24 GB -> 40 GB -> 48 GB -> 80 GB`.
- **Algorithm**:
  1. Determine target VRAM requirement $V_{req}$.
  2. Map $V_{req}$ to the smallest standard cloud VRAM tier $V_{tier} \ge V_{req}$.
  3. Query CB-Tumblebug with a bounded window $[0.95 \times V_{req}, 1.25 \times V_{tier}]$ to prevent OOM while eliminating excessive over-provisioning cost.

### 6.2 Power-of-Two ($2^k$) Topology Quantization (토폴로지 정량화)
- **Problem**: CSPs only offer VM instances with $1, 2, 4, 8$ GPUs. If an on-premise server has 3 GPUs or 6 GPUs (common with custom PCIe riser chassis), naive queries for exact card count fail.
- **Algorithm**:
  - Quantize requested card count $N$ up to the nearest valid power-of-two topology $N_{cloud} = 2^{\lceil \log_2 N \rceil}$ (e.g., $3 \to 4$, $5/6 \to 8$).
  - For large on-prem counts (e.g., 6 GPUs), also evaluate multi-node decomposition ($4 + 2$).

### 6.3 Legacy Architecture Modernization Control (`enableArchitectureModernization`, Default: `true`)
- **Problem**: On-premise servers often run legacy GPUs (NVIDIA Tesla K80, M60, V100). These are either deprecated, unavailable, or exorbitantly expensive on modern CSPs.
- **User-Controlled Strategy**:
  - `enableArchitectureModernization: true` (**Default**): **High-Availability Modernization**.
    - Ensures recommendation success without failures due to retired hardware.
    - If exact legacy model matching yields no results (or if target CSP lacks legacy instances), automatically maps to current-generation, cost-effective equivalents:
      - `Tesla K80 / M60` $\to$ Modern Tier: `T4` or `L4` (faster, 70%+ cheaper).
      - `Tesla V100` $\to$ Modern Tier: `A10G` or `L4` (superior FP32/Tensor core performance, half the cost).
  - `enableArchitectureModernization: false` (**Opt-Out**): **Strict Legacy Architecture Matching**.
    - Restricts search strictly to identical physical generations. Returns no recommendation if the exact generation is unavailable in the target region.

### 6.4 Telemetry-Driven Rightsizing Control (`enableRightsizing`, Default: `false`)
- **User-Controlled Strategy**: Rightsizing is **strictly opt-in** and configured via API parameter:
  - `enableRightsizing: false` (**Default**): **Hardware-Preserving Mirroring**.
    - Bases recommendation purely on installed physical capacity (`MemoryTotalGB` and card counts).
    - Guarantees workload stability and avoids under-provisioning caused by temporary idle periods during metrics collection.
  - `enableRightsizing: true` (**Opt-In**): **Telemetry-Driven Cost Optimization**.
    - Activated when the user explicitly requests cost optimization based on actual monitored telemetry.
    - Evaluates Honeybee's `MemoryUsedGB` telemetry with an added safety headroom ($V_{target} = \text{MemoryUsedGB} \times 1.25$).
    - Imposes a minimum safety floor (never downsize below 50% of original physical VRAM).

### 6.5 Summary of User-Configurable Recommendation Parameters

| Parameter | Type | Default | Description |
| :--- | :---: | :---: | :--- |
| `enableArchitectureModernization` | `bool` | **`true`** | Automatically maps deprecated/legacy GPUs (e.g. K80, V100) to modern equivalents (L4, A10G) to ensure recommendation success. |
| `enableRightsizing` | `bool` | **`false`** | Optimizes target spec using monitored telemetry (`MemoryUsedGB`) instead of physical installed capacity (`MemoryTotalGB`). |

### 6.6 Production Considerations for Cloud GPU Migration

1. **CPU Architecture Alignment (x86_64 vs. ARM64)**:
   - Cloud GPU instances may pair with ARM64 CPUs (e.g. AWS `g5g.xlarge` with Graviton2).
   - If the source server is x86_64, deploying to ARM64 causes binary incompatibilities (`Exec format error`). The query must enforce CPU architecture alignment (`node.CPU.Architecture`).
2. **OS Image & CUDA Driver Pre-requisites**:
   - Vanilla cloud OS images lack NVIDIA drivers and CUDA runtimes. Recommendation metadata attaches advisory notes recommending Deep Learning/GPU-optimized images.
3. **Cloud Service Quota Constraints**:
   - CSPs often set default GPU quotas to 0. Recommendation descriptions guide users to verify GPU service quotas.
4. **Storage Floor for Model Checkpoints**:
   - GPU workloads require significant disk for model weights; enforce a safe root disk minimum floor (e.g., 100 GB).

### 6.7 Core Implementation Principles
- **Simple is Best**: Keep logic concise, direct, and transparent.
- **No Convoluted Fallbacks**: Use deterministic single-shot targeted queries with clear bounds instead of complex multi-tiered retry fallbacks.
- **Avoid Simple Wrappers**: Directly invoke functions without introducing trivial one-line wrappers or unnecessary indirection layers.

---

## 7. Live CB-Tumblebug API Verification & Data Schema

Live validation of CB-Tumblebug's `/tumblebug/recommendSpecOptions` and `POST /tumblebug/recommendSpec` APIs confirmed the following real-world operational behaviors:

### 7.1 Filter Metrics & Catalog Options (`/tumblebug/recommendSpecOptions`)

- **Available Filter Metrics**: `id`, `providerName`, `regionName`, `cspSpecName`, `architecture`, `acceleratorModel`, `acceleratorType`, `description`, `vCPU`, `memoryGiB`, `acceleratorCount`, `acceleratorMemoryGB`, `costPerHour`, `evaluationScore01`
- **Available Prioritization Metrics**: `cost` (lowest cost first), `performance` (highest compute score first), `location`, `latency`, `random`
- **Normalized Accelerator Models in Catalog**:
  - **NVIDIA**: `"NVIDIA A10"`, `"NVIDIA A10G"`, `"NVIDIA A100"`, `"NVIDIA H100"`, `"NVIDIA L4"`, `"NVIDIA T4"`, `"NVIDIA T4G"`, `"NVIDIA TESLA A100"`, `"NVIDIA Tesla V100"`, `"NVIDIA B200"`, etc.
  - **AMD**: `"AMD INSTINCT"`, `"AMD MI300X"`, `"AMD RADEON PRO V520"`, `"AMD RADEON PRO V620"`, `"AMD RADEON PRO V710"`
  - **Intel**: `"Intel GAUDI3"`
  - **Google**: `"TPU7X"`, `"CT3"`, `"CT5L"`, `"CT5P"`

### 7.2 Verified Live Recommendations Across CSPs (`POST /tumblebug/recommendSpec`)

| CSP           | Region            | Returned Spec (`cspSpecName`) | Accelerator Model     | Count | VRAM (GB) | Cost/Hr |
| :------------ | :---------------- | :---------------------------- | :-------------------- | :---: | :-------: | :-----: |
| **AWS**       | `ap-northeast-2`  | `g5g.xlarge`                  | `NVIDIA T4G`          |   1   |   17 GB   | $0.5166 |
| **Azure**     | `koreacentral`    | `Standard_NC4as_T4_v3`        | `NVIDIA Tesla T4`     |   1   |   16 GB   | $0.6470 |
| **GCP**       | `asia-northeast3` | `g2-standard-4`               | `NVIDIA L4`           |   1   |   24 GB   | $0.9077 |
| **AWS (AMD)** | `us-east-1`       | `g4ad.xlarge`                 | `AMD RADEON PRO V520` |   1   |   9 GB    | $0.3785 |

> **Critical Performance Insight**: Including `regionName` in the recommendation query restricts search scope and returns in **4–6 seconds**. Omitting `regionName` forces Tumblebug to scan the CSP global catalog across dozens of regions, causing high response latency (~3–4 minutes).

---

## 8. Implementation Phasing

| Phase                                        | Tasks                                                                                                                                                                                                | Key Deliverables                                                                |
| :------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------ |
| **Phase 1: Foundation**                      | 1. Define typed protocol constants in Go.<br>2. Define `gpuVendorAliases` table matching `cpuVendorAliases` pattern.<br>3. Implement `detectGpuVendor()` with comprehensive unit tests. | Robust, self-contained vendor normalization without external config dependencies. |
| **Phase 2: Node Spec Pipeline**              | 1. Implement standalone `recommendGpuNodeSpec()` in `resource-node-spec-gpu.go`.<br>2. Implement discrete VRAM staircase search and $2^k$ topology right-sizing up (`rightSizeUpVramTier`, `rightSizeUpGpuCount`).<br>3. Implement `enableRightsizing` and `enableArchitectureModernization` logic.<br>4. Extract `recommendCpuNodeSpec()` in `resource-node-spec.go`.<br>5. Simplify `RecommendNodeSpecs()` dispatcher. | Clean separation of concerns (SRP) with specialized GPU algorithms at node level. |
| **Phase 3: Dedicated Infra Pipeline (`infra-gpu.go`)** | 1. Implement `DecomposeOnpremNodes()` for $1:K$ heterogeneous node expansion.<br>2. Implement `RecommendGpuInfraCandidates()` in dedicated `infra-gpu.go`.<br>3. Dispatch at API Handler level (`controller.RecommendInfraCandidates`) via `HasAnyGpu()`. | Full infrastructure topology recommendation with multi-node group decomposition. |
| **Phase 4: Integration & Testing**           | 1. Update unit tests in `resource-node-spec-gpu_test.go` and `infra-gpu_test.go`.<br>2. Run live integration tests against `cb-tumblebug` across AWS, Azure, GCP.<br>3. Validate multi-cluster card handling and decomposed node groups. | Verified end-to-end multi-cloud infrastructure recommendation workflow.         |

---

## 9. Verification & Success Criteria

1. **Automated Unit Tests**:
   - `go test -v ./pkg/core/recommendation/...` passes with 100% coverage of vendor matching, clustering, and right-sizing logic.
   - `go test ./imdl/... ./cmd/...` passes cleanly.
2. **Static Analysis**:
   - `go vet ./...` reports 0 issues.
   - `golangci-lint` passes without warnings.
3. **Live Tumblebug Compatibility**:
   - Querying AWS `ap-northeast-2` correctly matches `g5` or `g4dn` for NVIDIA nodes.
   - Querying AWS `us-east-1` correctly matches `g4ad` for AMD nodes.
   - Queries execute within 5–10 seconds by ensuring scoped `regionName` inclusion.

---

## 10. Future Plan: Host Resource (vCPU & Memory) Filter Optimization Based on Empirical Results

### 10.1 Background & Problem Statement

Currently, `recommendGpuNodeSpec` relies on a two-stage strategy:
1. **Catalog Query (Filter Stage)**: Queries CB-Tumblebug with GPU-centric constraints (`acceleratorCount >= targetCount`, `acceleratorMemoryGB >= targetVram`, `providerName`, `regionName`, `architecture`), deliberately omitting host `vCPU` and `memoryGiB` thresholds to avoid zero-result catalog queries caused by CSP-fixed GPU instance bundling (e.g., T4 cards rarely bundle with 64+ vCPUs).
2. **Post-Query Ranking (Ranking Stage)**: Sorts returned candidates by `gpuHostResourceProximity` (4th tie-break priority using Manhattan distance $|\Delta \text{vCPU}| + |\Delta \text{RAM}|$) to prefer instances closest to the source server's host capacity.

**Potential Risk**:
Because the Tumblebug query sorts by `"priority": [{"metric": "cost"}]`, the fetched candidate batch (e.g., top 15 by cost) could potentially be saturated with entry-level, low-vCPU instances (e.g., 4 vCPUs, 16 GB RAM). For heavy workloads running on large on-premise hosts, this may result in CPU or system RAM bottlenecks.

### 10.2 Empirical Validation & Next Steps

Following real-world test recommendations and empirical validation:
1. **Analyze Empirical Results**: Conduct end-to-end recommendation runs across diverse on-premise profiles (e.g., compute-heavy GPU nodes vs memory-heavy GPU nodes) to observe whether returned candidates exhibit host resource deficiency.
2. **Evaluate Filter Enhancements**:
   - **Option A (Safety Floor in Query)**: Introduce minimum host resource floors in the query template, e.g.:
     $$\text{minVCPU} = \min(\text{vcpusCalculated} \times 0.5, 8)$$
     $$\text{minMemoryGiB} = \min(\text{memory} \times 0.5, 16)$$
   - **Option B (Expanded Fetch & Post-Filter)**: Increase `fetchLimit` (e.g., 30) and prune candidates that fall below a tolerable host resource ratio before final ranking.
3. **Decide & Implement**: Finalize the host resource filtering strategy based on observed migration fidelity and user feedback.
