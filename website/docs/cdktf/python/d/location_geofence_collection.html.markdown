---
subcategory: "Location"
layout: "aws"
page_title: "AWS: aws_location_geofence_collection"
description: |-
    Retrieve information about a Location Service Geofence Collection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_location_geofence_collection

Retrieve information about a Location Service Geofence Collection.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_location_geofence_collection import DataAwsLocationGeofenceCollection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsLocationGeofenceCollection(self, "example",
            collection_name="example"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `collection_name` - (Required) Name of the geofence collection.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `collection_arn` - ARN for the geofence collection resource. Used when you need to specify a resource across all AWS.
* `create_time` - Timestamp for when the geofence collection resource was created in ISO 8601 format.
* `description` - Optional description of the geofence collection resource.
* `kms_key_id` - Key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
* `tags` - Key-value map of resource tags for the geofence collection.
* `update_time` - Timestamp for when the geofence collection resource was last updated in ISO 8601 format.

<!-- cache-key: cdktf-0.20.8 input-1fec0c9656e29e888310da17b35999e3c80074d9b82d341d42c7e44dd1dcab3d -->