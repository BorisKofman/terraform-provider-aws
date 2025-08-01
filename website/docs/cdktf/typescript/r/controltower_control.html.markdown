---
subcategory: "Control Tower"
layout: "aws"
page_title: "AWS: aws_controltower_control"
description: |-
  Allows the application of pre-defined controls to organizational units.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_controltower_control

Allows the application of pre-defined controls to organizational units. For more information on usage, please see the
[AWS Control Tower User Guide](https://docs.aws.amazon.com/controltower/latest/userguide/enable-guardrails.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ControltowerControl } from "./.gen/providers/aws/controltower-control";
import { DataAwsOrganizationsOrganization } from "./.gen/providers/aws/data-aws-organizations-organization";
import { DataAwsOrganizationsOrganizationalUnits } from "./.gen/providers/aws/data-aws-organizations-organizational-units";
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsOrganizationsOrganization(this, "example", {});
    const dataAwsOrganizationsOrganizationalUnitsExample =
      new DataAwsOrganizationsOrganizationalUnits(this, "example_1", {
        parentId: Token.asString(Fn.lookupNested(example.roots, ["0", "id"])),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsOrganizationsOrganizationalUnitsExample.overrideLogicalId("example");
    const current = new DataAwsRegion(this, "current", {});
    const awsControltowerControlExample = new ControltowerControl(
      this,
      "example_3",
      {
        controlIdentifier:
          "arn:aws:controltower:${" +
          current.region +
          "}::control/AWS-GR_EC2_VOLUME_INUSE_CHECK",
        parameters: [
          {
            key: "AllowedRegions",
            value: Token.asString(Fn.jsonencode(["us-east-1"])),
          },
        ],
        targetIdentifier: Token.asString(
          Fn.lookupNested(
            "${[ for x in ${" +
              dataAwsOrganizationsOrganizationalUnitsExample.children +
              '} : x.arn if x.name == "Infrastructure"]}',
            ["0"]
          )
        ),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsControltowerControlExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `controlIdentifier` - (Required) The ARN of the control. Only Strongly recommended and Elective controls are permitted, with the exception of the Region deny guardrail.
* `targetIdentifier` - (Required) The ARN of the organizational unit.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `parameters` - (Optional) Parameter values which are specified to configure the control when you enable it. See [Parameters](#parameters) for more details.

### Parameters

* `key` - (Required) The name of the parameter.
* `value` - (Required) The value of the parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the EnabledControl resource.
* `id` - The ARN of the organizational unit.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Control Tower Controls using their `organizational_unit_arn,control_identifier`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ControltowerControl } from "./.gen/providers/aws/controltower-control";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ControltowerControl.generateConfigForImport(
      this,
      "example",
      "arn:aws:organizations::123456789101:ou/o-qqaejywet/ou-qg5o-ufbhdtv3,arn:aws:controltower:us-east-1::control/WTDSMKDKDNLE"
    );
  }
}

```

Using `terraform import`, import Control Tower Controls using their `organizational_unit_arn/control_identifier`. For example:

```console
% terraform import aws_controltower_control.example arn:aws:organizations::123456789101:ou/o-qqaejywet/ou-qg5o-ufbhdtv3,arn:aws:controltower:us-east-1::control/WTDSMKDKDNLE
```

<!-- cache-key: cdktf-0.20.8 input-99f2dc97ffce7147a933e420c5b996d742c6d4b37110b3de09ef8f78c33f86a0 -->