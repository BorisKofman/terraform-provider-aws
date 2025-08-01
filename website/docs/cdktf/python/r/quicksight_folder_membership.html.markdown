---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_folder_membership"
description: |-
  Terraform resource for managing an AWS QuickSight Folder Membership.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_folder_membership

Terraform resource for managing an AWS QuickSight Folder Membership.

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
from imports.aws.quicksight_folder_membership import QuicksightFolderMembership
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightFolderMembership(self, "example",
            folder_id=Token.as_string(aws_quicksight_folder_example.folder_id),
            member_id=Token.as_string(aws_quicksight_data_set_example.data_set_id),
            member_type="DATASET"
        )
```

## Argument Reference

The following arguments are required:

* `folder_id` - (Required, Forces new resource) Identifier for the folder.
* `member_id` - (Required, Forces new resource) ID of the asset (the dashboard, analysis, or dataset).
* `member_type` - (Required, Forces new resource) Type of the member. Valid values are `ANALYSIS`, `DASHBOARD`, and `DATASET`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `aws_account_id` - (Optional, Forces new resource) AWS account ID.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A comma-delimited string joining AWS account ID, folder ID, member type, and member ID.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import QuickSight Folder Membership using the AWS account ID, folder ID, member type, and member ID separated by commas (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_folder_membership import QuicksightFolderMembership
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightFolderMembership.generate_config_for_import(self, "example", "123456789012,example-folder,DATASET,example-dataset")
```

Using `terraform import`, import QuickSight Folder Membership using the AWS account ID, folder ID, member type, and member ID separated by commas (`,`). For example:

```console
% terraform import aws_quicksight_folder_membership.example 123456789012,example-folder,DATASET,example-dataset
```

<!-- cache-key: cdktf-0.20.8 input-e89fcb0d6a3232546180eb5dd15790a94e0e56ca32d114446be9d09da1f7d770 -->