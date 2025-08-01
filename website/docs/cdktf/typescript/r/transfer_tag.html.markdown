---
subcategory: "Transfer Family"
layout: "aws"
page_title: "AWS: aws_transfer_tag"
description: |-
  Manages an individual Transfer Family resource tag
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_transfer_tag

Manages an individual Transfer Family resource tag. This resource should only be used in cases where Transfer Family resources are created outside Terraform (e.g., Servers without AWS Management Console) or the tag key has the `aws:` prefix.

~> **NOTE:** This tagging resource should not be combined with the Terraform resource for managing the parent resource. For example, using `aws_transfer_server` and `aws_transfer_tag` to manage tags of the same server will cause a perpetual difference where the `aws_transfer_server` resource will try to remove the tag being added by the `aws_transfer_tag` resource.

~> **NOTE:** This tagging resource does not use the [provider `ignoreTags` configuration](/docs/providers/aws/index.html#ignore_tags).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { TransferServer } from "./.gen/providers/aws/transfer-server";
import { TransferTag } from "./.gen/providers/aws/transfer-tag";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new TransferServer(this, "example", {
      identityProviderType: "SERVICE_MANAGED",
    });
    new TransferTag(this, "hostname", {
      key: "transfer:customHostname",
      resourceArn: example.arn,
      value: "example.com",
    });
    new TransferTag(this, "zone_id", {
      key: "transfer:route53HostedZoneId",
      resourceArn: example.arn,
      value: "/hostedzone/MyHostedZoneId",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resourceArn` - (Required) Amazon Resource Name (ARN) of the Transfer Family resource to tag.
* `key` - (Required) Tag name.
* `value` - (Required) Tag value.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Transfer Family resource identifier and key, separated by a comma (`,`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_transfer_tag` using the Transfer Family resource identifier and key, separated by a comma (`,`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { TransferTag } from "./.gen/providers/aws/transfer-tag";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    TransferTag.generateConfigForImport(
      this,
      "example",
      "arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name"
    );
  }
}

```

Using `terraform import`, import `aws_transfer_tag` using the Transfer Family resource identifier and key, separated by a comma (`,`). For example:

```console
% terraform import aws_transfer_tag.example arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name
```

<!-- cache-key: cdktf-0.20.8 input-dbff2d9c3045148139dd50b19e3b850f040e43d426f0e86074715746bb3c41ac -->