# GPU+NLB Infrastructure Recommendation Benchmark CLI (Workload-Agnostic)

This test CLI benchmarks the workload-agnostic GPU+NLB infrastructure transformation and recommendation engine (`Constraint-Aware Adaptive Cloud Topology Transformation`) through the single unified endpoint `POST /beetle/recommendation/infra` and verifies deployability via `POST /beetle/validation/ns/{nsId}/infra`.

## Test Scenarios (Static Hardware + Network Topology Only)

| Scenario | Case | Description |
| :--- | :--- | :--- |
| `S0` | Case 0 | Pure CPU NLB backend + standalone CPU server |
| `S1` | Case 1 | Mixed CPU + NVIDIA T4 GPU NLB backend + standalone CPU server |
| `S2` | Case 2 | Cross-vendor GPU (NVIDIA T4 + AMD Radeon Pro V520) NLB backend |
| `S3` | Case 3 | Heterogeneous same-vendor GPU (NVIDIA L4 + NVIDIA T4) NLB backend |
| `S4` | Case 4 | Homogeneous NVIDIA T4 NLB backend + multi-GPU standalone server (`T4 + L4`) |
| `S5` | Compound A | Single physical host inside NLB backend with heterogeneous GPUs (`L4 + T4`) |
| `S6` | Compound B | Multi-NLB (`N:M`) shared host overlap + mixed CPU/GPU backends |
| `S7` | Compound C (3-Tier) | 3-tier mixed CPU (`2x`) + NVIDIA T4 (`2x`) + NVIDIA L4 (`1x`) NLB backend + standalone CPU (`192` total benchmark runs with `8 CSPs x 8 scenarios x 3 policies`) |

## Policies Benchmarked

- `rightSizeUp`: Consolidates compatible NLB backend pools upward to a single homogeneous NodeGroup (returns HTTP `422` on cross-vendor GPUs)
- `dividing`: Partitions mixed or heterogeneous NLB backend pools into separate homogeneous NodeGroups with dedicated Managed NodeGroup-Bound 1:1 NLBs
- `auto`: Multi-Objective Candidate Topology Optimizer (`J(T)`) with Ordered Tier Partial Partition Optimization (`T_up`, `T_partial`, `T_div` subject to cross-vendor hard cuts), infeasible spec pruning, per-member LAU fitness evaluation, Hourly VM Compute Cost (`CostPerHour`), split overhead penalty, and Disjoint-AZ Multi-Subnet Allocation

## Outputs (`testresult/`)

- `patent-benchmark-summary.md`: Consolidated multi-CSP results, Hourly VM Compute Cost savings, `Workload-Aware Reference` comparison, and deterministic SHA-256 hash reproducibility tables
- `patent-benchmark-metrics.csv`: Machine-readable benchmark rows for every `(CSP, Region, Scenario, Policy)` evaluation
- `scenario-<id>-<csp>.md`: Detailed NodeGroup, NLB, cost, and validation breakdown per run

## Usage

```bash
make test-gpu-nlb-recommendation
```

To run with an explicit config or repeat count:

```bash
cd cmd/test-cli/gpu-nlb-recommendation
go run main.go -config testconf/test-config.yaml -repeat 5
```
