---
subcategory: "App Runner"
layout: "aws"
page_title: "AWS: aws_apprunner_deployment"
description: |-
  Manages an App Runner Deployment Operation.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apprunner_deployment

Manages an App Runner Deployment Operation.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApprunnerDeployment } from "./.gen/providers/aws/apprunner-deployment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ApprunnerDeployment(this, "example", {
      serviceArn: Token.asString(awsApprunnerServiceExample.arn),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `serviceArn` - (Required) The Amazon Resource Name (ARN) of the App Runner service to start the deployment for.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A unique identifier for the deployment.
* `operationId` - The unique ID of the operation associated with deployment.
* `status` - The current status of the App Runner service deployment.

<!-- cache-key: cdktf-0.20.8 input-fcfcb8c199313c9b7bda02684ebc1478b1a90d029a0d10e7dbf8264e2af0254c -->