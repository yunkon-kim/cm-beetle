# Managed RDBMS (MySQL) Test Report: GCP (us-central1)

- **Test Case:** GCP US-Central1 MySQL Test
- **Date & Time:** 2026-09-08 10:31:39
- **Namespace:** `default`
- **Total Duration:** 15m1.644s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** GCP
- **Target Region:** `us-central1`
- **Namespace:** `default`
- **Test Date:** 2026-09-08 10:31:39

### Scenario & Tested APIs
1. **Pre-flight Spec & Image Review**: `POST /tumblebug/specImagePairReview`
2. **Create Pre-requisite Infra (VNet/SG)**: `POST /tumblebug/ns/{nsId}/resources/vNet`, `POST /tumblebug/ns/{nsId}/resources/securityGroup`
3. **Get RDBMS Support Matrix**: `GET /beetle/recommendation/middleware/rdbms/support`
4. **Get Real-time Capability**: `GET /beetle/recommendation/middleware/rdbms/capability`
5. **Recommend Managed RDBMS**: `POST /beetle/recommendation/middleware/rdbms`
6. **Validate Recommendation**: `POST /beetle/recommendation/middleware/rdbms/validate`
7. **Migrate RDBMS (Provisioning)**: `POST /beetle/migration/middleware/ns/{nsId}/rdbms`
8. **Get RDBMS Info & List**: `GET /beetle/migration/middleware/ns/{nsId}/rdbms`
9. **Create Logical Database**: `POST /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database`
10. **External Data I/O**: Direct TCP/SQL connectivity test
11. **Internal Data I/O**: SQL execution via internal Runner VM (`POST /tumblebug/ns/{nsId}/infra`)
12. **Delete Logical Database**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database/{databaseName}`
13. **Delete RDBMS**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}`
14. **Delete Pre-requisite SG & VNet**: `DELETE /tumblebug/ns/{nsId}/resources/securityGroup/{sgId}`, `DELETE /tumblebug/ns/{nsId}/resources/vNet/{vNetId}`

## Execution Steps & API Traces

