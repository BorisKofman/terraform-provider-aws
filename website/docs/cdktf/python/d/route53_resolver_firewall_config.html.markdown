---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_firewall_config"
description: |-
    Provides details about a specific a Route 53 Resolver DNS Firewall config.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_route53_resolver_firewall_config

`aws_route53_resolver_firewall_config` provides details about a specific a Route 53 Resolver DNS Firewall config.

This data source allows to find a details about a specific a Route 53 Resolver DNS Firewall config.

## Example Usage

The following example shows how to get a firewall config using the VPC ID.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_route53_resolver_firewall_config import DataAwsRoute53ResolverFirewallConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsRoute53ResolverFirewallConfig(self, "example",
            resource_id="vpc-exampleid"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resource_id` - (Required) The ID of the VPC from Amazon VPC that the configuration is for.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `firewall_fail_open` - Determines how DNS Firewall operates during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply.
* `id` - The ID of the firewall configuration.
* `owner_id` - The Amazon Web Services account ID of the owner of the VPC that this firewall configuration applies to.

<!-- cache-key: cdktf-0.20.8 input-52d397caa6181334b1853429f784ba8bc2191ba6468885f3c3dc99c80f7f85d7 -->