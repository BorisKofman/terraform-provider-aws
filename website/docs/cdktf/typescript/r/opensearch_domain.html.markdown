---
subcategory: "OpenSearch"
layout: "aws"
page_title: "AWS: aws_opensearch_domain"
description: |-
  Terraform resource for managing an AWS OpenSearch Domain.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_opensearch_domain

Manages an Amazon OpenSearch Domain.

## Elasticsearch vs. OpenSearch

Amazon OpenSearch Service is the successor to Amazon Elasticsearch Service and supports OpenSearch and legacy Elasticsearch OSS (up to 7.10, the final open source version of the software).

OpenSearch Domain configurations are similar in many ways to Elasticsearch Domain configurations. However, there are important differences including these:

* OpenSearch has `engineVersion` while Elasticsearch has `elasticsearchVersion`
* Versions are specified differently - _e.g._, `Elasticsearch_7.10` with OpenSearch vs. `7.10` for Elasticsearch.
* `instanceType` argument values end in `search` for OpenSearch vs. `elasticsearch` for Elasticsearch (_e.g._, `t2.micro.search` vs. `t2.micro.elasticsearch`).
* The AWS-managed service-linked role for OpenSearch is called `AWSServiceRoleForAmazonOpenSearchService` instead of `AWSServiceRoleForAmazonElasticsearchService` for Elasticsearch.

There are also some potentially unexpected similarities in configurations:

* ARNs for both are prefaced with `arn:aws:es:`.
* Both OpenSearch and Elasticsearch use assume role policies that refer to the `Principal` `Service` as `es.amazonaws.com`.
* IAM policy actions, such as those you will find in `accessPolicies`, are prefaced with `es:` for both.

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
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new OpensearchDomain(this, "example", {
      clusterConfig: {
        instanceType: "r4.large.search",
      },
      domainName: "example",
      engineVersion: "Elasticsearch_7.10",
      tags: {
        Domain: "TestDomain",
      },
    });
  }
}

```

### Access Policy

-> See also: [`aws_opensearch_domain_policy` resource](/docs/providers/aws/r/opensearch_domain_policy.html)

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformVariable, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsCallerIdentity } from "./.gen/providers/aws/data-aws-caller-identity";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const domain = new TerraformVariable(this, "domain", {
      default: "tf-test",
    });
    const current = new DataAwsCallerIdentity(this, "current", {});
    const dataAwsRegionCurrent = new DataAwsRegion(this, "current_2", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsRegionCurrent.overrideLogicalId("current");
    const example = new DataAwsIamPolicyDocument(this, "example", {
      statement: [
        {
          actions: ["es:*"],
          condition: [
            {
              test: "IpAddress",
              values: ["66.193.100.22/32"],
              variable: "aws:SourceIp",
            },
          ],
          effect: "Allow",
          principals: [
            {
              identifiers: ["*"],
              type: "*",
            },
          ],
          resources: [
            "arn:aws:es:${" +
              dataAwsRegionCurrent.region +
              "}:${" +
              current.accountId +
              "}:domain/${" +
              domain.value +
              "}/*",
          ],
        },
      ],
    });
    const awsOpensearchDomainExample = new OpensearchDomain(this, "example_4", {
      accessPolicies: Token.asString(example.json),
      domainName: domain.stringValue,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsOpensearchDomainExample.overrideLogicalId("example");
  }
}

```

