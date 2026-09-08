# Managed RDBMS (MariaDB) Test Report: OPENSTACK (RegionOne)

- **Test Case:** OpenStack RegionOne MariaDB Test
- **Date & Time:** 2026-09-08 15:43:26
- **Namespace:** `default`
- **Total Duration:** 14m49.646s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** OPENSTACK
- **Target Region:** `RegionOne`
- **Namespace:** `default`
- **Test Date:** 2026-09-08 15:43:26

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
- **Duration:** 2.83s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "78d90dae-d21d-4606-a9dd-c1268e321864",
  "specId": "openstack+regionone+m1.large"
}
```
```json
// Response Body
{
  "connectionName": "openstack-regionone",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "openstack-regionone",
    "creationDate": "",
    "cspImageName": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "description": "",
    "details": [
      {
        "key": "ID",
        "value": "78d90dae-d21d-4606-a9dd-c1268e321864"
      },
      {
        "key": "Name",
        "value": "ubuntu-24.04"
      },
      {
        "key": "Status",
        "value": "active"
      },
      {
        "key": "ContainerFormat",
        "value": "bare"
      },
      {
        "key": "DiskFormat",
        "value": "qcow2"
      },
      {
        "key": "MinDiskGigabytes",
        "value": "0"
      },
      {
        "key": "MinRAMMegabytes",
        "value": "0"
      },
      {
        "key": "Owner",
        "value": "1b513759820b4fb497fac0661d81c347"
      },
      {
        "key": "Protected",
        "value": "false"
      },
      {
        "key": "Visibility",
        "value": "shared"
      },
      {
        "key": "Hidden",
        "value": "false"
      },
      {
        "key": "Checksum",
        "value": "77a77cff0d62968d5166322ea2b78008"
      },
      {
        "key": "CreatedAt",
        "value": "2025-09-16T05:28:43Z"
      },
      {
        "key": "UpdatedAt",
        "value": "2025-09-16T05:28:55Z"
      },
      {
        "key": "File",
        "value": "/v2/images/78d90dae-d21d-4606-a9dd-c1268e321864/file"
      },
      {
        "key": "Schema",
        "value": "/v2/schemas/image"
      },
      {
        "key": "VirtualSize",
        "value": "3758096384"
      }
    ],
    "fetchedTime": "2026.08.21 13:57:14 Fri",
    "id": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": false,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "namespace": "system",
    "osArchitecture": "na",
    "osDiskSizeGB": 0,
    "osDiskType": "NA",
    "osDistribution": "ubuntu-24.04",
    "osPlatform": "NA",
    "osType": "Ubuntu 24.04",
    "providerName": "openstack",
    "regionList": [
      "regionone"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbglac0fmf7lvmmuu60t"
  },
  "imageId": "78d90dae-d21d-4606-a9dd-c1268e321864",
  "imageValidation": {
    "cspResourceId": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "isAvailable": true,
    "resourceId": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "resourceName": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "openstack",
  "regionName": "regionone",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "openstack-regionone",
    "costPerHour": -1,
    "cspSpecName": "m1.large",
    "details": [
      {
        "key": "ID",
        "value": "d0e47e56-adbc-414c-932c-d734f6fb4acb"
      },
      {
        "key": "Disk",
        "value": "80"
      },
      {
        "key": "RAM",
        "value": "8192"
      },
      {
        "key": "Name",
        "value": "m1.large"
      },
      {
        "key": "RxTxFactor",
        "value": "1.00"
      },
      {
        "key": "VCPUs",
        "value": "4"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "Ephemeral",
        "value": "0"
      }
    ],
    "diskSizeGB": 80,
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
    "id": "openstack+regionone+m1.large",
    "infraType": "node",
    "memoryGiB": 8,
    "name": "openstack+regionone+m1.large",
    "namespace": "system",
    "providerName": "openstack",
    "regionLatitude": 36.3804,
    "regionLongitude": 127.365,
    "regionName": "regionone",
    "rootDiskSize": 80,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tb6if7bvagid4j8na7t1",
    "vCPU": 4
  },
  "specId": "openstack+regionone+m1.large",
  "specValidation": {
    "cspResourceId": "m1.large",
    "isAvailable": true,
    "resourceId": "openstack+regionone+m1.large",
    "resourceName": "m1.large",
    "status": "Available"
  },
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 35.633s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.6.0.0/16",
  "connectionName": "openstack-regionone",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-openstack",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.6.1.0/24",
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
  "cidrBlock": "OpenStack VPC does not support IPv4_CIDR",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T06:43:33Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T06:43:33Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T06:43:33Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "f151966d-2e0f-40b8-9af4-668c0c83e688",
  "cspResourceName": "tbl4dca47i5buvbsnga2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Network",
      "value": "{id:f151966d-2e0f-40b8-9af4-668c0c83e688,name:tbl4dca47i5buvbsnga2,description:,admin_state_up:true,status:ACTIVE,subnets:[67d6f411-2c88-400a-9bfd-614912f43141],tenant_id:1b513759820b4fb497fac0661d81c347,project_id:1b513759820b4fb497fac0661d81c347,shared:false,availability_zone_hints:[nova],tags:[],revision_number:2}"
    },
    {
      "key": "NetworkExternalExt",
      "value": "{router:external:false}"
    }
  ],
  "name": "test-rdbms-vnet-openstack",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T06:43:33Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T06:43:33Z",
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
      "cspResourceId": "67d6f411-2c88-400a-9bfd-614912f43141",
      "cspResourceName": "tb7jh1hfaaomjs07t0jj",
      "cspVNetId": "f151966d-2e0f-40b8-9af4-668c0c83e688",
      "cspVNetName": "tbl4dca47i5buvbsnga2",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.6.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "67d6f411-2c88-400a-9bfd-614912f43141"
        },
        {
          "key": "NetworkID",
          "value": "f151966d-2e0f-40b8-9af4-668c0c83e688"
        },
        {
          "key": "Name",
          "value": "tb7jh1hfaaomjs07t0jj"
        },
        {
          "key": "IPVersion",
          "value": "4"
        },
        {
          "key": "CIDR",
          "value": "10.6.1.0/24"
        },
        {
          "key": "GatewayIP",
          "value": "10.6.1.1"
        },
        {
          "key": "DNSNameservers",
          "value": "8.8.8.8"
        },
        {
          "key": "DNSPublishFixedIP",
          "value": "false"
        },
        {
          "key": "AllocationPools",
          "value": "{start:10.6.1.2,end:10.6.1.254}"
        },
        {
          "key": "EnableDHCP",
          "value": "true"
        },
        {
          "key": "TenantID",
          "value": "1b513759820b4fb497fac0661d81c347"
        },
        {
          "key": "ProjectID",
          "value": "1b513759820b4fb497fac0661d81c347"
        },
        {
          "key": "RevisionNumber",
          "value": "0"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb7jh1hfaaomjs07t0jj"
    }
  ],
  "systemLabel": "",
  "uid": "tbl4dca47i5buvbsnga2"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 16.761s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "openstack-regionone",
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
  "name": "test-rdbms-sg-openstack",
  "vNetId": "test-rdbms-vnet-openstack"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "b0865764-3cb9-496a-b7bf-afa70fcb111d",
  "cspResourceName": "tbvivtovu0gt2mijf37n",
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
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    },
    {
      "CIDR": "::/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    }
  ],
  "id": "test-rdbms-sg-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "IId",
      "value": "{NameId:tbvivtovu0gt2mijf37n,SystemId:b0865764-3cb9-496a-b7bf-afa70fcb111d}"
    },
    {
      "key": "VpcIID",
      "value": "{NameId:,SystemId:}"
    },
    {
      "key": "SecurityRules",
      "value": "[{Direction:inbound,IPProtocol:tcp,FromPort:3306,ToPort:3306,CIDR:0.0.0.0/0},{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:0.0.0.0/0},{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:::/0},{Direction:inbound,IPProtocol:tcp,FromPort:22,ToPort:22,CIDR:0.0.0.0/0}]"
    }
  ],
  "name": "test-rdbms-sg-openstack",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbvivtovu0gt2mijf37n",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=openstack`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "openstack": {
      "dbOperationMethod": "cspNativeApi",
      "note": "An earlier v0.12.44 test run requested MariaDB 10.6 and failed ('Datastore version 10.6 cannot be found'). CB-Spider's rdbms-mariadb-test README clarifies the correct reference version for this Trove environment is 10.4, not 10.6 — confirmed Available in a 2026-08-18 run (m1.small/20GB, 3m58s). The datastore was never actually missing; 10.6 was simply the wrong version to request.",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql",
        "mariadb"
      ]
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 4.402s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=openstack-regionone&dbEngine=mariadb`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "NA",
    "connectionName": "openstack-regionone",
    "dbEngine": "mariadb",
    "dbInstanceSpecOptions": [
      "m1.amphora",
      "m1.large",
      "m1.medium",
      "m1.small",
      "m1.tiny",
      "m1.xlarge"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "1024",
        "name": "m1.amphora",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "8192",
        "name": "m1.large",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "4096",
        "name": "m1.medium",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "m1.small",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "512",
        "name": "m1.tiny",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "16384",
        "name": "m1.xlarge",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
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
          "constraints": "Minimum 1GB storage.",
          "description": "Ceph RADOS Block Device storage (if available in deployment).",
          "displayName": "Ceph RBD",
          "minSize": 1,
          "recommendationLevel": "standard",
          "storageType": "RBD"
        },
        {
          "constraints": "Minimum 1GB storage.",
          "description": "Default storage type provided by OpenStack deployment. Actual backend depends on deployment configuration.",
          "displayName": "Default Storage",
          "minSize": 1,
          "recommendationLevel": "standard",
          "storageType": "__DEFAULT__"
        }
      ]
    },
    "providerName": "openstack",
    "regionName": "RegionOne",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
      "max": -1,
      "min": 1
    },
    "storageTypeOptions": [
      "RBD",
      "__DEFAULT__"
    ],
    "supportedVersions": [
      "10.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
    "supportsEncryption": false,
    "supportsHighAvailability": false,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": true,
    "supportsTag": false
  }
}
```

