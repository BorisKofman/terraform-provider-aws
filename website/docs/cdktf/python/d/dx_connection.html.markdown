---
subcategory: "Direct Connect"
layout: "aws"
page_title: "AWS: aws_dx_connection"
description: |-
  Retrieve information about a Direct Connect Connection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_dx_connection

Retrieve information about a Direct Connect Connection.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_dx_connection import DataAwsDxConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsDxConnection(self, "example",
            name="tf-dx-connection"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the connection to retrieve.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the connection.
* `aws_device` - Direct Connect endpoint on which the physical connection terminates.
* `bandwidth` - Bandwidth of the connection.
* `id` - ID of the connection.
* `location` - AWS Direct Connect location where the connection is located.
* `owner_account_id` - ID of the AWS account that owns the connection.
* `partner_name` - The name of the AWS Direct Connect service provider associated with the connection.
* `provider_name` - Name of the service provider associated with the connection.
* `state` - State of the connection.
* `tags` - Map of tags for the resource.
* `vlan_id` - The VLAN ID.

<!-- cache-key: cdktf-0.20.8 input-fc4168b2991511014961b38b3b097d7e1ee715f0f31ac7e2d0b2a0992ca4dac4 -->