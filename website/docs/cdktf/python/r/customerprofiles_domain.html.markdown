---
subcategory: "Connect Customer Profiles"
layout: "aws"
page_title: "AWS: aws_customerprofiles_domain"
description: |-
  Terraform resource for managing an Amazon Customer Profiles Domain.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_customerprofiles_domain

Terraform resource for managing an Amazon Customer Profiles Domain.
See the [Create Domain](https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html) for more information.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.customerprofiles_domain import CustomerprofilesDomain
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, defaultExpirationDays):
        super().__init__(scope, name)
        CustomerprofilesDomain(self, "example",
            domain_name="example",
            default_expiration_days=default_expiration_days
        )
```

### With SQS DLQ and KMS set

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.customerprofiles_domain import CustomerprofilesDomain
from imports.aws.kms_key import KmsKey
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_policy import S3BucketPolicy
from imports.aws.sqs_queue import SqsQueue
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = KmsKey(self, "example",
            deletion_window_in_days=10,
            description="example"
        )
        aws_s3_bucket_example = S3Bucket(self, "example_1",
            bucket="example",
            force_destroy=True
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_example.override_logical_id("example")
        aws_s3_bucket_policy_example = S3BucketPolicy(self, "example_2",
            bucket=Token.as_string(aws_s3_bucket_example.id),
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["s3:GetObject", "s3:PutObject", "s3:ListBucket"],
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "profile.amazonaws.com"
                        },
                        "Resource": [aws_s3_bucket_example.arn, "${" + aws_s3_bucket_example.arn + "}/*"
                        ],
                        "Sid": "Customer Profiles S3 policy"
                    }
                    ],
                    "Version": "2012-10-17"
                }))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_policy_example.override_logical_id("example")
        aws_sqs_queue_example = SqsQueue(self, "example_3",
            name="example",
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["sqs:SendMessage"],
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "profile.amazonaws.com"
                        },
                        "Resource": "*",
                        "Sid": "Customer Profiles SQS policy"
                    }
                    ],
                    "Version": "2012-10-17"
                }))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sqs_queue_example.override_logical_id("example")
        CustomerprofilesDomain(self, "test",
            dead_letter_queue_url=Token.as_string(aws_sqs_queue_example.id),
            default_encryption_key=example.arn,
            default_expiration_days=365,
            domain_name=example
        )
```

## Argument Reference

The following arguments are required:

