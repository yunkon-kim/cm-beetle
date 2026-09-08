# CM-Honeybee Assets & Configurations Sync Guide

> **Note:** These files originate from [CM-Honeybee](https://github.com/cloud-barista/cm-honeybee).

## Upstream Source

```
https://github.com/cloud-barista/cm-honeybee/tree/main/server/
```

CM-Beetle's copy lives at:

```
deployments/docker-compose/cm-honeybee/
```

## Structure & Mappings

| Local Path in CM-Beetle | Upstream Source Path in CM-Honeybee | Description |
| :--- | :--- | :--- |
| `openbao/openbao-config.hcl` | `server/openbao/openbao-config.hcl` | OpenBao configuration for Honeybee's dedicated secrets backend |

## v0.6.1 Sync (2026-09-08)

- Bumped `cm-honeybee` service image to `cloudbaristaorg/cm-honeybee:0.6.1`.
- Removed fixed agent port mapping (`8082:8082`) in `docker-compose.ui.yaml`, as the agent now dynamically allocates and binds to a kernel-chosen loopback port over SSH.
- Verified OpenBao configuration `cm-honeybee/openbao/openbao-config.hcl` remains identical with upstream v0.6.1.
