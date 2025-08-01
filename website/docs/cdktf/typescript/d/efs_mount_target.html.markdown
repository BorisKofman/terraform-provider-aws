---
subcategory: "EFS (Elastic File System)"
layout: "aws"
page_title: "AWS: aws_efs_mount_target"
description: |-
  Provides an Elastic File System Mount Target (EFS) data source.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_efs_mount_target

Provides information about an Elastic File System Mount Target (EFS).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { VariableType, TerraformVariable, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEfsMountTarget } from "./.gen/providers/aws/data-aws-efs-mount-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const mountTargetId = new TerraformVariable(this, "mount_target_id", {
      default: "",
      type: VariableType.STRING,
    });
    new DataAwsEfsMountTarget(this, "by_id", {
      mountTargetId: mountTargetId.stringValue,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `accessPointId` - (Optional) ID or ARN of the access point whose mount target that you want to find. It must be included if a `fileSystemId` and `mountTargetId` are not included.
* `fileSystemId` - (Optional) ID or ARN of the file system whose mount target that you want to find. It must be included if an `accessPointId` and `mountTargetId` are not included.
* `mountTargetId` - (Optional) ID or ARN of the mount target that you want to find. It must be included in your request if an `accessPointId` and `fileSystemId` are not included.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `fileSystemArn` - Amazon Resource Name of the file system for which the mount target is intended.
* `subnetId` - ID of the mount target's subnet.
* `ipAddress` - Address at which the file system may be mounted via the mount target.
* `securityGroups` - List of VPC security group IDs attached to the mount target.
* `dnsName` - DNS name for the EFS file system.
* `mountTargetDnsName` - The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
* `networkInterfaceId` - The ID of the network interface that Amazon EFS created when it created the mount target.
* `availabilityZoneName` - The name of the Availability Zone (AZ) that the mount target resides in.
* `availabilityZoneId` - The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
* `ownerId` - AWS account ID that owns the resource.

<!-- cache-key: cdktf-0.20.8 input-a4a76136977d94530467ac347e3eb3dcdee34d2341329168d290bff5f94117fc -->