---
subcategory: "WorkSpaces"
layout: "aws"
page_title: "AWS: aws_workspaces_workspace"
description: |-
  Provides a workspaces in AWS Workspaces Service.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_workspaces_workspace

Provides a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service

~> **NOTE:** AWS WorkSpaces service requires [`workspaces_DefaultRole`](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role) IAM role to operate normally.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsKmsKey } from "./.gen/providers/aws/data-aws-kms-key";
import { DataAwsWorkspacesBundle } from "./.gen/providers/aws/data-aws-workspaces-bundle";
import { WorkspacesWorkspace } from "./.gen/providers/aws/workspaces-workspace";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const workspaces = new DataAwsKmsKey(this, "workspaces", {
      keyId: "alias/aws/workspaces",
    });
    const valueWindows10 = new DataAwsWorkspacesBundle(
      this,
      "value_windows_10",
      {
        bundleId: "wsb-bh8rsxt14",
      }
    );
    new WorkspacesWorkspace(this, "example", {
      bundleId: Token.asString(valueWindows10.id),
      directoryId: Token.asString(awsWorkspacesDirectoryExample.id),
      rootVolumeEncryptionEnabled: true,
      tags: {
        Department: "IT",
      },
      userName: "john.doe",
      userVolumeEncryptionEnabled: true,
      volumeEncryptionKey: Token.asString(workspaces.arn),
      workspaceProperties: {
        computeTypeName: "VALUE",
        rootVolumeSizeGib: 80,
        runningMode: "AUTO_STOP",
        runningModeAutoStopTimeoutInMinutes: 60,
        userVolumeSizeGib: 10,
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `directoryId` - (Required) The ID of the directory for the WorkSpace.
* `bundleId` - (Required) The ID of the bundle for the WorkSpace.
* `userName` - (Required) The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
* `rootVolumeEncryptionEnabled` - (Optional) Indicates whether the data stored on the root volume is encrypted.
* `userVolumeEncryptionEnabled` - (Optional) Indicates whether the data stored on the user volume is encrypted.
* `volumeEncryptionKey` - (Optional) The ARN of a symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
* `tags` - (Optional) The tags for the WorkSpace. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `workspaceProperties` - (Optional) The WorkSpace properties.

`workspaceProperties` supports the following:

* `computeTypeName` - (Optional) The compute type. For more information, see [Amazon WorkSpaces Bundles](http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles). Valid values are `VALUE`, `STANDARD`, `PERFORMANCE`, `POWER`, `GRAPHICS`, `POWERPRO`, `GRAPHICSPRO`, `GRAPHICS_G4DN`, and `GRAPHICSPRO_G4DN`.
* `rootVolumeSizeGib` - (Optional) The size of the root volume.
* `runningMode` - (Optional) The running mode. For more information, see [Manage the WorkSpace Running Mode](https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html). Valid values are `AUTO_STOP` and `ALWAYS_ON`.
* `runningModeAutoStopTimeoutInMinutes` - (Optional) The time after a user logs off when WorkSpaces are automatically stopped. Configured in 60-minute intervals.
* `userVolumeSizeGib` - (Optional) The size of the user storage.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The workspaces ID.
* `ipAddress` - The IP address of the WorkSpace.
* `computerName` - The name of the WorkSpace, as seen by the operating system.
* `state` - The operational state of the WorkSpace.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `30m`)
- `update` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Workspaces using their ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WorkspacesWorkspace } from "./.gen/providers/aws/workspaces-workspace";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    WorkspacesWorkspace.generateConfigForImport(
      this,
      "example",
      "ws-9z9zmbkhv"
    );
  }
}

```

Using `terraform import`, import Workspaces using their ID. For example:

```console
% terraform import aws_workspaces_workspace.example ws-9z9zmbkhv
```

<!-- cache-key: cdktf-0.20.8 input-05f0cf9c25d811468de02289c29cb691ed2a953b366e6ce0167d6960201b6dbb -->