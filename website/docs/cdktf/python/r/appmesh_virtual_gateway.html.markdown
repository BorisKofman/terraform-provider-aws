---
subcategory: "App Mesh"
layout: "aws"
page_title: "AWS: aws_appmesh_virtual_gateway"
description: |-
  Provides an AWS App Mesh virtual gateway resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appmesh_virtual_gateway

Provides an AWS App Mesh virtual gateway resource.

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_gateway import AppmeshVirtualGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualGateway(self, "example",
            mesh_name="example-service-mesh",
            name="example-virtual-gateway",
            spec=AppmeshVirtualGatewaySpec(
                listener=[AppmeshVirtualGatewaySpecListener(
                    port_mapping=AppmeshVirtualGatewaySpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    )
                )
                ]
            ),
            tags={
                "Environment": "test"
            }
        )
```

### Access Logs and TLS

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_gateway import AppmeshVirtualGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualGateway(self, "example",
            mesh_name="example-service-mesh",
            name="example-virtual-gateway",
            spec=AppmeshVirtualGatewaySpec(
                listener=[AppmeshVirtualGatewaySpecListener(
                    port_mapping=AppmeshVirtualGatewaySpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    ),
                    tls=AppmeshVirtualGatewaySpecListenerTls(
                        certificate=AppmeshVirtualGatewaySpecListenerTlsCertificate(
                            acm=AppmeshVirtualGatewaySpecListenerTlsCertificateAcm(
                                certificate_arn=Token.as_string(aws_acm_certificate_example.arn)
                            )
                        ),
                        mode="STRICT"
                    )
                )
                ],
                logging=AppmeshVirtualGatewaySpecLogging(
                    access_log=AppmeshVirtualGatewaySpecLoggingAccessLog(
                        file=AppmeshVirtualGatewaySpecLoggingAccessLogFile(
                            path="/var/log/access.log"
                        )
                    )
                )
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name to use for the virtual gateway. Must be between 1 and 255 characters in length.
* `mesh_name` - (Required) Name of the service mesh in which to create the virtual gateway. Must be between 1 and 255 characters in length.
* `mesh_owner` - (Optional) AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider][1] is currently connected to.
* `spec` - (Required) Virtual gateway specification to apply.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

The `spec` object supports the following:

* `listener` - (Required) Listeners that the mesh endpoint is expected to receive inbound traffic from. You can specify one listener.
* `backend_defaults` - (Optional) Defaults for backends.
* `logging` - (Optional) Inbound and outbound access logging information for the virtual gateway.

The `backend_defaults` object supports the following:

* `client_policy` - (Optional) Default client policy for virtual gateway backends.

The `client_policy` object supports the following:

* `tls` - (Optional) Transport Layer Security (TLS) client policy.

The `tls` object supports the following:

* `certificate` (Optional) Virtual gateway's client's Transport Layer Security (TLS) certificate.
* `enforce` - (Optional) Whether the policy is enforced. Default is `true`.
* `ports` - (Optional) One or more ports that the policy is enforced for.
* `validation` - (Required) TLS validation context.

The `certificate` object supports the following:

* `file` - (Optional) Local file certificate.
* `sds` - (Optional) A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate chain for the certificate.
* `private_key` - (Required) Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.

The `validation` object supports the following:

* `subject_alternative_names` - (Optional) SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
* `trust` - (Required) TLS validation context trust.

The `subject_alternative_names` object supports the following:

* `match` - (Required) Criteria for determining a SAN's match.

The `match` object supports the following:

* `exact` - (Required) Values sent must match the specified values exactly.

The `trust` object supports the following:

* `acm` - (Optional) TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
* `file` - (Optional) TLS validation context trust for a local file certificate.
* `sds` - (Optional) TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `acm` object supports the following:

* `certificate_authority_arns` - (Required) One or more ACM ARNs.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret for a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.

The `listener` object supports the following:

* `port_mapping` - (Required) Port mapping information for the listener.
* `connection_pool` - (Optional) Connection pool information for the listener.
* `health_check` - (Optional) Health check information for the listener.
* `tls` - (Optional) Transport Layer Security (TLS) properties for the listener

The `logging` object supports the following:

* `access_log` - (Optional) Access log configuration for a virtual gateway.

The `access_log` object supports the following:

* `file` - (Optional) File object to send virtual gateway access logs to.

The `file` object supports the following:

* `format` - (Optional) The specified format for the logs.
* `path` - (Required) File path to write access logs to. You can use `/dev/stdout` to send access logs to standard out. Must be between 1 and 255 characters in length.

The `format` object supports the following:

* `json` - (Optional) The logging format for JSON.
* `text` - (Optional) The logging format for text. Must be between 1 and 1000 characters in length.

The `json` object supports the following:

* `key` - (Required) The specified key for the JSON. Must be between 1 and 100 characters in length.
* `value` - (Required) The specified value for the JSON. Must be between 1 and 100 characters in length.

The `port_mapping` object supports the following:

* `port` - (Required) Port used for the port mapping.
* `protocol` - (Required) Protocol used for the port mapping. Valid values are `http`, `http2`, `tcp` and `grpc`.

The `connection_pool` object supports the following:

* `grpc` - (Optional) Connection pool information for gRPC listeners.
* `http` - (Optional) Connection pool information for HTTP listeners.
* `http2` - (Optional) Connection pool information for HTTP2 listeners.

The `grpc` connection pool object supports the following:

* `max_requests` - (Required) Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of `1`.

The `http` connection pool object supports the following:

* `max_connections` - (Required) Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of `1`.
* `max_pending_requests` - (Optional) Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster. Minimum value of `1`.

The `http2` connection pool object supports the following:

* `max_requests` - (Required) Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of `1`.

The `health_check` object supports the following:

* `healthy_threshold` - (Required) Number of consecutive successful health checks that must occur before declaring listener healthy.
* `interval_millis`- (Required) Time period in milliseconds between each health check execution.
* `protocol` - (Required) Protocol for the health check request. Valid values are `http`, `http2`, and `grpc`.
* `timeout_millis` - (Required) Amount of time to wait when receiving a response from the health check, in milliseconds.
* `unhealthy_threshold` - (Required) Number of consecutive failed health checks that must occur before declaring a virtual gateway unhealthy.
* `path` - (Optional) Destination path for the health check request. This is only required if the specified protocol is `http` or `http2`.
* `port` - (Optional) Destination port for the health check request. This port must match the port defined in the `port_mapping` for the listener.

The `tls` object supports the following:

* `certificate` - (Required) Listener's TLS certificate.
* `mode`- (Required) Listener's TLS mode. Valid values: `DISABLED`, `PERMISSIVE`, `STRICT`.
* `validation`- (Optional) Listener's Transport Layer Security (TLS) validation context.

The `certificate` object supports the following:

* `acm` - (Optional) An AWS Certificate Manager (ACM) certificate.
* `file` - (Optional) Local file certificate.
* `sds` - (Optional) A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `acm` object supports the following:

* `certificate_arn` - (Required) ARN for the certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate chain for the certificate. Must be between 1 and 255 characters in length.
* `private_key` - (Required) Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.

The `validation` object supports the following:

* `subject_alternative_names` - (Optional) SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
* `trust` - (Required) TLS validation context trust.

The `subject_alternative_names` object supports the following:

* `match` - (Required) Criteria for determining a SAN's match.

The `match` object supports the following:

* `exact` - (Required) Values sent must match the specified values exactly.

The `trust` object supports the following:

* `file` - (Optional) TLS validation context trust for a local file certificate.
* `sds` - (Optional) TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret for a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the virtual gateway.
* `arn` - ARN of the virtual gateway.
* `created_date` - Creation date of the virtual gateway.
* `last_updated_date` - Last update date of the virtual gateway.
* `resource_owner` - Resource owner's AWS account ID.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import App Mesh virtual gateway using `mesh_name` together with the virtual gateway's `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_gateway import AppmeshVirtualGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualGateway.generate_config_for_import(self, "example", "mesh/gw1")
```

Using `terraform import`, import App Mesh virtual gateway using `mesh_name` together with the virtual gateway's `name`. For example:

```console
% terraform import aws_appmesh_virtual_gateway.example mesh/gw1
```

[1]: /docs/providers/aws/index.html

<!-- cache-key: cdktf-0.20.8 input-81f4fa640ffb488186c672189cd26a0c662bd2c1bc09764fea8d2992db63a101 -->