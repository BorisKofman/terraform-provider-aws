---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_nat_gateway"
description: |-
  Provides a resource to create a VPC NAT Gateway.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_nat_gateway

Provides a resource to create a VPC NAT Gateway.

!> **WARNING:** You should not use the `aws_nat_gateway` resource that has `secondary_allocation_ids` in conjunction with an [`aws_nat_gateway_eip_association`](nat_gateway_eip_association.html) resource. Doing so may cause perpetual differences, and result in associations being overwritten.

## Example Usage

### Public NAT

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.nat_gateway import NatGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NatGateway(self, "example",
            allocation_id=Token.as_string(aws_eip_example.id),
            depends_on=[aws_internet_gateway_example],
            subnet_id=Token.as_string(aws_subnet_example.id),
            tags={
                "Name": "gw NAT"
            }
        )
```

### Public NAT with Secondary Private IP Addresses

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.nat_gateway import NatGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NatGateway(self, "example",
            allocation_id=Token.as_string(aws_eip_example.id),
            secondary_allocation_ids=[secondary.id],
            secondary_private_ip_addresses=["10.0.1.5"],
            subnet_id=Token.as_string(aws_subnet_example.id)
        )
```

### Private NAT

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.nat_gateway import NatGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NatGateway(self, "example",
            connectivity_type="private",
            subnet_id=Token.as_string(aws_subnet_example.id)
        )
```

### Private NAT with Secondary Private IP Addresses

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.nat_gateway import NatGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NatGateway(self, "example",
            connectivity_type="private",
            secondary_private_ip_address_count=7,
            subnet_id=Token.as_string(aws_subnet_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `allocation_id` - (Optional) The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivity_type` of `public`.
* `connectivity_type` - (Optional) Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
* `private_ip` - (Optional) The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `subnet_id` - (Required) The Subnet ID of the subnet in which to place the NAT Gateway.
* `secondary_allocation_ids` - (Optional) A list of secondary allocation EIP IDs for this NAT Gateway. To remove all secondary allocations an empty list should be specified.
* `secondary_private_ip_address_count` - (Optional) [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
* `secondary_private_ip_addresses` - (Optional) A list of secondary private IPv4 addresses to assign to the NAT Gateway. To remove all secondary private addresses an empty list should be specified.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `association_id` - The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivity_type` is `public`.
* `id` - The ID of the NAT Gateway.
* `network_interface_id` - The ID of the network interface associated with the NAT Gateway.
* `public_ip` - The Elastic IP address associated with the NAT Gateway.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `update` - (Default `10m`)
- `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import NAT Gateways using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.nat_gateway import NatGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NatGateway.generate_config_for_import(self, "privateGw", "nat-05dba92075d71c408")
```

Using `terraform import`, import NAT Gateways using the `id`. For example:

```console
% terraform import aws_nat_gateway.private_gw nat-05dba92075d71c408
```

<!-- cache-key: cdktf-0.20.8 input-684318b90065d6b0158a7a90514e79167993b7abe86500e585533d0c98ecee64 -->