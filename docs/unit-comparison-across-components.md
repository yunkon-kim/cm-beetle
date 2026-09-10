# Unit Comparison Across Components & GPU Memory Normalization

This document outlines the unit standards across the Cloud-Barista sub-systems (CB-Spider, CB-Tumblebug, CM-Honeybee, CM-Beetle) and explains why GPU memory is normalized to **`GB`** in `imdl`.

---

## 1. Why `imdl` Uses `GB` for GPU Memory

### The Discrepancy
Source hardware inspection tools report GPU memory in vendor-specific units:
* **NVIDIA (`nvidia-smi`)**: Reports in **`MiB`** (e.g., `40960 MiB`).
* **AMD ROCm (`rocm-smi`)**: Reports in **`Bytes`** (e.g., `34359738368`).
* **AMD ROCm (`amd-smi`)**: Reports in **`MB`**.

### Architectural Rationale
1. **Collector-Side Normalization (Honeybee → `imdl`)**:
   CM-Honeybee's role is to absorb vendor-specific CLI differences (NVIDIA `MiB`, AMD `Bytes`/`MB`) and provide a unified representation. This follows the existing refinement pattern in Honeybee (e.g., CPU MHz → GHz, RAM MiB → GiB).
2. **Target Cloud Alignment**:
   Cloud providers (AWS, Azure, GCP) and CB-Tumblebug catalog VM accelerators in nominal **`GB`** (`acceleratorMemoryGB: 40`). Normalizing to `GB` (`float32`) enables direct matching without conversion discrepancies.
3. **Why Not Pass `MiB` to Beetle?**:
   Passing raw `MiB` would be necessary only for micro-architectural memory packing (e.g., CUDA kernel scheduling or LLM KV-cache sizing). For cloud VM migration and recommendation, provisioning occurs on discrete catalog tiers (16, 24, 40, 80 GB). Passing `MiB` would leak hardware-specific quirks into Beetle and require downstream consumers to duplicate vendor-handling logic.

---

## 2. Unit Comparison Matrix

| Resource Metric | CB-Spider | CB-Tumblebug | CM-Honeybee (Agent) | CM-Beetle / `imdl` | Beetle UX Lab (UI) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Host System RAM** | **`MiB`** (`MemSizeMiB`) | **`GiB`** (`MemoryGiB`) | Raw OS pages / `MiB` | **`GiB`** (`TotalSize`) | **`GiB`** |
| **Storage / Disk Size** | **`GB`** (`DiskSizeGB`) | **`GB`** (`DiskSizeGB`, `RootDiskSize`) | `df -h` blocks / `GB` | **`GB`** (`TotalSize`) | **`GB`** |
| **Total Storage Volume** | N/A | **`TiB`** (`MaxTotalStorageTiB`) | N/A | N/A | **`TiB`** |
| **GPU Memory (VRAM)** | **`GB`** (`MemSizeGB`) | **`GB`** (`AcceleratorMemoryGB`) | `MiB` (NVIDIA) / `Bytes` (AMD) | **`GB`** (`TotalMemoryGB`, `MemoryTotalGB`) | **`GB`** |
| **CPU Clock Speed** | **`GHz`** (`ClockGHz`) | N/A | `MHz` (`/proc/cpuinfo`) | **`GHz`** (`MaxSpeed`) | **`GHz`** |

---

## 3. Essential Struct References

```go
// CB-Spider: cloud-driver/interfaces/resources/VMSpecHandler.go
type VMSpecInfo struct {
    MemSizeMiB string    `json:"MemSizeMiB"` // Unit: MiB
    DiskSizeGB string    `json:"DiskSizeGB"` // Unit: GB
    Gpu        []GpuInfo `json:"Gpu,omitempty"`
}
type GpuInfo struct {
    MemSizeGB string `json:"MemSizeGB"` // Unit: GB
}

// CB-Tumblebug: src/core/model/spec.go
type TbSpecInfo struct {
    MemoryGiB           float32 `json:"memoryGiB,omitempty"`           // Unit: GiB
    DiskSizeGB          float32 `json:"diskSizeGB,omitempty"`          // Unit: GB
    AcceleratorMemoryGB float32 `json:"acceleratorMemoryGB,omitempty"` // Unit: GB
}

// CM-Beetle: imdl/on-premise-model/node.go
type NodeProperty struct {
    // ...
    GPUCards []GpuCardProperty `json:"gpuCards,omitempty"`
}

type GpuCardProperty struct {
    DriverIndex   string  `json:"driverIndex,omitempty"`
    Vendor        string  `json:"vendor,omitempty"`
    Model         string  `json:"model,omitempty"`
    Slot          string  `json:"slot,omitempty"`
    PciBusId      string  `json:"pciBusId,omitempty"`
    MemoryTotalGB float32 `json:"memoryTotalGb,omitempty"` // Unit: GB
    MemoryFreeGB  float32 `json:"memoryFreeGb,omitempty"`  // Unit: GB
    MemoryUsedGB  float32 `json:"memoryUsedGb,omitempty"`  // Unit: GB
}
```

---

## 4. Key Takeaways

* **Host RAM**: Tracked in binary units (**`MiB`** in Spider, **`GiB`** in Tumblebug and `imdl`).
* **Disk & GPU Memory**: Tracked in decimal units (**`GB`** across Spider, Tumblebug, and `imdl`).
* **Conversion Boundary**: CM-Honeybee converts vendor-specific telemetry (NVIDIA `MiB / 1024`, AMD `Bytes / 1024^3`) into `imdl` `GB` (`float32`), allowing CM-Beetle to cleanly match CB-Tumblebug specs without driver-specific branches.
