# Managed RDBMS (MySQL) Test Report: OPENSTACK (RegionOne)

- **Test Case:** OpenStack RegionOne MySQL Test
- **Date & Time:** 2026-09-08 14:26:43
- **Namespace:** `default`
- **Total Duration:** 16m52.72s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** OPENSTACK
- **Target Region:** `RegionOne`
- **Namespace:** `default`
- **Test Date:** 2026-09-08 14:26:43

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
- **Duration:** 4.271s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "78d90dae-d21d-4606-a9dd-c1268e321864",
  "specId": "openstack+regionone+m1.xlarge"
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
    "cspSpecName": "m1.xlarge",
    "details": [
      {
        "key": "ID",
        "value": "61b7f2ff-c21c-4216-b5ca-dbda45b551da"
      },
      {
        "key": "Disk",
        "value": "160"
      },
      {
        "key": "RAM",
        "value": "16384"
      },
      {
        "key": "Name",
        "value": "m1.xlarge"
      },
      {
        "key": "RxTxFactor",
        "value": "1.00"
      },
      {
        "key": "VCPUs",
        "value": "8"
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
    "diskSizeGB": 160,
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
    "id": "openstack+regionone+m1.xlarge",
    "infraType": "node",
    "memoryGiB": 16,
    "name": "openstack+regionone+m1.xlarge",
    "namespace": "system",
    "providerName": "openstack",
    "regionLatitude": 36.3804,
    "regionLongitude": 127.365,
    "regionName": "regionone",
    "rootDiskSize": 160,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbvjd3kh694j58q8qflg",
    "vCPU": 8
  },
  "specId": "openstack+regionone+m1.xlarge",
  "specValidation": {
    "cspResourceId": "m1.xlarge",
    "isAvailable": true,
    "resourceId": "openstack+regionone+m1.xlarge",
    "resourceName": "m1.xlarge",
    "status": "Available"
  },
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 21.339s
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
      "lastTransitionTime": "2026-09-08T05:26:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T05:26:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T05:26:53Z",
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
  "cspResourceId": "4d1c8b58-2bb6-4973-954c-11bd139a6b34",
  "cspResourceName": "tbkvkgq0tarc0e1505ag",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Network",
      "value": "{id:4d1c8b58-2bb6-4973-954c-11bd139a6b34,name:tbkvkgq0tarc0e1505ag,description:,admin_state_up:true,status:ACTIVE,subnets:[c1c4a150-d1f2-4fea-9c82-95ae61358465],tenant_id:1b513759820b4fb497fac0661d81c347,project_id:1b513759820b4fb497fac0661d81c347,shared:false,availability_zone_hints:[nova],tags:[],revision_number:2}"
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
          "lastTransitionTime": "2026-09-08T05:26:53Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:26:53Z",
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
      "cspResourceId": "c1c4a150-d1f2-4fea-9c82-95ae61358465",
      "cspResourceName": "tbjr6rsj27df99r5fkm0",
      "cspVNetId": "4d1c8b58-2bb6-4973-954c-11bd139a6b34",
      "cspVNetName": "tbkvkgq0tarc0e1505ag",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.6.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "c1c4a150-d1f2-4fea-9c82-95ae61358465"
        },
        {
          "key": "NetworkID",
          "value": "4d1c8b58-2bb6-4973-954c-11bd139a6b34"
        },
        {
          "key": "Name",
          "value": "tbjr6rsj27df99r5fkm0"
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
      "uid": "tbjr6rsj27df99r5fkm0"
    }
  ],
  "systemLabel": "",
  "uid": "tbkvkgq0tarc0e1505ag"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 13.539s
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
  "cspResourceId": "b189669c-3e1b-4d20-963a-061d4edccd3e",
  "cspResourceName": "tbpfhlvkmpve7l8uanl4",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "::/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    },
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
  "id": "test-rdbms-sg-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "IId",
      "value": "{NameId:tbpfhlvkmpve7l8uanl4,SystemId:b189669c-3e1b-4d20-963a-061d4edccd3e}"
    },
    {
      "key": "VpcIID",
      "value": "{NameId:,SystemId:}"
    },
    {
      "key": "SecurityRules",
      "value": "[{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:::/0},{Direction:inbound,IPProtocol:tcp,FromPort:3306,ToPort:3306,CIDR:0.0.0.0/0},{Direction:inbound,IPProtocol:tcp,FromPort:22,ToPort:22,CIDR:0.0.0.0/0},{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:0.0.0.0/0}]"
    }
  ],
  "name": "test-rdbms-sg-openstack",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbpfhlvkmpve7l8uanl4",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 52ms
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
- **Duration:** 2.605s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=openstack-regionone&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "NA",
    "connectionName": "openstack-regionone",
    "dbEngine": "mysql",
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
      "5.7.29"
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
- **Duration:** 13ms
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
        "engine": "mysql",
        "engineVersion": "8.0",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "5.7.29",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-openstack",
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
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '5.7.29' (supported versions: [5.7.29]).",
    "Storage type 'SSD' requested for instance 'Source MySQL 01' is not directly offered on openstack; recommended closest available storage type 'RBD' (available: [RBD __DEFAULT__])."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 24ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "openstack-regionone",
  "dbEngine": "mysql",
  "dbEngineVersion": "5.7.29",
  "dbInstanceSpec": "m1.medium",
  "name": "rdbms-openstack",
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
    "dbEngine": "mysql",
    "dbEngineVersion": "5.7.29",
    "dbInstanceSpec": "m1.medium",
    "name": "rdbms-openstack",
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
- **Duration:** 4m28.214s
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
      "dbEngine": "mysql",
      "dbEngineVersion": "5.7.29",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-openstack",
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
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '5.7.29' (supported versions: [5.7.29]).",
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
- **Duration:** 3.486s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-openstack`
```json
// Response Body
{
  "adminUserName": "admin",
  "backupTime": "NA",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T05:31:20Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T05:31:20Z",
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
  "cspResourceId": "c42f210d-ab82-4aa4-ac9c-38b411179db0",
  "cspResourceName": "tbuhibbvt08qje8v0hul",
  "dbEngine": "mysql",
  "dbEngineVersion": "5.7.29",
  "dbInstanceSpec": "m1.medium",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "183.111.177.136:3306",
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
  "uid": "tbuhibbvt08qje8v0hul",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 13ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "admin",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T05:27:26Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:27:26Z",
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
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
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
      "uid": "tbbl4kdjt619j4klbvjc",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "adminUserName": "admin",
      "backupTime": "NA",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T05:31:20Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:31:20Z",
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
      "cspResourceId": "c42f210d-ab82-4aa4-ac9c-38b411179db0",
      "cspResourceName": "tbuhibbvt08qje8v0hul",
      "dbEngine": "mysql",
      "dbEngineVersion": "5.7.29",
      "dbInstanceSpec": "m1.medium",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "183.111.177.136:3306",
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
      "uid": "tbuhibbvt08qje8v0hul",
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 4.538s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-openstack/database`
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
- **Duration:** 8.026s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-openstack/database`
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
- **Duration:** 75ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 8m45.029s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 39.026s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-openstack/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m51.651s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-openstack?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.704s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 7.112s

