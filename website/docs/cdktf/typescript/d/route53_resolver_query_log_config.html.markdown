---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_query_log_config"
description: |-
  Provides details about a specific Route53 Resolver Query Logging Configuration.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_route53_resolver_query_log_config

`aws_route53_resolver_query_log_config` provides details about a specific Route53 Resolver Query Logging Configuration.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRoute53ResolverQueryLogConfig } from "./.gen/providers/aws/data-aws-route53-resolver-query-log-config";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRoute53ResolverQueryLogConfig(this, "example", {
      resolverQueryLogConfigId: "rqlc-1abc2345ef678g91h",
    });
  }
}

```

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRoute53ResolverQueryLogConfig } from "./.gen/providers/aws/data-aws-route53-resolver-query-log-config";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRoute53ResolverQueryLogConfig(this, "example", {
      filter: [
        {
          name: "Name",
          values: ["shared-query-log-config"],
        },
        {
          name: "ShareStatus",
          values: ["SHARED_WITH_ME"],
        },
      ],
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resolverQueryLogConfigId` - (Optional) ID of the Route53 Resolver Query Logging Configuration.
* `filter` - (Optional) One or more name/value pairs to use as filters. There are
several valid keys, for a full reference, check out
[Route53resolver Filter value in the AWS API reference][1].

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - The ID for the query logging configuration.
* `arn` - Computed ARN of the Route53 Resolver Query Logging Configuration.
* `destinationArn` - The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group or a Kinesis Data Firehose delivery stream.
* `name` - The name of the query logging configuration.
* `ownerId` - The AWS account ID for the account that created the query logging configuration.
* `shareStatus` - An indication of whether the query logging configuration is shared with other AWS accounts or was shared with the current account by another AWS account.
* `tags` - Map of tags to assign to the service.

[1]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_Filter.html

<!-- cache-key: cdktf-0.20.8 input-caea4099ca24d011a0e3ff384fbd5c0efa2c1ef75935cfe3c4e0ac56cb99845a -->