### 1. Tumblebug POST /specImagePairReview (Pre-flight Spec & Image Review) [✅ SUCCESS]
- **Duration:** 4.809s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
  "specId": "gcp+us-central1+e2-standard-2"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "e2-standard-2",
    "provider": "gcp",
    "queriedAt": "2026-09-08T01:31:43.830881068Z",
    "region": "us-central1",
    "source": "gcp:machineTypes.aggregatedList",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-a"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-f"
      }
    ]
  },
  "connectionName": "gcp-us-central1",
  "estimatedCost": "$0.0670/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "gcp-africa-south1",
    "creationDate": "",
    "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "description": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17",
    "details": [
      {
        "key": "Architecture",
        "value": "X86_64"
      },
      {
        "key": "ArchiveSizeBytes",
        "value": "50232915200"
      },
      {
        "key": "CreationTimestamp",
        "value": "2026-08-18T02:42:50.563-07:00"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17"
      },
      {
        "key": "DiskSizeGb",
        "value": "10"
      },
      {
        "key": "EnableConfidentialCompute",
        "value": "false"
      },
      {
        "key": "Family",
        "value": "ubuntu-minimal-2404-lts-amd64"
      },
      {
        "key": "GuestOsFeatures",
        "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:SNP_SVSM_CAPABLE}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
      },
      {
        "key": "Id",
        "value": "5467262061071931381"
      },
      {
        "key": "Kind",
        "value": "compute#image"
      },
      {
        "key": "LabelFingerprint",
        "value": "iNBmVNCFF9w="
      },
      {
        "key": "Labels",
        "value": "{public-image:true}"
      },
      {
        "key": "LicenseCodes",
        "value": "6508311393003325021"
      },
      {
        "key": "Licenses",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-minimal-2404-lts"
      },
      {
        "key": "Name",
        "value": "ubuntu-minimal-2404-noble-amd64-v20260817"
      },
      {
        "key": "RawDisk",
        "value": "{containerType:TAR}"
      },
      {
        "key": "SatisfiesPzi",
        "value": "false"
      },
      {
        "key": "SatisfiesPzs",
        "value": "false"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817"
      },
      {
        "key": "SourceType",
        "value": "RAW"
      },
      {
        "key": "Status",
        "value": "READY"
      },
      {
        "key": "StorageLocations",
        "value": "me-central2; asia-northeast3; asia-east1; us-central2; europe-north1; europe-west15; europe-southwest1; asia-northeast1; northamerica-northeast2; me-west1; australia-southeast2; us-west1; europe-west10; australia-southeast1; africa-south1; europe-west3; us-west2; asia-southeast3; europe-central2; asia-northeast2; us-central1; northamerica-south1; us-west8; us-east7; asia-east2; us-west3; southamerica-west1; europe-west12; us-west4; europe-north2; asia; us; me-central1; us-east1; europe-west2; us-south1; asia-south1; asia-south2; europe-west4; europe-west1; asia-southeast1; us-east4; europe-west9; europe-west6; southamerica-east1; northamerica-northeast1; europe-west8; us-east5; asia-southeast2; eu"
      }
    ],
    "fetchedTime": "2026.08.21 13:58:59 Fri",
    "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 10,
    "osDiskType": "NA",
    "osDistribution": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "gcp",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbljac0fg9pr918114et"
  },
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
  "imageValidation": {
    "cspResourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "isAvailable": true,
    "resourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "resourceName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "gcp",
  "regionName": "us-central1",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "gcp-us-central1",
    "costPerHour": 0.067011,
    "cspSpecName": "e2-standard-2",
    "details": [
      {
        "key": "CreationTimestamp",
        "value": "1969-12-31T16:00:00.000-08:00"
      },
      {
        "key": "Description",
        "value": "Efficient Instance, 2 vCPUs, 8 GB RAM"
      },
      {
        "key": "GuestCpus",
        "value": "2"
      },
      {
        "key": "Id",
        "value": "335002"
      },
      {
        "key": "ImageSpaceGb",
        "value": "0"
      },
      {
        "key": "IsSharedCpu",
        "value": "false"
      },
      {
        "key": "Kind",
        "value": "compute#machineType"
      },
      {
        "key": "MaximumPersistentDisks",
        "value": "128"
      },
      {
        "key": "MaximumPersistentDisksSizeGb",
        "value": "263168"
      },
      {
        "key": "MemoryMb",
        "value": "8192"
      },
      {
        "key": "Name",
        "value": "e2-standard-2"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/us-central1-a/machineTypes/e2-standard-2"
      },
      {
        "key": "Zone",
        "value": "us-central1-a"
      }
    ],
    "diskSizeGB": -1,
    "evaluationScore01": -1,
    "evaluationScore02": -1,
    "evaluationScore03": -1,
    "evaluationScore04": -1,
    "evaluationScore05": -1,
    "evaluationScore06": -1,
    "evaluationScore07": -1,
    "evaluationScore08": -1,
    "evaluationScore09": -1,
    "evaluationScore10": -1,
    "id": "gcp+us-central1+e2-standard-2",
    "infraType": "node",
    "memoryGiB": 7.8125,
    "name": "gcp+us-central1+e2-standard-2",
    "namespace": "system",
    "providerName": "gcp",
    "regionLatitude": 41.2522,
    "regionLongitude": -95.8575,
    "regionName": "us-central1",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbe0e9dnnuunhi69h1tq",
    "vCPU": 2
  },
  "specId": "gcp+us-central1+e2-standard-2",
  "specValidation": {
    "cspResourceId": "e2-standard-2",
    "isAvailable": true,
    "resourceId": "gcp+us-central1+e2-standard-2",
    "resourceName": "e2-standard-2",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "us-central1-b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 26.473s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.2.0.0/16",
  "connectionName": "gcp-us-central1",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-gcp",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.2.1.0/24",
      "name": "subnet-1",
      "zone": "us-central1-a"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.2.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:32:08Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:32:08Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:32:08Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
    }
  ],
  "connectionConfig": {
    "configName": "gcp-us-central1",
    "credentialHolder": "admin",
    "credentialName": "gcp",
    "driverName": "gcp-driver-v1.0.so",
    "providerName": "gcp",
    "regionDetail": {
      "description": "Council Bluffs Iowa  USA",
      "location": {
        "display": "Council Bluffs Iowa USA",
        "latitude": 41.2522,
        "longitude": -95.8575
      },
      "regionId": "us-central1",
      "regionName": "us-central1",
      "zones": [
        "us-central1-a",
        "us-central1-b",
        "us-central1-c",
        "us-central1-f"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "us-central1",
      "assignedZone": "us-central1-a"
    },
    "regionZoneInfoName": "gcp-us-central1",
    "verified": true
  },
  "connectionName": "gcp-us-central1",
  "cspResourceId": "tbdhatv48245si6evr8k",
  "cspResourceName": "tbdhatv48245si6evr8k",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-gcp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "AutoCreateSubnetworks",
      "value": "false"
    },
    {
      "key": "CreationTimestamp",
      "value": "2026-09-07T18:31:43.471-07:00"
    },
    {
      "key": "EnableUlaInternalIpv6",
      "value": "false"
    },
    {
      "key": "Id",
      "value": "7457596897276124048"
    },
    {
      "key": "Kind",
      "value": "compute#network"
    },
    {
      "key": "Mtu",
      "value": "0"
    },
    {
      "key": "Name",
      "value": "tbdhatv48245si6evr8k"
    },
    {
      "key": "NetworkFirewallPolicyEnforcementOrder",
      "value": "AFTER_CLASSIC_FIREWALL"
    },
    {
      "key": "RoutingConfig",
      "value": "{bgpBestPathSelectionMode:LEGACY,routingMode:REGIONAL}"
    },
    {
      "key": "SelfLink",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k"
    },
    {
      "key": "SelfLinkWithId",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/7457596897276124048"
    },
    {
      "key": "Subnetworks",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbvp2k698vd27tqi5rg0"
    }
  ],
  "name": "test-rdbms-vnet-gcp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:08Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:08Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "gcp-us-central1",
        "credentialHolder": "admin",
        "credentialName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "providerName": "gcp",
        "regionDetail": {
          "description": "Council Bluffs Iowa  USA",
          "location": {
            "display": "Council Bluffs Iowa USA",
            "latitude": 41.2522,
            "longitude": -95.8575
          },
          "regionId": "us-central1",
          "regionName": "us-central1",
          "zones": [
            "us-central1-a",
            "us-central1-b",
            "us-central1-c",
            "us-central1-f"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "us-central1",
          "assignedZone": "us-central1-a"
        },
        "regionZoneInfoName": "gcp-us-central1",
        "verified": true
      },
      "connectionName": "gcp-us-central1",
      "cspResourceId": "tbvp2k698vd27tqi5rg0",
      "cspResourceName": "tbvp2k698vd27tqi5rg0",
      "cspVNetId": "tbdhatv48245si6evr8k",
      "cspVNetName": "tbdhatv48245si6evr8k",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.2.1.0/24",
      "keyValueList": [
        {
          "key": "AllowSubnetCidrRoutesOverlap",
          "value": "false"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-09-07T18:31:55.791-07:00"
        },
        {
          "key": "EnableFlowLogs",
          "value": "false"
        },
        {
          "key": "Fingerprint",
          "value": "fJxwwWfohVA="
        },
        {
          "key": "GatewayAddress",
          "value": "10.2.1.1"
        },
        {
          "key": "Id",
          "value": "7851289890438437732"
        },
        {
          "key": "IpCidrRange",
          "value": "10.2.1.0/24"
        },
        {
          "key": "Kind",
          "value": "compute#subnetwork"
        },
        {
          "key": "Name",
          "value": "tbvp2k698vd27tqi5rg0"
        },
        {
          "key": "Network",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k"
        },
        {
          "key": "PrivateIpGoogleAccess",
          "value": "false"
        },
        {
          "key": "PrivateIpv6GoogleAccess",
          "value": "DISABLE_GOOGLE_ACCESS"
        },
        {
          "key": "Purpose",
          "value": "PRIVATE"
        },
        {
          "key": "Region",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbvp2k698vd27tqi5rg0"
        },
        {
          "key": "StackType",
          "value": "IPV4_ONLY"
        },
        {
          "key": "region",
          "value": "us-central1"
        },
        {
          "key": "subnet",
          "value": "tbvp2k698vd27tqi5rg0"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbvp2k698vd27tqi5rg0",
      "zone": "us-central1-a"
    }
  ],
  "systemLabel": "",
  "uid": "tbdhatv48245si6evr8k"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 33.972s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "gcp-us-central1",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Ports": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Ports": "22",
      "Protocol": "TCP"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "vNetId": "test-rdbms-vnet-gcp"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
  "connectionConfig": {
    "configName": "gcp-us-central1",
    "credentialHolder": "admin",
    "credentialName": "gcp",
    "driverName": "gcp-driver-v1.0.so",
    "providerName": "gcp",
    "regionDetail": {
      "description": "Council Bluffs Iowa  USA",
      "location": {
        "display": "Council Bluffs Iowa USA",
        "latitude": 41.2522,
        "longitude": -95.8575
      },
      "regionId": "us-central1",
      "regionName": "us-central1",
      "zones": [
        "us-central1-a",
        "us-central1-b",
        "us-central1-c",
        "us-central1-f"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "us-central1",
      "assignedZone": "us-central1-a"
    },
    "regionZoneInfoName": "gcp-us-central1",
    "verified": true
  },
  "connectionName": "gcp-us-central1",
  "cspResourceId": "tb295v640jrdqjcgdoog",
  "cspResourceName": "tb295v640jrdqjcgdoog",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    }
  ],
  "id": "test-rdbms-sg-gcp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Items",
      "value": "{allowed:[{IPProtocol:tcp,ports:[3306]}],creationTimestamp:2026-09-07T18:32:24.493-07:00,direction:INGRESS,id:6527963891915479879,kind:compute#firewall,logConfig:{},name:tb295v640jrdqjcgdoog-i-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb295v640jrdqjcgdoog-i-001,sourceRanges:[0.0.0.0/0],targetTags:[tb295v640jrdqjcgdoog]}; {allowed:[{IPProtocol:tcp,ports:[22]}],creationTimestamp:2026-09-07T18:32:30.972-07:00,direction:INGRESS,id:1355187267243793217,kind:compute#firewall,logConfig:{},name:tb295v640jrdqjcgdoog-i-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb295v640jrdqjcgdoog-i-002,sourceRanges:[0.0.0.0/0],targetTags:[tb295v640jrdqjcgdoog]}; {creationTimestamp:2026-09-07T18:32:07.579-07:00,denied:[{IPProtocol:all}],destinationRanges:[0.0.0.0/0],direction:EGRESS,id:8433617383384603512,kind:compute#firewall,logConfig:{},name:tb295v640jrdqjcgdoog-o-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k,priority:65535,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb295v640jrdqjcgdoog-o-001,targetTags:[tb295v640jrdqjcgdoog]}; {allowed:[{IPProtocol:all}],creationTimestamp:2026-09-07T18:32:14.053-07:00,destinationRanges:[0.0.0.0/0],direction:EGRESS,id:4993194755034897266,kind:compute#firewall,logConfig:{},name:tb295v640jrdqjcgdoog-o-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbdhatv48245si6evr8k,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb295v640jrdqjcgdoog-o-002,targetTags:[tb295v640jrdqjcgdoog]}"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tb295v640jrdqjcgdoog",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=gcp`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "gcp": {
      "dbOperationMethod": "cspNativeApi",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql"
      ],
      "supportsTag": true
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 1.282s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=gcp-us-central1&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "1-7",
    "connectionName": "gcp-us-central1",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "db-c4a-highmem-16",
      "db-c4a-highmem-2",
      "db-c4a-highmem-32",
      "db-c4a-highmem-4",
      "db-c4a-highmem-48",
      "db-c4a-highmem-64",
      "db-c4a-highmem-72",
      "db-c4a-highmem-8",
      "db-f1-micro",
      "db-g1-small",
      "db-memory-optimized-N-16",
      "db-memory-optimized-N-4",
      "db-memory-optimized-N-8",
      "db-n1-highmem-16",
      "db-n1-highmem-2",
      "db-n1-highmem-32",
      "db-n1-highmem-4",
      "db-n1-highmem-64",
      "db-n1-highmem-8",
      "db-n1-highmem-96",
      "db-n1-standard-1",
      "db-n1-standard-16",
      "db-n1-standard-2",
      "db-n1-standard-32",
      "db-n1-standard-4",
      "db-n1-standard-64",
      "db-n1-standard-8",
      "db-n1-standard-96",
      "db-perf-optimized-N-128",
      "db-perf-optimized-N-16",
      "db-perf-optimized-N-2",
      "db-perf-optimized-N-32",
      "db-perf-optimized-N-4",
      "db-perf-optimized-N-48",
      "db-perf-optimized-N-64",
      "db-perf-optimized-N-8",
      "db-perf-optimized-N-80",
      "db-perf-optimized-N-96"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "131072",
        "name": "db-c4a-highmem-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "16384",
        "name": "db-c4a-highmem-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-c4a-highmem-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "32768",
        "name": "db-c4a-highmem-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db-c4a-highmem-48",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-c4a-highmem-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "589824",
        "name": "db-c4a-highmem-72",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "72"
      },
      {
        "memSizeMiB": "65536",
        "name": "db-c4a-highmem-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "614",
        "name": "db-f1-micro",
        "storageSizeRangeGB": {
          "max": 3279,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "1741",
        "name": "db-g1-small",
        "storageSizeRangeGB": {
          "max": 3279,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-memory-optimized-N-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "131072",
        "name": "db-memory-optimized-N-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-memory-optimized-N-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "106496",
        "name": "db-n1-highmem-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "13312",
        "name": "db-n1-highmem-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "212992",
        "name": "db-n1-highmem-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "26624",
        "name": "db-n1-highmem-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "425984",
        "name": "db-n1-highmem-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "53248",
        "name": "db-n1-highmem-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "638976",
        "name": "db-n1-highmem-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "3840",
        "name": "db-n1-standard-1",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "61440",
        "name": "db-n1-standard-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "7680",
        "name": "db-n1-standard-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "122880",
        "name": "db-n1-standard-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "15360",
        "name": "db-n1-standard-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "245760",
        "name": "db-n1-standard-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "30720",
        "name": "db-n1-standard-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "368640",
        "name": "db-n1-standard-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "884736",
        "name": "db-perf-optimized-N-128",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "131072",
        "name": "db-perf-optimized-N-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "16384",
        "name": "db-perf-optimized-N-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-perf-optimized-N-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "32768",
        "name": "db-perf-optimized-N-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "393216",
        "name": "db-perf-optimized-N-48",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-perf-optimized-N-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "65536",
        "name": "db-perf-optimized-N-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "655360",
        "name": "db-perf-optimized-N-80",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "786432",
        "name": "db-perf-optimized-N-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      }
    ],
    "dbOperationMethod": "",
    "defaultStorageType": "",
    "liveSupportedEngines": [
      "mysql"
    ],
    "notes": {
      "storageTypes": [
        {
          "constraints": "Minimum 20GB storage. Only available on machine series: C4A, N4.",
          "description": "High-performance next-generation storage. Minimum 20 GB. Only available on C4A (db-c4a-highmem-*) and N4 (db-custom-N4-*) machine series.",
          "displayName": "Hyperdisk Balanced",
          "maxSize": 65536,
          "minSize": 20,
          "recommendationLevel": "premium",
          "storageType": "HYPERDISK_BALANCED"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Cost-effective HDD storage for less I/O-intensive workloads. Available on Shared/Dedicated core instances.",
          "displayName": "Persistent Disk HDD",
          "maxSize": 65536,
          "minSize": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_HDD"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Standard SSD storage. Automatically selected for N2 machine series (db-perf-optimized-N-*) and Shared/Dedicated core instances.",
          "displayName": "Persistent Disk SSD",
          "maxSize": 65536,
          "minSize": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_SSD"
        }
      ]
    },
    "providerName": "gcp",
    "regionName": "us-central1",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
      "max": 70369,
      "min": 10
    },
    "storageTypeOptions": [
      "HYPERDISK_BALANCED",
      "PD_HDD",
      "PD_SSD"
    ],
    "supportedVersions": [
      "5.1",
      "5.5",
      "5.6",
      "5.7",
      "8.0",
      "8.0.18",
      "8.0.26",
      "8.0.27",
      "8.0.28",
      "8.0.29",
      "8.0.30",
      "8.0.31",
      "8.0.32",
      "8.0.33",
      "8.0.34",
      "8.0.35",
      "8.0.36",
      "8.0.37",
      "8.0.39",
      "8.0.40",
      "8.0.41",
      "8.0.42",
      "8.0.43",
      "8.0.44",
      "8.0.45",
      "8.0.46",
      "8.4",
      "9.7"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": true,
    "supportsHighAvailability": true,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": true,
    "supportsTag": true
  }
}
```

### 6. Beetle POST Recommend RDBMS [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "gcp",
    "region": "us-central1"
  },
  "sourceRDBMSInstances": [
    {
      "backupRetentionDays": 7,
      "databases": [
        {
          "characterSet": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "engine": "mysql",
      "engineVersion": "8.0",
      "instanceName": "source-mysql-01",
      "iops": 3000,
      "machineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "memoryMb": 4096,
      "port": 3306,
      "publicAccess": true,
      "storageSizeGb": 100,
      "storageType": "SSD",
      "vcpu": 2
    }
  ]
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for gcp (us-central1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "gcp",
    "region": "us-central1"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-gcp",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-gcp"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 14ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "gcp-us-central1",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "name": "rdbms-gcp",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-gcp"
  ],
  "storageSize": 100,
  "storageType": "PD_SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-gcp"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "gcp-us-central1",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "db-n1-standard-2",
    "name": "rdbms-gcp",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-gcp"
    ],
    "storageSize": 100,
    "storageType": "PD_SSD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-gcp"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 3m33.111s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for gcp (us-central1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "gcp",
    "region": "us-central1"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-gcp",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-gcp"
    }
  ]
}
```
```json
// Response Body
{
  "message": "Managed RDBMS instances created successfully",
  "success": true
}
```

