# Managed RDBMS (MySQL) Test Report: AZURE (koreacentral)

- **Test Case:** Azure KoreaCentral MySQL Test
- **Date & Time:** 2026-09-08 10:31:39
- **Namespace:** `default`
- **Total Duration:** 24m59.505s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AZURE
- **Target Region:** `koreacentral`
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
- **Duration:** 4.387s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
  "specId": "azure+koreacentral+standard_d2ps_v6"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "Standard_D2ps_v6",
    "provider": "azure",
    "queriedAt": "2026-09-08T01:31:43.143087681Z",
    "region": "koreacentral",
    "source": "azure:CheckSpecAvailability"
  },
  "connectionName": "azure-koreacentral",
  "estimatedCost": "$0.0914/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "azure-australiacentral",
    "creationDate": "",
    "cspImageName": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608270",
    "description": "",
    "details": [
      {
        "key": "Location",
        "value": "australiacentral"
      },
      {
        "key": "Publisher",
        "value": "Canonical"
      },
      {
        "key": "Offer",
        "value": "ubuntu-24_04-lts"
      },
      {
        "key": "SKU",
        "value": "minimal-arm64"
      },
      {
        "key": "Version",
        "value": "24.04.202608100"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/ubuntu-24_04-lts/Skus/minimal-arm64/Versions/24.04.202608100"
      },
      {
        "key": "HyperVGeneration",
        "value": "V2"
      },
      {
        "key": "Features",
        "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
      },
      {
        "key": "FeatureCount",
        "value": "4"
      },
      {
        "key": "ImageDeprecationState",
        "value": "Active"
      }
    ],
    "fetchedTime": "2026.08.21 14:01:58 Fri",
    "id": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "default",
    "osDistribution": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "azure",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tb65d9cf6b2qanmqmpf1"
  },
  "imageId": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
  "imageValidation": {
    "cspResourceId": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608270",
    "isAvailable": true,
    "resourceId": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
    "resourceName": "Canonical:ubuntu-24_04-lts:minimal-arm64:24.04.202608100",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "azure",
  "regionName": "koreacentral",
  "specDetails": {
    "architecture": "arm64",
    "connectionName": "azure-koreacentral",
    "costPerHour": 0.0914,
    "cspSpecName": "Standard_D2ps_v6",
    "details": [
      {
        "key": "MaxDataDiskCount",
        "value": "8"
      },
      {
        "key": "MemoryInMB",
        "value": "8192"
      },
      {
        "key": "Name",
        "value": "Standard_D2ps_v6"
      },
      {
        "key": "NumberOfCores",
        "value": "2"
      },
      {
        "key": "OSDiskSizeInMB",
        "value": "1047552"
      },
      {
        "key": "ResourceDiskSizeInMB",
        "value": "0"
      },
      {
        "key": "MaxResourceVolumeMB",
        "value": "0"
      },
      {
        "key": "OSVhdSizeMB",
        "value": "1047552"
      },
      {
        "key": "vCPUs",
        "value": "2"
      },
      {
        "key": "MemoryPreservingMaintenanceSupported",
        "value": "True"
      },
      {
        "key": "HyperVGenerations",
        "value": "V2"
      },
      {
        "key": "DiskControllerTypes",
        "value": "SCSI"
      },
      {
        "key": "SupportedCapacityReservationTypes",
        "value": "Open,Targeted"
      },
      {
        "key": "MemoryGB",
        "value": "8"
      },
      {
        "key": "MaxDataDiskCount",
        "value": "8"
      },
      {
        "key": "CpuArchitectureType",
        "value": "Arm64"
      },
      {
        "key": "LowPriorityCapable",
        "value": "True"
      },
      {
        "key": "PremiumIO",
        "value": "True"
      },
      {
        "key": "VMDeploymentTypes",
        "value": "IaaS"
      },
      {
        "key": "vCPUsConstraintsAllowed",
        "value": "1, 2"
      },
      {
        "key": "vCPUsAvailable",
        "value": "2"
      },
      {
        "key": "vCPUsPerCore",
        "value": "1"
      },
      {
        "key": "CombinedTempDiskAndCachedIOPS",
        "value": "9000"
      },
      {
        "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
        "value": "125000000"
      },
      {
        "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
        "value": "125000000"
      },
      {
        "key": "UncachedDiskIOPS",
        "value": "3750"
      },
      {
        "key": "UncachedDiskBytesPerSecond",
        "value": "106000000"
      },
      {
        "key": "EphemeralOSDiskSupported",
        "value": "False"
      },
      {
        "key": "EncryptionAtHostSupported",
        "value": "True"
      },
      {
        "key": "CapacityReservationSupported",
        "value": "True"
      },
      {
        "key": "AcceleratedNetworkingEnabled",
        "value": "True"
      },
      {
        "key": "RdmaEnabled",
        "value": "False"
      },
      {
        "key": "MaxNetworkInterfaces",
        "value": "2"
      },
      {
        "key": "UltraSSDAvailable",
        "value": "False"
      },
      {
        "key": "LocationInfo_0_Location",
        "value": "KoreaCentral"
      },
      {
        "key": "LocationInfo_0_Zone_0",
        "value": "2"
      },
      {
        "key": "LocationInfo_0_Zone_1",
        "value": "3"
      },
      {
        "key": "LocationInfo_0_Zone_2",
        "value": "1"
      },
      {
        "key": "Family",
        "value": "StandardDpsv6Family"
      },
      {
        "key": "Tier",
        "value": "Standard"
      },
      {
        "key": "Size",
        "value": "D2ps_v6"
      },
      {
        "key": "ResourceType",
        "value": "virtualMachines"
      }
    ],
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
    "id": "azure+koreacentral+standard_d2ps_v6",
    "infraType": "node",
    "memoryGiB": 7.8125,
    "name": "azure+koreacentral+standard_d2ps_v6",
    "namespace": "system",
    "providerName": "azure",
    "regionLatitude": 37.5665,
    "regionLongitude": 126.978,
    "regionName": "koreacentral",
    "rootDiskSize": 0,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbahqjlrbu2ab7hk6kqr",
    "vCPU": 2
  },
  "specId": "azure+koreacentral+standard_d2ps_v6",
  "specValidation": {
    "cspResourceId": "Standard_D2ps_v6",
    "isAvailable": true,
    "resourceId": "azure+koreacentral+standard_d2ps_v6",
    "resourceName": "Standard_D2ps_v6",
    "status": "Available"
  },
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 9.296s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.1.0.0/16",
  "connectionName": "azure-koreacentral",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-azure",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.1.1.0/24",
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
  "cidrBlock": "10.1.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:31:50Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:50Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:50Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875",
  "cspResourceName": "tboae11frt922f888875",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{addressSpace:{addressPrefixes:[10.1.0.0/16]},enableDdosProtection:false,privateEndpointVNetPolicies:Disabled,provisioningState:Succeeded,resourceGuid:3aa26f01-5540-4427-94f9-579427b87e2f,subnets:[{etag:W/\\cc777b90-bfbc-4936-9bff-57e3554201f8\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875/subnets/tbu0qsdri99sd33et3up,name:tbu0qsdri99sd33et3up,properties:{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]},type:Microsoft.Network/virtualNetworks/subnets}],virtualNetworkPeerings:[]}"
    },
    {
      "key": "Etag",
      "value": "W/\\cc777b90-bfbc-4936-9bff-57e3554201f8\\"
    },
    {
      "key": "Name",
      "value": "tboae11frt922f888875"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/virtualNetworks"
    }
  ],
  "name": "test-rdbms-vnet-azure",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:31:50Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:31:50Z",
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
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875/subnets/tbu0qsdri99sd33et3up",
      "cspResourceName": "tbu0qsdri99sd33et3up",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875",
      "cspVNetName": "tboae11frt922f888875",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.1.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tboae11frt922f888875/subnets/tbu0qsdri99sd33et3up"
        },
        {
          "key": "Name",
          "value": "tbu0qsdri99sd33et3up"
        },
        {
          "key": "Properties",
          "value": "{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]}"
        },
        {
          "key": "Type",
          "value": "Microsoft.Network/virtualNetworks/subnets"
        },
        {
          "key": "Etag",
          "value": "W/\\cc777b90-bfbc-4936-9bff-57e3554201f8\\"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbu0qsdri99sd33et3up"
    }
  ],
  "systemLabel": "",
  "uid": "tboae11frt922f888875"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 8.399s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "azure-koreacentral",
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
  "name": "test-rdbms-sg-azure",
  "vNetId": "test-rdbms-vnet-azure"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp",
  "cspResourceName": "tbitplqmtpshib5072tp",
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
  "id": "test-rdbms-sg-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{defaultSecurityRules:[{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/AllowVnetInBound,name:AllowVnetInBound,properties:{access:Allow,description:Allow inbound traffic from all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/AllowAzureLoadBalancerInBound,name:AllowAzureLoadBalancerInBound,properties:{access:Allow,description:Allow inbound traffic from azure load balancer,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:AzureLoadBalancer,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/DenyAllInBound,name:DenyAllInBound,properties:{access:Deny,description:Deny all inbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/AllowVnetOutBound,name:AllowVnetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/AllowInternetOutBound,name:AllowInternetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to Internet,destinationAddressPrefix:Internet,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/defaultSecurityRules/DenyAllOutBound,name:DenyAllOutBound,properties:{access:Deny,description:Deny all outbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules}],provisioningState:Succeeded,resourceGuid:fcbea45b-331b-40a3-8816-d8164d6735e0,securityRules:[{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/securityRules/inbound-rules-58084-3306-3306-TCP,name:inbound-rules-58084-3306-3306-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:3306,destinationPortRanges:[],direction:Inbound,priority:100,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/securityRules/inbound-rules-35692-22-22-TCP,name:inbound-rules-35692-22-22-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:22,destinationPortRanges:[],direction:Inbound,priority:101,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/securityRules/deny-outbound,name:deny-outbound,properties:{access:Deny,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:4096,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbitplqmtpshib5072tp/securityRules/allow-outbound,name:allow-outbound,properties:{access:Allow,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:101,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules}]}"
    },
    {
      "key": "Etag",
      "value": "W/\\9a10c288-57cc-4797-a6df-e18c3089d0a5\\"
    },
    {
      "key": "Name",
      "value": "tbitplqmtpshib5072tp"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/networkSecurityGroups"
    }
  ],
  "name": "test-rdbms-sg-azure",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbitplqmtpshib5072tp",
  "vNetId": "test-rdbms-vnet-azure"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=azure`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "azure": {
      "dbOperationMethod": "cspNativeApi",
      "note": "Storage type selection not supported. Storage SKU is read-only and automatically set by Azure based on compute tier (Premium SSD for General Purpose/Memory Optimized tiers, locally redundant storage for Burstable tier). dbOperationMethod is cspNativeApi (armmysqlfs.DatabasesClient). Azure uses SubnetNames only in VPC-private mode (PublicAccess=false); when PublicAccess=true, subnet is not used.",
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
- **Duration:** 6.52s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=azure-koreacentral&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "1-35",
    "connectionName": "azure-koreacentral",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "Standard_B12ms",
      "Standard_B16ms",
      "Standard_B1ms",
      "Standard_B20ms",
      "Standard_B2ms",
      "Standard_B2s",
      "Standard_B4ms",
      "Standard_B8ms",
      "Standard_D16ads_v6",
      "Standard_D16ds_v4",
      "Standard_D16ds_v6",
      "Standard_D2ads_v6",
      "Standard_D2ds_v4",
      "Standard_D2ds_v6",
      "Standard_D32ads_v6",
      "Standard_D32ds_v4",
      "Standard_D32ds_v6",
      "Standard_D48ads_v6",
      "Standard_D48ds_v4",
      "Standard_D48ds_v6",
      "Standard_D4ads_v6",
      "Standard_D4ds_v4",
      "Standard_D4ds_v6",
      "Standard_D64ads_v6",
      "Standard_D64ds_v4",
      "Standard_D64ds_v6",
      "Standard_D8ads_v6",
      "Standard_D8ds_v4",
      "Standard_D8ds_v6",
      "Standard_D96ads_v6",
      "Standard_D96ds_v6",
      "Standard_E16ads_v6",
      "Standard_E16ds_v4",
      "Standard_E16ds_v5",
      "Standard_E16ds_v6",
      "Standard_E20ads_v6",
      "Standard_E20ds_v4",
      "Standard_E20ds_v5",
      "Standard_E20ds_v6",
      "Standard_E2ads_v6",
      "Standard_E2ds_v4",
      "Standard_E2ds_v5",
      "Standard_E2ds_v6",
      "Standard_E32ads_v6",
      "Standard_E32ds_v4",
      "Standard_E32ds_v5",
      "Standard_E32ds_v6",
      "Standard_E48ads_v6",
      "Standard_E48ds_v4",
      "Standard_E48ds_v5",
      "Standard_E48ds_v6",
      "Standard_E4ads_v6",
      "Standard_E4ds_v4",
      "Standard_E4ds_v5",
      "Standard_E4ds_v6",
      "Standard_E64ads_v6",
      "Standard_E64ds_v4",
      "Standard_E64ds_v5",
      "Standard_E64ds_v6",
      "Standard_E80ids_v4",
      "Standard_E8ads_v6",
      "Standard_E8ds_v4",
      "Standard_E8ds_v5",
      "Standard_E8ds_v6",
      "Standard_E96ads_v6",
      "Standard_E96ds_v5",
      "Standard_E96ds_v6"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "49152",
        "name": "Standard_B12ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "12"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_B16ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "2048",
        "name": "Standard_B1ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "81920",
        "name": "Standard_B20ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_B2ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "Standard_B2s",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_B4ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_B8ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_D96ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_D96ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "516080",
        "name": "Standard_E80ids_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "688128",
        "name": "Standard_E96ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "688128",
        "name": "Standard_E96ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "786432",
        "name": "Standard_E96ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
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
          "description": "Storage SKU is automatically determined by Azure based on compute tier. User cannot specify storage type.",
          "displayName": "Automatic (Azure-managed)",
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "azure",
    "regionName": "koreacentral",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
      "max": 35184,
      "min": 21
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "5.7",
      "8.0.21",
      "8.4",
      "9.5"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
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
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "azure",
    "region": "koreacentral"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for azure (koreacentral)",
  "status": "recommended",
  "targetCloud": {
    "csp": "azure",
    "region": "koreacentral"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "azureuser",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-azure",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-azure"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (azure); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 18ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "azureuser",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "azure-koreacentral",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.21",
  "dbInstanceSpec": "Standard_B2s",
  "name": "rdbms-azure",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-azure"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-azure"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "azureuser",
    "adminUserPassword": "******",
    "connectionName": "azure-koreacentral",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0.21",
    "dbInstanceSpec": "Standard_B2s",
    "name": "rdbms-azure",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-azure"
    ],
    "storageSize": 100,
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-azure"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 16m12.819s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for azure (koreacentral)",
  "status": "recommended",
  "targetCloud": {
    "csp": "azure",
    "region": "koreacentral"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "azureuser",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-azure",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-azure"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (azure); requested storage type 'SSD' for instance 'source-mysql-01' will be managed automatically by the provider."
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
- **Duration:** 3.295s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure`
```json
// Response Body
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
  "tagList": [
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.uid",
      "value": "tbtbfiifbfn8uvco6qhr"
    },
    {
      "key": "sys.connectionName",
      "value": "azure-koreacentral"
    },
    {
      "key": "sys.cspResourceId",
      "value": "tbtbfiifbfn8uvco6qhr"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbtbfiifbfn8uvco6qhr"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.name",
      "value": "test-rdbms-azure"
    }
  ],
  "uid": "tbtbfiifbfn8uvco6qhr",
  "vNetId": "test-rdbms-vnet-azure"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
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
- **Duration:** 23.125s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database`
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
- **Duration:** 7.385s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database`
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
- **Duration:** 208ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m29.852s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 53.985s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 41.582s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 8.516s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 30.102s

