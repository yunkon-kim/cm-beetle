# Managed RDBMS (MySQL) Test Report: TENCENT (ap-seoul)

- **Test Case:** Tencent AP-Seoul MySQL Test
- **Date & Time:** 2026-09-08 10:31:39
- **Namespace:** `default`
- **Total Duration:** 12m26.534s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** TENCENT
- **Target Region:** `ap-seoul`
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
- **Duration:** 3.99s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "img-487zeit5",
  "specId": "tencent+ap-seoul+sa2.medium4"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "SA2.MEDIUM4",
    "provider": "tencent",
    "queriedAt": "2026-09-08T01:31:43.042699969Z",
    "region": "ap-seoul",
    "source": "tencent:DescribeZoneInstanceConfigInfos",
    "zones": [
      {
        "available": true,
        "status": "SELL",
        "zoneId": "ap-seoul-1"
      },
      {
        "available": true,
        "status": "SELL",
        "zoneId": "ap-seoul-2"
      }
    ]
  },
  "connectionName": "tencent-ap-seoul",
  "estimatedCost": "$0.0400/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "tencent-sa-saopaulo",
    "creationDate": "",
    "cspImageName": "img-487zeit5",
    "description": "",
    "details": [
      {
        "key": "ImageId",
        "value": "img-487zeit5"
      },
      {
        "key": "OsName",
        "value": "Ubuntu Server 22.04 LTS 64bit"
      },
      {
        "key": "ImageType",
        "value": "PUBLIC_IMAGE"
      },
      {
        "key": "ImageName",
        "value": "Ubuntu Server 22.04 LTS 64bit"
      },
      {
        "key": "ImageDescription",
        "value": "Ubuntu Server 22.04 LTS 64bit"
      },
      {
        "key": "ImageSize",
        "value": "20"
      },
      {
        "key": "Architecture",
        "value": "x86_64"
      },
      {
        "key": "ImageState",
        "value": "NORMAL"
      },
      {
        "key": "Platform",
        "value": "Ubuntu"
      },
      {
        "key": "ImageSource",
        "value": "OFFICIAL"
      },
      {
        "key": "IsSupportCloudinit",
        "value": "true"
      },
      {
        "key": "ImageDeprecated",
        "value": "false"
      }
    ],
    "fetchedTime": "2026.08.21 13:57:36 Fri",
    "id": "img-487zeit5",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "img-487zeit5",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 20,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu Server 22.04 LTS 64bit",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 22.04",
    "providerName": "tencent",
    "regionList": [
      "ap-bangkok",
      "ap-beijing",
      "ap-chengdu",
      "ap-chongqing",
      "ap-guangzhou",
      "ap-hongkong",
      "ap-jakarta",
      "ap-nanjing",
      "ap-seoul",
      "ap-shanghai",
      "ap-singapore",
      "ap-tokyo",
      "eu-frankfurt",
      "me-saudi-arabia",
      "na-ashburn",
      "na-siliconvalley",
      "sa-saopaulo"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tb0lihs2v8sp4akmc3ss"
  },
  "imageId": "img-487zeit5",
  "imageValidation": {
    "cspResourceId": "img-487zeit5",
    "isAvailable": true,
    "resourceId": "img-487zeit5",
    "resourceName": "img-487zeit5",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "tencent",
  "regionName": "ap-seoul",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "tencent-ap-seoul",
    "costPerHour": 0.04,
    "cspSpecName": "SA2.MEDIUM4",
    "details": [
      {
        "key": "Zone",
        "value": "ap-seoul-1"
      },
      {
        "key": "InstanceType",
        "value": "SA2.MEDIUM4"
      },
      {
        "key": "InstanceChargeType",
        "value": "SPOTPAID"
      },
      {
        "key": "NetworkCard",
        "value": "25"
      },
      {
        "key": "Externals",
        "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
      },
      {
        "key": "Cpu",
        "value": "2"
      },
      {
        "key": "Memory",
        "value": "4"
      },
      {
        "key": "InstanceFamily",
        "value": "SA2"
      },
      {
        "key": "TypeName",
        "value": "SA2"
      },
      {
        "key": "Status",
        "value": "SELL"
      },
      {
        "key": "Price",
        "value": "{UnitPrice:0.04,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.008,UnitPriceSecondStep:0.04,UnitPriceDiscountSecondStep:0.008,UnitPriceThirdStep:0.04,UnitPriceDiscountThirdStep:0.008}"
      },
      {
        "key": "InstanceBandwidth",
        "value": "1.5"
      },
      {
        "key": "InstancePps",
        "value": "30"
      },
      {
        "key": "StorageBlockAmount",
        "value": "0"
      },
      {
        "key": "CpuType",
        "value": "AMD EPYC™ Rome"
      },
      {
        "key": "Gpu",
        "value": "0"
      },
      {
        "key": "Fpga",
        "value": "0"
      },
      {
        "key": "GpuCount",
        "value": "0"
      },
      {
        "key": "Frequency",
        "value": "2.6GHz/3.3GHz"
      },
      {
        "key": "StatusCategory",
        "value": "UnderStock"
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
    "id": "tencent+ap-seoul+sa2.medium4",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "tencent+ap-seoul+sa2.medium4",
    "namespace": "system",
    "providerName": "tencent",
    "regionLatitude": 37.566536,
    "regionLongitude": 126.977966,
    "regionName": "ap-seoul",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbus5den09dtgg280ngk",
    "vCPU": 2
  },
  "specId": "tencent+ap-seoul+sa2.medium4",
  "specValidation": {
    "cspResourceId": "SA2.MEDIUM4",
    "isAvailable": true,
    "resourceId": "tencent+ap-seoul+sa2.medium4",
    "resourceName": "SA2.MEDIUM4",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "ap-seoul-1"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 34.196s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.4.0.0/16",
  "connectionName": "tencent-ap-seoul",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-tencent",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.4.1.0/24",
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
  "cidrBlock": "10.4.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:31:45Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:45Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:45Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "vpc-9an88tk6",
  "cspResourceName": "tbgl3afd8uopjl21tn1c",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-tencent",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcName",
      "value": "tbgl3afd8uopjl21tn1c"
    },
    {
      "key": "VpcId",
      "value": "vpc-9an88tk6"
    },
    {
      "key": "CidrBlock",
      "value": "10.4.0.0/16"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "EnableMulticast",
      "value": "false"
    },
    {
      "key": "CreatedTime",
      "value": "2026-09-08 09:31:41"
    },
    {
      "key": "DnsServerSet",
      "value": "183.60.83.19; 183.60.82.98"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-lv0pb60n"
    },
    {
      "key": "EnableDhcp",
      "value": "true"
    },
    {
      "key": "EnableRouteVpcPublish",
      "value": "false"
    },
    {
      "key": "EnableRouteVpcPublishIpv6",
      "value": "false"
    }
  ],
  "name": "test-rdbms-vnet-tencent",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:31:45Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:31:45Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "subnet-91lwoni1",
      "cspResourceName": "tboa9phqvb1374hbn3ee",
      "cspVNetId": "vpc-9an88tk6",
      "cspVNetName": "tbgl3afd8uopjl21tn1c",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.4.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-9an88tk6"
        },
        {
          "key": "SubnetId",
          "value": "subnet-91lwoni1"
        },
        {
          "key": "SubnetName",
          "value": "tboa9phqvb1374hbn3ee"
        },
        {
          "key": "CidrBlock",
          "value": "10.4.1.0/24"
        },
        {
          "key": "IsDefault",
          "value": "false"
        },
        {
          "key": "EnableBroadcast",
          "value": "false"
        },
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "RouteTableId",
          "value": "rtb-2w78k96t"
        },
        {
          "key": "CreatedTime",
          "value": "2026-09-08 09:31:44"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "253"
        },
        {
          "key": "IsRemoteVpcSnat",
          "value": "false"
        },
        {
          "key": "TotalIpAddressCount",
          "value": "253"
        },
        {
          "key": "IsCdcSubnet",
          "value": "0"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tboa9phqvb1374hbn3ee"
    }
  ],
  "systemLabel": "",
  "uid": "tbgl3afd8uopjl21tn1c"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 13.161s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "tencent-ap-seoul",
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
  "name": "test-rdbms-sg-tencent",
  "vNetId": "test-rdbms-vnet-tencent"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "sg-qyxs45ob",
  "cspResourceName": "tbhpt13ie4tj507mv1rs",
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
  "id": "test-rdbms-sg-tencent",
  "isAutoGenerated": false,
  "keyValueList": null,
  "name": "test-rdbms-sg-tencent",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbhpt13ie4tj507mv1rs",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=tencent`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "tencent": {
      "dbOperationMethod": "cspNativeApi",
      "note": "SecurityGroup is optional (shares the VM's security group if provided). Other CSPs except AWS ignore securityGroupIds.",
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
- **Duration:** 3.131s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=tencent-ap-seoul&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-1830",
    "connectionName": "tencent-ap-seoul",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "1000",
      "2000",
      "4000",
      "8000",
      "12000",
      "16000",
      "24000",
      "32000",
      "48000",
      "64000",
      "72000",
      "96000",
      "128000",
      "144000",
      "160000",
      "192000",
      "224000",
      "244000",
      "256000",
      "288000",
      "320000",
      "384000",
      "448000",
      "488000",
      "512000",
      "690000",
      "720000"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "1024",
        "name": "1000",
        "storageSizeRangeGB": {
          "max": 3000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "2048",
        "name": "2000",
        "storageSizeRangeGB": {
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "4096",
        "name": "4000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "8000",
        "storageSizeRangeGB": {
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "12288",
        "name": "12000",
        "storageSizeRangeGB": {
          "max": 4000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "16000",
        "storageSizeRangeGB": {
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "24576",
        "name": "24000",
        "storageSizeRangeGB": {
          "max": 6000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "32000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "49152",
        "name": "48000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "65536",
        "name": "64000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "73728",
        "name": "72000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "12"
      },
      {
        "memSizeMiB": "98304",
        "name": "96000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "131072",
        "name": "128000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "147456",
        "name": "144000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "163840",
        "name": "160000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "196608",
        "name": "192000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "229376",
        "name": "224000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "112"
      },
      {
        "memSizeMiB": "249856",
        "name": "244000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "262144",
        "name": "256000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "294912",
        "name": "288000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "327680",
        "name": "320000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "393216",
        "name": "384000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "458752",
        "name": "448000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "112"
      },
      {
        "memSizeMiB": "499712",
        "name": "488000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "512000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "706560",
        "name": "690000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "737280",
        "name": "720000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "90"
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
          "constraints": "Minimum 50GB storage.",
          "description": "Cloud-based enhanced SSD with high performance. Minimum 50 GB.",
          "displayName": "Cloud Enhanced SSD",
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_HSSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Premium cloud storage with good balance. Minimum 50 GB.",
          "displayName": "Cloud Premium",
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_PREMIUM"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Standard cloud SSD storage. Minimum 50 GB.",
          "displayName": "Cloud SSD",
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_SSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "High-performance local SSD storage. Minimum 50 GB.",
          "displayName": "Local SSD",
          "minSize": 50,
          "recommendationLevel": "premium",
          "storageType": "local_ssd"
        }
      ]
    },
    "providerName": "tencent",
    "regionName": "ap-seoul",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 32000,
      "min": 20
    },
    "storageTypeOptions": [
      "CLOUD_HSSD",
      "CLOUD_PREMIUM",
      "CLOUD_SSD",
      "local_ssd"
    ],
    "supportedVersions": [
      "5.6",
      "5.7",
      "8.0",
      "8.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
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
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for tencent (ap-seoul)",
  "status": "recommended",
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-tencent",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "local_ssd",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 16ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "root",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "tencent-ap-seoul",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "4000",
  "name": "rdbms-tencent",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-tencent"
  ],
  "storageSize": 100,
  "storageType": "local_ssd",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-tencent"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "root",
    "adminUserPassword": "******",
    "connectionName": "tencent-ap-seoul",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "4000",
    "name": "rdbms-tencent",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-tencent"
    ],
    "storageSize": 100,
    "storageType": "local_ssd",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-tencent"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 4m48.323s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for tencent (ap-seoul)",
  "status": "recommended",
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-tencent",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "local_ssd",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-tencent"
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
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent`
```json
// Response Body
{
  "adminUserName": "root",
  "backupRetentionDays": 7,
  "backupTime": "00:00-12:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:36:41Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:36:41Z",
      "reason": "Available",
      "status": "True",
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
  "cspResourceId": "cdb-ctp6vws2",
  "cspResourceName": "tb7s8uhr7rhep52rtc7p",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "4000",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "endpoint": "kr-cdb-ctp6vws2.sql.tencentcdb.com:24848",
  "highAvailability": false,
  "id": "test-rdbms-tencent",
  "name": "test-rdbms-tencent",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-tencent"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "local_ssd",
  "subnetIds": [
    "subnet-1"
  ],
  "tagList": [
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.connectionName",
      "value": "tencent-ap-seoul"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tb7s8uhr7rhep52rtc7p"
    },
    {
      "key": "sys.cspResourceId",
      "value": "cdb-ctp6vws2"
    },
    {
      "key": "sys.uid",
      "value": "tb7s8uhr7rhep52rtc7p"
    },
    {
      "key": "sys.name",
      "value": "test-rdbms-tencent"
    },
    {
      "key": "sys.id",
      "value": "test-rdbms-tencent"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    }
  ],
  "uid": "tb7s8uhr7rhep52rtc7p",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 19ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:36:34Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:36:34Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Deleting",
      "storageSize": 500,
      "storageType": "cloud_essd2",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbv2od4971fhfkmsn3nt",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
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
      "endpoint": "tblunavhu1of0ntjighq.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
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
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:36:47Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:36:47Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "tbtbfiifbfn8uvco6qhr",
      "cspResourceName": "tbtbfiifbfn8uvco6qhr",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "dbInstanceType": "Burstable",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "tbtbfiifbfn8uvco6qhr.mysql.database.azure.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-azure",
      "name": "test-rdbms-azure",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "Premium_LRS",
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
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:36:41Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:36:41Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "cdb-ctp6vws2",
      "cspResourceName": "tb7s8uhr7rhep52rtc7p",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "kr-cdb-ctp6vws2.sql.tencentcdb.com:24848",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Available",
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
- **Duration:** 1.73s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database`
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
- **Duration:** 3.115s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database`
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
- **Duration:** 380ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m18.404s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 12.54s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m0.161s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 2.128s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 5.224s

