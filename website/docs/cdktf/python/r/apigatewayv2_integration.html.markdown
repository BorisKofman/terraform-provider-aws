---
subcategory: "API Gateway V2"
layout: "aws"
page_title: "AWS: aws_apigatewayv2_integration"
description: |-
  Manages an Amazon API Gateway Version 2 integration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apigatewayv2_integration

Manages an Amazon API Gateway Version 2 integration.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Apigatewayv2Integration(self, "example",
            api_id=Token.as_string(aws_apigatewayv2_api_example.id),
            integration_type="MOCK"
        )
```

### Lambda Integration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
from imports.aws.lambda_function import LambdaFunction
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = LambdaFunction(self, "example",
            filename="example.zip",
            function_name="Example",
            handler="index.handler",
            role=Token.as_string(aws_iam_role_example.arn),
            runtime="nodejs20.x"
        )
        aws_apigatewayv2_integration_example = Apigatewayv2Integration(self, "example_1",
            api_id=Token.as_string(aws_apigatewayv2_api_example.id),
            connection_type="INTERNET",
            content_handling_strategy="CONVERT_TO_TEXT",
            description="Lambda example",
            integration_method="POST",
            integration_type="AWS_PROXY",
            integration_uri=example.invoke_arn,
            passthrough_behavior="WHEN_NO_MATCH"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_apigatewayv2_integration_example.override_logical_id("example")
```

### AWS Service Integration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Apigatewayv2Integration(self, "example",
            api_id=Token.as_string(aws_apigatewayv2_api_example.id),
            credentials_arn=Token.as_string(aws_iam_role_example.arn),
            description="SQS example",
            integration_subtype="SQS-SendMessage",
            integration_type="AWS_PROXY",
            request_parameters={
                "MessageBody": "$request.body.message",
                "QueueUrl": "$request.header.queueUrl"
            }
        )
```

### Private Integration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Apigatewayv2Integration(self, "example",
            api_id=Token.as_string(aws_apigatewayv2_api_example.id),
            connection_id=Token.as_string(aws_apigatewayv2_vpc_link_example.id),
            connection_type="VPC_LINK",
            credentials_arn=Token.as_string(aws_iam_role_example.arn),
            description="Example with a load balancer",
            integration_method="ANY",
            integration_type="HTTP_PROXY",
            integration_uri=Token.as_string(aws_lb_listener_example.arn),
            request_parameters={
                "append:header.authforintegration": "$context.authorizer.authorizerResponse",
                "overwrite:path": "staticValueForIntegration"
            },
            response_parameters=[Apigatewayv2IntegrationResponseParameters(
                mappings={
                    "append:header.auth": "$context.authorizer.authorizerResponse"
                },
                status_code=Token.as_string(403)
            ), Apigatewayv2IntegrationResponseParameters(
                mappings={
                    "overwrite:statuscode": "204"
                },
                status_code=Token.as_string(200)
            )
            ],
            tls_config=Apigatewayv2IntegrationTlsConfig(
                server_name_to_verify="example.com"
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `api_id` - (Required) API identifier.
* `integration_type` - (Required) Integration type of an integration.
Valid values: `AWS` (supported only for WebSocket APIs), `AWS_PROXY`, `HTTP` (supported only for WebSocket APIs), `HTTP_PROXY`, `MOCK` (supported only for WebSocket APIs). For an HTTP API private integration, use `HTTP_PROXY`.
* `connection_id` - (Optional) ID of the [VPC link](apigatewayv2_vpc_link.html) for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
* `connection_type` - (Optional) Type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
* `content_handling_strategy` - (Optional) How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
* `credentials_arn` - (Optional) Credentials required for the integration, if any.
* `description` - (Optional) Description of the integration.
* `integration_method` - (Optional) Integration's HTTP method. Must be specified if `integration_type` is not `MOCK`.
* `integration_subtype` - (Optional) AWS service action to invoke. Supported only for HTTP APIs when `integration_type` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
* `integration_uri` - (Optional) URI of the Lambda function for a Lambda proxy integration, when `integration_type` is `AWS_PROXY`.
For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
* `passthrough_behavior` - (Optional) Pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `request_templates` attribute.
Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
* `payload_format_version` - (Optional) The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
* `request_parameters` - (Optional) For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
For HTTP APIs with a specified `integration_subtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
For HTTP APIs without a specified `integration_subtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
* `request_templates` - (Optional) Map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
* `response_parameters` - (Optional) Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
* `template_selection_expression` - (Optional) The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
* `timeout_milliseconds` - (Optional) Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
Terraform will only perform drift detection of its value when present in a configuration.
* `tls_config` - (Optional) TLS configuration for a private integration. Supported only for HTTP APIs.

The `response_parameters` object supports the following:

* `status_code` - (Required) HTTP status code in the range 200-599.
* `mappings` - (Required) Key-value map. The key of this map identifies the location of the request parameter to change, and how to change it. The corresponding value specifies the new data for the parameter.
See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.

The `tls_config` object supports the following:

* `server_name_to_verify` - (Optional) If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Integration identifier.
* `integration_response_selection_expression` - The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_apigatewayv2_integration` using the API identifier and integration identifier. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Apigatewayv2Integration.generate_config_for_import(self, "example", "aabbccddee/1122334")
```

Using `terraform import`, import `aws_apigatewayv2_integration` using the API identifier and integration identifier. For example:

```console
% terraform import aws_apigatewayv2_integration.example aabbccddee/1122334
```

-> **Note:** The API Gateway managed integration created as part of [_quick_create_](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html#apigateway-definition-quick-create) cannot be imported.

<!-- cache-key: cdktf-0.20.8 input-19a8a0a3db8d36fc6fe256c58d6d4dc3f257fc738bc47f1b5cdc00513386c442 -->