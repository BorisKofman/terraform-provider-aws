---
subcategory: "CloudWatch Observability Access Manager"
layout: "aws"
page_title: "AWS: aws_oam_link"
description: |-
  Terraform data source for managing an AWS CloudWatch Observability Access Manager Link.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_oam_link

Terraform data source for managing an AWS CloudWatch Observability Access Manager Link.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsOamLink } from "./.gen/providers/aws/data-aws-oam-link";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsOamLink(this, "example", {
      linkIdentifier:
        "arn:aws:oam:us-west-1:111111111111:link/abcd1234-a123-456a-a12b-a123b456c789",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `linkIdentifier` - (Required) ARN of the link.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the link.
* `id` - ARN of the link. Use `arn` instead.
* `label` - Label that is assigned to this link.
* `labelTemplate` - Human-readable name used to identify this source account when you are viewing data from it in the monitoring account.
* `linkConfiguration` - Configuration for creating filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account. See [`linkConfiguration` Block](#link_configuration-block) for details.
* `linkId` - ID string that AWS generated as part of the link ARN.
* `resourceTypes` - Types of data that the source account shares with the monitoring account.
* `sinkArn` - ARN of the sink that is used for this link.

### `linkConfiguration` Block

The `linkConfiguration` configuration block supports the following arguments:

* `logGroupConfiguration` - Configuration for filtering which log groups are to send log events from the source account to the monitoring account. See [`logGroupConfiguration` Block](#log_group_configuration-block) for details.
* `metricConfiguration` - Configuration for filtering which metric namespaces are to be shared from the source account to the monitoring account. See [`metricConfiguration` Block](#metric_configuration-block) for details.

### `logGroupConfiguration` Block

The `logGroupConfiguration` configuration block supports the following arguments:

* `filter` - Filter string that specifies which log groups are to share their log events with the monitoring account. See [LogGroupConfiguration](https://docs.aws.amazon.com/OAM/latest/APIReference/API_LogGroupConfiguration.html) for details.

### `metricConfiguration` Block

The `metricConfiguration` configuration block supports the following arguments:

* `filter` - Filter string that specifies  which metrics are to be shared with the monitoring account. See [MetricConfiguration](https://docs.aws.amazon.com/OAM/latest/APIReference/API_MetricConfiguration.html) for details.

<!-- cache-key: cdktf-0.20.8 input-588c692473632284f33e4e0eacd5d730c0ceb61d523df67ed52dcde62289d36d -->