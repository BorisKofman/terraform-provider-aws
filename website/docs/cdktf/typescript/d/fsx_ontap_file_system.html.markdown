---
subcategory: "FSx"
layout: "aws"
page_title: "AWS: aws_fsx_ontap_file_system"
description: |-
  Retrieve information on FSx ONTAP File System.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_fsx_ontap_file_system

Retrieve information on FSx ONTAP File System.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsFsxOntapFileSystem } from "./.gen/providers/aws/data-aws-fsx-ontap-file-system";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsFsxOntapFileSystem(this, "example", {
      id: "fs-12345678",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `id` - (Required) Identifier of the file system (e.g. `fs-12345678`).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name of the file system.
* `automaticBackupRetentionDays` - The number of days to retain automatic backups.
* `dailyAutomaticBackupStartTime` - The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
* `deploymentType` - The file system deployment type.
* `diskIopsConfiguration` - The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system, specifying the number of provisioned IOPS and the provision mode. See [Disk IOPS](#disk-iops) Below.
* `dnsName` - DNS name for the file system.

  **Note:** This attribute does not apply to FSx for ONTAP file systems and is consequently not set. You can access your FSx for ONTAP file system and volumes via a [Storage Virtual Machine (SVM)](fsx_ontap_storage_virtual_machine.html) using its DNS name or IP address.
* `endpointIpAddressRange` - (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system exist.
* `endpoints` - The Management and Intercluster FileSystemEndpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See [FileSystemEndpoints](#file-system-endpoints) below.
* `haPairs` - The number of HA pairs for the file system.
* `id` - Identifier of the file system (e.g. `fs-12345678`).
* `kmsKeyId` - ARN for the KMS Key to encrypt the file system at rest.
* `networkInterfaceIds` - The IDs of the elastic network interfaces from which a specific file system is accessible.
* `ownerId` - AWS account identifier that created the file system.
* `preferredSubnetId` - Specifies the subnet in which you want the preferred file server to be located.
* `routeTableIds` - (Multi-AZ only) The VPC route tables in which your file system's endpoints exist.
* `storageCapacity` - The storage capacity of the file system in gibibytes (GiB).
* `storageType` - The type of storage the file system is using. If set to `SSD`, the file system uses solid state drive storage. If set to `HDD`, the file system uses hard disk drive storage.
* `subnetIds` - Specifies the IDs of the subnets that the file system is accessible from. For the MULTI_AZ_1 file system deployment type, there are two subnet IDs, one for the preferred file server and one for the standby file server. The preferred file server subnet identified in the `preferredSubnetId` property.
* `tags` - The tags associated with the file system.
* `throughputCapacity` - The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps). If the file system uses multiple HA pairs this will equal throuthput_capacity_per_ha_pair x ha_pairs
* `throughputCapacityPerHaPair` - The sustained throughput of each HA pair for an Amazon FSx file system in Megabytes per second (MBps).
* `vpcId` - The ID of the primary virtual private cloud (VPC) for the file system.
* `weeklyMaintenanceStartTime` - The preferred start time (in `D:HH:MM` format) to perform weekly maintenance, in the UTC time zone.

### Disk IOPS

* `iops` - The total number of SSD IOPS provisioned for the file system.
* `mode` - Specifies whether the file system is using the `AUTOMATIC` setting of SSD IOPS of 3 IOPS per GB of storage capacity, or if it using a `USER_PROVISIONED` value.

### File System Endpoints

* `intercluster` - A FileSystemEndpoint for managing your file system by setting up NetApp SnapMirror with other ONTAP systems. See [FileSystemEndpoint](#file-system-endpoint) below.
* `management` - A FileSystemEndpoint for managing your file system using the NetApp ONTAP CLI and NetApp ONTAP API. See [FileSystemEndpoint](#file-system-endpoint) below.

### File System Endpoint

* `DNSName` - The file system's DNS name. You can mount your file system using its DNS name.
* `IpAddresses` - IP addresses of the file system endpoint.

<!-- cache-key: cdktf-0.20.8 input-17c51b20fd22b46b4f19504556832c4775ff8652c471859ecb2278f8af4f3f81 -->