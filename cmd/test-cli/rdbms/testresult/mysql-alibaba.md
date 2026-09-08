# Managed RDBMS (MySQL) Test Report: ALIBABA (ap-northeast-2)

- **Test Case:** Alibaba AP-Northeast-2 (Seoul) MySQL Test
- **Date & Time:** 2026-09-08 11:30:25
- **Namespace:** `default`
- **Total Duration:** 21m20.497s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** ALIBABA
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-09-08 11:30:25

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
- **Duration:** 4.075s
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
    "queriedAt": "2026-09-08T02:30:26.836522189Z",
    "region": "ap-northeast-2",
    "source": "alibaba:DescribeAvailableResource",
    "zones": [
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_essd_entry",
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2a"
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
  "suggestedSystemDisk": "cloud_auto",
  "suggestedZone": "ap-northeast-2b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 7.44s
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
      "zone": ""
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
      "lastTransitionTime": "2026-09-08T02:30:36Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T02:30:36Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T02:30:36Z",
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
  "cspResourceId": "vpc-mj7sy7k1yfqgketm0pr7p",
  "cspResourceName": "tb6olom16v7p7o3fpp63",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreationTime",
      "value": "2026-09-08T02:30:28Z"
    },
    {
      "key": "Status",
      "value": "Available"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7sy7k1yfqgketm0pr7p"
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
      "value": "tb6olom16v7p7o3fpp63"
    },
    {
      "key": "VRouterId",
      "value": "vrt-mj7i8dreprslopmoqwuti"
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
      "value": "{VSwitchId:[vsw-mj7t5owdc32gn4s3rglg8]}"
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
      "value": "{RouterTableIds:[vtb-mj7us4upbebr13ymdj3lz]}"
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
          "lastTransitionTime": "2026-09-08T02:30:36Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:30:36Z",
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
      "cspResourceId": "vsw-mj7t5owdc32gn4s3rglg8",
      "cspResourceName": "tbpv3i3ivgrmcvjjmdfh",
      "cspVNetId": "vpc-mj7sy7k1yfqgketm0pr7p",
      "cspVNetName": "tb6olom16v7p7o3fpp63",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.3.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-mj7sy7k1yfqgketm0pr7p"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "CreationTime",
          "value": "2026-09-08T02:30:31Z"
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
          "value": "vsw-mj7t5owdc32gn4s3rglg8"
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
          "value": "tbpv3i3ivgrmcvjjmdfh"
        },
        {
          "key": "EnabledIpv6",
          "value": "false"
        },
        {
          "key": "RouteTable",
          "value": "{ResourceGroupId:,CreationTime:,Status:,RouteTableType:System,VRouterId:,RouteTableId:vtb-mj7us4upbebr13ymdj3lz,VSwitchIds:{VSwitchId:null},RouteEntrys:{RouteEntry:null}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbpv3i3ivgrmcvjjmdfh",
      "zone": "ap-northeast-2a"
    }
  ],
  "systemLabel": "",
  "uid": "tb6olom16v7p7o3fpp63"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 4.815s
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
  "cspResourceId": "sg-mj73lk0dvuqndjvvuk5h",
  "cspResourceName": "tba277jjh06dmg20kife",
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
      "value": "sg-mj73lk0dvuqndjvvuk5h"
    },
    {
      "key": "SecurityGroupName",
      "value": "tba277jjh06dmg20kife"
    },
    {
      "key": "Description",
      "value": "tba277jjh06dmg20kife"
    },
    {
      "key": "SecurityGroupType",
      "value": "enterprise"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7sy7k1yfqgketm0pr7p"
    },
    {
      "key": "CreationTime",
      "value": "2026-09-08T02:30:35Z"
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
  "uid": "tba277jjh06dmg20kife",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 5ms
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
- **Duration:** 2m56.7s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=alibaba-ap-northeast-2&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-730",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "myduck.x2.2xlarge.xc",
      "myduck.x2.4xlarge.xc",
      "myduck.x2.8xlarge.xc",
      "myduck.x2.large.xc",
      "myduck.x2.xlarge.xc",
      "myduck.x4.2xlarge.xc",
      "myduck.x4.4xlarge.xc",
      "myduck.x4.8xlarge.xc",
      "myduck.x4.large.xc",
      "myduck.x4.xlarge.xc",
      "myduck.x8.2xlarge.xc",
      "myduck.x8.4xlarge.xc",
      "myduck.x8.8xlarge.xc",
      "myduck.x8.large.xc",
      "myduck.x8.xlarge.xc",
      "mysql.n2.large.xc",
      "mysql.n2.medium.2c",
      "mysql.n2.medium.xc",
      "mysql.n2.small.2c",
      "mysql.n2.xlarge.xc",
      "mysql.n2e.large.xc",
      "mysql.n2e.medium.xc",
      "mysql.n2e.small.xc",
      "mysql.n2e.xlarge.xc",
      "mysql.n2m.large.2c",
      "mysql.n2m.medium.2c",
      "mysql.n2m.small.2c",
      "mysql.n2m.xlarge.2c",
      "mysql.n4.large.xc",
      "mysql.n4.medium.xc",
      "mysql.n4.xlarge.xc",
      "mysql.n4e.large.xc",
      "mysql.n4e.medium.xc",
      "mysql.n4e.xlarge.xc",
      "mysql.n4m.large.2c",
      "mysql.n4m.medium.2c",
      "mysql.n4m.xlarge.2c",
      "mysql.n8.large.xc",
      "mysql.n8.medium.xc",
      "mysql.n8.xlarge.xc",
      "mysql.n8e.large.xc",
      "mysql.n8e.medium.xc",
      "mysql.n8e.xlarge.xc",
      "mysql.n8m.large.2c",
      "mysql.n8m.medium.2c",
      "mysql.n8m.xlarge.2c",
      "mysql.x2.13large.2c",
      "mysql.x2.13large.xc",
      "mysql.x2.13xlarge.2c",
      "mysql.x2.13xlarge.xc",
      "mysql.x2.2xlarge.2c",
      "mysql.x2.2xlarge.xc",
      "mysql.x2.3large.2c",
      "mysql.x2.3large.xc",
      "mysql.x2.3xlarge.2c",
      "mysql.x2.3xlarge.xc",
      "mysql.x2.4xlarge.2c",
      "mysql.x2.4xlarge.xc",
      "mysql.x2.6xlarge.2c",
      "mysql.x2.8xlarge.2c",
      "mysql.x2.8xlarge.xc",
      "mysql.x2.large.2c",
      "mysql.x2.large.xc",
      "mysql.x2.medium.2c",
      "mysql.x2.medium.xc",
      "mysql.x2.xlarge.2c",
      "mysql.x2.xlarge.xc",
      "mysql.x2e.2xlarge.xc",
      "mysql.x2e.4xlarge.xc",
      "mysql.x2e.8xlarge.xc",
      "mysql.x2e.large.xc",
      "mysql.x2e.medium.xc",
      "mysql.x2e.xlarge.xc",
      "mysql.x2m.2xlarge.2c",
      "mysql.x2m.4xlarge.2c",
      "mysql.x2m.large.2c",
      "mysql.x2m.medium.2c",
      "mysql.x2m.xlarge.2c",
      "mysql.x4.13large.2c",
      "mysql.x4.13large.xc",
      "mysql.x4.13xlarge.2c",
      "mysql.x4.13xlarge.xc",
      "mysql.x4.16xlarge.2c",
      "mysql.x4.2xlarge.2c",
      "mysql.x4.2xlarge.xc",
      "mysql.x4.3large.2c",
      "mysql.x4.3large.xc",
      "mysql.x4.3xlarge.2c",
      "mysql.x4.3xlarge.xc",
      "mysql.x4.4xlarge.2c",
      "mysql.x4.4xlarge.xc",
      "mysql.x4.6xlarge.2c",
      "mysql.x4.8xlarge.2c",
      "mysql.x4.8xlarge.xc",
      "mysql.x4.large.2c",
      "mysql.x4.large.xc",
      "mysql.x4.medium.2c",
      "mysql.x4.medium.xc",
      "mysql.x4.xlarge.2c",
      "mysql.x4.xlarge.xc",
      "mysql.x4e.2xlarge.xc",
      "mysql.x4e.4xlarge.xc",
      "mysql.x4e.8xlarge.xc",
      "mysql.x4e.large.xc",
      "mysql.x4e.medium.xc",
      "mysql.x4e.xlarge.xc",
      "mysql.x4m.2xlarge.2c",
      "mysql.x4m.4xlarge.2c",
      "mysql.x4m.8xlarge.2c",
      "mysql.x4m.large.2c",
      "mysql.x4m.medium.2c",
      "mysql.x4m.xlarge.2c",
      "mysql.x8.13large.2c",
      "mysql.x8.13large.xc",
      "mysql.x8.13xlarge.2c",
      "mysql.x8.13xlarge.xc",
      "mysql.x8.2xlarge.2c",
      "mysql.x8.2xlarge.xc",
      "mysql.x8.3large.2c",
      "mysql.x8.3large.xc",
      "mysql.x8.3xlarge.2c",
      "mysql.x8.3xlarge.xc",
      "mysql.x8.4xlarge.2c",
      "mysql.x8.4xlarge.xc",
      "mysql.x8.6xlarge.2c",
      "mysql.x8.8xlarge.2c",
      "mysql.x8.8xlarge.xc",
      "mysql.x8.large.2c",
      "mysql.x8.large.xc",
      "mysql.x8.medium.2c",
      "mysql.x8.medium.xc",
      "mysql.x8.xlarge.2c",
      "mysql.x8.xlarge.xc",
      "mysql.x8e.2xlarge.xc",
      "mysql.x8e.4xlarge.xc",
      "mysql.x8e.8xlarge.xc",
      "mysql.x8e.large.xc",
      "mysql.x8e.medium.xc",
      "mysql.x8e.xlarge.xc",
      "mysql.x8m.2xlarge.2c",
      "mysql.x8m.4xlarge.2c",
      "mysql.x8m.8xlarge.2c",
      "mysql.x8m.large.2c",
      "mysql.x8m.medium.2c",
      "mysql.x8m.xlarge.2c",
      "myduck.n2.2xlarge.1",
      "myduck.n2.large.1",
      "myduck.n2.xlarge.1",
      "myduck.n4.2xlarge.1",
      "myduck.n4.large.1",
      "myduck.n4.xlarge.1",
      "myduck.n8.2xlarge.1",
      "myduck.n8.large.1",
      "myduck.n8.xlarge.1",
      "mysql.n1.micro.1",
      "mysql.n1e.medium.1",
      "mysql.n1e.small.1",
      "mysql.n2.large.1",
      "mysql.n2.medium.1",
      "mysql.n2.small.1",
      "mysql.n2.xlarge.1",
      "mysql.n2e.medium.1",
      "mysql.n2e.small.1",
      "mysql.n4.large.1",
      "mysql.n4.medium.1",
      "mysql.n4.xlarge.1"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "32768",
        "name": "myduck.n2.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "8192",
        "name": "myduck.n2.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "myduck.n2.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "myduck.n4.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "16384",
        "name": "myduck.n4.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "myduck.n4.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "myduck.n8.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "32768",
        "name": "myduck.n8.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "65536",
        "name": "myduck.n8.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "2048",
        "name": "mysql.n1e.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "mysql.n1e.small.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "8192",
        "name": "mysql.n2.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "4096",
        "name": "mysql.n2.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "mysql.n2.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "4096",
        "name": "mysql.n2e.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "mysql.n2e.small.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "16384",
        "name": "mysql.n4.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "8192",
        "name": "mysql.n4.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "mysql.n4.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
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
          "description": "Alibaba Cloud automatically selects optimal SSD storage (General ESSD / PL1) based on region and instance specification. Recommended for simplicity.",
          "displayName": "Auto-selected Storage Type (SSD)",
          "maxSize": 32768,
          "minSize": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "cloud_auto"
        },
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
      "max": 64000,
      "min": 20
    },
    "storageTypeOptions": [
      "cloud_auto",
      "cloud_essd",
      "cloud_essd2",
      "cloud_essd3",
      "cloud_ssd"
    ],
    "supportedVersions": [
      "5.7",
      "8.0",
      "8.4"
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
- **Duration:** 5ms
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_auto",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 1m36.416s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "alibaba-ap-northeast-2",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "mysql.n2.medium.1",
  "name": "rdbms-alibaba",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "storageSize": 100,
  "storageType": "cloud_auto",
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
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "mysql.n2.medium.1",
    "name": "rdbms-alibaba",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-alibaba"
    ],
    "storageSize": 100,
    "storageType": "cloud_auto",
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
- **Duration:** 4m14.46s
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_auto",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
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
- **Duration:** 10ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba`
```json
// Response Body
{
  "adminUserName": "dbadmin",
  "backupRetentionDays": 7,
  "backupTime": "20:00Z-21:00Z",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T02:38:42Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T02:38:42Z",
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
  "cspResourceId": "rm-mj7c42v27wo27vuc2",
  "cspResourceName": "tbmabhhcvjuvqvfsk58r",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "mysql.n2.medium.1",
  "dbInstanceType": "Basic",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "endpoint": "43.108.66.51:3306",
  "highAvailability": false,
  "id": "test-rdbms-alibaba",
  "name": "test-rdbms-alibaba",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "general_essd",
  "subnetIds": [
    "subnet-1"
  ],
  "tagList": [
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.name",
      "value": "test-rdbms-alibaba"
    },
    {
      "key": "sys.uid",
      "value": "tbmabhhcvjuvqvfsk58r"
    },
    {
      "key": "sys.cspResourceId",
      "value": "rm-mj7c42v27wo27vuc2"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbmabhhcvjuvqvfsk58r"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.id",
      "value": "test-rdbms-alibaba"
    },
    {
      "key": "sys.connectionName",
      "value": "alibaba-ap-northeast-2"
    }
  ],
  "uid": "tbmabhhcvjuvqvfsk58r",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 8ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "backupTime": "20:00Z-21:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T02:38:42Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:38:42Z",
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
      "cspResourceId": "rm-mj7c42v27wo27vuc2",
      "cspResourceName": "tbmabhhcvjuvqvfsk58r",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "dbInstanceType": "Basic",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "43.108.66.51:3306",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "general_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbmabhhcvjuvqvfsk58r",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "adminUserName": "admin",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T02:31:11Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:31:11Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "ibm-us-south",
        "credentialHolder": "admin",
        "credentialName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "providerName": "ibm",
        "regionDetail": {
          "description": "us-south",
          "location": {
            "display": "Dallas USA",
            "latitude": 32.81248,
            "longitude": -96.77619
          },
          "regionId": "us-south",
          "regionName": "us-south",
          "zones": [
            "us-south-1",
            "us-south-2",
            "us-south-3"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "us-south",
          "assignedZone": "us-south-1"
        },
        "regionZoneInfoName": "ibm-us-south",
        "verified": true
      },
      "connectionName": "ibm-us-south",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "b3c.4x16.encrypted",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-ibm",
      "name": "test-rdbms-ibm",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "status": "Creating",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tblv8isdiqhulslqdbtk",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "adminUserName": "admin",
      "backupTime": "NA",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T02:35:23Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:35:23Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "d3f03149-a472-47e3-8840-2e7c44bf0723",
      "cspResourceName": "tbdvg8am5bhagcijql1b",
      "dbEngine": "mysql",
      "dbEngineVersion": "5.7.29",
      "dbInstanceSpec": "m1.medium",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "183.111.177.152:3306",
      "highAvailability": false,
      "id": "test-rdbms-openstack",
      "name": "test-rdbms-openstack",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "NA",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbdvg8am5bhagcijql1b",
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.477s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database`
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
- **Duration:** 2.693s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database`
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
- **Duration:** 6ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m10.179s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 9.508s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 8m42.953s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.203s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 8.544s