### 9. Beetle GET RDBMS Info [✅ SUCCESS]
- **Duration:** 1.335s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp`
```json
// Response Body
{
  "adminUserName": "admin",
  "backupRetentionDays": 7,
  "backupTime": "00:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:35:49Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:35:49Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    }
  ],
  "connectionConfig": {
    "configName": "gcp-us-central1",
    "credentialHolder": "admin",
    "credentialName": "gcp",
    "driverName": "gcp-driver-v1.0.so",
    "providerName": "gcp",
    "regionDetail": {
      "description": "Council Bluffs Iowa  USA",
      "location": {
        "display": "Council Bluffs Iowa USA",
        "latitude": 41.2522,
        "longitude": -95.8575
      },
      "regionId": "us-central1",
      "regionName": "us-central1",
      "zones": [
        "us-central1-a",
        "us-central1-b",
        "us-central1-c",
        "us-central1-f"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "us-central1",
      "assignedZone": "us-central1-a"
    },
    "regionZoneInfoName": "gcp-us-central1",
    "verified": true
  },
  "connectionName": "gcp-us-central1",
  "cspResourceId": "tbvajigo0eprcvoca9pk",
  "cspResourceName": "tbvajigo0eprcvoca9pk",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "dbInstanceType": "ZONAL",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "encryption": true,
  "endpoint": "34.68.70.53:3306",
  "highAvailability": false,
  "id": "test-rdbms-gcp",
  "name": "test-rdbms-gcp",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-gcp"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "PD_SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbvajigo0eprcvoca9pk",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "root",
      "backupRetentionDays": 7,
      "backupTime": "19:02-19:32",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:21Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:21Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "aws-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "aws",
        "driverName": "aws-driver-v1.0.so",
        "providerName": "aws",
        "regionDetail": {
          "description": "Asia Pacific (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c",
            "ap-northeast-2d"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "aws-ap-northeast-2",
        "verified": true
      },
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "tblunavhu1of0ntjighq",
      "cspResourceName": "tblunavhu1of0ntjighq",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-aws",
      "iops": "3000",
      "name": "test-rdbms-aws",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "tagList": [
        {
          "key": "Name",
          "value": "tblunavhu1of0ntjighq"
        }
      ],
      "uid": "tblunavhu1of0ntjighq",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "adminUserName": "azureuser",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:05Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:05Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "credentialHolder": "admin",
        "credentialName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "providerName": "azure",
        "regionDetail": {
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": ""
        },
        "regionZoneInfoName": "azure-koreacentral",
        "verified": true
      },
      "connectionName": "azure-koreacentral",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-azure",
      "name": "test-rdbms-azure",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "status": "Creating",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbtbfiifbfn8uvco6qhr",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "adminUserName": "admin",
      "backupRetentionDays": 7,
      "backupTime": "00:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:35:49Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:35:49Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "gcp-us-central1",
        "credentialHolder": "admin",
        "credentialName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "providerName": "gcp",
        "regionDetail": {
          "description": "Council Bluffs Iowa  USA",
          "location": {
            "display": "Council Bluffs Iowa USA",
            "latitude": 41.2522,
            "longitude": -95.8575
          },
          "regionId": "us-central1",
          "regionName": "us-central1",
          "zones": [
            "us-central1-a",
            "us-central1-b",
            "us-central1-c",
            "us-central1-f"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "us-central1",
          "assignedZone": "us-central1-a"
        },
        "regionZoneInfoName": "gcp-us-central1",
        "verified": true
      },
      "connectionName": "gcp-us-central1",
      "cspResourceId": "tbvajigo0eprcvoca9pk",
      "cspResourceName": "tbvajigo0eprcvoca9pk",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "encryption": true,
      "endpoint": "34.68.70.53:3306",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "name": "test-rdbms-gcp",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbvajigo0eprcvoca9pk",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "adminUserName": "",
      "backupRetentionDays": 7,
      "backupTime": "00:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:51Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:51Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "ncp-kr",
        "credentialHolder": "admin",
        "credentialName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "providerName": "ncp",
        "regionDetail": {
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "regionId": "KR",
          "regionName": "kr",
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionZoneInfoName": "ncp-kr",
        "verified": true
      },
      "connectionName": "ncp-kr",
      "cspResourceId": "145107438",
      "cspResourceName": "tb7k7td81l8pt2p0nhml",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "dbInstanceType": "Stand Alone",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": ":3306",
      "highAvailability": false,
      "id": "test-rdbms-ncp",
      "name": "test-rdbms-ncp",
      "publicAccess": false,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "status": "Creating",
      "storageSize": 10,
      "storageType": "SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb7k7td81l8pt2p0nhml",
      "vNetId": "test-rdbms-vnet-ncp"
    },
    {
      "adminUserName": "myadmin",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:23Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:23Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "nhn-kr1",
        "credentialHolder": "admin",
        "credentialName": "nhn",
        "driverName": "nhn-driver-v1.0.so",
        "providerName": "nhn",
        "regionDetail": {
          "description": "Pangyo (South Korea)",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "regionId": "KR1",
          "regionName": "kr1",
          "zones": [
            "kr-pub-a",
            "kr-pub-b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "KR1",
          "assignedZone": "kr-pub-a"
        },
        "regionZoneInfoName": "nhn-kr1",
        "verified": true
      },
      "connectionName": "nhn-kr1",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-nhn",
      "name": "test-rdbms-nhn",
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbb3k63114huhljjepms",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "adminUserName": "root",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:28Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:28Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "tencent-ap-seoul",
        "credentialHolder": "admin",
        "credentialName": "tencent",
        "driverName": "tencent-driver-v1.0.so",
        "providerName": "tencent",
        "regionDetail": {
          "description": "Seoul",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "regionId": "ap-seoul",
          "regionName": "ap-seoul",
          "zones": [
            "ap-seoul-1",
            "ap-seoul-2"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-seoul",
          "assignedZone": "ap-seoul-1"
        },
        "regionZoneInfoName": "tencent-ap-seoul",
        "verified": true
      },
      "connectionName": "tencent-ap-seoul",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "local_ssd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb7s8uhr7rhep52rtc7p",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 2.439s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database`
```json
// Request Body
{
  "adminUserPassword": "******",
  "databaseName": "sampledb_dyn"
}
```
```json
// Response Body
{
  "message": "Logical database 'sampledb_dyn' created successfully",
  "success": true
}
```

### 12. Beetle GET List Logical Databases [✅ SUCCESS]
- **Duration:** 3.121s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database`
```json
// Response Body
{
  "databases": [
    "mysql",
    "information_schema",
    "performance_schema",
    "sys",
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 191ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m46.714s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 13.991s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m44.398s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 37.891s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 1m51.885s