### Log publishing to CloudWatch Logs

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchLogGroup } from "./.gen/providers/aws/cloudwatch-log-group";
import { CloudwatchLogResourcePolicy } from "./.gen/providers/aws/cloudwatch-log-resource-policy";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
interface MyConfig {
  domainName: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    const example = new CloudwatchLogGroup(this, "example", {
      name: "example",
    });
    const awsOpensearchDomainExample = new OpensearchDomain(this, "example_1", {
      logPublishingOptions: [
        {
          cloudwatchLogGroupArn: example.arn,
          logType: "INDEX_SLOW_LOGS",
        },
      ],
      domainName: config.domainName,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsOpensearchDomainExample.overrideLogicalId("example");
    const dataAwsIamPolicyDocumentExample = new DataAwsIamPolicyDocument(
      this,
      "example_2",
      {
        statement: [
          {
            actions: [
              "logs:PutLogEvents",
              "logs:PutLogEventsBatch",
              "logs:CreateLogStream",
            ],
            effect: "Allow",
            principals: [
              {
                identifiers: ["es.amazonaws.com"],
                type: "Service",
              },
            ],
            resources: ["arn:aws:logs:*"],
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsIamPolicyDocumentExample.overrideLogicalId("example");
    const awsCloudwatchLogResourcePolicyExample =
      new CloudwatchLogResourcePolicy(this, "example_3", {
        policyDocument: Token.asString(dataAwsIamPolicyDocumentExample.json),
        policyName: "example",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsCloudwatchLogResourcePolicyExample.overrideLogicalId("example");
  }
}

```

### VPC based OpenSearch

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformVariable, Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsCallerIdentity } from "./.gen/providers/aws/data-aws-caller-identity";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { DataAwsSubnets } from "./.gen/providers/aws/data-aws-subnets";
import { DataAwsVpc } from "./.gen/providers/aws/data-aws-vpc";
import { IamServiceLinkedRole } from "./.gen/providers/aws/iam-service-linked-role";
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
import { SecurityGroup } from "./.gen/providers/aws/security-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const domain = new TerraformVariable(this, "domain", {
      default: "tf-test",
    });
    const vpc = new TerraformVariable(this, "vpc", {});
    const example = new IamServiceLinkedRole(this, "example", {
      awsServiceName: "opensearchservice.amazonaws.com",
    });
    const current = new DataAwsCallerIdentity(this, "current", {});
    const dataAwsRegionCurrent = new DataAwsRegion(this, "current_4", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsRegionCurrent.overrideLogicalId("current");
    const dataAwsVpcExample = new DataAwsVpc(this, "example_5", {
      tags: {
        Name: vpc.stringValue,
      },
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsVpcExample.overrideLogicalId("example");
    const awsSecurityGroupExample = new SecurityGroup(this, "example_6", {
      description: "Managed by Terraform",
      ingress: [
        {
          cidrBlocks: [Token.asString(dataAwsVpcExample.cidrBlock)],
          fromPort: 443,
          protocol: "tcp",
          toPort: 443,
        },
      ],
      name: "${" + vpc.value + "}-opensearch-${" + domain.value + "}",
      vpcId: Token.asString(dataAwsVpcExample.id),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSecurityGroupExample.overrideLogicalId("example");
    const dataAwsIamPolicyDocumentExample = new DataAwsIamPolicyDocument(
      this,
      "example_7",
      {
        statement: [
          {
            actions: ["es:*"],
            effect: "Allow",
            principals: [
              {
                identifiers: ["*"],
                type: "*",
              },
            ],
            resources: [
              "arn:aws:es:${" +
                dataAwsRegionCurrent.region +
                "}:${" +
                current.accountId +
                "}:domain/${" +
                domain.value +
                "}/*",
            ],
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsIamPolicyDocumentExample.overrideLogicalId("example");
    const dataAwsSubnetsExample = new DataAwsSubnets(this, "example_8", {
      filter: [
        {
          name: "vpc-id",
          values: [Token.asString(dataAwsVpcExample.id)],
        },
      ],
      tags: {
        Tier: "private",
      },
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsSubnetsExample.overrideLogicalId("example");
    const awsOpensearchDomainExample = new OpensearchDomain(this, "example_9", {
      accessPolicies: Token.asString(dataAwsIamPolicyDocumentExample.json),
      advancedOptions: {
        "rest.action.multi.allow_explicit_index": "true",
      },
      clusterConfig: {
        instanceType: "m4.large.search",
        zoneAwarenessEnabled: true,
      },
      dependsOn: [example],
      domainName: domain.stringValue,
      engineVersion: "OpenSearch_1.0",
      tags: {
        Domain: "TestDomain",
      },
      vpcOptions: {
        securityGroupIds: [Token.asString(awsSecurityGroupExample.id)],
        subnetIds: [
          Token.asString(Fn.lookupNested(dataAwsSubnetsExample.ids, ["0"])),
          Token.asString(Fn.lookupNested(dataAwsSubnetsExample.ids, ["1"])),
        ],
      },
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsOpensearchDomainExample.overrideLogicalId("example");
  }
}

```

### Enabling fine-grained access control on an existing domain

This example shows two configurations: one to create a domain without fine-grained access control and the second to modify the domain to enable fine-grained access control. For more information, see [Enabling fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html).

#### First apply

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new OpensearchDomain(this, "example", {
      advancedSecurityOptions: {
        anonymousAuthEnabled: true,
        enabled: false,
        internalUserDatabaseEnabled: true,
        masterUserOptions: {
          masterUserName: "example",
          masterUserPassword: "Barbarbarbar1!",
        },
      },
      clusterConfig: {
        instanceType: "r5.large.search",
      },
      domainEndpointOptions: {
        enforceHttps: true,
        tlsSecurityPolicy: "Policy-Min-TLS-1-2-2019-07",
      },
      domainName: "ggkitty",
      ebsOptions: {
        ebsEnabled: true,
        volumeSize: 10,
      },
      encryptAtRest: {
        enabled: true,
      },
      engineVersion: "Elasticsearch_7.1",
      nodeToNodeEncryption: {
        enabled: true,
      },
    });
  }
}

```

#### Second apply

Notice that the only change is `advanced_security_options.0.enabled` is now set to `true`.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new OpensearchDomain(this, "example", {
      advancedSecurityOptions: {
        anonymousAuthEnabled: true,
        enabled: true,
        internalUserDatabaseEnabled: true,
        masterUserOptions: {
          masterUserName: "example",
          masterUserPassword: "Barbarbarbar1!",
        },
      },
      clusterConfig: {
        instanceType: "r5.large.search",
      },
      domainEndpointOptions: {
        enforceHttps: true,
        tlsSecurityPolicy: "Policy-Min-TLS-1-2-2019-07",
      },
      domainName: "ggkitty",
      ebsOptions: {
        ebsEnabled: true,
        volumeSize: 10,
      },
      encryptAtRest: {
        enabled: true,
      },
      engineVersion: "Elasticsearch_7.1",
      nodeToNodeEncryption: {
        enabled: true,
      },
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `domainName` - (Required) Name of the domain.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `accessPolicies` - (Optional) IAM policy document specifying the access policies for the domain.
* `advancedOptions` - (Optional) Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing Terraform to want to recreate your OpenSearch domain on every apply.
* `advancedSecurityOptions` - (Optional) Configuration block for [fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html). Detailed below.
* `autoTuneOptions` - (Optional) Configuration block for the Auto-Tune options of the domain. Detailed below.
* `clusterConfig` - (Optional) Configuration block for the cluster of the domain. Detailed below.
* `cognitoOptions` - (Optional) Configuration block for authenticating dashboard with Cognito. Detailed below.
* `domainEndpointOptions` - (Optional) Configuration block for domain endpoint HTTP(S) related options. Detailed below.
* `ebsOptions` - (Optional) Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/opensearch-service/pricing/). Detailed below.
* `engineVersion` - (Optional) Either `Elasticsearch_X.Y` or `OpenSearch_X.Y` to specify the engine version for the Amazon OpenSearch Service domain. For example, `OpenSearch_1.0` or `Elasticsearch_7.9`.
  See [Creating and managing Amazon OpenSearch Service domains](http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
  Defaults to the lastest version of OpenSearch.
* `ipAddressType` - (Optional) The IP address type for the endpoint. Valid values are `ipv4` and `dualstack`.
* `encryptAtRest` - (Optional) Configuration block for encrypt at rest options. Only available for [certain instance types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html). Detailed below.
* `logPublishingOptions` - (Optional) Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
* `nodeToNodeEncryption` - (Optional) Configuration block for node-to-node encryption options. Detailed below.
* `snapshotOptions` - (Optional) Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
* `softwareUpdateOptions` - (Optional) Software update options for the domain. Detailed below.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpcOptions` - (Optional) Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)). Detailed below.
* `offPeakWindowOptions` - (Optional) Configuration to add Off Peak update options. ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)). Detailed below.

### advanced_security_options

* `anonymousAuthEnabled` - (Optional) Whether Anonymous auth is enabled. Enables fine-grained access control on an existing domain. Ignored unless `advancedSecurityOptions` are enabled. _Can only be enabled on an existing domain._
* `enabled` - (Required, Forces new resource when changing from `true` to `false`) Whether advanced security is enabled.
* `internalUserDatabaseEnabled` - (Optional) Whether the internal user database is enabled. Default is `false`.
* `masterUserOptions` - (Optional) Configuration block for the main user. Detailed below.

#### master_user_options

* `masterUserArn` - (Optional) ARN for the main user. Only specify if `internalUserDatabaseEnabled` is not set or set to `false`.
* `masterUserName` - (Optional) Main user's username, which is stored in the Amazon OpenSearch Service domain's internal database. Only specify if `internalUserDatabaseEnabled` is set to `true`.
* `masterUserPassword` - (Optional) Main user's password, which is stored in the Amazon OpenSearch Service domain's internal database. Only specify if `internalUserDatabaseEnabled` is set to `true`.

### auto_tune_options

* `desiredState` - (Required) Auto-Tune desired state for the domain. Valid values: `ENABLED` or `DISABLED`.
* `maintenanceSchedule` - (Required if `rollbackOnDisable` is set to `DEFAULT_ROLLBACK`) Configuration block for Auto-Tune maintenance windows. Can be specified multiple times for each maintenance window. Detailed below.

