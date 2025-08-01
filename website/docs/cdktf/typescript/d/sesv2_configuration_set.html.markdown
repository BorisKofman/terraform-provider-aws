---
subcategory: "SESv2 (Simple Email V2)"
layout: "aws"
page_title: "AWS: aws_sesv2_configuration_set"
description: |-
  Terraform data source for managing an AWS SESv2 (Simple Email V2) Configuration Set.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_sesv2_configuration_set

Terraform data source for managing an AWS SESv2 (Simple Email V2) Configuration Set.

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
import { DataAwsSesv2ConfigurationSet } from "./.gen/providers/aws/data-aws-sesv2-configuration-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsSesv2ConfigurationSet(this, "example", {
      configurationSetName: "example",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `configurationSetName` - (Required) The name of the configuration set.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `deliveryOptions` - An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
    * `maxDeliverySeconds` - The maximum amount of time, in seconds, that Amazon SES API v2 will attempt delivery of email. If specified, the value must greater than or equal to 300 seconds (5 minutes) and less than or equal to 50400 seconds (840 minutes).
    * `sendingPoolName` - The name of the dedicated IP pool to associate with the configuration set.
    * `tlsPolicy` - Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
* `reputationOptions` - An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
    * `lastFreshStart` - The date and time (in Unix time) when the reputation metrics were last given a fresh start.
    * `reputationMetricsEnabled` - Specifies whether tracking of reputation metrics is enabled.
* `sendingOptions` - An object that defines whether or not Amazon SES can send email that you send using the configuration set.
    * `sendingEnabled` - Specifies whether email sending is enabled.
* `suppressionOptions` - An object that contains information about the suppression list preferences for your account.
    * `suppressedReasons` - A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
* `tags` - Key-value map of resource tags for the container recipe.
* `trackingOptions` - An object that defines the open and click tracking options for emails that you send using the configuration set.
    * `customRedirectDomain` - The domain to use for tracking open and click events.
    * `httpsPolicy`: The https policy to use for tracking open and click events. Valid values are `REQUIRE`, `REQUIRE_OPEN_ONLY` or `OPTIONAL`.
* `vdmOptions` - An object that contains information about the VDM preferences for your configuration set.
    * `dashboardOptions` - Specifies additional settings for your VDM configuration as applicable to the Dashboard.
        * `engagementMetrics` - Specifies the status of your VDM engagement metrics collection.
    * `guardianOptions` - Specifies additional settings for your VDM configuration as applicable to the Guardian.
        * `optimizedSharedDelivery` - Specifies the status of your VDM optimized shared delivery.

<!-- cache-key: cdktf-0.20.8 input-839396ade58efc075f49eac4bb59e33f8b473a746f424a518dff974740b9d98c -->