---
subcategory: "ElastiCache"
layout: "aws"
page_title: "AWS: aws_elasticache_cluster"
description: |-
  Provides an ElastiCache Cluster resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_elasticache_cluster

Provides an ElastiCache Cluster resource, which manages a Memcached cluster, a single-node Redis instance,
or a read replica in a Redis (Cluster Mode Enabled) replication group. For more information, refer to
the AWS document [What is Amazon ElastiCache?](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/WhatIs.html).

For working with Redis (Cluster Mode Enabled) replication groups, see the
[`aws_elasticache_replication_group` resource](/docs/providers/aws/r/elasticache_replication_group.html).

~> **Note:** When you change an attribute, such as `numCacheNodes`, by default
it is applied in the next maintenance window. Because of this, Terraform may report
a difference in its planning phase because the actual modification has not yet taken
place. You can use the `applyImmediately` flag to instruct the service to apply the
change immediately. Using `applyImmediately` can result in a brief downtime as the server reboots.
See the "Changes take effect" section of the "Details" column in the AWS Documentation on Engine specific parameters for
[ElastiCache for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Engine.html#ParameterGroups.Memcached) or
[ElastiCache for Valkey and Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Engine.html#ParameterGroups.Redis)
for more information.

~> **Note:** Any attribute changes that re-create the resource will be applied immediately, regardless of the value of `applyImmediately`.

## Example Usage

### Memcached Cluster

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheCluster(this, "example", {
      clusterId: "cluster-example",
      engine: "memcached",
      nodeType: "cache.m4.large",
      numCacheNodes: 2,
      parameterGroupName: "default.memcached1.4",
      port: 11211,
    });
  }
}

```

### Redis Instance

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheCluster(this, "example", {
      clusterId: "cluster-example",
      engine: "redis",
      engineVersion: "3.2.10",
      nodeType: "cache.m4.large",
      numCacheNodes: 1,
      parameterGroupName: "default.redis3.2",
      port: 6379,
    });
  }
}

```

### Redis Cluster Mode Disabled Read Replica Instance

These inherit their settings from the replication group.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheCluster(this, "replica", {
      clusterId: "cluster-example",
      replicationGroupId: example.id,
    });
  }
}

```

### Redis Log Delivery configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheCluster(this, "test", {
      applyImmediately: true,
      clusterId: "mycluster",
      engine: "redis",
      logDeliveryConfiguration: [
        {
          destination: example.name,
          destinationType: "cloudwatch-logs",
          logFormat: "text",
          logType: "slow-log",
        },
        {
          destination: Token.asString(
            awsKinesisFirehoseDeliveryStreamExample.name
          ),
          destinationType: "kinesis-firehose",
          logFormat: "json",
          logType: "engine-log",
        },
      ],
      nodeType: "cache.t3.micro",
      numCacheNodes: 1,
      port: 6379,
    });
  }
}

```

