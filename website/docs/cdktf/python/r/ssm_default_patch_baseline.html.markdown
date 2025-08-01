---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_default_patch_baseline"
description: |-
  Terraform resource for managing an AWS Systems Manager Default Patch Baseline.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_default_patch_baseline

Terraform resource for registering an AWS Systems Manager Default Patch Baseline.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssm_default_patch_baseline import SsmDefaultPatchBaseline
from imports.aws.ssm_patch_baseline import SsmPatchBaseline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = SsmPatchBaseline(self, "example",
            approved_patches=["KB123456"],
            name="example"
        )
        aws_ssm_default_patch_baseline_example = SsmDefaultPatchBaseline(self, "example_1",
            baseline_id=example.id,
            operating_system=example.operating_system
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssm_default_patch_baseline_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `baseline_id` - (Required) ID of the patch baseline.
  Can be an ID or an ARN.
  When specifying an AWS-provided patch baseline, must be the ARN.
* `operating_system` - (Required) The operating system the patch baseline applies to.
  Valid values are
  `AMAZON_LINUX`,
  `AMAZON_LINUX_2`,
  `AMAZON_LINUX_2022`,
  `CENTOS`,
  `DEBIAN`,
  `MACOS`,
  `ORACLE_LINUX`,
  `RASPBIAN`,
  `REDHAT_ENTERPRISE_LINUX`,
  `ROCKY_LINUX`,
  `SUSE`,
  `UBUNTU`, and
  `WINDOWS`.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the Systems Manager Default Patch Baseline using the patch baseline ID, patch baseline ARN, or the operating system value. For example:

Using the patch baseline ID:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssm_default_patch_baseline import SsmDefaultPatchBaseline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsmDefaultPatchBaseline.generate_config_for_import(self, "example", "pb-1234567890abcdef1")
```

Using the patch baseline ARN:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssm_default_patch_baseline import SsmDefaultPatchBaseline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsmDefaultPatchBaseline.generate_config_for_import(self, "example", "arn:aws:ssm:us-west-2:123456789012:patchbaseline/pb-1234567890abcdef1")
```

Using the operating system value:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssm_default_patch_baseline import SsmDefaultPatchBaseline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsmDefaultPatchBaseline.generate_config_for_import(self, "example", "CENTOS")
```

**Using `terraform import` to import** the Systems Manager Default Patch Baseline using the patch baseline ID, patch baseline ARN, or the operating system value. For example:

Using the patch baseline ID:

```console
% terraform import aws_ssm_default_patch_baseline.example pb-1234567890abcdef1
```

Using the patch baseline ARN:

```console
% terraform import aws_ssm_default_patch_baseline.example arn:aws:ssm:us-west-2:123456789012:patchbaseline/pb-1234567890abcdef1
```

Using the operating system value:

```console
% terraform import aws_ssm_default_patch_baseline.example CENTOS
```

<!-- cache-key: cdktf-0.20.8 input-da392fd56f381716ef47e028e9f89a0c63aaa0d943becc4823c05a5044b1dbcf -->