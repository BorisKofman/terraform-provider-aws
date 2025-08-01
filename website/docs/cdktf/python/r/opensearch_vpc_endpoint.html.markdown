---
subcategory: "OpenSearch"
layout: "aws"
page_title: "AWS: aws_opensearch_vpc_endpoint"
description: |-
  Terraform resource for managing an AWS OpenSearch VPC Endpoint.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_opensearch_vpc_endpoint

Manages an [AWS Opensearch VPC Endpoint](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_CreateVpcEndpoint.html). Creates an Amazon OpenSearch Service-managed VPC endpoint.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.opensearch_vpc_endpoint import OpensearchVpcEndpoint
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OpensearchVpcEndpoint(self, "foo",
            domain_arn=domain1.arn,
            vpc_options=OpensearchVpcEndpointVpcOptions(
                security_group_ids=[test.id, test2.id],
                subnet_ids=[
                    Token.as_string(aws_subnet_test.id),
                    Token.as_string(aws_subnet_test2.id)
                ]
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `domain_arn` - (Required, Forces new resource) Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
* `vpc_options` - (Required) Options to specify the subnets and security groups for the endpoint.

### vpc_options

* `security_group_ids` - (Optional) The list of security group IDs associated with the VPC endpoints for the domain. If you do not provide a security group ID, OpenSearch Service uses the default security group for the VPC.
* `subnet_ids` - (Required) A list of subnet IDs associated with the VPC endpoints for the domain. If your domain uses multiple Availability Zones, you need to provide two subnet IDs, one per zone. Otherwise, provide only one.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The unique identifier of the endpoint.
* `endpoint` - The connection endpoint ID for connecting to the domain.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)
* `update` - (Default `60m`)
* `delete` - (Default `90m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OpenSearch VPC endpoint connections using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws. import OpensearchVpcEndpointConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OpensearchVpcEndpointConnection.generate_config_for_import(self, "example", "endpoint-id")
```

Using `terraform import`, import OpenSearch VPC endpoint connections using the `id`. For example:

```console
% terraform import aws_opensearch_vpc_endpoint_connection.example endpoint-id
```

<!-- cache-key: cdktf-0.20.8 input-5a744203f3493a60bc75c408ac5c67446963991f758c602973d17bdc5aea9324 -->