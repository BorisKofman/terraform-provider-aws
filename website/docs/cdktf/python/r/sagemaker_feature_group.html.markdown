---
subcategory: "SageMaker AI"
layout: "aws"
page_title: "AWS: aws_sagemaker_feature_group"
description: |-
  Provides a SageMaker AI Feature Group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sagemaker_feature_group

Provides a SageMaker AI Feature Group resource.

## Example Usage

Basic usage:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sagemaker_feature_group import SagemakerFeatureGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SagemakerFeatureGroup(self, "example",
            event_time_feature_name="example",
            feature_definition=[SagemakerFeatureGroupFeatureDefinition(
                feature_name="example",
                feature_type="String"
            )
            ],
            feature_group_name="example",
            online_store_config=SagemakerFeatureGroupOnlineStoreConfig(
                enable_online_store=True
            ),
            record_identifier_feature_name="example",
            role_arn=test.arn
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `feature_group_name` - (Required) The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
* `record_identifier_feature_name` - (Required) The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
* `event_time_feature_name` - (Required) The name of the feature that stores the EventTime of a Record in a Feature Group.
* `description` (Optional) - A free-form description of a Feature Group.
* `role_arn` (Required) - The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offline_store_config` is provided.
* `feature_definition` (Optional) - A list of Feature names and types. See [Feature Definition](#feature-definition) Below.
* `offline_store_config` (Optional) - The Offline Feature Store Configuration. See [Offline Store Config](#offline-store-config) Below.
* `online_store_config` (Optional) - The Online Feature Store Configuration. See [Online Store Config](#online-store-config) Below.
* `tags` - (Optional) Map of resource tags for the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### Feature Definition

* `feature_name` - (Required) The name of a feature. `feature_name` cannot be any of the following: `is_deleted`, `write_time`, `api_invocation_time`.
* `feature_type` - (Required) The value type of a feature. Valid values are `Integral`, `Fractional`, or `String`.

### Offline Store Config

* `enable_online_store` - (Optional) Set to `true` to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
* `s3_storage_config` - (Required) The Amazon Simple Storage (Amazon S3) location of OfflineStore. See [S3 Storage Config](#s3-storage-config) Below.
* `data_catalog_config` - (Optional) The meta data of the Glue table that is autogenerated when an OfflineStore is created. See [Data Catalog Config](#data-catalog-config) Below.
* `table_format` - (Optional) Format for the offline store table. Supported formats are `Glue` (Default) and Apache `Iceberg` (https://iceberg.apache.org/).

### Online Store Config

* `disable_glue_table_creation` - (Optional) Set to `true` to turn Online Store On.
* `security_config` - (Required) Security config for at-rest encryption of your OnlineStore. See [Security Config](#security-config) Below.
* `storage_type` - (Optional) Option for different tiers of low latency storage for real-time data retrieval. Valid values are `Standard`, or `InMemory`.
* `ttl_duration` - (Optional) Time to live duration, where the record is hard deleted after the expiration time is reached; ExpiresAt = EventTime + TtlDuration.. See [TTl Duration](#ttl-duration) Below.

#### S3 Storage Config

* `kms_key_id` - (Optional) The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
* `s3_uri` - (Required) The S3 URI, or location in Amazon S3, of OfflineStore.
* `resolved_output_s3_uri` - (Optional) The S3 path where offline records are written.

#### Data Catalog Config

* `catalog` - (Optional) The name of the Glue table catalog.
* `database` - (Optional) The name of the Glue table database.
* `table_name` - (Optional) The name of the Glue table.

#### Security Config

* `kms_key_id` - (Optional) The ID of the AWS Key Management Service (AWS KMS) key that SageMaker AI Feature Store uses to encrypt the Amazon S3 objects at rest using Amazon S3 server-side encryption.

#### TTl Duration

* `unit` - (Optional) TtlDuration time unit. Valid values are `Seconds`, `Minutes`, `Hours`, `Days`, or `Weeks`.
* `value` - (Optional) TtlDuration time value.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `name` - The name of the Feature Group.
* `arn` - The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Feature Groups using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sagemaker_feature_group import SagemakerFeatureGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SagemakerFeatureGroup.generate_config_for_import(self, "testFeatureGroup", "feature_group-foo")
```

Using `terraform import`, import Feature Groups using the `name`. For example:

```console
% terraform import aws_sagemaker_feature_group.test_feature_group feature_group-foo
```

<!-- cache-key: cdktf-0.20.8 input-78b465d67681f4bcd1ec14afda11841888d2991277f7f48fe1bfac8b5f8c2412 -->