# Managed RDBMS (MySQL) Test Report: NCP (kr)

- **Test Case:** NCP Korea MySQL Test
- **Date & Time:** 2026-09-08 10:31:39
- **Namespace:** `default`
- **Total Duration:** 30m57.089s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NCP
- **Target Region:** `kr`
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
- **Duration:** 19.388s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "23214590",
  "specId": "s2-g3"
}
```
```json
// Response Body
{
  "connectionName": "ncp-jpn",
  "errors": [
    "Image '23214590' not available: image '23214590' (CSP id: 23214590) is registered for region(s) [kr] but connection 'ncp-jpn' targets region 'jpn'; pick an image registered for region 'jpn'"
  ],
  "estimatedCost": "$0.0922/hour",
  "imageId": "23214590",
  "imageValidation": {
    "cspResourceId": "23214590",
    "isAvailable": false,
    "message": "image '23214590' (CSP id: 23214590) is registered for region(s) [kr] but connection 'ncp-jpn' targets region 'jpn'; pick an image registered for region 'jpn'",
    "resourceId": "23214590",
    "status": "Unavailable"
  },
  "isValid": false,
  "message": "Image '23214590' is not available",
  "providerName": "ncp",
  "regionName": "jpn",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "ncp-jpn",
    "costPerHour": 0.0922,
    "cspSpecName": "s2-g3",
    "details": [
      {
        "key": "ServerSpecCode",
        "value": "s2-g3"
      },
      {
        "key": "GenerationCode",
        "value": "G3"
      },
      {
        "key": "CpuCount",
        "value": "2"
      },
      {
        "key": "MemorySize",
        "value": "8589934592"
      },
      {
        "key": "HypervisorType",
        "value": "{code:KVM,codeName:KVM}"
      },
      {
        "key": "CpuArchitectureType",
        "value": "{code:X86_64,codeName:x86 64bit}"
      },
      {
        "key": "BlockStorageMaxCount",
        "value": "20"
      },
      {
        "key": "BlockStorageMaxIops",
        "value": "4725"
      },
      {
        "key": "BlockStorageMaxThroughput",
        "value": "84934656"
      },
      {
        "key": "NetworkPerformance",
        "value": "1000000000"
      },
      {
        "key": "NetworkInterfaceMaxCount",
        "value": "3"
      },
      {
        "key": "ServerProductCode",
        "value": "SVR.VSVR.STAND.C002.M008.G003"
      },
      {
        "key": "ServerSpecDescription",
        "value": "vCPU 2EA, Memory 8GB"
      },
      {
        "key": "ServerSpecNo",
        "value": "283"
      },
      {
        "key": "CorrespondingImageIds",
        "value": "142902717,108645046,106703137,104027588,26905053,24075285,22224744,22224742,22224738,22224737"
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
    "id": "ncp+jpn+s2-g3",
    "infraType": "node",
    "memoryGiB": 8,
    "name": "ncp+jpn+s2-g3",
    "namespace": "system",
    "providerName": "ncp",
    "regionLatitude": 35.51742,
    "regionLongitude": 136.80291,
    "regionName": "jpn",
    "rootDiskSize": 0,
    "rootDiskType": "default",
    "systemLabel": "from-assets",
    "uid": "tb0tuig9ihijn201tn1c",
    "vCPU": 2
  },
  "specId": "s2-g3",
  "specValidation": {
    "cspResourceId": "s2-g3",
    "isAvailable": true,
    "resourceId": "s2-g3",
    "resourceName": "s2-g3",
    "status": "Available"
  },
  "status": "Error"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 31.858s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.7.0.0/16",
  "connectionName": "ncp-kr",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-ncp",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.7.1.0/24",
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
  "cidrBlock": "10.7.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:32:25Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:32:25Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:32:25Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "147648",
  "cspResourceName": "tb6m2rlplp263mgt6o0p",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-ncp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcNo",
      "value": "147648"
    },
    {
      "key": "VpcName",
      "value": "tb6m2rlplp263mgt6o0p"
    },
    {
      "key": "Ipv4CidrBlock",
      "value": "10.7.0.0/16"
    },
    {
      "key": "VpcStatus",
      "value": "{code:RUN,codeName:운영중}"
    },
    {
      "key": "RegionCode",
      "value": "KR"
    },
    {
      "key": "CreateDate",
      "value": "2026-09-08T10:31:56+0900"
    }
  ],
  "name": "test-rdbms-vnet-ncp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:32:25Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:32:25Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "321884",
      "cspResourceName": "tb5gm53s8ml4j8ogj9m6",
      "cspVNetId": "147648",
      "cspVNetName": "tb6m2rlplp263mgt6o0p",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.7.1.0/24",
      "keyValueList": [
        {
          "key": "SubnetNo",
          "value": "321884"
        },
        {
          "key": "VpcNo",
          "value": "147648"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "SubnetName",
          "value": "tb5gm53s8ml4j8ogj9m6"
        },
        {
          "key": "Subnet",
          "value": "10.7.1.0/24"
        },
        {
          "key": "SubnetStatus",
          "value": "{code:RUN,codeName:운영중}"
        },
        {
          "key": "CreateDate",
          "value": "2026-09-08T10:32:12+0900"
        },
        {
          "key": "SubnetType",
          "value": "{code:PUBLIC,codeName:Public}"
        },
        {
          "key": "UsageType",
          "value": "{code:GEN,codeName:General}"
        },
        {
          "key": "NetworkAclNo",
          "value": "198576"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb5gm53s8ml4j8ogj9m6",
      "zone": "KR-1"
    }
  ],
  "systemLabel": "",
  "uid": "tb6m2rlplp263mgt6o0p"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 25.003s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "ncp-kr",
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
  "name": "test-rdbms-sg-ncp",
  "vNetId": "test-rdbms-vnet-ncp"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "393359",
  "cspResourceName": "tbkm4bk21dnepd6q81hi",
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
      "Protocol": "ICMP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "1-65535",
      "Protocol": "UDP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "1-65535",
      "Protocol": "TCP"
    }
  ],
  "id": "test-rdbms-sg-ncp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "AccessControlGroupNo",
      "value": "393359"
    },
    {
      "key": "AccessControlGroupName",
      "value": "tbkm4bk21dnepd6q81hi"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "VpcNo",
      "value": "147648"
    },
    {
      "key": "AccessControlGroupStatus",
      "value": "{code:RUN,codeName:운영중}"
    }
  ],
  "name": "test-rdbms-sg-ncp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbkm4bk21dnepd6q81hi",
  "vNetId": "test-rdbms-vnet-ncp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 11ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=ncp`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "ncp": {
      "dbOperationMethod": "cspNativeApi",
      "note": "Storage type and storage size configuration not supported. NCP G3 generation automatically applies SSD storage, starts at 10GB, and auto-scales by 10GB increments up to 6000GB.",
      "storageTypeSelectable": false,
      "supported": true,
      "supportedDBEngines": [
        "mysql"
      ]
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 2.933s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ncp-kr&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "connectionName": "ncp-kr",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C002.M008.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C004.M008.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C004.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C008.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C002.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C008.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C016.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C004.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C016.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C032.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C008.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C048.M096.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C032.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C064.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C016.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C048.M192.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C064.M256.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C032.M256.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C048.M384.NET.SSD.B050.G003"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "4096",
        "name": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "SVR.VDBAS.AMD.HICPU.C004.M008.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.HICPU.C008.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.HICPU.C016.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.HICPU.C032.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "98304",
        "name": "SVR.VDBAS.AMD.HICPU.C048.M096.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.HICPU.C064.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.HIMEM.C002.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.HIMEM.C004.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.HIMEM.C008.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.HIMEM.C016.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "SVR.VDBAS.AMD.HIMEM.C032.M256.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "393216",
        "name": "SVR.VDBAS.AMD.HIMEM.C048.M384.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "8192",
        "name": "SVR.VDBAS.AMD.STAND.C002.M008.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.STAND.C004.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.STAND.C008.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.STAND.C016.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.STAND.C032.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "196608",
        "name": "SVR.VDBAS.AMD.STAND.C048.M192.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "SVR.VDBAS.AMD.STAND.C064.M256.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
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
          "constraints": "Minimum 10GB storage.",
          "description": "NCP G3 automatically applies SSD storage (starts at 10GB and auto-scales by 10GB increments up to 6000GB). Storage size and type are not user-configurable.",
          "displayName": "SSD (automatic)",
          "maxSize": 6000,
          "minSize": 10,
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ncp",
    "regionName": "KR",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 6000,
      "min": 10
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "8.4.8",
      "8.4.6",
      "8.0.45",
      "8.0.42",
      "8.0.40",
      "8.0.36"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": false,
    "supportsHighAvailability": true,
    "supportsPublicAccess": false,
    "supportsStorageSizeConfiguration": false,
    "supportsStorageTypeSelection": false,
    "supportsTag": false
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
    "csp": "ncp",
    "region": "kr"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ncp (kr)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
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
      "dbEngineVersion": "8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "highAvailability": false,
      "publicAccess": false,
      "rdbmsName": "rdbms-ncp",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ncp"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (ncp); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider.",
    "NCP Cloud DB does not provide external public IP by default; instance 'source-mysql-01' will be created within private VPC."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 14ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "ncp-kr",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.45",
  "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
  "name": "rdbms-ncp",
  "securityGroupIds": [
    "test-rdbms-sg-ncp"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-ncp"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "ncp-kr",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0.45",
    "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
    "name": "rdbms-ncp",
    "securityGroupIds": [
      "test-rdbms-sg-ncp"
    ],
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-ncp"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 12m46.043s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ncp (kr)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
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
      "dbEngineVersion": "8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "highAvailability": false,
      "publicAccess": false,
      "rdbmsName": "rdbms-ncp",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ncp"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (ncp); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider.",
    "NCP Cloud DB does not provide external public IP by default; instance 'source-mysql-01' will be created within private VPC."
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
- **Duration:** 1.384s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp`
```json
// Response Body
{
  "adminUserName": "dbadmin",
  "backupRetentionDays": 7,
  "backupTime": "06:30",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:44:29Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:44:29Z",
      "reason": "Available",
      "status": "True",
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
  "endpoint": "db-4acagh.vpc-cdb.ntruss.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-ncp",
  "name": "test-rdbms-ncp",
  "publicAccess": false,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-ncp"
  ],
  "status": "Available",
  "storageSize": 10,
  "storageType": "SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tb7k7td81l8pt2p0nhml",
  "vNetId": "test-rdbms-vnet-ncp"
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
          "lastTransitionTime": "2026-09-08T01:44:00Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:40:15Z",
          "reason": "Available",
          "status": "True",
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
      "status": "Deleting",
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
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "backupTime": "06:30",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:44:29Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:44:29Z",
          "reason": "Available",
          "status": "True",
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
      "endpoint": "db-4acagh.vpc-cdb.ntruss.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-ncp",
      "name": "test-rdbms-ncp",
      "publicAccess": false,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "status": "Available",
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
      "backupTime": "03:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:41:12Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:41:12Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "74046c69-3e7c-4cb4-9078-b78b76ce6824",
      "cspResourceName": "tbb3k63114huhljjepms",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "endpoint": "1d3a2c9f-9a78-49a7-b26e-5be59cd9c05d.external.kr1.mysql.rds.nhncloudservice.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-nhn",
      "name": "test-rdbms-nhn",
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbb3k63114huhljjepms",
      "vNetId": "test-rdbms-vnet-nhn"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 8.826s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database`
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
- **Duration:** 4.684s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database`
```json
// Response Body
{
  "databases": [
    "mydb",
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Data I/O Test (External Remote) [⚪ SKIPPED (Skipped: Public Access is not supported/configured for this CSP (VPC private endpoint only))]
- **Duration:** 0s

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 10m57.914s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 30.11s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 4m37.256s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 4.71s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 26.944s