  **NOTE:** Maintenance windows are deprecated and have been replaced with [off-peak windows](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html). Consequently, `maintenanceSchedule` configuration blocks cannot be specified when `useOffPeakWindow` is set to `true`.
* `rollbackOnDisable` - (Optional) Whether to roll back to default Auto-Tune settings when disabling Auto-Tune. Valid values: `DEFAULT_ROLLBACK` or `NO_ROLLBACK`.
* `useOffPeakWindow` - (Optional) Whether to schedule Auto-Tune optimizations that require blue/green deployments during the domain's configured daily off-peak window. Defaults to `false`.

#### maintenance_schedule

* `startAt` - (Required) Date and time at which to start the Auto-Tune maintenance schedule in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
* `duration` - (Required) Configuration block for the duration of the Auto-Tune maintenance window. Detailed below.
* `cronExpressionForRecurrence` - (Required) A cron expression specifying the recurrence pattern for an Auto-Tune maintenance schedule.

##### duration

* `value` - (Required) An integer specifying the value of the duration of an Auto-Tune maintenance window.
* `unit` - (Required) Unit of time specifying the duration of an Auto-Tune maintenance window. Valid values: `HOURS`.

### cluster_config

* `coldStorageOptions` - (Optional) Configuration block containing cold storage configuration. Detailed below.
* `dedicatedMasterCount` - (Optional) Number of dedicated main nodes in the cluster.
* `dedicatedMasterEnabled` - (Optional) Whether dedicated main nodes are enabled for the cluster.
* `dedicatedMasterType` - (Optional) Instance type of the dedicated main nodes in the cluster.
* `instanceCount` - (Optional) Number of instances in the cluster.
* `instanceType` - (Optional) Instance type of data nodes in the cluster.
* `multiAzWithStandbyEnabled` - (Optional) Whether a multi-AZ domain is turned on with a standby AZ. For more information, see [Configuring a multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html).
* `nodeOptions` - (Optional) List of node options for the domain.
* `warmCount` - (Optional) Number of warm nodes in the cluster. Valid values are between `2` and `150`. `warmCount` can be only and must be set when `warmEnabled` is set to `true`.
* `warmEnabled` - (Optional) Whether to enable warm storage.
* `warmType` - (Optional) Instance type for the OpenSearch cluster's warm nodes. Valid values are `ultrawarm1.medium.search`, `ultrawarm1.large.search` and `ultrawarm1.xlarge.search`. `warmType` can be only and must be set when `warmEnabled` is set to `true`.
* `zoneAwarenessConfig` - (Optional) Configuration block containing zone awareness settings. Detailed below.
* `zoneAwarenessEnabled` - (Optional) Whether zone awareness is enabled, set to `true` for multi-az deployment. To enable awareness with three Availability Zones, the `availabilityZoneCount` within the `zoneAwarenessConfig` must be set to `3`.

#### node_options

Container object to specify configuration for a node type.

* `nodeConfig` - (Optional) Container to specify sizing of a node type.
* `nodeType` - (Optional) Type of node this configuration describes. Valid values: `coordinator`.

#### node_config

* `count` - (Optional) Number of nodes of a particular node type in the cluster.
* `enabled` - (Optional) Whether a particular node type is enabled.
* `type` - (Optional) The instance type of a particular node type in the cluster.

#### cold_storage_options

* `enabled` - (Optional) Boolean to enable cold storage for an OpenSearch domain. Defaults to `false`. Master and ultrawarm nodes must be enabled for cold storage.

#### zone_awareness_config

* `availabilityZoneCount` - (Optional) Number of Availability Zones for the domain to use with `zoneAwarenessEnabled`. Defaults to `2`. Valid values: `2` or `3`.

### cognito_options

AWS documentation: [Amazon Cognito Authentication for Dashboard](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/es-cognito-auth.html)

* `enabled` - (Optional) Whether Amazon Cognito authentication with Dashboard is enabled or not. Default is `false`.
* `identityPoolId` - (Required) ID of the Cognito Identity Pool to use.
* `roleArn` - (Required) ARN of the IAM role that has the AmazonOpenSearchServiceCognitoAccess policy attached.
* `userPoolId` - (Required) ID of the Cognito User Pool to use.

### domain_endpoint_options

* `customEndpointCertificateArn` - (Optional) ACM certificate ARN for your custom endpoint.
* `customEndpointEnabled` - (Optional) Whether to enable custom endpoint for the OpenSearch domain.
* `customEndpoint` - (Optional) Fully qualified domain for your custom endpoint.
* `enforceHttps` - (Optional) Whether or not to require HTTPS. Defaults to `true`.
* `tlsSecurityPolicy` - (Optional) Name of the TLS security policy that needs to be applied to the HTTPS endpoint. For valid values, refer to the [AWS documentation](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainEndpointOptions.html#opensearchservice-Type-DomainEndpointOptions-TLSSecurityPolicy). Terraform will only perform drift detection if a configuration value is provided.

### ebs_options

* `ebsEnabled` - (Required) Whether EBS volumes are attached to data nodes in the domain.
* `iops` - (Optional) Baseline input/output (I/O) performance of EBS volumes attached to data nodes. Applicable only for the GP3 and Provisioned IOPS EBS volume types.
* `throughput` - (Required if `volumeType` is set to `gp3`) Specifies the throughput (in MiB/s) of the EBS volumes attached to data nodes. Applicable only for the gp3 volume type.
* `volumeSize` - (Required if `ebsEnabled` is set to `true`.) Size of EBS volumes attached to data nodes (in GiB).
* `volumeType` - (Optional) Type of EBS volumes attached to data nodes.

### encrypt_at_rest

~> **Note:** You can enable `encryptAtRest` _in place_ for an existing, unencrypted domain only if you are using OpenSearch or your Elasticsearch version is 6.7 or greater. For other versions, if you enable `encryptAtRest`, Terraform with recreate the domain, potentially causing data loss. For any version, if you disable `encryptAtRest` for an existing, encrypted domain, Terraform will recreate the domain, potentially causing data loss. If you change the `kmsKeyId`, Terraform will also recreate the domain, potentially causing data loss.

* `enabled` - (Required) Whether to enable encryption at rest. If the `encryptAtRest` block is not provided then this defaults to `false`. Enabling encryption on new domains requires an `engineVersion` of `OpenSearch_X.Y` or `Elasticsearch_5.1` or greater.
* `kmsKeyId` - (Optional) KMS key ARN to encrypt the Elasticsearch domain with. If not specified then it defaults to using the `aws/es` service KMS key. Note that KMS will accept a KMS key ID but will return the key ARN. To prevent Terraform detecting unwanted changes, use the key ARN instead.

### log_publishing_options

* `cloudwatchLogGroupArn` - (Required) ARN of the Cloudwatch log group to which log needs to be published.
* `enabled` - (Optional, Default: true) Whether given log publishing option is enabled or not.
* `logType` - (Required) Type of OpenSearch log. Valid values: `INDEX_SLOW_LOGS`, `SEARCH_SLOW_LOGS`, `ES_APPLICATION_LOGS`, `AUDIT_LOGS`.

### node_to_node_encryption

~> **Note:** You can enable `nodeToNodeEncryption` _in place_ for an existing, unencrypted domain only if you are using OpenSearch or your Elasticsearch version is 6.7 or greater. For other versions, if you enable `nodeToNodeEncryption`, Terraform will recreate the domain, potentially causing data loss. For any version, if you disable `nodeToNodeEncryption` for an existing, node-to-node encrypted domain, Terraform will recreate the domain, potentially causing data loss.

* `enabled` - (Required) Whether to enable node-to-node encryption. If the `nodeToNodeEncryption` block is not provided then this defaults to `false`. Enabling node-to-node encryption of a new domain requires an `engineVersion` of `OpenSearch_X.Y` or `Elasticsearch_6.0` or greater.

### snapshot_options

* `automatedSnapshotStartHour` - (Required) Hour during which the service takes an automated daily snapshot of the indices in the domain.

### software_update_options

* `autoSoftwareUpdateEnabled` - (Optional) Whether automatic service software updates are enabled for the domain. Defaults to `false`.

### vpc_options

AWS documentation: [VPC Support for Amazon OpenSearch Service Domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/es-vpc.html)

~> **Note:** You must have created the service linked role for the OpenSearch service to use `vpcOptions`. If you need to create the service linked role at the same time as the OpenSearch domain then you must use `dependsOn` to make sure that the role is created before the OpenSearch domain. See the [VPC based ES domain example](#vpc-based-opensearch) above.

-> Security Groups and Subnets referenced in these attributes must all be within the same VPC. This determines what VPC the endpoints are created in.

* `securityGroupIds` - (Optional) List of VPC Security Group IDs to be applied to the OpenSearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
* `subnetIds` - (Required) List of VPC Subnet IDs for the OpenSearch domain endpoints to be created in.

### off_peak_window_options

AWS documentation: [Off Peak Hours Support for Amazon OpenSearch Service Domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)

* `enabled` - (Optional) Enabled disabled toggle for off-peak update window.
* `offPeakWindow` - (Optional)
    * `windowStartTime` - (Optional) 10h window for updates
        * `hours` - (Required) Starting hour of the 10-hour window for updates
        * `minutes` - (Required) Starting minute of the 10-hour window for updates

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the domain.
* `domainEndpointV2HostedZoneId` -  Dual stack hosted zone ID for the domain.
* `domainId` - Unique identifier for the domain.
* `domainName` - Name of the OpenSearch domain.
* `endpoint` - Domain-specific endpoint used to submit index, search, and data upload requests.
* `endpointV2` - V2 domain endpoint that works with both IPv4 and IPv6 addresses, used to submit index, search, and data upload requests.
* `dashboardEndpoint` - Domain-specific endpoint for Dashboard without https scheme.
* `dashboardEndpointV2` - V2 domain endpoint for Dashboard that works with both IPv4 and IPv6 addresses, without https scheme.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
* `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `90m`)
* `update` - (Default `180m`)
* `delete` - (Default `90m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OpenSearch domains using the `domainName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { OpensearchDomain } from "./.gen/providers/aws/opensearch-domain";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    OpensearchDomain.generateConfigForImport(this, "example", "domain_name");
  }
}

```

Using `terraform import`, import OpenSearch domains using the `domainName`. For example:

```console
% terraform import aws_opensearch_domain.example domain_name
```

<!-- cache-key: cdktf-0.20.8 input-11e323fefac14842b321e31a1ee1224c5019dc60471c25e358c03918db0afcc7 -->