# Managed RDBMS (MariaDB) Test Report: ALIBABA (ap-northeast-2)

- **Test Case:** Alibaba AP-Northeast-2 (Seoul) MariaDB Test
- **Date & Time:** 2026-09-08 10:07:29
- **Namespace:** `default`
- **Total Duration:** 16m45.355s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** ALIBABA
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-09-08 10:07:29

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
- **Duration:** 3.539s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
  "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.large"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "ecs.e-c1m2.large",
    "provider": "alibaba",
    "queriedAt": "2026-09-08T01:07:31.251162473Z",
    "region": "ap-northeast-2",
    "source": "alibaba:DescribeAvailableResource",
    "zones": [
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_essd_entry",
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2a"
      },
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2b"
      }
    ]
  },
  "connectionName": "alibaba-ap-northeast-2",
  "estimatedCost": "$0.0356/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "alibaba-us-west-1",
    "creationDate": "",
    "cspImageName": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "description": "Kernel version is 6.8.0-137-generic, 2026.8.13",
    "details": [
      {
        "key": "BootMode",
        "value": "UEFI-Preferred"
      },
      {
        "key": "ImageId",
        "value": "ubuntu_24_04_x64_20G_alibase_20260810.vhd"
      },
      {
        "key": "ImageOwnerAlias",
        "value": "system"
      },
      {
        "key": "OSName",
        "value": "Ubuntu  24.04 64位"
      },
      {
        "key": "OSNameEn",
        "value": "Ubuntu  24.04 64 bit"
      },
      {
        "key": "ImageFamily",
        "value": "acs:ubuntu_24_04_x64"
      },
      {
        "key": "Architecture",
        "value": "x86_64"
      },
      {
        "key": "IsSupportIoOptimized",
        "value": "true"
      },
      {
        "key": "Size",
        "value": "20"
      },
      {
        "key": "Description",
        "value": "Kernel version is 6.8.0-137-generic, 2026.8.13"
      },
      {
        "key": "Usage",
        "value": "instance"
      },
      {
        "key": "IsCopied",
        "value": "false"
      },
      {
        "key": "LoginAsNonRootSupported",
        "value": "true"
      },
      {
        "key": "ImageVersion",
        "value": "v2026.8.13"
      },
      {
        "key": "OSType",
        "value": "linux"
      },
      {
        "key": "IsSubscribed",
        "value": "false"
      },
      {
        "key": "IsSupportCloudinit",
        "value": "true"
      },
      {
        "key": "CreationTime",
        "value": "2026-08-13T01:49:37Z"
      },
      {
        "key": "Progress",
        "value": "100%"
      },
      {
        "key": "Platform",
        "value": "Ubuntu"
      },
      {
        "key": "ImageName",
        "value": "ubuntu_24_04_x64_20G_alibase_20260810.vhd"
      },
      {
        "key": "Status",
        "value": "Available"
      },
      {
        "key": "ImageOwnerId",
        "value": "0"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "DetectionOptions",
        "value": "{Status:,Items:{Item:null}}"
      },
      {
        "key": "Features",
        "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
      },
      {
        "key": "Tags",
        "value": "{Tag:[]}"
      },
      {
        "key": "DiskDeviceMappings",
        "value": "{DiskDeviceMapping:[]}"
      }
    ],
    "fetchedTime": "2026.08.21 13:58:28 Fri",
    "id": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 20,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu  24.04 64 bit",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "alibaba",
    "regionList": [
      "ap-northeast-1",
      "ap-northeast-2",
      "ap-southeast-1",
      "ap-southeast-3",
      "ap-southeast-5",
      "ap-southeast-6",
      "ap-southeast-7",
      "ap-southeast-8",
      "cn-beijing",
      "cn-chengdu",
      "cn-fuzhou",
      "cn-guangzhou",
      "cn-hangzhou",
      "cn-heyuan",
      "cn-hongkong",
      "cn-huhehaote",
      "cn-nanjing",
      "cn-qingdao",
      "cn-shanghai",
      "cn-shenzhen",
      "cn-wuhan-lr",
      "cn-wulanchabu",
      "cn-zhangjiakou",
      "cn-zhongwei",
      "eu-central-1",
      "eu-west-1",
      "eu-west-2",
      "me-central-1",
      "me-east-1",
      "na-south-1",
      "us-east-1",
      "us-west-1"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbm30neucqi9cs1ci7ut"
  },
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
  "imageValidation": {
    "cspResourceId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "isAvailable": true,
    "resourceId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "resourceName": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "alibaba",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "alibaba-ap-northeast-2",
    "costPerHour": 0.0356,
    "cspSpecName": "ecs.e-c1m2.large",
    "details": [
      {
        "key": "CpuArchitecture",
        "value": "X86"
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
    "id": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "namespace": "system",
    "providerName": "alibaba",
    "regionLatitude": 37.36,
    "regionLongitude": 126.78,
    "regionName": "ap-northeast-2",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tba6uee340r6ln51e0hd",
    "vCPU": 2
  },
  "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
  "specValidation": {
    "cspResourceId": "ecs.e-c1m2.large",
    "isAvailable": true,
    "resourceId": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "resourceName": "ecs.e-c1m2.large",
    "status": "Available"
  },
  "status": "OK",
  "suggestedSystemDisk": "cloud_essd_entry",
  "suggestedZone": "ap-northeast-2a"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 7.196s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.3.0.0/16",
  "connectionName": "alibaba-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-alibaba",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.3.1.0/24",
      "name": "subnet-1",
      "zone": "ap-northeast-2a"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.3.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:07:40Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:07:40Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:07:40Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
    }
  ],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "vpc-mj7t88h10n92d9wb39xq8",
  "cspResourceName": "tb6rfvg6v0669mu4of3p",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreationTime",
      "value": "2026-09-08T01:07:32Z"
    },
    {
      "key": "Status",
      "value": "Available"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7t88h10n92d9wb39xq8"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "AdvancedResource",
      "value": "false"
    },
    {
      "key": "OwnerId",
      "value": "5469257408566579"
    },
    {
      "key": "RegionId",
      "value": "ap-northeast-2"
    },
    {
      "key": "VpcName",
      "value": "tb6rfvg6v0669mu4of3p"
    },
    {
      "key": "VRouterId",
      "value": "vrt-mj7sxr5gppjs5t4k6qxhd"
    },
    {
      "key": "CidrBlock",
      "value": "10.3.0.0/16"
    },
    {
      "key": "NetworkAclNum",
      "value": "0"
    },
    {
      "key": "SupportAdvancedFeature",
      "value": "false"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "CenStatus",
      "value": "Detached"
    },
    {
      "key": "EnabledIpv6",
      "value": "false"
    },
    {
      "key": "DnsHostnameStatus",
      "value": "DISABLED"
    },
    {
      "key": "VSwitchIds",
      "value": "{VSwitchId:[vsw-mj7iris6e0fz5a8ugl1c4]}"
    },
    {
      "key": "SecondaryCidrBlocks",
      "value": "{SecondaryCidrBlock:[]}"
    },
    {
      "key": "UserCidrs",
      "value": "{UserCidr:[]}"
    },
    {
      "key": "NatGatewayIds",
      "value": "{NatGatewayIds:[]}"
    },
    {
      "key": "RouterTableIds",
      "value": "{RouterTableIds:[vtb-mj7lr2fct6lubgf0gimf6]}"
    },
    {
      "key": "Tags",
      "value": "{Tag:null}"
    },
    {
      "key": "Ipv6CidrBlocks",
      "value": "{Ipv6CidrBlock:null}"
    }
  ],
  "name": "test-rdbms-vnet-alibaba",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:07:40Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:07:40Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "providerName": "alibaba",
        "regionDetail": {
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "verified": true
      },
      "connectionName": "alibaba-ap-northeast-2",
      "cspResourceId": "vsw-mj7iris6e0fz5a8ugl1c4",
      "cspResourceName": "tboc3bq88vfronrs92ee",
      "cspVNetId": "vpc-mj7t88h10n92d9wb39xq8",
      "cspVNetName": "tb6rfvg6v0669mu4of3p",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.3.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-mj7t88h10n92d9wb39xq8"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "CreationTime",
          "value": "2026-09-08T01:07:35Z"
        },
        {
          "key": "IsDefault",
          "value": "false"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "252"
        },
        {
          "key": "OwnerId",
          "value": "5469257408566579"
        },
        {
          "key": "VSwitchId",
          "value": "vsw-mj7iris6e0fz5a8ugl1c4"
        },
        {
          "key": "CidrBlock",
          "value": "10.3.1.0/24"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "VSwitchName",
          "value": "tboc3bq88vfronrs92ee"
        },
        {
          "key": "EnabledIpv6",
          "value": "false"
        },
        {
          "key": "RouteTable",
          "value": "{ResourceGroupId:,CreationTime:,Status:,RouteTableType:System,VRouterId:,RouteTableId:vtb-mj7lr2fct6lubgf0gimf6,VSwitchIds:{VSwitchId:null},RouteEntrys:{RouteEntry:null}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tboc3bq88vfronrs92ee",
      "zone": "ap-northeast-2a"
    }
  ],
  "systemLabel": "",
  "uid": "tb6rfvg6v0669mu4of3p"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 3.663s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "alibaba-ap-northeast-2",
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
  "name": "test-rdbms-sg-alibaba",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "sg-mj74dlpc6t12slkuncb6",
  "cspResourceName": "tbq01oi3197iglb1nr51",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    }
  ],
  "id": "test-rdbms-sg-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "SecurityGroupId",
      "value": "sg-mj74dlpc6t12slkuncb6"
    },
    {
      "key": "SecurityGroupName",
      "value": "tbq01oi3197iglb1nr51"
    },
    {
      "key": "Description",
      "value": "tbq01oi3197iglb1nr51"
    },
    {
      "key": "SecurityGroupType",
      "value": "enterprise"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7t88h10n92d9wb39xq8"
    },
    {
      "key": "CreationTime",
      "value": "2026-09-08T01:07:38Z"
    },
    {
      "key": "EcsCount",
      "value": "0"
    },
    {
      "key": "AvailableInstanceAmount",
      "value": "0"
    },
    {
      "key": "ServiceManaged",
      "value": "false"
    },
    {
      "key": "ServiceID",
      "value": "0"
    },
    {
      "key": "RuleCount",
      "value": "3"
    },
    {
      "key": "GroupToGroupRuleCount",
      "value": "0"
    },
    {
      "key": "Tags",
      "value": "{Tag:[]}"
    }
  ],
  "name": "test-rdbms-sg-alibaba",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbq01oi3197iglb1nr51",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=alibaba`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "alibaba": {
      "dbOperationMethod": "cspNativeApi",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql",
        "mariadb"
      ],
      "supportsTag": true
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 26.398s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=alibaba-ap-northeast-2&dbEngine=mariadb`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-730",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbInstanceSpecOptions": [
      "mariadb.n2.medium.2c",
      "mariadb.n2.small.2c",
      "mariadb.x2.2xlarge.2c",
      "mariadb.x2.large.2c",
      "mariadb.x2.xlarge.2c",
      "mariadb.x4.2xlarge.2c",
      "mariadb.x4.4xlarge.2c",
      "mariadb.x4.8xlarge.2c",
      "mariadb.x4.large.2c",
      "mariadb.x4.xlarge.2c",
      "mariadb.x8.2xlarge.2c",
      "mariadb.x8.4xlarge.2c",
      "mariadb.x8.8xlarge.2c"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "4096",
        "name": "mariadb.n2.medium.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "mariadb.n2.small.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "32768",
        "name": "mariadb.x2.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "8192",
        "name": "mariadb.x2.large.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "mariadb.x2.xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "mariadb.x4.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "mariadb.x4.4xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "229376",
        "name": "mariadb.x4.8xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "56 "
      },
      {
        "memSizeMiB": "16384",
        "name": "mariadb.x4.large.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "mariadb.x4.xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "mariadb.x8.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "mariadb.x8.4xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "491520",
        "name": "mariadb.x8.8xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "56 "
      }
    ],
    "dbOperationMethod": "",
    "defaultStorageType": "",
    "liveSupportedEngines": [
      "mysql",
      "mariadb"
    ],
    "notes": {
      "storageTypes": [
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Standard enhanced SSD with good balance of performance and cost. Performance Level 1.",
          "displayName": "Enhanced SSD (ESSD PL1)",
          "maxSize": 32768,
          "minSize": 20,
          "recommendationLevel": "standard",
          "storageType": "cloud_essd"
        },
        {
          "constraints": "Minimum 500GB storage. Not compatible with dbInstanceSpec(s): mysql.n4.*.",
          "description": "High-performance enhanced SSD. Performance Level 2. Minimum 500 GB storage. Not compatible with mysql.n4.* instance specifications.",
          "displayName": "Enhanced SSD (ESSD PL2)",
          "incompatibleSpecs": [
            "mysql.n4.*"
          ],
          "maxSize": 32768,
          "minSize": 500,
          "recommendationLevel": "premium",
          "storageType": "cloud_essd2"
        },
        {
          "constraints": "Minimum 1500GB storage. Not compatible with dbInstanceSpec(s): mysql.n4.*.",
          "description": "Ultra-high-performance enhanced SSD. Performance Level 3. Minimum 1,500 GB storage. Not compatible with mysql.n4.* instance specifications.",
          "displayName": "Enhanced SSD (ESSD PL3)",
          "incompatibleSpecs": [
            "mysql.n4.*"
          ],
          "maxSize": 32768,
          "minSize": 1500,
          "recommendationLevel": "premium",
          "storageType": "cloud_essd3"
        },
        {
          "description": "Storage type details not yet documented.",
          "displayName": "cloud_ssd",
          "storageType": "cloud_ssd"
        }
      ]
    },
    "providerName": "alibaba",
    "regionName": "ap-northeast-2",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 32000,
      "min": 20
    },
    "storageTypeOptions": [
      "cloud_essd",
      "cloud_essd2",
      "cloud_essd3",
      "cloud_ssd"
    ],
    "supportedVersions": [
      "10.3",
      "10.6"
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
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
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
      "engine": "mariadb",
      "engineVersion": "10.6",
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for alibaba (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 1500,
      "storageType": "cloud_essd3",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ],
  "warnings": [
    "Adjusted storage size for instance 'source-mysql-01' from 100GB to minimum 1500GB required by target cloud (cloud_essd3)."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 31.138s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "alibaba-ap-northeast-2",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6",
  "dbInstanceSpec": "mariadb.n2.medium.2c",
  "name": "rdbms-mariadb-alibaba",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "storageSize": 1500,
  "storageType": "cloud_essd3",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-alibaba"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbEngineVersion": "10.6",
    "dbInstanceSpec": "mariadb.n2.medium.2c",
    "name": "rdbms-mariadb-alibaba",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-alibaba"
    ],
    "storageSize": 1500,
    "storageType": "cloud_essd3",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-alibaba"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 3m9.09s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for alibaba (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 1500,
      "storageType": "cloud_essd3",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ],
  "warnings": [
    "Adjusted storage size for instance 'source-mysql-01' from 100GB to minimum 1500GB required by target cloud (cloud_essd3)."
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
- **Duration:** 17ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba`
```json
// Response Body
{
  "adminUserName": "dbadmin",
  "backupRetentionDays": 7,
  "backupTime": "18:00Z-19:00Z",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:11:26Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:11:26Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    }
  ],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "rm-mj78nw56y668pn2xm",
  "cspResourceName": "tbitr22omg4oanvd5bat",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6",
  "dbInstanceSpec": "mariadb.n2.medium.2c",
  "dbInstanceType": "HighAvailability",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "endpoint": "43.108.66.87:3306",
  "highAvailability": true,
  "id": "test-rdbms-mariadb-alibaba",
  "name": "test-rdbms-mariadb-alibaba",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "status": "Available",
  "storageSize": 1500,
  "storageType": "cloud_essd3",
  "subnetIds": [
    "subnet-1"
  ],
  "tagList": [
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.id",
      "value": "test-rdbms-mariadb-alibaba"
    },
    {
      "key": "sys.cspResourceId",
      "value": "rm-mj78nw56y668pn2xm"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbitr22omg4oanvd5bat"
    },
    {
      "key": "sys.connectionName",
      "value": "alibaba-ap-northeast-2"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.name",
      "value": "test-rdbms-mariadb-alibaba"
    },
    {
      "key": "sys.uid",
      "value": "tbitr22omg4oanvd5bat"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    }
  ],
  "uid": "tbitr22omg4oanvd5bat",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 17ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "backupTime": "18:00Z-19:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:11:26Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:11:26Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "providerName": "alibaba",
        "regionDetail": {
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "verified": true
      },
      "connectionName": "alibaba-ap-northeast-2",
      "cspResourceId": "rm-mj78nw56y668pn2xm",
      "cspResourceName": "tbitr22omg4oanvd5bat",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "43.108.66.87:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
      "name": "test-rdbms-mariadb-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Available",
      "storageSize": 1500,
      "storageType": "cloud_essd3",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbitr22omg4oanvd5bat",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "adminUserName": "root",
      "backupRetentionDays": 7,
      "backupTime": "20:15-20:45",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:09:05Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:09:05Z",
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
      "cspResourceId": "tbvjs61lnopefchkhfmj",
      "cspResourceName": "tbvjs61lnopefchkhfmj",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.27",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-aws",
      "iops": "3000",
      "name": "test-rdbms-mariadb-aws",
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
          "value": "tbvjs61lnopefchkhfmj"
        }
      ],
      "uid": "tbvjs61lnopefchkhfmj",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "adminUserName": "myadmin",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:08:13Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:08:13Z",
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
      "dbEngine": "mariadb",
      "dbEngineVersion": "MARIADB_V101118",
      "dbInstanceSpec": "m2.c2m4",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-nhn",
      "name": "test-rdbms-mariadb-nhn",
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
      "uid": "tbi7jnt1v41387icedng",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "adminUserName": "admin",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:08:12Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:08:12Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "openstack-regionone",
        "credentialHolder": "admin",
        "credentialName": "openstack",
        "driverName": "openstack-driver-v1.0.so",
        "providerName": "openstack",
        "regionDetail": {
          "description": "Korea Daejeon (Internal)",
          "location": {
            "display": "South Korea (Daejeon)",
            "latitude": 36.3804,
            "longitude": 127.365
          },
          "regionId": "RegionOne",
          "regionName": "regionone",
          "zones": [
            "nova"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "RegionOne",
          "assignedZone": "nova"
        },
        "regionZoneInfoName": "openstack-regionone",
        "verified": true
      },
      "connectionName": "openstack-regionone",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.4",
      "dbInstanceSpec": "m1.medium",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-openstack",
      "name": "test-rdbms-mariadb-openstack",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbr9fm05505msvq0ppj4",
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.151s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database`
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
- **Duration:** 2.009s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database`
```json
// Response Body
{
  "databases": [
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 7ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m21.642s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 7.25s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 8m42.591s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.471s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 8.165s

