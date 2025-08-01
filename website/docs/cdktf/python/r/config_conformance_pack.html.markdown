---
subcategory: "Config"
layout: "aws"
page_title: "AWS: aws_config_conformance_pack"
description: |-
  Manages a Config Conformance Pack
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_config_conformance_pack

Manages a Config Conformance Pack. More information about this collection of Config rules and remediation actions can be found in the
[Conformance Packs](https://docs.aws.amazon.com/config/latest/developerguide/conformance-packs.html) documentation.
Sample Conformance Pack templates may be found in the
[AWS Config Rules Repository](https://github.com/awslabs/aws-config-rules/tree/master/aws-config-conformance-packs).

~> **NOTE:** The account must have a Configuration Recorder with proper IAM permissions before the Conformance Pack will
successfully create or update. See also the
[`aws_config_configuration_recorder` resource](/docs/providers/aws/r/config_configuration_recorder.html).

## Example Usage

### Template Body

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.config_conformance_pack import ConfigConformancePack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ConfigConformancePack(self, "example",
            depends_on=[aws_config_configuration_recorder_example],
            input_parameter=[ConfigConformancePackInputParameter(
                parameter_name="AccessKeysRotatedParameterMaxAccessKeyAge",
                parameter_value="90"
            )
            ],
            name="example",
            template_body="Parameters:\n  AccessKeysRotatedParameterMaxAccessKeyAge:\n    Type: String\nResources:\n  IAMPasswordPolicy:\n    Properties:\n      ConfigRuleName: IAMPasswordPolicy\n      Source:\n        Owner: AWS\n        SourceIdentifier: IAM_PASSWORD_POLICY\n    Type: AWS::Config::ConfigRule\n\n"
        )
```

### Template S3 URI

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.config_conformance_pack import ConfigConformancePack
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_object import S3Object
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="example"
        )
        aws_s3_object_example = S3Object(self, "example_1",
            bucket=example.id,
            content="Resources:\n  IAMPasswordPolicy:\n    Properties:\n      ConfigRuleName: IAMPasswordPolicy\n      Source:\n        Owner: AWS\n        SourceIdentifier: IAM_PASSWORD_POLICY\n    Type: AWS::Config::ConfigRule\n\n",
            key="example-key"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_object_example.override_logical_id("example")
        aws_config_conformance_pack_example = ConfigConformancePack(self, "example_2",
            depends_on=[aws_config_configuration_recorder_example],
            name="example",
            template_s3_uri="s3://${" + example.bucket + "}/${" + aws_s3_object_example.key + "}"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_config_conformance_pack_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required, Forces new resource) The name of the conformance pack. Must begin with a letter and contain from 1 to 256 alphanumeric characters and hyphens.
* `delivery_s3_bucket` - (Optional) Amazon S3 bucket where AWS Config stores conformance pack templates. Maximum length of 63.
* `delivery_s3_key_prefix` - (Optional) The prefix for the Amazon S3 bucket. Maximum length of 1024.
* `input_parameter` - (Optional) Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `template_body` or in the template stored in Amazon S3 if using `template_s3_uri`.
* `template_body` - (Optional, required if `template_s3_uri` is not provided) A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
* `template_s3_uri` - (Optional, required if `template_body` is not provided) Location of file, e.g., `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.

~> **Note:** If both `template_body` and `template_s3_uri` are specified, AWS Config uses the `template_s3_uri` and ignores the `template_body`.

### input_parameter Argument Reference

The `input_parameter` configuration block supports the following arguments:

* `parameter_name` - (Required) The input key.
* `parameter_value` - (Required) The input value.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the conformance pack.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Config Conformance Packs using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.config_conformance_pack import ConfigConformancePack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ConfigConformancePack.generate_config_for_import(self, "example", "example")
```

Using `terraform import`, import Config Conformance Packs using the `name`. For example:

```console
% terraform import aws_config_conformance_pack.example example
```

<!-- cache-key: cdktf-0.20.8 input-8f440865510b05bee3086dfd9b75b90d596a862eaf06570b709426526910546a -->