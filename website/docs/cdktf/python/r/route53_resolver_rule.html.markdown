---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_rule"
description: |-
  Provides a Route53 Resolver rule.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_route53_resolver_rule

Provides a Route53 Resolver rule.

## Example Usage

### System rule

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_rule import Route53ResolverRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ResolverRule(self, "sys",
            domain_name="subdomain.example.com",
            rule_type="SYSTEM"
        )
```

### Forward rule

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_rule import Route53ResolverRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ResolverRule(self, "fwd",
            domain_name="example.com",
            name="example",
            resolver_endpoint_id=foo.id,
            rule_type="FORWARD",
            tags={
                "Environment": "Prod"
            },
            target_ip=[Route53ResolverRuleTargetIp(
                ip="123.45.67.89"
            )
            ]
        )
```

### IPv6 Forward rule

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_rule import Route53ResolverRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ResolverRule(self, "fwd",
            domain_name="example.com",
            name="example",
            resolver_endpoint_id=foo.id,
            rule_type="FORWARD",
            tags={
                "Environment": "Prod"
            },
            target_ip=[Route53ResolverRuleTargetIp(
                ipv6="2600:1f18:1686:2000:4e60:6e3e:258:da36"
            )
            ]
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `domain_name` - (Required) DNS queries for this domain name are forwarded to the IP addresses that are specified using `target_ip`.
* `rule_type` - (Required) Rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
* `name` - (Optional) Friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
* `resolver_endpoint_id` (Optional) ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `target_ip`.
This argument should only be specified for `FORWARD` type rules.
* `target_ip` - (Optional) Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
This argument should only be specified for `FORWARD` type rules.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

The `target_ip` object supports the following:

* `ip` - (Optional) One IPv4 address that you want to forward DNS queries to.
* `ipv6` - (Optional) One IPv6 address that you want to forward DNS queries to.
* `port` - (Optional) Port at `ip` that you want to forward DNS queries to. Default value is `53`.
* `protocol` - (Optional) Protocol for the resolver endpoint. Valid values can be found in the [AWS documentation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_TargetAddress.html). Default value is `Do53`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the resolver rule.
* `arn` - ARN (Amazon Resource Name) for the resolver rule.
* `owner_id` - When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
* `share_status` - Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Route53 Resolver rules using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_rule import Route53ResolverRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ResolverRule.generate_config_for_import(self, "sys", "rslvr-rr-0123456789abcdef0")
```

Using `terraform import`, import Route53 Resolver rules using the `id`. For example:

```console
% terraform import aws_route53_resolver_rule.sys rslvr-rr-0123456789abcdef0
```

<!-- cache-key: cdktf-0.20.8 input-f7db1efb6857f8ae80e97adfdfe925c84a0260d4a051f989209b82bb2fa2684a -->