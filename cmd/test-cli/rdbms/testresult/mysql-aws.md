# Managed RDBMS (MySQL) Test Report: AWS (ap-northeast-2)

- **Test Case:** AWS AP-Northeast-2 (Seoul) MySQL Test
- **Date & Time:** 2026-09-08 10:31:39
- **Namespace:** `default`
- **Total Duration:** 16m10.396s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AWS
- **Target Region:** `ap-northeast-2`
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
- **Duration:** 2.506s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ami-0821dda66c52ff1ed",
  "specId": "aws+ap-northeast-2+t4g.medium"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "t4g.medium",
    "provider": "aws",
    "queriedAt": "2026-09-08T01:31:41.924224743Z",
    "region": "ap-northeast-2",
    "source": "aws:DescribeInstanceTypeOfferings",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2d"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2a"
      }
    ]
  },
  "connectionName": "aws-ap-northeast-2",
  "estimatedCost": "$0.0416/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "aws-ap-northeast-2",
    "creationDate": "2026-07-14T15:00:54.000Z",
    "cspImageName": "ami-0821dda66c52ff1ed",
    "description": "Canonical, Ubuntu Minimal, 24.04, arm64 noble image",
    "details": [
      {
        "key": "Architecture",
        "value": "arm64"
      },
      {
        "key": "BlockDeviceMappings",
        "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-072843fb772e92178,Throughput:null,VolumeSize:8,VolumeType:gp3},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
      },
      {
        "key": "BootMode",
        "value": "uefi"
      },
      {
        "key": "CreationDate",
        "value": "2026-07-14T15:00:54.000Z"
      },
      {
        "key": "DeprecationTime",
        "value": "2028-07-14T15:00:54.000Z"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu Minimal, 24.04, arm64 noble image"
      },
      {
        "key": "EnaSupport",
        "value": "true"
      },
      {
        "key": "Hypervisor",
        "value": "xen"
      },
      {
        "key": "ImageId",
        "value": "ami-0821dda66c52ff1ed"
      },
      {
        "key": "ImageLocation",
        "value": "amazon/ubuntu-minimal/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-minimal-20260714"
      },
      {
        "key": "ImageOwnerAlias",
        "value": "amazon"
      },
      {
        "key": "ImageType",
        "value": "machine"
      },
      {
        "key": "Name",
        "value": "ubuntu-minimal/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-minimal-20260714"
      },
      {
        "key": "OwnerId",
        "value": "099720109477"
      },
      {
        "key": "PlatformDetails",
        "value": "Linux/UNIX"
      },
      {
        "key": "Public",
        "value": "true"
      },
      {
        "key": "RootDeviceName",
        "value": "/dev/sda1"
      },
      {
        "key": "RootDeviceType",
        "value": "ebs"
      },
      {
        "key": "SriovNetSupport",
        "value": "simple"
      },
      {
        "key": "State",
        "value": "available"
      },
      {
        "key": "UsageOperation",
        "value": "RunInstances"
      },
      {
        "key": "VirtualizationType",
        "value": "hvm"
      }
    ],
    "fetchedTime": "2026.08.21 14:20:40 Fri",
    "id": "ami-0821dda66c52ff1ed",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ami-0821dda66c52ff1ed",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "ebs",
    "osDistribution": "ubuntu-minimal/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-minimal-20260714",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "aws",
    "regionList": [
      "ap-northeast-2"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbd7vkkmglhda3jfcjb1"
  },
  "imageId": "ami-0821dda66c52ff1ed",
  "imageValidation": {
    "cspResourceId": "ami-0821dda66c52ff1ed",
    "isAvailable": true,
    "resourceId": "ami-0821dda66c52ff1ed",
    "resourceName": "ami-0821dda66c52ff1ed",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "aws",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "arm64",
    "connectionName": "aws-ap-northeast-2",
    "costPerHour": 0.0416,
    "cspSpecName": "t4g.medium",
    "details": [
      {
        "key": "AutoRecoverySupported",
        "value": "true"
      },
      {
        "key": "BareMetal",
        "value": "false"
      },
      {
        "key": "BurstablePerformanceSupported",
        "value": "true"
      },
      {
        "key": "CurrentGeneration",
        "value": "true"
      },
      {
        "key": "DedicatedHostsSupported",
        "value": "false"
      },
      {
        "key": "EbsInfo",
        "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:347,BaselineIops:2000,BaselineThroughputInMBps:43.375,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
      },
      {
        "key": "FreeTierEligible",
        "value": "false"
      },
      {
        "key": "HibernationSupported",
        "value": "true"
      },
      {
        "key": "Hypervisor",
        "value": "nitro"
      },
      {
        "key": "InstanceStorageSupported",
        "value": "false"
      },
      {
        "key": "InstanceType",
        "value": "t4g.medium"
      },
      {
        "key": "MemoryInfo",
        "value": "{SizeInMiB:4096}"
      },
      {
        "key": "NetworkInfo",
        "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:6,Ipv6AddressesPerInterface:6,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
      },
      {
        "key": "PlacementGroupInfo",
        "value": "{SupportedStrategies:[partition,spread]}"
      },
      {
        "key": "ProcessorInfo",
        "value": "{SupportedArchitectures:[arm64],SustainedClockSpeedInGhz:2.5}"
      },
      {
        "key": "SupportedBootModes",
        "value": "uefi"
      },
      {
        "key": "SupportedRootDeviceTypes",
        "value": "ebs"
      },
      {
        "key": "SupportedUsageClasses",
        "value": "on-demand; spot"
      },
      {
        "key": "SupportedVirtualizationTypes",
        "value": "hvm"
      },
      {
        "key": "VCpuInfo",
        "value": "{DefaultCores:2,DefaultThreadsPerCore:1,DefaultVCpus:2,ValidCores:[1,2],ValidThreadsPerCore:[1]}"
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
    "id": "aws+ap-northeast-2+t4g.medium",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "aws+ap-northeast-2+t4g.medium",
    "namespace": "system",
    "providerName": "aws",
    "regionLatitude": 37.36,
    "regionLongitude": 126.78,
    "regionName": "ap-northeast-2",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbk5s7k571ba0vqjtqkg",
    "vCPU": 2
  },
  "specId": "aws+ap-northeast-2+t4g.medium",
  "specValidation": {
    "cspResourceId": "t4g.medium",
    "isAvailable": true,
    "resourceId": "aws+ap-northeast-2+t4g.medium",
    "resourceName": "t4g.medium",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "ap-northeast-2b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 3.551s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.0.0.0/16",
  "connectionName": "aws-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-aws",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.0.1.0/24",
      "name": "subnet-1",
      "zone": "ap-northeast-2a"
    },
    {
      "ipv4_CIDR": "10.0.2.0/24",
      "name": "subnet-2",
      "zone": "ap-northeast-2c"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.0.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:31:44Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:44Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-09-08T01:31:44Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "vpc-0e082105bec4893e7",
  "cspResourceName": "tb26tngju18ljuco5p6s",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CidrBlock",
      "value": "10.0.0.0/16"
    },
    {
      "key": "CidrBlockAssociationSet",
      "value": "{AssociationId:vpc-cidr-assoc-03be1524b478f5cf0,CidrBlock:10.0.0.0/16,CidrBlockState:{State:associated,StatusMessage:null}}"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-fa6b9492"
    },
    {
      "key": "InstanceTenancy",
      "value": "default"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "OwnerId",
      "value": "635484366616"
    },
    {
      "key": "State",
      "value": "pending"
    },
    {
      "key": "Tags",
      "value": "{Key:Name,Value:tb26tngju18ljuco5p6s}"
    },
    {
      "key": "VpcId",
      "value": "vpc-0e082105bec4893e7"
    }
  ],
  "name": "test-rdbms-vnet-aws",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:31:44Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:31:44Z",
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
      "cspResourceId": "subnet-078473082152506bb",
      "cspResourceName": "tbh8ai15ctp47ss0oc3l",
      "cspVNetId": "vpc-0e082105bec4893e7",
      "cspVNetName": "tb26tngju18ljuco5p6s",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.0.1.0/24",
      "keyValueList": [
        {
          "key": "AssignIpv6AddressOnCreation",
          "value": "false"
        },
        {
          "key": "AvailabilityZone",
          "value": "ap-northeast-2a"
        },
        {
          "key": "AvailabilityZoneId",
          "value": "apne2-az1"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "251"
        },
        {
          "key": "CidrBlock",
          "value": "10.0.1.0/24"
        },
        {
          "key": "DefaultForAz",
          "value": "false"
        },
        {
          "key": "MapCustomerOwnedIpOnLaunch",
          "value": "false"
        },
        {
          "key": "MapPublicIpOnLaunch",
          "value": "false"
        },
        {
          "key": "OwnerId",
          "value": "635484366616"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "SubnetArn",
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-078473082152506bb"
        },
        {
          "key": "SubnetId",
          "value": "subnet-078473082152506bb"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbh8ai15ctp47ss0oc3l}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0e082105bec4893e7"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbh8ai15ctp47ss0oc3l",
      "zone": "ap-northeast-2a"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-09-08T01:31:44Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-09-08T01:31:44Z",
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
      "cspResourceId": "subnet-094f97ca90c4d7577",
      "cspResourceName": "tbu34kdtvhplognill29",
      "cspVNetId": "vpc-0e082105bec4893e7",
      "cspVNetName": "tb26tngju18ljuco5p6s",
      "description": "",
      "id": "subnet-2",
      "ipv4_CIDR": "10.0.2.0/24",
      "keyValueList": [
        {
          "key": "AssignIpv6AddressOnCreation",
          "value": "false"
        },
        {
          "key": "AvailabilityZone",
          "value": "ap-northeast-2c"
        },
        {
          "key": "AvailabilityZoneId",
          "value": "apne2-az3"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "251"
        },
        {
          "key": "CidrBlock",
          "value": "10.0.2.0/24"
        },
        {
          "key": "DefaultForAz",
          "value": "false"
        },
        {
          "key": "MapCustomerOwnedIpOnLaunch",
          "value": "false"
        },
        {
          "key": "MapPublicIpOnLaunch",
          "value": "false"
        },
        {
          "key": "OwnerId",
          "value": "635484366616"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "SubnetArn",
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-094f97ca90c4d7577"
        },
        {
          "key": "SubnetId",
          "value": "subnet-094f97ca90c4d7577"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbu34kdtvhplognill29}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0e082105bec4893e7"
        }
      ],
      "name": "subnet-2",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbu34kdtvhplognill29",
      "zone": "ap-northeast-2c"
    }
  ],
  "systemLabel": "",
  "uid": "tb26tngju18ljuco5p6s"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.491s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "aws-ap-northeast-2",
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
  "name": "test-rdbms-sg-aws",
  "vNetId": "test-rdbms-vnet-aws"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "sg-02965863fb6ab1cf2",
  "cspResourceName": "tbck7n5mi36aateo197i",
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
  "id": "test-rdbms-sg-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "GroupName",
      "value": "tbck7n5mi36aateo197i"
    },
    {
      "key": "VpcID",
      "value": "vpc-0e082105bec4893e7"
    },
    {
      "key": "OwnerID",
      "value": "635484366616"
    },
    {
      "key": "Description",
      "value": "tbck7n5mi36aateo197i"
    }
  ],
  "name": "test-rdbms-sg-aws",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbck7n5mi36aateo197i",
  "vNetId": "test-rdbms-vnet-aws"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 3ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=aws`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "aws": {
      "dbOperationMethod": "sqlFallback",
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
- **Duration:** 35.511s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=aws-ap-northeast-2&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "0-35",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "db.m5.12xlarge",
      "db.m5.16xlarge",
      "db.m5.24xlarge",
      "db.m5.2xlarge",
      "db.m5.4xlarge",
      "db.m5.8xlarge",
      "db.m5.large",
      "db.m5.xlarge",
      "db.m5d.12xlarge",
      "db.m5d.16xlarge",
      "db.m5d.24xlarge",
      "db.m5d.2xlarge",
      "db.m5d.4xlarge",
      "db.m5d.8xlarge",
      "db.m5d.large",
      "db.m5d.xlarge",
      "db.m6g.12xlarge",
      "db.m6g.16xlarge",
      "db.m6g.2xlarge",
      "db.m6g.4xlarge",
      "db.m6g.8xlarge",
      "db.m6g.large",
      "db.m6g.xlarge",
      "db.m6gd.12xlarge",
      "db.m6gd.16xlarge",
      "db.m6gd.2xlarge",
      "db.m6gd.4xlarge",
      "db.m6gd.8xlarge",
      "db.m6gd.large",
      "db.m6gd.xlarge",
      "db.m6i.12xlarge",
      "db.m6i.16xlarge",
      "db.m6i.24xlarge",
      "db.m6i.2xlarge",
      "db.m6i.32xlarge",
      "db.m6i.4xlarge",
      "db.m6i.8xlarge",
      "db.m6i.large",
      "db.m6i.xlarge",
      "db.m7g.12xlarge",
      "db.m7g.16xlarge",
      "db.m7g.2xlarge",
      "db.m7g.4xlarge",
      "db.m7g.8xlarge",
      "db.m7g.large",
      "db.m7g.xlarge",
      "db.m7i.12xlarge",
      "db.m7i.16xlarge",
      "db.m7i.24xlarge",
      "db.m7i.2xlarge",
      "db.m7i.48xlarge",
      "db.m7i.4xlarge",
      "db.m7i.8xlarge",
      "db.m7i.large",
      "db.m7i.xlarge",
      "db.m8g.12xlarge",
      "db.m8g.16xlarge",
      "db.m8g.24xlarge",
      "db.m8g.2xlarge",
      "db.m8g.48xlarge",
      "db.m8g.4xlarge",
      "db.m8g.8xlarge",
      "db.m8g.large",
      "db.m8g.xlarge",
      "db.r5.12xlarge",
      "db.r5.16xlarge",
      "db.r5.24xlarge",
      "db.r5.2xlarge",
      "db.r5.4xlarge",
      "db.r5.8xlarge",
      "db.r5.large",
      "db.r5.xlarge",
      "db.r5d.12xlarge",
      "db.r5d.16xlarge",
      "db.r5d.24xlarge",
      "db.r5d.2xlarge",
      "db.r5d.4xlarge",
      "db.r5d.8xlarge",
      "db.r5d.large",
      "db.r5d.xlarge",
      "db.r6g.12xlarge",
      "db.r6g.16xlarge",
      "db.r6g.2xlarge",
      "db.r6g.4xlarge",
      "db.r6g.8xlarge",
      "db.r6g.large",
      "db.r6g.xlarge",
      "db.r6i.12xlarge",
      "db.r6i.16xlarge",
      "db.r6i.24xlarge",
      "db.r6i.2xlarge",
      "db.r6i.32xlarge",
      "db.r6i.4xlarge",
      "db.r6i.8xlarge",
      "db.r6i.large",
      "db.r6i.xlarge",
      "db.r7g.12xlarge",
      "db.r7g.16xlarge",
      "db.r7g.2xlarge",
      "db.r7g.4xlarge",
      "db.r7g.8xlarge",
      "db.r7g.large",
      "db.r7g.xlarge",
      "db.r7i.12xlarge",
      "db.r7i.16xlarge",
      "db.r7i.24xlarge",
      "db.r7i.2xlarge",
      "db.r7i.48xlarge",
      "db.r7i.4xlarge",
      "db.r7i.8xlarge",
      "db.r7i.large",
      "db.r7i.xlarge",
      "db.r8g.12xlarge",
      "db.r8g.16xlarge",
      "db.r8g.24xlarge",
      "db.r8g.2xlarge",
      "db.r8g.48xlarge",
      "db.r8g.4xlarge",
      "db.r8g.8xlarge",
      "db.r8g.large",
      "db.r8g.xlarge",
      "db.r8gd.12xlarge",
      "db.r8gd.16xlarge",
      "db.r8gd.24xlarge",
      "db.r8gd.2xlarge",
      "db.r8gd.48xlarge",
      "db.r8gd.4xlarge",
      "db.r8gd.8xlarge",
      "db.r8gd.large",
      "db.r8gd.xlarge",
      "db.t3.2xlarge",
      "db.t3.large",
      "db.t3.medium",
      "db.t3.micro",
      "db.t3.small",
      "db.t3.xlarge",
      "db.t4g.2xlarge",
      "db.t4g.large",
      "db.t4g.medium",
      "db.t4g.micro",
      "db.t4g.small",
      "db.t4g.xlarge"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "196608",
        "name": "db.m5.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m5.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m5.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m5.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m5.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m5.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m5.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m5.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m5d.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m5d.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m5d.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m5d.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m5d.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m5d.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m5d.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m5d.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6gd.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6gd.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6gd.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6gd.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6gd.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6gd.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6gd.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m6i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.m6i.32xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "128"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m7g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m7g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m7g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m7g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m7g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m7g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m7g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m7i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m7i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m7i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m7i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.m7i.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m7i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m7i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m7i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m7i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m8g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m8g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m8g.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m8g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.m8g.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m8g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m8g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m8g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m8g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r5.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r5.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r5.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r5.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r5.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r5.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r5.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r5.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r5d.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r5d.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r5d.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r5d.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r5d.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r5d.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r5d.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r5d.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r6g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r6g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r6g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r6g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r6g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r6g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r6g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r6i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r6i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r6i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r6i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1048576",
        "name": "db.r6i.32xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "128"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r6i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r6i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r6i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r6i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r7g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r7g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r7g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r7g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r7g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r7g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r7g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r7i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r7i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r7i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r7i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r7i.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r7i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r7i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r7i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r7i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r8g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r8g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r8g.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r8g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r8g.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r8g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r8g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r8g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r8g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r8gd.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r8gd.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r8gd.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r8gd.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r8gd.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r8gd.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r8gd.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r8gd.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r8gd.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.t3.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.t3.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "db.t3.medium",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "db.t3.micro",
        "storageSizeRangeGB": {
          "max": 6597,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "db.t3.small",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.t3.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.t4g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.t4g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "db.t4g.medium",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "db.t4g.micro",
        "storageSizeRangeGB": {
          "max": 6597,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "db.t4g.small",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.t4g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
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
          "description": "Legacy general-purpose SSD storage. Consider gp3 for better cost/performance.",
          "displayName": "General Purpose SSD (gp2)",
          "maxSize": 65536,
          "minSize": 20,
          "recommendationLevel": "legacy",
          "storageType": "gp2"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Latest generation general-purpose SSD with 3,000 baseline IOPS and 125 MiB/s throughput. Recommended for most workloads.",
          "displayName": "General Purpose SSD (gp3)",
          "maxSize": 65536,
          "minSize": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "gp3"
        },
        {
          "constraints": "Requires 'iops' parameter (range: 1000-64000). Minimum 100GB storage.",
          "description": "High-performance SSD for I/O-intensive workloads. Requires 'iops' parameter (e.g., '3000'). Minimum 100 GB storage.",
          "displayName": "Provisioned IOPS SSD (io1)",
          "iopsRange": {
            "max": 64000,
            "min": 1000
          },
          "maxSize": 65536,
          "minSize": 100,
          "recommendationLevel": "premium",
          "requiresIops": true,
          "storageType": "io1"
        },
        {
          "constraints": "Requires 'iops' parameter (range: 1000-256000). Minimum 100GB storage.",
          "description": "Next-generation high-performance SSD with higher durability (99.999%). Requires 'iops' parameter. Minimum 100 GB storage.",
          "displayName": "Provisioned IOPS SSD (io2)",
          "iopsRange": {
            "max": 256000,
            "min": 1000
          },
          "maxSize": 65536,
          "minSize": 100,
          "recommendationLevel": "premium",
          "requiresIops": true,
          "storageType": "io2"
        }
      ]
    },
    "providerName": "aws",
    "regionName": "ap-northeast-2",
    "requiresSecurityGroup": true,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 70369,
      "min": 5
    },
    "storageTypeOptions": [
      "gp2",
      "gp3",
      "io1",
      "io2"
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
- **Duration:** 8ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "aws",
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-aws",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-aws"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 3.724s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "root",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "aws-ap-northeast-2",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db.t3.medium",
  "name": "rdbms-aws",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-aws"
  ],
  "storageSize": 100,
  "storageType": "gp3",
  "subnetIds": [
    "subnet-1",
    "subnet-2"
  ],
  "vNetId": "test-rdbms-vnet-aws"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "root",
    "adminUserPassword": "******",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "db.t3.medium",
    "name": "rdbms-aws",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-aws"
    ],
    "storageSize": 100,
    "storageType": "gp3",
    "subnetIds": [
      "subnet-1",
      "subnet-2"
    ],
    "vNetId": "test-rdbms-vnet-aws"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 8m42.373s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-aws",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-aws"
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
- **Duration:** 11ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws`
```json
// Response Body
{
  "adminUserName": "root",
  "backupRetentionDays": 7,
  "backupTime": "19:02-19:32",
  "conditions": [
    {
      "lastTransitionTime": "2026-09-08T01:40:15Z",
      "reason": "Available",
      "status": "True",
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
  "status": "Available",
  "storageSize": 100,
  "storageType": "gp3",
  "subnetIds": [
    "subnet-1",
    "subnet-2"
  ],
  "tagList": [
    {
      "key": "sys.connectionName",
      "value": "aws-ap-northeast-2"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.cspResourceId",
      "value": "tblunavhu1of0ntjighq"
    },
    {
      "key": "sys.id",
      "value": "test-rdbms-aws"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tblunavhu1of0ntjighq"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.uid",
      "value": "tblunavhu1of0ntjighq"
    },
    {
      "key": "sys.name",
      "value": "test-rdbms-aws"
    },
    {
      "key": "Name",
      "value": "tblunavhu1of0ntjighq"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    }
  ],
  "uid": "tblunavhu1of0ntjighq",
  "vNetId": "test-rdbms-vnet-aws"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 11ms
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
          "lastTransitionTime": "2026-09-08T01:40:15Z",
          "reason": "Available",
          "status": "True",
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
      "status": "Available",
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
      "endpoint": "db-4acagh.vpc-cdb.ntruss.com:3306",
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
- **Duration:** 678ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database`
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
- **Duration:** 530ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database`
```json
// Response Body
{
  "databases": [
    "information_schema",
    "mysql",
    "performance_schema",
    "sampledb",
    "sampledb_dyn",
    "sys"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 107ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m34.146s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 19.165s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m22.181s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 2.457s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 21.943s

