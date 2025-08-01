---
subcategory: "VPC IPAM (IP Address Manager)"
layout: "aws"
page_title: "AWS: aws_vpc_ipam"
description: |-
  Provides an IPAM resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_ipam

Provides an IPAM resource.

## Example Usage

Basic usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { VpcIpam } from "./.gen/providers/aws/vpc-ipam";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const current = new DataAwsRegion(this, "current", {});
    new VpcIpam(this, "main", {
      description: "My IPAM",
      operatingRegions: [
        {
          regionName: Token.asString(current.region),
        },
      ],
      tags: {
        Test: "Main",
      },
    });
  }
}

```

Shared with multiple operating_regions:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import {
  VariableType,
  TerraformVariable,
  Fn,
  Token,
  TerraformIterator,
  TerraformStack,
} from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { VpcIpam } from "./.gen/providers/aws/vpc-ipam";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const ipamRegions = new TerraformVariable(this, "ipam_regions", {
      default: ["us-east-1", "us-west-2"],
      type: VariableType.ANY,
    });
    const current = new DataAwsRegion(this, "current", {});
    const allIpamRegions = Fn.distinct(
      Token.asAny(Fn.concat([[current.region], ipamRegions.value]))
    );
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const mainDynamicIterator0 = TerraformIterator.fromList(
      Token.asAny(allIpamRegions)
    );
    new VpcIpam(this, "main", {
      description: "multi region ipam",
      operatingRegions: mainDynamicIterator0.dynamic({
        region_name: mainDynamicIterator0.value,
      }),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cascade` - (Optional) Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
* `description` - (Optional) A description for the IPAM.
* `enablePrivateGua` - (Optional) Enable this option to use your own GUA ranges as private IPv6 addresses. Default: `false`.
* `operatingRegions` - (Required) Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the [region_name](#operating_regions) parameter. You **must** set your provider block region as an operating_region.
* `tier` - (Optional) specifies the IPAM tier. Valid options include `free` and `advanced`. Default is `advanced`.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### operating_regions

* `regionName` - (Required) The name of the Region you want to add to the IPAM.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of IPAM
* `id` - The ID of the IPAM
* `defaultResourceDiscoveryId` - The IPAM's default resource discovery ID.
* `defaultResourceDiscoveryAssociationId` - The IPAM's default resource discovery association ID.
* `privateDefaultScopeId` - The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
* `publicDefaultScopeId` - The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
IP space. The public scope is intended for all internet-routable IP space.
* `scopeCount` - The number of scopes in the IPAM.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IPAMs using the IPAM `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcIpam } from "./.gen/providers/aws/vpc-ipam";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcIpam.generateConfigForImport(this, "example", "ipam-0178368ad2146a492");
  }
}

```

Using `terraform import`, import IPAMs using the IPAM `id`. For example:

```console
% terraform import aws_vpc_ipam.example ipam-0178368ad2146a492
```

<!-- cache-key: cdktf-0.20.8 input-41a160287db88dd545c3f1a2c18b23ba0f7bd115ca91f86152dc636cf70f7e23 -->