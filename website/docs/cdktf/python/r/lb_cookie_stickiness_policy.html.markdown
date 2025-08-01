---
subcategory: "ELB Classic"
layout: "aws"
page_title: "AWS: aws_lb_cookie_stickiness_policy"
description: |-
  Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lb_cookie_stickiness_policy

Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elb import Elb
from imports.aws.lb_cookie_stickiness_policy import LbCookieStickinessPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        lb = Elb(self, "lb",
            availability_zones=["us-east-1a"],
            listener=[ElbListener(
                instance_port=8000,
                instance_protocol="http",
                lb_port=80,
                lb_protocol="http"
            )
            ],
            name="test-lb"
        )
        LbCookieStickinessPolicy(self, "foo",
            cookie_expiration_period=600,
            lb_port=80,
            load_balancer=lb.id,
            name="foo-policy"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the stickiness policy.
* `load_balancer` - (Required) The load balancer to which the policy
  should be attached.
* `lb_port` - (Required) The load balancer port to which the policy
  should be applied. This must be an active listener on the load
balancer.
* `cookie_expiration_period` - (Optional) The time period after which
  the session cookie should be considered stale, expressed in seconds.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the policy.
* `name` - The name of the stickiness policy.
* `load_balancer` - The load balancer to which the policy is attached.
* `lb_port` - The load balancer port to which the policy is applied.
* `cookie_expiration_period` - The time period after which the session cookie is considered stale, expressed in seconds.

<!-- cache-key: cdktf-0.20.8 input-d638009f7c029a079b861184055cfc5cbead4e7243716af93ac00abb11b1bc08 -->