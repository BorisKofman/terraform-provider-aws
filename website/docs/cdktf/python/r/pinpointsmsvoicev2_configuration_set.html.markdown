---
subcategory: "End User Messaging SMS"
layout: "aws"
page_title: "AWS: aws_pinpointsmsvoicev2_configuration_set"
description: |-
  Manages an AWS End User Messaging SMS Configuration Set.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_pinpointsmsvoicev2_configuration_set

Manages an AWS End User Messaging SMS Configuration Set.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.pinpointsmsvoicev2_configuration_set import Pinpointsmsvoicev2ConfigurationSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Pinpointsmsvoicev2ConfigurationSet(self, "example",
            default_message_type="TRANSACTIONAL",
            default_sender_id="example",
            name="example-configuration-set"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the configuration set.
* `default_sender_id` - (Optional) The default sender ID to use for this configuration set.
* `default_message_type` - (Optional) The default message type. Must either be "TRANSACTIONAL" or "PROMOTIONAL"
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the configuration set.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import -out lists using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.pinpointsmsvoicev2_configuration_set import Pinpointsmsvoicev2ConfigurationSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Pinpointsmsvoicev2ConfigurationSet.generate_config_for_import(self, "example", "example-configuration-set")
```

Using `terraform import`, import configuration sets using the `name`. For example:

```console
% terraform import aws_pinpointsmsvoicev2_configuration_set.example example-configuration-set
```

<!-- cache-key: cdktf-0.20.8 input-14d675df8a71afe20e7eacfb81a406c4347ed2d6c5e2e19f6ed572a177d80a25 -->