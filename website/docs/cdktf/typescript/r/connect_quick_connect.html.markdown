---
subcategory: "Connect"
layout: "aws"
page_title: "AWS: aws_connect_quick_connect"
description: |-
  Provides details about a specific Amazon Quick Connect
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_connect_quick_connect

Provides an Amazon Connect Quick Connect resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ConnectQuickConnect } from "./.gen/providers/aws/connect-quick-connect";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ConnectQuickConnect(this, "test", {
      description: "quick connect phone number",
      instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
      name: "Example Name",
      quickConnectConfig: {
        phoneConfig: [
          {
            phoneNumber: "+12345678912",
          },
        ],
        quickConnectType: "PHONE_NUMBER",
      },
      tags: {
        Name: "Example Quick Connect",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Specifies the description of the Quick Connect.
* `instanceId` - (Required) Specifies the identifier of the hosting Amazon Connect Instance.
* `name` - (Required) Specifies the name of the Quick Connect.
* `quickConnectConfig` - (Required) A block that defines the configuration information for the Quick Connect: `quickConnectType` and one of `phoneConfig`, `queueConfig`, `userConfig` . The Quick Connect Config block is documented below.
* `tags` - (Optional) Tags to apply to the Quick Connect. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

A `quickConnectConfig` block supports the following arguments:

* `quickConnectType` - (Required) Specifies the configuration type of the quick connect. valid values are `PHONE_NUMBER`, `QUEUE`, `USER`.
* `phoneConfig` - (Optional) Specifies the phone configuration of the Quick Connect. This is required only if `quickConnectType` is `PHONE_NUMBER`. The `phoneConfig` block is documented below.
* `queueConfig` - (Optional) Specifies the queue configuration of the Quick Connect. This is required only if `quickConnectType` is `QUEUE`. The `queueConfig` block is documented below.
* `userConfig` - (Optional) Specifies the user configuration of the Quick Connect. This is required only if `quickConnectType` is `USER`. The `userConfig` block is documented below.

A `phoneConfig` block supports the following arguments:

* `phoneNumber` - (Required) Specifies the phone number in in E.164 format.

A `queueConfig` block supports the following arguments:

* `contactFlowId` - (Required) Specifies the identifier of the contact flow.
* `queueId` - (Required) Specifies the identifier for the queue.

A `userConfig` block supports the following arguments:

* `contactFlowId` - (Required) Specifies the identifier of the contact flow.
* `userId` - (Required) Specifies the identifier for the user.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the Quick Connect.
* `quickConnectId` - The identifier for the Quick Connect.
* `id` - The identifier of the hosting Amazon Connect Instance and identifier of the Quick Connect separated by a colon (`:`).
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Amazon Connect Quick Connects using the `instanceId` and `quickConnectId` separated by a colon (`:`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ConnectQuickConnect } from "./.gen/providers/aws/connect-quick-connect";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ConnectQuickConnect.generateConfigForImport(
      this,
      "example",
      "f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5"
    );
  }
}

```

Using `terraform import`, import Amazon Connect Quick Connects using the `instanceId` and `quickConnectId` separated by a colon (`:`). For example:

```console
% terraform import aws_connect_quick_connect.example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```

<!-- cache-key: cdktf-0.20.8 input-9f7839e35ec969fabb5d333a0754d70757944b1bbf4000bd41318d7c8cafc097 -->