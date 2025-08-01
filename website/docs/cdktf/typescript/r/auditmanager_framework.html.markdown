---
subcategory: "Audit Manager"
layout: "aws"
page_title: "AWS: aws_auditmanager_framework"
description: |-
  Terraform resource for managing an AWS Audit Manager Framework.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_auditmanager_framework

Terraform resource for managing an AWS Audit Manager Framework.

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
import { AuditmanagerFramework } from "./.gen/providers/aws/auditmanager-framework";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AuditmanagerFramework(this, "test", {
      controlSets: [
        {
          controls: [
            {
              id: test1.id,
            },
            {
              id: test2.id,
            },
          ],
          name: "example",
        },
      ],
      name: "example",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the framework.
* `controlSets` - (Required) Configuration block(s) for the control sets that are associated with the framework. See [`controlSets` Block](#control_sets-block) below for details.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `complianceType` - (Optional) Compliance type that the new custom framework supports, such as `CIS` or `HIPAA`.
* `description` - (Optional) Description of the framework.
* `tags` - (Optional) A map of tags to assign to the framework. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `controlSets` Block

The `controlSets` configuration block supports the following arguments:

* `name` - (Required) Name of the control set.
* `controls` - (Required) Configuration block(s) for the controls within the control set. See [`controls` Block](#controls-block) below for details.

### `controls` Block

The `controls` configuration block supports the following arguments:

* `id` - (Required) Unique identifier of the control.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the framework.
* `control_sets[*].id` - Unique identifier for the framework control set.
* `id` - Unique identifier for the framework.
* `frameworkType` - Framework type, such as a custom framework or a standard framework.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Audit Manager Framework using the framework `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AuditmanagerFramework } from "./.gen/providers/aws/auditmanager-framework";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AuditmanagerFramework.generateConfigForImport(
      this,
      "example",
      "abc123-de45"
    );
  }
}

```

Using `terraform import`, import Audit Manager Framework using the framework `id`. For example:

```console
% terraform import aws_auditmanager_framework.example abc123-de45
```

<!-- cache-key: cdktf-0.20.8 input-4fbba115e5fa9a433093d92d64cd1e11bd17ef8ec02771d17f0c49debceebcaa -->