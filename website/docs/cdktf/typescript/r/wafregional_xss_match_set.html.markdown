---
subcategory: "WAF Classic Regional"
layout: "aws"
page_title: "AWS: aws_wafregional_xss_match_set"
description: |-
  Provides an AWS WAF Regional XSS Match Set resource for use with ALB.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_wafregional_xss_match_set

Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WafregionalXssMatchSet } from "./.gen/providers/aws/wafregional-xss-match-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new WafregionalXssMatchSet(this, "xss_match_set", {
      name: "xss_match_set",
      xssMatchTuple: [
        {
          fieldToMatch: {
            type: "URI",
          },
          textTransformation: "NONE",
        },
        {
          fieldToMatch: {
            type: "QUERY_STRING",
          },
          textTransformation: "NONE",
        },
      ],
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the set
* `xssMatchTuple` - (Optional) The parts of web requests that you want to inspect for cross-site scripting attacks.

### Nested fields

#### `xssMatchTuple`

* `fieldToMatch` - (Required) Specifies where in a web request to look for cross-site scripting attacks.
* `textTransformation` - (Required) Which text transformation, if any, to perform on the web request before inspecting the request for cross-site scripting attacks.

#### `fieldToMatch`

* `data` - (Optional) When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
* `type` - (Required) The part of the web request that you want AWS WAF to search for a specified stringE.g., `HEADER` or `METHOD`

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the Regional WAF XSS Match Set.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AWS WAF Regional XSS Match using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WafregionalXssMatchSet } from "./.gen/providers/aws/wafregional-xss-match-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    WafregionalXssMatchSet.generateConfigForImport(
      this,
      "example",
      "12345abcde"
    );
  }
}

```

Using `terraform import`, import AWS WAF Regional XSS Match using the `id`. For example:

```console
% terraform import aws_wafregional_xss_match_set.example 12345abcde
```

<!-- cache-key: cdktf-0.20.8 input-8932cc99a9375ad6ac76feb468466f1581535196463ce51c2861c96097ecebd6 -->