### Elasticache Cluster in Outpost

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsOutpostsOutpost } from "./.gen/providers/aws/data-aws-outposts-outpost";
import { DataAwsOutpostsOutposts } from "./.gen/providers/aws/data-aws-outposts-outposts";
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
import { ElasticacheSubnetGroup } from "./.gen/providers/aws/elasticache-subnet-group";
import { Subnet } from "./.gen/providers/aws/subnet";
import { Vpc } from "./.gen/providers/aws/vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Vpc(this, "example", {
      cidrBlock: "10.0.0.0/16",
    });
    const dataAwsOutpostsOutpostsExample = new DataAwsOutpostsOutposts(
      this,
      "example_1",
      {}
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsOutpostsOutpostsExample.overrideLogicalId("example");
    const awsSubnetExample = new Subnet(this, "example_2", {
      cidrBlock: "10.0.1.0/24",
      tags: {
        Name: "my-subnet",
      },
      vpcId: example.id,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSubnetExample.overrideLogicalId("example");
    const dataAwsOutpostsOutpostExample = new DataAwsOutpostsOutpost(
      this,
      "example_3",
      {
        id: Token.asString(
          Fn.lookupNested(Fn.tolist(dataAwsOutpostsOutpostsExample.ids), ["0"])
        ),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsOutpostsOutpostExample.overrideLogicalId("example");
    const awsElasticacheSubnetGroupExample = new ElasticacheSubnetGroup(
      this,
      "example_4",
      {
        name: "my-cache-subnet",
        subnetIds: [Token.asString(awsSubnetExample.id)],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsElasticacheSubnetGroupExample.overrideLogicalId("example");
    const awsElasticacheClusterExample = new ElasticacheCluster(
      this,
      "example_5",
      {
        clusterId: "cluster-example",
        engine: "memcached",
        nodeType: "cache.r5.large",
        numCacheNodes: 2,
        outpostMode: "single-outpost",
        parameterGroupName: "default.memcached1.4",
        port: 11211,
        preferredOutpostArn: Token.asString(dataAwsOutpostsOutpostExample.arn),
        subnetGroupName: Token.asString(awsElasticacheSubnetGroupExample.name),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsElasticacheClusterExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `clusterId` - (Required) Group identifier. ElastiCache converts this name to lowercase. Changing this value will re-create the resource.
* `engine` - (Optional, Required if `replicationGroupId` is not specified) Name of the cache engine to be used for this cache cluster. Valid values are `memcached`, `redis` and `valkey`.
* `nodeType` - (Required unless `replicationGroupId` is provided) The instance class used.
  See AWS documentation for information on [supported node types for Valkey or Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.CurrentGen) and [guidance on selecting node types for Valkey or Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SelectSize.html#CacheNodes.SelectSize.redis).
  See AWS documentation for information on [supported node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.CurrentGen-Memcached) and [guidance on selecting node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SelectSize.html#CacheNodes.SelectSize.Mem).
  For Memcached, changing this value will re-create the resource.
* `numCacheNodes` - (Required unless `replicationGroupId` is provided) The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
* `parameterGroupName` - (Required unless `replicationGroupId` is provided) The name of the parameter group to associate with this cache cluster.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `applyImmediately` - (Optional) Whether any database modifications are applied immediately, or during the next maintenance window. Default is `false`. See [Amazon ElastiCache Documentation for more information](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html).
* `autoMinorVersionUpgrade` - (Optional) Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
  Only supported for engine type `"redis"` and if the engine version is 6 or higher.
  Defaults to `true`.
* `availabilityZone` - (Optional) Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
* `azMode` - (Optional, Memcached only) Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`.
* `engineVersion` - (Optional) Version number of the cache engine to be used.
  If not set, defaults to the latest version.
  See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
  When `engine` is `redis` and the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
  When the version is 6, the major and minor version can be set, e.g., `6.2`,
  or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
  Otherwise, specify the full version desired, e.g., `5.0.6`.
  The actual engine version used is returned in the attribute `engineVersionActual`, see [Attribute Reference](#attribute-reference) below. Cannot be provided with `replication_group_id.`
* `finalSnapshotIdentifier` - (Optional, Redis only) Name of your final cluster snapshot. If omitted, no final snapshot will be made.
* `ipDiscovery` - (Optional) The IP version to advertise in the discovery protocol. Valid values are `ipv4` or `ipv6`.
* `logDeliveryConfiguration` - (Optional, Redis only) Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Log_Delivery.html). See [Log Delivery Configuration](#log-delivery-configuration) below for more details.
* `maintenanceWindow` - (Optional) Specifies the weekly time range for when maintenance
on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`.
* `networkType` - (Optional) The IP versions for cache cluster connections. IPv6 is supported with Redis engine `6.2` onword or Memcached version `1.6.6` for all [Nitro system](https://aws.amazon.com/ec2/nitro/) instances. Valid values are `ipv4`, `ipv6` or `dual_stack`.
* `notificationTopicArn` - (Optional) ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`.
* `outpostMode` - (Optional) Specify the outpost mode that will apply to the cache cluster creation. Valid values are `"single-outpost"` and `"cross-outpost"`, however AWS currently only supports `"single-outpost"` mode.
* `port` - (Optional) The port number on which each of the cache nodes will accept connections. For Memcached the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`. Changing this value will re-create the resource.
* `preferredAvailabilityZones` - (Optional, Memcached only) List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
* `preferredOutpostArn` - (Optional, Required if `outpostMode` is specified) The outpost ARN in which the cache cluster will be created.
* `replicationGroupId` - (Optional, Required if `engine` is not specified) ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
* `securityGroupIds` - (Optional, VPC only) One or more VPC security groups associated with the cache cluster. Cannot be provided with `replication_group_id.`
* `snapshotArns` - (Optional, Redis only) Single-element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing `snapshotArns` forces a new resource.
* `snapshotName` - (Optional, Redis only) Name of a snapshot from which to restore data into the new node group. Changing `snapshotName` forces a new resource.
* `snapshotRetentionLimit` - (Optional, Redis only) Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro cache nodes
* `snapshotWindow` - (Optional, Redis only) Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
* `subnetGroupName` - (Optional, VPC only) Name of the subnet group to be used for the cache cluster. Changing this value will re-create the resource. Cannot be provided with `replication_group_id.`
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `transitEncryptionEnabled` - (Optional) Enable encryption in-transit. Supported with Memcached versions `1.6.12` and later, Valkey `7.2` and later, Redis OSS versions `3.2.6`, `4.0.10` and later, running in a VPC. See the [ElastiCache in-transit encryption documentation](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/in-transit-encryption.html#in-transit-encryption-constraints) for more details.

### Log Delivery Configuration

The `logDeliveryConfiguration` block allows the streaming of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Log_Delivery.html#Log_contents-engine-log) to CloudWatch Logs or Kinesis Data Firehose. Max of 2 blocks.

* `destination` - Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
* `destinationType` - For CloudWatch Logs use `cloudwatch-logs` or for Kinesis Data Firehose use `kinesis-firehose`.
* `logFormat` - Valid values are `json` or `text`
* `logType` - Valid values are  `slow-log` or `engine-log`. Max 1 of each.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the created ElastiCache Cluster.
* `engineVersionActual` - Because ElastiCache pulls the latest minor or patch for a version, this attribute returns the running version of the cache engine.
* `cacheNodes` - List of node objects including `id`, `address`, `port` and `availabilityZone`.
* `clusterAddress` - (Memcached only) DNS name of the cache cluster without the port appended.
* `configurationEndpoint` - (Memcached only) Configuration endpoint to allow host discovery.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `40m`)
- `update` - (Default `80m`)
- `delete` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ElastiCache Clusters using the `clusterId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheCluster } from "./.gen/providers/aws/elasticache-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ElasticacheCluster.generateConfigForImport(this, "myCluster", "my_cluster");
  }
}

```

Using `terraform import`, import ElastiCache Clusters using the `clusterId`. For example:

```console
% terraform import aws_elasticache_cluster.my_cluster my_cluster
```

<!-- cache-key: cdktf-0.20.8 input-8d8c23e5e710e1004b7d477916fd27bb248f9f84ac8f1b8e0446bb995f1447da -->