* `domain_name` - The name for your Customer Profile domain. It must be unique for your AWS account.
* `default_expiration_days` - The default number of days until the data within the domain expires.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `dead_letter_queue_url` - The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
* `default_encryption_key` - The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified. It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
* `matching` - A block that specifies the process of matching duplicate profiles. [Documented below](#matching).
* `rule_based_matching` - A block that specifies the process of matching duplicate profiles using the Rule-Based matching. [Documented below](#rule_based_matching).
* `tags` - Tags to apply to the domain. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `matching`

The `matching` configuration block supports the following attributes:

* `enabled` - (Required) The flag that enables the matching process of duplicate profiles.
* `auto_merging` - (Optional) A block that specifies the configuration about the auto-merging process. [Documented below](#auto_merging).
* `exporting_config` - (Optional) A block that specifies the configuration for exporting Identity Resolution results. [Documented below](#exporting_config).
* `job_schedule` - (Optional) A block that specifies the day and time when you want to start the Identity Resolution Job every week. [Documented below](#job_schedule).

### `rule_based_matching`

The `rule_based_matching` configuration block supports the following attributes:

* `enabled` - (Required) The flag that enables the rule-based matching process of duplicate profiles.
* `attribute_types_selector` - (Optional) A block that configures information about the `AttributeTypesSelector` where the rule-based identity resolution uses to match profiles. [Documented below](#attribute_types_selector).
* `conflict_resolution` - (Optional) A block that specifies how the auto-merging process should resolve conflicts between different profiles. [Documented below](#conflict_resolution).
* `exporting_config` - (Optional) A block that specifies the configuration for exporting Identity Resolution results. [Documented below](#exporting_config).
* `matching_rules` - (Optional) A block that configures how the rule-based matching process should match profiles. You can have up to 15 `rule` in the `natching_rules`. [Documented below](#matching_rules).
* `max_allowed_rule_level_for_matching` - (Optional) Indicates the maximum allowed rule level for matching.
* `max_allowed_rule_level_for_merging` - (Optional) Indicates the maximum allowed rule level for merging.

### `auto_merging`

The `auto_merging` configuration block supports the following attributes:

* `enabled` - (Required) The flag that enables the auto-merging of duplicate profiles.
* `conflict_resolution` - (Optional) A block that specifies how the auto-merging process should resolve conflicts between different profiles. [Documented below](#conflict_resolution).
* `consolidation` - (Optional) A block that specifies a list of matching attributes that represent matching criteria. If two profiles meet at least one of the requirements in the matching attributes list, they will be merged. [Documented below](#consolidation).
* `min_allowed_confidence_score_for_merging ` - (Optional) A number between 0 and 1 that represents the minimum confidence score required for profiles within a matching group to be merged during the auto-merge process. A higher score means higher similarity required to merge profiles.

### `conflict_resolution`

The `conflict_resolution` configuration block supports the following attributes:

* `conflict_resolving_model` - (Required) How the auto-merging process should resolve conflicts between different profiles. Valid values are `RECENCY` and `SOURCE`
* `source_name` - (Optional) The `ObjectType` name that is used to resolve profile merging conflicts when choosing `SOURCE` as the `ConflictResolvingModel`.

### `consolidation`

The `consolidation` configuration block supports the following attributes:

* `matching_attributes_list` - (Required) A list of matching criteria.

### `exporting_config`

The `exporting_config` configuration block supports the following attributes:

* `s3_exporting_config` - (Optional) A block that specifies the S3 location where Identity Resolution Jobs write result files. [Documented below](#s3_exporting_config).

### `s3_exporting_config`

The `s3_exporting_config` configuration block supports the following attributes:

* `s3_bucket_name` - (Required) The name of the S3 bucket where Identity Resolution Jobs write result files.
* `s3_key_name` - (Optional) The S3 key name of the location where Identity Resolution Jobs write result files.

### `job_schedule`

The `job_schedule` configuration block supports the following attributes:

* `day_of_the_week` - (Required) The day when the Identity Resolution Job should run every week.
* `time` - (Required) The time when the Identity Resolution Job should run every week.

### `attribute_types_selector`

The `attribute_types_selector` configuration block supports the following attributes:

* `attribute_matching_model` - (Required) Configures the `AttributeMatchingModel`, you can either choose `ONE_TO_ONE` or `MANY_TO_MANY`.
* `address` - (Optional) The `Address` type. You can choose from `Address`, `BusinessAddress`, `MaillingAddress`, and `ShippingAddress`.
* `email_address` - (Optional) The `Email` type. You can choose from `EmailAddress`, `BusinessEmailAddress` and `PersonalEmailAddress`.
* `phone_number` - (Optional) The `PhoneNumber` type. You can choose from `PhoneNumber`, `HomePhoneNumber`, and `MobilePhoneNumber`.

### `matching_rules`

The `matching_rules` configuration block supports the following attributes:

* `rule` - (Required) A single rule level of the `match_rules`. Configures how the rule-based matching process should match profiles.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the Customer Profiles Domain.
* `id` - The identifier of the Customer Profiles Domain.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Amazon Customer Profiles Domain using the resource `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.customerprofiles_domain import CustomerprofilesDomain
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CustomerprofilesDomain.generate_config_for_import(self, "example", "e6f777be-22d0-4b40-b307-5d2720ef16b2")
```

Using `terraform import`, import Amazon Customer Profiles Domain using the resource `id`. For example:

```console
% terraform import aws_customerprofiles_domain.example e6f777be-22d0-4b40-b307-5d2720ef16b2
```

<!-- cache-key: cdktf-0.20.8 input-778f067702a609868bebc7da14c3bdc692fe349b9697617b103b4ad9cbfcba27 -->