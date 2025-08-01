---
subcategory: "DynamoDB Accelerator (DAX)"
layout: "aws"
page_title: "AWS: aws_dax_parameter_group"
description: |-
  Provides an DAX Parameter Group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dax_parameter_group

Provides a DAX Parameter Group resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.dax_parameter_group import DaxParameterGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DaxParameterGroup(self, "example",
            name="example",
            parameters=[DaxParameterGroupParameters(
                name="query-ttl-millis",
                value="100000"
            ), DaxParameterGroupParameters(
                name="record-ttl-millis",
                value="100000"
            )
            ]
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the parameter group.
* `description` - (Optional, ForceNew) A description of the parameter group.
* `parameters` - (Optional) The parameters of the parameter group.

## parameters

`parameters` supports the following:

* `name` - (Required) The name of the parameter.
* `value` - (Required) The value for the parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the parameter group.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DAX Parameter Group using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.dax_parameter_group import DaxParameterGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DaxParameterGroup.generate_config_for_import(self, "example", "my_dax_pg")
```

Using `terraform import`, import DAX Parameter Group using the `name`. For example:

```console
% terraform import aws_dax_parameter_group.example my_dax_pg
```

<!-- cache-key: cdktf-0.20.8 input-e24f534e16d285d3b69705c60b61368fc5bcbb86119ac7fdadff64841c637882 -->