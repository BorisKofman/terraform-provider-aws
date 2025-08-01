---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_endpoint_route_table_association"
description: |-
  Manages a VPC Endpoint Route Table Association
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_endpoint_route_table_association

Manages a VPC Endpoint Route Table Association

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_endpoint_route_table_association import VpcEndpointRouteTableAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcEndpointRouteTableAssociation(self, "example",
            route_table_id=Token.as_string(aws_route_table_example.id),
            vpc_endpoint_id=Token.as_string(aws_vpc_endpoint_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `route_table_id` - (Required) Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
* `vpc_endpoint_id` - (Required) Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A hash of the EC2 Route Table and VPC Endpoint identifiers.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPC Endpoint Route Table Associations using `vpc_endpoint_id` together with `route_table_id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_endpoint_route_table_association import VpcEndpointRouteTableAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcEndpointRouteTableAssociation.generate_config_for_import(self, "example", "vpce-aaaaaaaa/rtb-bbbbbbbb")
```

Using `terraform import`, import VPC Endpoint Route Table Associations using `vpc_endpoint_id` together with `route_table_id`. For example:

```console
% terraform import aws_vpc_endpoint_route_table_association.example vpce-aaaaaaaa/rtb-bbbbbbbb
```

<!-- cache-key: cdktf-0.20.8 input-642bffdb6fac82c79353d54280251de2b233ae4e54c6477e8500b23989e28b95 -->