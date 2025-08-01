---
subcategory: "ElastiCache"
layout: "aws"
page_title: "AWS: aws_elasticache_parameter_group"
description: |-
  Provides an ElastiCache parameter group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_elasticache_parameter_group

Provides an ElastiCache parameter group resource.

~> **NOTE:** Attempting to remove the `reserved-memory` parameter when `family` is set to `redis2.6` or `redis2.8` may show a perpetual difference in Terraform due to an ElastiCache API limitation. Leave that parameter configured with any value to workaround the issue.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheParameterGroup } from "./.gen/providers/aws/elasticache-parameter-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheParameterGroup(this, "default", {
      family: "redis2.8",
      name: "cache-params",
      parameter: [
        {
          name: "activerehashing",
          value: "yes",
        },
        {
          name: "min-slaves-to-write",
          value: "2",
        },
      ],
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the ElastiCache parameter group.
* `family` - (Required) The family of the ElastiCache parameter group.
* `description` - (Optional) The description of the ElastiCache parameter group. Defaults to "Managed by Terraform".
* `parameter` - (Optional) A list of ElastiCache parameters to apply.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

Parameter blocks support the following:

* `name` - (Required) The name of the ElastiCache parameter.
* `value` - (Required) The value of the ElastiCache parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ElastiCache parameter group name.
* `arn` - The AWS ARN associated with the parameter group.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ElastiCache Parameter Groups using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheParameterGroup } from "./.gen/providers/aws/elasticache-parameter-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ElasticacheParameterGroup.generateConfigForImport(
      this,
      "default",
      "redis-params"
    );
  }
}

```

Using `terraform import`, import ElastiCache Parameter Groups using the `name`. For example:

```console
% terraform import aws_elasticache_parameter_group.default redis-params
```

<!-- cache-key: cdktf-0.20.8 input-7d9203044f895d86b1a7977fe072d089f9a1c38d58185976712414c53f792c4b -->