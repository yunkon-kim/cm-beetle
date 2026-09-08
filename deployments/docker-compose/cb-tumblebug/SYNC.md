# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:

   ```bash
   cd /path/to/cb-tumblebug
   git checkout v0.13.2
   ```

2. Discover all changed files systematically:

   ```bash
   TB=/path/to/cb-tumblebug
   BEETLE=deployments/docker-compose/cb-tumblebug

   # Detect changed files in each directory
   for dir in assets conf init scripts; do
     echo "=== Checking $dir/ ==="
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Files.*differ$"
   done

   # Check MCP files (different path in TB)
   diff -qr $BEETLE/interface/mcp $TB/src/interface/mcp

   # Detect new files (TB only)
   for dir in assets conf init scripts; do
     echo "=== Checking new files in $dir/ ==="
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Only in $TB"
   done

   # Note: assets/spider has been deprecated and removed
   # Note: assets/rdbmsinfo.yaml is maintained directly in cm-beetle for managed RDBMS feature
   ```

3. Review individual file changes and copy updated files:

   ```bash
   cp $TB/assets/*.yaml $BEETLE/assets/
   cp $TB/assets/*.csv $BEETLE/assets/
   cp $TB/assets/assets.dump.gz* $BEETLE/assets/
   cp -r $TB/src/interface/mcp/* $BEETLE/interface/mcp/
   ```

4. For binary assets (`assets.dump.gz`), compare checksums:
   ```bash
   md5sum $BEETLE/assets/assets.dump.gz $TB/assets/assets.dump.gz
   ```

---

## v0.13.3 Sync (2026-09-08)

Based on CB-Tumblebug tag `v0.13.3` (`af9ba2056d6f075e35e08c2ab32a0d90fd22b71a`). Upgrade path: **v0.13.2 &rarr; v0.13.3**.

### Model Changes (v0.13.2 &rarr; v0.13.3)

| Change | Description |
| --- | --- |
| **`ExecCredentialStatus`**: `ExpirationTimestamp` | Added optional `ExpirationTimestamp *string` relaying token expiry (RFC3339) from CB-Spider for Kubernetes exec-based authentication |
| **Model Version Headers** | Updated `copied-tb-model.go` and `copied-tb-k8s-model.go` version headers to CB-Tumblebug `v0.13.3` (`af9ba2056d6f075e35e08c2ab32a0d90fd22b71a`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/cloud-model/`)** | |
| `copied-tb-model.go` | **Updated** — Synced version header to `v0.13.3` (`af9ba205`) |
| `copied-tb-k8s-model.go` | **Updated** — Synced with v0.13.3 (`af9ba205`): added `ExecCredentialStatus.ExpirationTimestamp`, version header |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** — `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.2` to `v0.13.3` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** — Synced service images: `cb-spider:0.13.4`, `cb-mapui:0.13.7`, `cb-tumblebug:0.13.3` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/spider/` | **Removed** — No longer maintained per project decision |
| `assets/assets.dump.gz` | **No change** — Verified MD5 `dba8da8e89b5ebdd203daf7c4480147b` (matches upstream) |
| `assets/assets.dump.gz.info` | **No change** — Verified matches upstream sidecar manifest |
| `assets/rdbmsinfo.yaml` | **Maintained** — Preserved latest CM-Beetle RDBMS metadata as instructed |
| `assets/cloudinfo.yaml`, `diskinfo.yaml`, `k8sclusterinfo.yaml`, etc. | **No change** — Verified matches upstream provider configurations |
| `init/templates/infra-usecase-llm-bench.json` | **Updated** — Synced multi-CSP GPU node group benchmarking template (NVIDIA A10, L20, L4, L40S, AMD V710) |
| `scripts/*` | **No change** — Verified matches upstream asset utilities and lib scripts |
| `interface/mcp/tb-mcp.py` | **Updated** — Synced upstream command template expansions (`GetNodeIds`) |

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
