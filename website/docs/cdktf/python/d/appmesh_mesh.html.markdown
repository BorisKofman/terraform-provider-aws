---
subcategory: "App Mesh"
layout: "aws"
page_title: "AWS: aws_appmesh_mesh"
description: |-
    Terraform data source for managing an AWS App Mesh Mesh.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_appmesh_mesh

The App Mesh Mesh data source allows details of an App Mesh Mesh to be retrieved by its name and optionally the mesh_owner.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_appmesh_mesh import DataAwsAppmeshMesh
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsAppmeshMesh(self, "simple",
            name="simpleapp"
        )
```

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_appmesh_mesh import DataAwsAppmeshMesh
from imports.aws.data_aws_caller_identity import DataAwsCallerIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        current = DataAwsCallerIdentity(self, "current")
        DataAwsAppmeshMesh(self, "simple",
            mesh_owner=Token.as_string(current.account_id),
            name="simpleapp"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the service mesh.
* `mesh_owner` - (Optional) AWS account ID of the service mesh's owner.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the service mesh.
* `created_date` - Creation date of the service mesh.
* `last_updated_date` - Last update date of the service mesh.
* `resource_owner` - Resource owner's AWS account ID.
* `spec` - Service mesh specification. See the [`aws_appmesh_mesh`](/docs/providers/aws/r/appmesh_mesh.html#spec) resource for details.
* `tags` - Map of tags.

<!-- cache-key: cdktf-0.20.8 input-3a18ad9bf7d5c0b46b9eb98c059bc3a993630967b0965d5dfa2cc9efca1af211 -->