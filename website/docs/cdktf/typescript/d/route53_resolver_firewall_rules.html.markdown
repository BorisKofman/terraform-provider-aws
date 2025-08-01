---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_firewall_rules"
description: |-
    Provides details about rules in a specific Route53 Resolver Firewall rule group.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_route53_resolver_firewall_rules

`aws_route53_resolver_firewall_rules` Provides details about rules in a specific Route53 Resolver Firewall rule group.

## Example Usage

The following example shows how to get Route53 Resolver Firewall rules based on its associated firewall group id.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRoute53ResolverFirewallRules } from "./.gen/providers/aws/data-aws-route53-resolver-firewall-rules";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRoute53ResolverFirewallRules(this, "example", {
      firewallRuleGroupId: Token.asString(
        awsRoute53ResolverFirewallRuleGroupExample.id
      ),
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `firewallRuleGroupId` - (Required) The unique identifier of the firewall rule group that you want to retrieve the rules for.
* `action` - (Optional) The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list.
* `priority` - (Optional) The setting that determines the processing order of the rules in a rule group.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `firewallRules` - List with information about the firewall rules. See details below.

### provisioning_artifact_details

* `blockOverrideDnsType` - The DNS record's type.
* `blockOverrideDomain` - The custom DNS record to send back in response to the query.
* `blockOverrideTtl` - The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
* `blockResponse` - The way that you want DNS Firewall to block the request.
* `creationTime` - The date and time that the rule was created, in Unix time format and Coordinated Universal Time (UTC).
* `creatorRequestId` - A unique string defined by you to identify the request.
* `firewallDomainListId` - The ID of the domain list that's used in the rule.
* `modificationTime` - The date and time that the rule was last modified, in Unix time format and Coordinated Universal Time (UTC).
* `name` - The name of the rule.

<!-- cache-key: cdktf-0.20.8 input-5ada4df21cd5a4fef0f5b602cfec7f0ae8ac12bcdc7c3e26d4c269400414e394 -->