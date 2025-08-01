---
subcategory: "Service Catalog AppRegistry"
layout: "aws"
page_title: "AWS: aws_servicecatalogappregistry_application"
description: |-
  Terraform data source for managing an AWS Service Catalog AppRegistry Application.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_servicecatalogappregistry_application

Terraform data source for managing an AWS Service Catalog AppRegistry Application.

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
import { DataAwsServicecatalogappregistryApplication } from "./.gen/providers/aws/data-aws-servicecatalogappregistry-application";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsServicecatalogappregistryApplication(this, "example", {
      id: "application-1234",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `id` - (Required) Application identifier.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `applicationTag` - A map with a single tag key-value pair used to associate resources with the application.
* `arn` - ARN (Amazon Resource Name) of the application.
* `description` - Description of the application.
* `name` - Name of the application.
* `tags` - A map of tags assigned to the Application. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

<!-- cache-key: cdktf-0.20.8 input-877263c109c14b327ad0130e26e2eb86ae347a9fcb6b01c4f44d58c555e17f3a -->