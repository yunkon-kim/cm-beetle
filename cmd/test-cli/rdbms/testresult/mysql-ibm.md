# Managed RDBMS (MySQL) Test Report: IBM (us-south)

- **Test Case:** IBM US-South MySQL Test
- **Date & Time:** 2026-09-08 14:26:43
- **Namespace:** `default`
- **Total Duration:** 25m19.054s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** IBM
- **Target Region:** `us-south`
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
- **Duration:** 5.548s
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
- **Duration:** 22.752s
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
      "lastTransitionTime": "2026-09-08T05:27:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T05:27:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T05:27:11Z",
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
  "cspResourceId": "r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375",
  "cspResourceName": "tb4u3bq4enj35se6qmjp",
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
      "value": "2026-09-08T05:26:53.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375"
    },
    {
      "key": "CseSourceIps",
      "value": "{ip:{address:10.22.219.126},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}}; {ip:{address:10.22.227.78},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-2,name:us-south-2}}; {ip:{address:10.12.170.128},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-3,name:us-south-3}}"
    },
    {
      "key": "DefaultNetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,id:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,name:deforest-cylinder-arrive-aspire}"
    },
    {
      "key": "DefaultRoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/r006-4e9ea780-3bd5-4300-8772-ad3936973775,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/routing_tables/r006-4e9ea780-3bd5-4300-8772-ad3936973775,id:r006-4e9ea780-3bd5-4300-8772-ad3936973775,name:rekindle-speak-enigmatic-stir,resource_type:routing_table}"
    },
    {
      "key": "DefaultSecurityGroup",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-e29fcff8-cc11-405d-812d-b37f8833c89c,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e29fcff8-cc11-405d-812d-b37f8833c89c,id:r006-e29fcff8-cc11-405d-812d-b37f8833c89c,name:tamer-simple-snowbird-rewire}"
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
      "value": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375"
    },
    {
      "key": "ID",
      "value": "r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375"
    },
    {
      "key": "Name",
      "value": "tb4u3bq4enj35se6qmjp"
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
      "value": "2026-09-08T05:27:03.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-f211ebed-8639-469b-b555-0725f00f0225"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-f211ebed-8639-469b-b555-0725f00f0225"
    },
    {
      "key": "ID",
      "value": "0717-f211ebed-8639-469b-b555-0725f00f0225"
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
      "value": "tbi0bha8lil4pk9s433s"
    },
    {
      "key": "NetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,id:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,name:deforest-cylinder-arrive-aspire}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/r006-4e9ea780-3bd5-4300-8772-ad3936973775,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/routing_tables/r006-4e9ea780-3bd5-4300-8772-ad3936973775,id:r006-4e9ea780-3bd5-4300-8772-ad3936973775,name:rekindle-speak-enigmatic-stir,resource_type:routing_table}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,id:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,name:tb4u3bq4enj35se6qmjp,resource_type:vpc}"
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
          "lastTransitionTime": "2026-09-08T05:27:11Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:27:11Z",
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
      "cspResourceId": "0717-f211ebed-8639-469b-b555-0725f00f0225",
      "cspResourceName": "tbi0bha8lil4pk9s433s",
      "cspVNetId": "r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375",
      "cspVNetName": "tb4u3bq4enj35se6qmjp",
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
          "value": "2026-09-08T05:27:03.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-f211ebed-8639-469b-b555-0725f00f0225"
        },
        {
          "key": "Href",
          "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-f211ebed-8639-469b-b555-0725f00f0225"
        },
        {
          "key": "ID",
          "value": "0717-f211ebed-8639-469b-b555-0725f00f0225"
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
          "value": "tbi0bha8lil4pk9s433s"
        },
        {
          "key": "NetworkACL",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,id:r006-8bd83e84-e5b4-4cf6-9a1b-06ac8f118689,name:deforest-cylinder-arrive-aspire}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/r006-4e9ea780-3bd5-4300-8772-ad3936973775,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375/routing_tables/r006-4e9ea780-3bd5-4300-8772-ad3936973775,id:r006-4e9ea780-3bd5-4300-8772-ad3936973775,name:rekindle-speak-enigmatic-stir,resource_type:routing_table}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,id:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,name:tb4u3bq4enj35se6qmjp,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbi0bha8lil4pk9s433s",
      "zone": "us-south-1"
    }
  ],
  "systemLabel": "",
  "uid": "tb4u3bq4enj35se6qmjp"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 8.131s
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
  "cspResourceId": "r006-45a38175-1603-4e8b-b635-67a5b830755f",
  "cspResourceName": "tbqvt6hla7fm31gpfrsk",
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
      "value": "2026-09-08T05:27:13.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-45a38175-1603-4e8b-b635-67a5b830755f"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-45a38175-1603-4e8b-b635-67a5b830755f"
    },
    {
      "key": "ID",
      "value": "r006-45a38175-1603-4e8b-b635-67a5b830755f"
    },
    {
      "key": "Name",
      "value": "tbqvt6hla7fm31gpfrsk"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "Rules",
      "value": "{direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-45a38175-1603-4e8b-b635-67a5b830755f/rules/r006-35d4b0ce-64c8-4d64-b5a4-2f771bc1ed90,id:r006-35d4b0ce-64c8-4d64-b5a4-2f771bc1ed90,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:braid-stainless-crystal-dropkick,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:3306,port_min:3306,protocol:tcp}; {direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-45a38175-1603-4e8b-b635-67a5b830755f/rules/r006-a6a18ce1-abc3-478a-8405-6ff8e70eac9b,id:r006-a6a18ce1-abc3-478a-8405-6ff8e70eac9b,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:explore-jugum-trade-linseed,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:22,port_min:22,protocol:tcp}; {direction:outbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-45a38175-1603-4e8b-b635-67a5b830755f/rules/r006-9e9248e1-83ec-4aee-93b2-980dfdfbef7a,id:r006-9e9248e1-83ec-4aee-93b2-980dfdfbef7a,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:laundry-zip-stooped-squiggly,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,protocol:any}"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,id:r006-201c49a7-94ae-4ca1-af79-1ec2a35ef375,name:tb4u3bq4enj35se6qmjp,resource_type:vpc}"
    }
  ],
  "name": "test-rdbms-sg-ibm",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbqvt6hla7fm31gpfrsk",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 20ms
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
- **Duration:** 8.931s
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
- **Duration:** 16ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "ibm",
    "region": "us-south"
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
      "sourceInstanceName": "Source MySQL 01",
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
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 31ms
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
- **Duration:** 18m23.913s
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
      "sourceInstanceName": "Source MySQL 01",
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
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
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
- **Duration:** 4.646s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm`
```json
// Response Body
{
  "adminUserName": "admin",
  "backupRetentionDays": 30,
  "backupTime": "AUTO",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T05:44:17Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T05:44:17Z",
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
  "cspResourceId": "3a08a972-d52b-48c4-9703-bf0a5d53877f",
  "cspResourceName": "tbbl4kdjt619j4klbvjc",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.4",
  "dbInstanceSpec": "b3c.4x16.encrypted",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "encryption": true,
  "endpoint": "3a08a972-d52b-48c4-9703-bf0a5d53877f.c9v3nfod0e3fgcbd1oug.databases.appdomain.cloud:30556",
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
  "uid": "tbbl4kdjt619j4klbvjc",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "adminUserName": "dbadmin",
      "backupRetentionDays": 7,
      "backupTime": "05:00Z-06:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T05:37:17Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:33:57Z",
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
      "cspResourceId": "rm-mj7e5k8a1l626j1lb",
      "cspResourceName": "tbn0suh6d0ggvjt3g8gr",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "dbInstanceType": "Basic",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
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
      "uid": "tbn0suh6d0ggvjt3g8gr",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "adminUserName": "admin",
      "backupRetentionDays": 30,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T05:44:17Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T05:44:17Z",
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
      "cspResourceId": "3a08a972-d52b-48c4-9703-bf0a5d53877f",
      "cspResourceName": "tbbl4kdjt619j4klbvjc",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "b3c.4x16.encrypted",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "3a08a972-d52b-48c4-9703-bf0a5d53877f.c9v3nfod0e3fgcbd1oug.databases.appdomain.cloud:30556",
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
      "uid": "tbbl4kdjt619j4klbvjc",
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 6.614s
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
- **Duration:** 11.588s
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
- **Duration:** 250ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m11.659s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 43.656s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 25.386s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.163s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 22.746s

