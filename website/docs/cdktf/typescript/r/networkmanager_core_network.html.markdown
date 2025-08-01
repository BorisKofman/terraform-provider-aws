---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_core_network"
description: |-
  Manages a Network Manager Core Network.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkmanager_core_network

Manages a Network Manager Core Network.

Use this resource to create and manage a core network within a global network.

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
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkmanagerCoreNetwork(this, "example", {
      globalNetworkId: Token.asString(awsNetworkmanagerGlobalNetworkExample.id),
    });
  }
}

```

### With description

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkmanagerCoreNetwork(this, "example", {
      description: "example",
      globalNetworkId: Token.asString(awsNetworkmanagerGlobalNetworkExample.id),
    });
  }
}

```

### With tags

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkmanagerCoreNetwork(this, "example", {
      globalNetworkId: Token.asString(awsNetworkmanagerGlobalNetworkExample.id),
      tags: {
        hello: "world",
      },
    });
  }
}

```

### With VPC Attachment (Single Region)

The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `terraform apply` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument. There are 2 options to implement this:

- Option 1: Use the `basePolicyDocument` argument that allows the most customizations to a base policy. Use this to customize the `edgeLocations` `asn`. In the example below, `us-west-2` and ASN `65500` are used in the base policy.
- Option 2: Use the `createBasePolicy` argument only. This creates a base policy in the region specified in the `provider` block.

#### Option 1 - using base_policy_document

If you require a custom ASN for the edge location, please use the `basePolicyDocument` argument to pass a specific ASN. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNetworkmanagerCoreNetworkPolicyDocument } from "./.gen/providers/aws/data-aws-networkmanager-core-network-policy-document";
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
import { NetworkmanagerCoreNetworkPolicyAttachment } from "./.gen/providers/aws/networkmanager-core-network-policy-attachment";
import { NetworkmanagerGlobalNetwork } from "./.gen/providers/aws/networkmanager-global-network";
import { NetworkmanagerVpcAttachment } from "./.gen/providers/aws/networkmanager-vpc-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new NetworkmanagerGlobalNetwork(this, "example", {});
    const base = new DataAwsNetworkmanagerCoreNetworkPolicyDocument(
      this,
      "base",
      {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                asn: "65500",
                location: "us-west-2",
              },
            ],
          },
        ],
        segments: [
          {
            name: "segment",
          },
        ],
      }
    );
    const awsNetworkmanagerCoreNetworkExample = new NetworkmanagerCoreNetwork(
      this,
      "example_2",
      {
        basePolicyDocument: Token.asString(base.json),
        createBasePolicy: true,
        globalNetworkId: example.id,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkExample.overrideLogicalId("example");
    const awsNetworkmanagerVpcAttachmentExample =
      new NetworkmanagerVpcAttachment(this, "example_3", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExample, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExample.arn),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerVpcAttachmentExample.overrideLogicalId("example");
    const dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample =
      new DataAwsNetworkmanagerCoreNetworkPolicyDocument(this, "example_4", {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                asn: "65500",
                location: "us-west-2",
              },
            ],
          },
        ],
        segmentActions: [
          {
            action: "create-route",
            destinationCidrBlocks: ["0.0.0.0/0"],
            destinations: [
              Token.asString(awsNetworkmanagerVpcAttachmentExample.id),
            ],
            segment: "segment",
          },
        ],
        segments: [
          {
            name: "segment",
          },
        ],
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.overrideLogicalId(
      "example"
    );
    const awsNetworkmanagerCoreNetworkPolicyAttachmentExample =
      new NetworkmanagerCoreNetworkPolicyAttachment(this, "example_5", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        policyDocument: Token.asString(
          dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.json
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkPolicyAttachmentExample.overrideLogicalId(
      "example"
    );
  }
}

```

#### Option 2 - create_base_policy only

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNetworkmanagerCoreNetworkPolicyDocument } from "./.gen/providers/aws/data-aws-networkmanager-core-network-policy-document";
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
import { NetworkmanagerCoreNetworkPolicyAttachment } from "./.gen/providers/aws/networkmanager-core-network-policy-attachment";
import { NetworkmanagerGlobalNetwork } from "./.gen/providers/aws/networkmanager-global-network";
import { NetworkmanagerVpcAttachment } from "./.gen/providers/aws/networkmanager-vpc-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new NetworkmanagerGlobalNetwork(this, "example", {});
    const awsNetworkmanagerCoreNetworkExample = new NetworkmanagerCoreNetwork(
      this,
      "example_1",
      {
        createBasePolicy: true,
        globalNetworkId: example.id,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkExample.overrideLogicalId("example");
    const awsNetworkmanagerVpcAttachmentExample =
      new NetworkmanagerVpcAttachment(this, "example_2", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExample, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExample.arn),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerVpcAttachmentExample.overrideLogicalId("example");
    const dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample =
      new DataAwsNetworkmanagerCoreNetworkPolicyDocument(this, "example_3", {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                location: "us-west-2",
              },
            ],
          },
        ],
        segmentActions: [
          {
            action: "create-route",
            destinationCidrBlocks: ["0.0.0.0/0"],
            destinations: [
              Token.asString(awsNetworkmanagerVpcAttachmentExample.id),
            ],
            segment: "segment",
          },
        ],
        segments: [
          {
            name: "segment",
          },
        ],
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.overrideLogicalId(
      "example"
    );
    const awsNetworkmanagerCoreNetworkPolicyAttachmentExample =
      new NetworkmanagerCoreNetworkPolicyAttachment(this, "example_4", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        policyDocument: Token.asString(
          dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.json
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkPolicyAttachmentExample.overrideLogicalId(
      "example"
    );
  }
}