### 6. Beetle POST Recommend RDBMS [✅ SUCCESS]
- **Duration:** 12ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "openstack",
    "region": "RegionOne"
  },
  "sourceRDBMSInstances": [
    {
      "dbEngine": {
        "engine": "mariadb",
        "engineVersion": "10.4",
        "port": 3306,
        "role": "primary"
      },
      "dbNode": {
        "cpu": {
          "cores": 2,
          "cpus": 1,
          "maxSpeed": 2.4,
          "threads": 2
        },
        "dataDisks": [
          {
            "label": "/data",
            "totalSize": 50,
            "type": "SSD"
          }
        ],
        "hostname": "db-server-01",
        "machineId": "node-550e8400-e29b-41d4-a716-446655440000",
        "memory": {
          "totalSize": 4
        },
        "rootDisk": {
          "label": "/",
          "totalSize": 50,
          "type": "SSD"
        }
      },
      "description": "Production database instance discovered from on-premise node",
      "displayName": "Source MySQL 01",
      "innerDatabases": [
        {
          "characterSet": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ]
    }
  ],
  "targetPreferences": {
    "adminUserName": "admin",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for openstack (regionone)",
  "status": "recommended",
  "targetCloud": {
    "csp": "openstack",
    "region": "regionone"
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
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.4",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-openstack",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ],
  "warnings": [
    "Storage type 'SSD' requested for instance 'Source MySQL 01' is not directly offered on openstack; recommended closest available storage type 'RBD' (available: [RBD __DEFAULT__])."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 28ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "openstack-regionone",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.4",
  "dbInstanceSpec": "m1.medium",
  "name": "rdbms-mariadb-openstack",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-openstack"
  ],
  "storageSize": 100,
  "storageType": "RBD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-openstack"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "openstack-regionone",
    "dbEngine": "mariadb",
    "dbEngineVersion": "10.4",
    "dbInstanceSpec": "m1.medium",
    "name": "rdbms-mariadb-openstack",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-openstack"
    ],
    "storageSize": 100,
    "storageType": "RBD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-openstack"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 4m22.061s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for openstack (regionone)",
  "status": "recommended",
  "targetCloud": {
    "csp": "openstack",
    "region": "regionone"
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
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.4",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-openstack",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ],
  "warnings": [
    "Storage type 'SSD' requested for instance 'Source MySQL 01' is not directly offered on openstack; recommended closest available storage type 'RBD' (available: [RBD __DEFAULT__])."
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
- **Duration:** 4.211s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack`
```json
// Response Body
{
  "adminUserName": "admin",
  "backupTime": "NA",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T06:48:13Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T06:48:13Z",
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
  "cspResourceId": "d2328034-37e8-44e8-8a45-aa85c7e74c8b",
  "cspResourceName": "tbnvbqr3580esmlmfp3q",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.4",
  "dbInstanceSpec": "m1.medium",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "183.111.177.148:3306",
  "highAvailability": false,
  "id": "test-rdbms-mariadb-openstack",
  "name": "test-rdbms-mariadb-openstack",
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
  "uid": "tbnvbqr3580esmlmfp3q",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "backupTime": "17:00Z-18:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T06:47:12Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T06:47:12Z",
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
      "cspResourceId": "rm-mj7on13jbb97vg9bq",
      "cspResourceName": "tbeo4ttttknb54tko9dp",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.51:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
      "name": "test-rdbms-mariadb-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "cloud_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbeo4ttttknb54tko9dp",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "adminUserName": "root",
      "backupRetentionDays": 7,
      "backupTime": "15:59-16:29",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T06:45:07Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T06:45:07Z",
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
      "cspResourceId": "tb73s3vujjccjaid9hl8",
      "cspResourceName": "tb73s3vujjccjaid9hl8",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.27",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
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
          "value": "tb73s3vujjccjaid9hl8"
        }
      ],
      "uid": "tb73s3vujjccjaid9hl8",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "adminUserName": "myadmin",
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T06:44:50Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T06:44:50Z",
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
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
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
      "uid": "tb5orqrpujtia47vn84d",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "adminUserName": "admin",
      "backupTime": "NA",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T06:48:13Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T06:48:13Z",
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
      "cspResourceId": "d2328034-37e8-44e8-8a45-aa85c7e74c8b",
      "cspResourceName": "tbnvbqr3580esmlmfp3q",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.4",
      "dbInstanceSpec": "m1.medium",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "183.111.177.148:3306",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-openstack",
      "name": "test-rdbms-mariadb-openstack",
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
      "uid": "tbnvbqr3580esmlmfp3q",
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 5.678s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database`
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
- **Duration:** 9.344s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database`
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
- **Duration:** 74ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 6m1.084s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 44.823s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m55.506s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 12.083s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 15.106s

