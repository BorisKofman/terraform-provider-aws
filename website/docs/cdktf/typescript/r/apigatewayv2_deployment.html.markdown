---
subcategory: "API Gateway V2"
layout: "aws"
page_title: "AWS: aws_apigatewayv2_deployment"
description: |-
  Manages an Amazon API Gateway Version 2 deployment.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apigatewayv2_deployment

Manages an Amazon API Gateway Version 2 deployment.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

-> **Note:** Creating a deployment for an API requires at least one `aws_apigatewayv2_route` resource associated with that API. To avoid race conditions when all resources are being created together, you need to add implicit resource references via the `triggers` argument or explicit resource references using the [resource `dependsOn` meta-argument](https://www.terraform.io/docs/configuration/meta-arguments/depends_on.html).

-> Enable the [resource `lifecycle` configuration block `create_before_destroy` argument](https://www.terraform.io/language/meta-arguments/lifecycle#create_before_destroy) in this resource configuration to properly order redeployments in Terraform.

## Example Usage

### Basic

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Apigatewayv2Deployment } from "./.gen/providers/aws/apigatewayv2-deployment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Apigatewayv2Deployment(this, "example", {
      apiId: Token.asString(awsApigatewayv2ApiExample.id),
      description: "Example deployment",
      lifecycle: {
        createBeforeDestroy: true,
      },
    });
  }
}

```

### Redeployment Triggers

-> **NOTE:** This is an optional and Terraform 0.12 (or later) advanced configuration that shows calculating a hash of the API's Terraform resources to determine changes that should trigger a new deployment. This value will change after the first Terraform apply of new resources, triggering an immediate redeployment, however it will stabilize afterwards except for resource changes. The `triggers` map can also be configured in other, more complex ways to fit the environment, avoiding the immediate redeployment issue.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Apigatewayv2Deployment } from "./.gen/providers/aws/apigatewayv2-deployment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Apigatewayv2Deployment(this, "example", {
      apiId: Token.asString(awsApigatewayv2ApiExample.id),
      description: "Example deployment",
      lifecycle: {
        createBeforeDestroy: true,
      },
      triggers: {
        redeployment: Token.asString(
          Fn.sha1(
            Token.asString(
              Fn.join(
                ",",
                Token.asList(
                  Fn.tolist([
                    Fn.jsonencode(awsApigatewayv2IntegrationExample),
                    Fn.jsonencode(awsApigatewayv2RouteExample),
                  ])
                )
              )
            )
          )
        ),
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `apiId` - (Required) API identifier.
* `description` - (Optional) Description for the deployment resource. Must be less than or equal to 1024 characters in length.
* `triggers` - (Optional) Map of arbitrary keys and values that, when changed, will trigger a redeployment. To force a redeployment without changing these keys/values, use the [`terraform taint` command](https://www.terraform.io/docs/commands/taint.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Deployment identifier.
* `autoDeployed` - Whether the deployment was automatically released.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_apigatewayv2_deployment` using the API identifier and deployment identifier. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Apigatewayv2Deployment } from "./.gen/providers/aws/apigatewayv2-deployment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Apigatewayv2Deployment.generateConfigForImport(
      this,
      "example",
      "aabbccddee/1122334"
    );
  }
}

```

Using `terraform import`, import `aws_apigatewayv2_deployment` using the API identifier and deployment identifier. For example:

```console
% terraform import aws_apigatewayv2_deployment.example aabbccddee/1122334
```

The `triggers` argument cannot be imported.

<!-- cache-key: cdktf-0.20.8 input-c7450fa8c59e1a68ada8776a1a5bec2daa0f662e44c8964a62d6583102b52e81 -->