```

### With VPC Attachment (Multi-Region)

The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument of the [`aws_networkmanager_core_network` resource](/docs/providers/aws/r/networkmanager_core_network.html) to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `terraform apply` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument. For multi-region in a core network that does not yet have a `LIVE` policy, there are 2 options:

- Option 1: Use the `basePolicyDocument` argument that allows the most customizations to a base policy. Use this to customize the `edgeLocations` `asn`. In the example below, `us-west-2`, `us-east-1` and specific ASNs are used in the base policy.
- Option 2: Pass a list of regions to the `aws_networkmanager_core_network` `basePolicyRegions` argument. In the example below, `us-west-2` and `us-east-1` are specified in the base policy.

#### Option 1 - using base_policy_document

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNetworkmanagerCoreNetworkPolicyDocument } from "./.gen/providers/aws/data-aws-networkmanager-core-network-policy-document";
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
import { NetworkmanagerCoreNetworkPolicyAttachment } from "./.gen/providers/aws/networkmanager-core-network-policy-attachment";
import { NetworkmanagerGlobalNetwork } from "./.gen/providers/aws/networkmanager-global-network";
import { NetworkmanagerVpcAttachment } from "./.gen/providers/aws/networkmanager-vpc-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new NetworkmanagerGlobalNetwork(this, "example", {});
    const base = new DataAwsNetworkmanagerCoreNetworkPolicyDocument(
      this,
      "base",
      {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                asn: "65500",
                location: "us-west-2",
              },
              {
                asn: "65501",
                location: "us-east-1",
              },
            ],
          },
        ],
        segments: [
          {
            name: "segment",
          },
        ],
      }
    );
    const awsNetworkmanagerCoreNetworkExample = new NetworkmanagerCoreNetwork(
      this,
      "example_2",
      {
        basePolicyDocument: Token.asString(base.json),
        createBasePolicy: true,
        globalNetworkId: example.id,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkExample.overrideLogicalId("example");
    const exampleUsEast1 = new NetworkmanagerVpcAttachment(
      this,
      "example_us_east_1",
      {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        provider: "alternate",
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExampleUsEast1, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExampleUsEast1.arn),
      }
    );
    const exampleUsWest2 = new NetworkmanagerVpcAttachment(
      this,
      "example_us_west_2",
      {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExampleUsWest2, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExampleUsWest2.arn),
      }
    );
    const dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample =
      new DataAwsNetworkmanagerCoreNetworkPolicyDocument(this, "example_5", {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                asn: "65500",
                location: "us-west-2",
              },
              {
                asn: "65501",
                location: "us-east-1",
              },
            ],
          },
        ],
        segmentActions: [
          {
            action: "create-route",
            destinationCidrBlocks: ["10.0.0.0/16"],
            destinations: [exampleUsWest2.id],
            segment: "segment",
          },
          {
            action: "create-route",
            destinationCidrBlocks: ["10.1.0.0/16"],
            destinations: [exampleUsEast1.id],
            segment: "segment",
          },
        ],
        segments: [
          {
            name: "segment",
          },
          {
            name: "segment2",
          },
        ],
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.overrideLogicalId(
      "example"
    );
    const awsNetworkmanagerCoreNetworkPolicyAttachmentExample =
      new NetworkmanagerCoreNetworkPolicyAttachment(this, "example_6", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        policyDocument: Token.asString(
          dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.json
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkPolicyAttachmentExample.overrideLogicalId(
      "example"
    );
  }
}

```

