# Managed RDBMS (MySQL) Test Report: IBM (us-south)

- **Test Case:** IBM US-South MySQL Test
- **Date & Time:** 2026-09-08 11:30:25
- **Namespace:** `default`
- **Total Duration:** 19m23.49s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** IBM
- **Target Region:** `us-south`
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
- **Duration:** 4.326s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
  "specId": "ibm+us-south+cxf-2x4"
}
```
```json
// Response Body
{
  "connectionName": "ibm-us-south",
  "estimatedCost": "$0.0850/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "ibm-us-south",
    "creationDate": "",
    "cspImageName": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "description": "",
    "details": [
      {
        "key": "AllowedUse",
        "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
      },
      {
        "key": "CatalogOffering",
        "value": "{managed:false}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-07-21T05:06:10.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:us-south:a/811f8abfbd32425597dc7ba40da98fa6::image:r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "Encryption",
        "value": "none"
      },
      {
        "key": "File",
        "value": "{checksums:{sha256:576fcfb94804e51dd910b51dba7925c078f00a2a268912ae1120af5bca05e4a1},size:2}"
      },
      {
        "key": "Href",
        "value": "https://us-south.iaas.cloud.ibm.com/v1/images/r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "ID",
        "value": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "MinimumProvisionedSize",
        "value": "10"
      },
      {
        "key": "Name",
        "value": "ibm-ubuntu-24-04-4-minimal-amd64-6"
      },
      {
        "key": "OperatingSystem",
        "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 24.04 LTS Noble Numbat Minimal Install (amd64),family:Ubuntu Linux,href:https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-24-04-amd64,name:ubuntu-24-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:24.04 LTS Noble Numbat Minimal Install}"
      },
      {
        "key": "Remote",
        "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
      },
      {
        "key": "ResourceGroup",
        "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
      },
      {
        "key": "ResourceType",
        "value": "image"
      },
      {
        "key": "Status",
        "value": "available"
      },
      {
        "key": "UserDataFormat",
        "value": "cloud_init"
      },
      {
        "key": "Visibility",
        "value": "public"
      }
    ],
    "fetchedTime": "2026.08.21 14:00:57 Fri",
    "id": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": -1,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu Linux 24.04 LTS Noble Numbat Minimal Install (amd64)",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "ibm",
    "regionList": [
      "us-south"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbfri1sv9qegh0ro7vmr"
  },
  "imageId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
  "imageValidation": {
    "cspResourceId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "isAvailable": true,
    "resourceId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "resourceName": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "ibm",
  "regionName": "us-south",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "ibm-us-south",
    "costPerHour": 0.085,
    "cspSpecName": "cxf-2x4",
    "details": [
      {
        "key": "AvailabilityClass",
        "value": "{default:standard,type:enum,values:[standard,spot]}"
      },
      {
        "key": "Bandwidth",
        "value": "{type:fixed,value:4000}"
      },
      {
        "key": "ClusterNetworkAttachmentCount",
        "value": "{type:enum,values:[0]}"
      },
      {
        "key": "ConfidentialComputeModes",
        "value": "{default:disabled,type:enum,values:[disabled]}"
      },
      {
        "key": "Family",
        "value": "compute"
      },
      {
        "key": "Href",
        "value": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/cxf-2x4"
      },
      {
        "key": "Memory",
        "value": "{type:fixed,value:4}"
      },
      {
        "key": "Name",
        "value": "cxf-2x4"
      },
      {
        "key": "NetworkAttachmentCount",
        "value": "{max:1,min:1,type:range}"
      },
      {
        "key": "NetworkBandwidthMode",
        "value": "{type:fixed,value:divided}"
      },
      {
        "key": "NetworkInterfaceCount",
        "value": "{max:1,min:1,type:range}"
      },
      {
        "key": "NumaCount",
        "value": "{type:fixed,value:1}"
      },
      {
        "key": "OsArchitecture",
        "value": "{default:amd64,type:enum,values:[amd64]}"
      },
      {
        "key": "PortSpeed",
        "value": "{type:fixed,value:25000}"
      },
      {
        "key": "ReservationTerms",
        "value": "{type:enum,values:[one_year,three_year]}"
      },
      {
        "key": "ResourceType",
        "value": "instance_profile"
      },
      {
        "key": "SecureBootModes",
        "value": "{default:false,type:enum,values:[false]}"
      },
      {
        "key": "Status",
        "value": "current"
      },
      {
        "key": "TotalVolumeBandwidth",
        "value": "{type:range,default:1000,max:3500,min:500,step:1}"
      },
      {
        "key": "VcpuArchitecture",
        "value": "{type:fixed,value:amd64}"
      },
      {
        "key": "VcpuBurstLimit",
        "value": "{type:fixed,value:200}"
      },
      {
        "key": "VcpuCount",
        "value": "{type:fixed,value:2}"
      },
      {
        "key": "VcpuManufacturer",
        "value": "{type:dependent}"
      },
      {
        "key": "VcpuPercentage",
        "value": "{default:100,type:enum,values:[25,50,100]}"
      },
      {
        "key": "VolumeBandwidthQosModes",
        "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
      },
      {
        "key": "Zones",
        "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-1,name:us-south-1}; {href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-2,name:us-south-2}; {href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-3,name:us-south-3}"
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
    "id": "ibm+us-south+cxf-2x4",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "ibm+us-south+cxf-2x4",
    "namespace": "system",
    "providerName": "ibm",
    "regionLatitude": 32.81248,
    "regionLongitude": -96.77619,
    "regionName": "us-south",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbfa0ks787p6sdeq4b30",
    "vCPU": 2
  },
  "specId": "ibm+us-south+cxf-2x4",
  "specValidation": {
    "cspResourceId": "cxf-2x4",
    "isAvailable": true,
    "resourceId": "ibm+us-south+cxf-2x4",
    "resourceName": "cxf-2x4",
    "status": "Available"
  },
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 26.214s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.5.0.0/16",
  "connectionName": "ibm-us-south",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-ibm",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.5.1.0/24",
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
  "cidrBlock": "10.5.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T02:30:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T02:30:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T02:30:53Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92",
  "cspResourceName": "tbk8ig7cmroqnv59r95o",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-ibm",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ClassicAccess",
      "value": "false"
    },
    {
      "key": "CreatedAt",
      "value": "2026-09-08T02:30:33.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92"
    },
    {
      "key": "CseSourceIps",
      "value": "{ip:{address:10.249.192.94},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}}; {ip:{address:10.12.162.235},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-2,name:us-south-2}}; {ip:{address:10.12.169.32},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-3,name:us-south-3}}"
    },
    {
      "key": "DefaultNetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,id:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,name:canyon-handpick-dairy-emu}"
    },
    {
      "key": "DefaultRoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/routing_tables/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,id:r006-212f1e92-8d31-474d-8eb1-73ebe811616e,name:custody-tingly-face-pastrami,resource_type:routing_table}"
    },
    {
      "key": "DefaultSecurityGroup",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-bbe8fcf5-7086-4933-9ef6-1d3bb6eb3152,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-bbe8fcf5-7086-4933-9ef6-1d3bb6eb3152,id:r006-bbe8fcf5-7086-4933-9ef6-1d3bb6eb3152,name:negator-aluminum-chop-mobilization}"
    },
    {
      "key": "Dns",
      "value": "{enable_hub:false,resolution_binding_count:0,resolver:{servers:[{address:161.26.0.10},{address:161.26.0.11}],type:system,configuration:default}}"
    },
    {
      "key": "HealthState",
      "value": "ok"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92"
    },
    {
      "key": "ID",
      "value": "r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92"
    },
    {
      "key": "Name",
      "value": "tbk8ig7cmroqnv59r95o"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "ResourceType",
      "value": "vpc"
    },
    {
      "key": "Status",
      "value": "available"
    },
    {
      "key": "AvailableIpv4AddressCount",
      "value": "251"
    },
    {
      "key": "CreatedAt",
      "value": "2026-09-08T02:30:47.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
    },
    {
      "key": "ID",
      "value": "0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
    },
    {
      "key": "IPVersion",
      "value": "ipv4"
    },
    {
      "key": "Ipv4CIDRBlock",
      "value": "10.5.1.0/24"
    },
    {
      "key": "Name",
      "value": "tb3f71ona5iforbi2cd7"
    },
    {
      "key": "NetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,id:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,name:canyon-handpick-dairy-emu}"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "ResourceType",
      "value": "subnet"
    },
    {
      "key": "RoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/routing_tables/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,id:r006-212f1e92-8d31-474d-8eb1-73ebe811616e,name:custody-tingly-face-pastrami,resource_type:routing_table}"
    },
    {
      "key": "Status",
      "value": "available"
    },
    {
      "key": "TotalIpv4AddressCount",
      "value": "256"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,id:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,name:tbk8ig7cmroqnv59r95o,resource_type:vpc}"
    },
    {
      "key": "Zone",
      "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
    }
  ],
  "name": "test-rdbms-vnet-ibm",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T02:30:53Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:30:53Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c",
      "cspResourceName": "tb3f71ona5iforbi2cd7",
      "cspVNetId": "r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92",
      "cspVNetName": "tbk8ig7cmroqnv59r95o",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.5.1.0/24",
      "keyValueList": [
        {
          "key": "AvailableIpv4AddressCount",
          "value": "251"
        },
        {
          "key": "CreatedAt",
          "value": "2026-09-08T02:30:47.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
        },
        {
          "key": "Href",
          "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
        },
        {
          "key": "ID",
          "value": "0717-ca9b62fd-aed9-46ef-8993-8fde04f83e2c"
        },
        {
          "key": "IPVersion",
          "value": "ipv4"
        },
        {
          "key": "Ipv4CIDRBlock",
          "value": "10.5.1.0/24"
        },
        {
          "key": "Name",
          "value": "tb3f71ona5iforbi2cd7"
        },
        {
          "key": "NetworkACL",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,id:r006-2bdbffc0-6f2e-4b1b-8352-e59f5b57b38a,name:canyon-handpick-dairy-emu}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "subnet"
        },
        {
          "key": "RoutingTable",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92/routing_tables/r006-212f1e92-8d31-474d-8eb1-73ebe811616e,id:r006-212f1e92-8d31-474d-8eb1-73ebe811616e,name:custody-tingly-face-pastrami,resource_type:routing_table}"
        },
        {
          "key": "Status",
          "value": "available"
        },
        {
          "key": "TotalIpv4AddressCount",
          "value": "256"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,id:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,name:tbk8ig7cmroqnv59r95o,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb3f71ona5iforbi2cd7",
      "zone": "us-south-1"
    }
  ],
  "systemLabel": "",
  "uid": "tbk8ig7cmroqnv59r95o"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 10.525s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "ibm-us-south",
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
  "name": "test-rdbms-sg-ibm",
  "vNetId": "test-rdbms-vnet-ibm"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1",
  "cspResourceName": "tbj7vn9va2ug95bol8tf",
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
  "id": "test-rdbms-sg-ibm",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreatedAt",
      "value": "2026-09-08T02:30:57.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1"
    },
    {
      "key": "ID",
      "value": "r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1"
    },
    {
      "key": "Name",
      "value": "tbj7vn9va2ug95bol8tf"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "Rules",
      "value": "{direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1/rules/r006-fc367837-e2df-41ac-bf6a-65382b7821bc,id:r006-fc367837-e2df-41ac-bf6a-65382b7821bc,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:mostly-unwashed-culture-steadier,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:3306,port_min:3306,protocol:tcp}; {direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1/rules/r006-7a6b9678-4c52-46ce-9f89-9eb83a940a0a,id:r006-7a6b9678-4c52-46ce-9f89-9eb83a940a0a,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:tamale-freezable-unbridle-pluck,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:22,port_min:22,protocol:tcp}; {direction:outbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-b21dae59-7eb9-44f2-b0b3-4d4fadefb5b1/rules/r006-53684a2a-db19-45be-bfff-274ec29e88b7,id:r006-53684a2a-db19-45be-bfff-274ec29e88b7,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:compactor-lurch-cure-unfairly,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,protocol:any}"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,id:r006-8f2a5d5d-c7c0-4901-834b-7627724a8a92,name:tbk8ig7cmroqnv59r95o,resource_type:vpc}"
    }
  ],
  "name": "test-rdbms-sg-ibm",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbj7vn9va2ug95bol8tf",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=ibm`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "ibm": {
      "dbOperationMethod": "sqlFallback",
      "note": "Storage type selection not supported. IBM Cloud Databases manages storage automatically.",
      "storageTypeSelectable": false,
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
- **Duration:** 8.025s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ibm-us-south&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "NA",
    "connectionName": "ibm-us-south",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "b3c.16x64.encrypted",
      "b3c.32x128.encrypted",
      "b3c.4x16.encrypted",
      "b3c.8x32.encrypted",
      "m3c.30x240.encrypted",
      "m3c.8x64.encrypted",
      "multitenant"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "65536",
        "name": "b3c.16x64.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "b3c.32x128.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "b3c.4x16.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "b3c.8x32.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "245760",
        "name": "m3c.30x240.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "30"
      },
      {
        "memSizeMiB": "65536",
        "name": "m3c.8x64.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "12288",
        "name": "multitenant",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "0"
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
          "description": "Storage is managed automatically by IBM Cloud Databases. User cannot specify storage type.",
          "displayName": "Automatic (IBM-managed)",
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ibm",
    "regionName": "us-south",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
      "max": 13194,
      "min": 32
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "8.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": true,
    "supportsHighAvailability": true,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": false,
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
    "csp": "ibm",
    "region": "us-south"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ibm (us-south)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ibm",
    "region": "us-south"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "b3c.4x16.encrypted",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-ibm",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ],
  "warnings": [
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '8.4' (supported versions: [8.4]).",
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 15ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "ibm-us-south",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.4",
  "dbInstanceSpec": "b3c.4x16.encrypted",
  "name": "rdbms-ibm",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-ibm"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-ibm"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "ibm-us-south",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.4",
    "dbInstanceSpec": "b3c.4x16.encrypted",
    "name": "rdbms-ibm",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-ibm"
    ],
    "storageSize": 100,
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-ibm"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 11m51.256s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ibm (us-south)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ibm",
    "region": "us-south"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "b3c.4x16.encrypted",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-ibm",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ],
  "warnings": [
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '8.4' (supported versions: [8.4]).",
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider."
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
- **Duration:** 5.533s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm`
```json
// Response Body
{
  "adminUserName": "admin",
  "backupRetentionDays": 30,
  "backupTime": "AUTO",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T02:42:00Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T02:42:00Z",
      "reason": "Available",
      "status": "True",
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
  "cspResourceId": "4dc8184d-4e69-4bbc-b67a-dc69cb310d32",
  "cspResourceName": "tblv8isdiqhulslqdbtk",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.4",
  "dbInstanceSpec": "b3c.4x16.encrypted",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "encryption": true,
  "endpoint": "4dc8184d-4e69-4bbc-b67a-dc69cb310d32.c9v3nfod0e3fgcbd1oug.databases.appdomain.cloud:32140",
  "highAvailability": false,
  "id": "test-rdbms-ibm",
  "name": "test-rdbms-ibm",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-ibm"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "NA",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tblv8isdiqhulslqdbtk",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 12ms
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
          "lastTransitionTime": "2026-09-08T02:41:55Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
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
      "backupRetentionDays": 30,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T02:42:00Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T02:42:00Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "4dc8184d-4e69-4bbc-b67a-dc69cb310d32",
      "cspResourceName": "tblv8isdiqhulslqdbtk",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "b3c.4x16.encrypted",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "encryption": true,
      "endpoint": "4dc8184d-4e69-4bbc-b67a-dc69cb310d32.c9v3nfod0e3fgcbd1oug.databases.appdomain.cloud:32140",
      "highAvailability": false,
      "id": "test-rdbms-ibm",
      "name": "test-rdbms-ibm",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "NA",
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
          "lastTransitionTime": "2026-09-08T02:41:37Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
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
- **Duration:** 6.535s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database`
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
- **Duration:** 10.954s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database`
```json
// Response Body
{
  "databases": [
    "ibmclouddb",
    "information_schema",
    "meta",
    "mysql",
    "performance_schema",
    "sampledb",
    "sampledb_dyn",
    "sys"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 322ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m45.268s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 47.139s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 26.503s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 4.21s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 16.643s