#### Option 2 - using base_policy_regions

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNetworkmanagerCoreNetworkPolicyDocument } from "./.gen/providers/aws/data-aws-networkmanager-core-network-policy-document";
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
import { NetworkmanagerCoreNetworkPolicyAttachment } from "./.gen/providers/aws/networkmanager-core-network-policy-attachment";
import { NetworkmanagerGlobalNetwork } from "./.gen/providers/aws/networkmanager-global-network";
import { NetworkmanagerVpcAttachment } from "./.gen/providers/aws/networkmanager-vpc-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new NetworkmanagerGlobalNetwork(this, "example", {});
    const awsNetworkmanagerCoreNetworkExample = new NetworkmanagerCoreNetwork(
      this,
      "example_1",
      {
        basePolicyRegions: ["us-west-2", "us-east-1"],
        createBasePolicy: true,
        globalNetworkId: example.id,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkExample.overrideLogicalId("example");
    const exampleUsEast1 = new NetworkmanagerVpcAttachment(
      this,
      "example_us_east_1",
      {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        provider: "alternate",
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExampleUsEast1, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExampleUsEast1.arn),
      }
    );
    const exampleUsWest2 = new NetworkmanagerVpcAttachment(
      this,
      "example_us_west_2",
      {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        subnetArns: Token.asList(
          Fn.lookupNested(awsSubnetExampleUsWest2, ["*", "arn"])
        ),
        vpcArn: Token.asString(awsVpcExampleUsWest2.arn),
      }
    );
    const dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample =
      new DataAwsNetworkmanagerCoreNetworkPolicyDocument(this, "example_4", {
        coreNetworkConfiguration: [
          {
            asnRanges: ["65022-65534"],
            edgeLocations: [
              {
                location: "us-west-2",
              },
              {
                location: "us-east-1",
              },
            ],
          },
        ],
        segmentActions: [
          {
            action: "create-route",
            destinationCidrBlocks: ["10.0.0.0/16"],
            destinations: [exampleUsWest2.id],
            segment: "segment",
          },
          {
            action: "create-route",
            destinationCidrBlocks: ["10.1.0.0/16"],
            destinations: [exampleUsEast1.id],
            segment: "segment",
          },
        ],
        segments: [
          {
            name: "segment",
          },
          {
            name: "segment2",
          },
        ],
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.overrideLogicalId(
      "example"
    );
    const awsNetworkmanagerCoreNetworkPolicyAttachmentExample =
      new NetworkmanagerCoreNetworkPolicyAttachment(this, "example_5", {
        coreNetworkId: Token.asString(awsNetworkmanagerCoreNetworkExample.id),
        policyDocument: Token.asString(
          dataAwsNetworkmanagerCoreNetworkPolicyDocumentExample.json
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsNetworkmanagerCoreNetworkPolicyAttachmentExample.overrideLogicalId(
      "example"
    );
  }
}

```

## Argument Reference

The following arguments are required:

* `globalNetworkId` - (Required) ID of the global network that a core network will be a part of.

The following arguments are optional:

* `basePolicyDocument` - (Optional, conflicts with `basePolicyRegions`) Sets the base policy document for the core network. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
* `basePolicyRegions` - (Optional, conflicts with `basePolicyDocument`) List of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
* `createBasePolicy` - (Optional) Whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the [`aws_networkmanager_core_network_policy_attachment` resource](/docs/providers/aws/r/networkmanager_core_network_policy_attachment.html). This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Terraform snippet can be found above [for VPC Attachment in a single region](#with-vpc-attachment-single-region) and [for VPC Attachment multi-region](#with-vpc-attachment-multi-region). An example base policy is shown below. This base policy is overridden with the policy that you specify in the [`aws_networkmanager_core_network_policy_attachment` resource](/docs/providers/aws/r/networkmanager_core_network_policy_attachment.html).

```json
{
  "version": "2021.12",
  "core-network-configuration": {
    "asn-ranges": [
      "64512-65534"
    ],
    "vpn-ecmp-support": false,
    "edge-locations": [
      {
        "location": "us-east-1"
      }
    ]
  },
  "segments": [
    {
      "name": "segment",
      "description": "base-policy",
      "isolate-attachments": false,
      "require-attachment-acceptance": false
    }
  ]
}
```

* `description` - (Optional) Description of the Core Network.
* `tags` - (Optional) Key-value tags for the Core Network. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Core Network ARN.
* `createdAt` - Timestamp when a core network was created.
* `edges` - One or more blocks detailing the edges within a core network. [Detailed below](#edges).
* `id` - Core Network ID.
* `segments` - One or more blocks detailing the segments within a core network. [Detailed below](#segments).
* `state` - Current state of a core network.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

### `edges`

The `edges` configuration block supports the following arguments:

* `asn` - ASN of a core network edge.
* `edgeLocation` - Region where a core network edge is located.
* `insideCidrBlocks` - Inside IP addresses used for core network edges.

### `segments`

The `segments` configuration block supports the following arguments:

* `edgeLocations` - Regions where the edges are located.
* `name` - Name of a core network segment.
* `shared_segments` - Shared segments of a core network.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)
* `update` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_core_network` using the core network ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkmanagerCoreNetwork } from "./.gen/providers/aws/networkmanager-core-network";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    NetworkmanagerCoreNetwork.generateConfigForImport(
      this,
      "example",
      "core-network-0d47f6t230mz46dy4"
    );
  }
}

```

Using `terraform import`, import `aws_networkmanager_core_network` using the core network ID. For example:

```console
% terraform import aws_networkmanager_core_network.example core-network-0d47f6t230mz46dy4
```

<!-- cache-key: cdktf-0.20.8 input-a8f63cb5ad0be1f33999c52038490ce409174d1bad497bbe3c33f417d8bd250a -->