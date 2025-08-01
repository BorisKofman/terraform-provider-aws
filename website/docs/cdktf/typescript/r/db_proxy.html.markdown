---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_db_proxy"
description: |-
  Provides an RDS DB proxy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_db_proxy

Provides an RDS DB proxy resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html).

~> **Note:** Not all Availability Zones (AZs) support DB proxies. Specifying `vpcSubnetIds` for AZs that do not support proxies will not trigger an error as long as at least one `vpc_subnet_id` is valid. However, this will cause Terraform to continuously detect differences between the configuration and the actual infrastructure. Refer to the [Unsupported Availability Zones](#unsupported-availability-zones) section below for potential workarounds.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbProxy } from "./.gen/providers/aws/db-proxy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DbProxy(this, "example", {
      auth: [
        {
          authScheme: "SECRETS",
          description: "example",
          iamAuth: "DISABLED",
          secretArn: Token.asString(awsSecretsmanagerSecretExample.arn),
        },
      ],
      debugLogging: false,
      engineFamily: "MYSQL",
      idleClientTimeout: 1800,
      name: "example",
      requireTls: true,
      roleArn: Token.asString(awsIamRoleExample.arn),
      tags: {
        Key: "value",
        Name: "example",
      },
      vpcSecurityGroupIds: [Token.asString(awsSecurityGroupExample.id)],
      vpcSubnetIds: [Token.asString(awsSubnetExample.id)],
    });
  }
}

```

### Unsupported Availability Zones

Terraform may report constant differences if you use `vpcSubnetIds` that correspond to Availability Zones (AZs) that do not support a DB proxy. While this typically does not result in an error, AWS only returns `vpcSubnetIds` for AZs that support DB proxies. As a result, Terraform detects a mismatch between your configuration and the actual infrastructure, leading it to report that changes are required. Below are some ways to avoid this issue.

One solution is to exclude AZs that do not support DB proxies by using the [`aws_availability_zones` data source](/docs/providers/aws/d/availability_zones.html). The example below demonstrates how to configure this for the `us-east-1` region, excluding the `use1-az3` AZ. (Keep in mind that AZ names can vary between accounts, while AZ IDs remain consistent.) If the `us-east-1` region has six AZs in total and you aim to configure the maximum number of subnets, you would exclude one AZ and configure five subnets:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformCount, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsAvailabilityZones } from "./.gen/providers/aws/data-aws-availability-zones";
import { DbProxy } from "./.gen/providers/aws/db-proxy";
import { Subnet } from "./.gen/providers/aws/subnet";
import { Vpc } from "./.gen/providers/aws/vpc";
interface MyConfig {
  auth: any;
  engineFamily: any;
  roleArn: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    const example = new Vpc(this, "example", {
      cidrBlock: "10.0.0.0/16",
    });
    const available = new DataAwsAvailabilityZones(this, "available", {
      excludeZoneIds: ["use1-az3"],
      filter: [
        {
          name: "opt-in-status",
          values: ["opt-in-not-required"],
        },
      ],
      state: "available",
    });
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const exampleCount = TerraformCount.of(Token.asNumber("5"));
    const awsSubnetExample = new Subnet(this, "example_2", {
      availabilityZone: Token.asString(
        Fn.lookupNested(available.names, [exampleCount.index])
      ),
      cidrBlock: Token.asString(
        Fn.cidrsubnet(example.cidrBlock, 8, Token.asNumber(exampleCount.index))
      ),
      vpcId: example.id,
      count: exampleCount,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSubnetExample.overrideLogicalId("example");
    const awsDbProxyExample = new DbProxy(this, "example_3", {
      name: "example",
      vpcSubnetIds: [Token.asString(awsSubnetExample.id)],
      auth: config.auth,
      engineFamily: config.engineFamily,
      roleArn: config.roleArn,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDbProxyExample.overrideLogicalId("example");
  }
}

```

Another approach is to use the [`lifecycle` `ignore_changes`](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes) meta-argument. With this method, Terraform will stop detecting differences for the `vpcSubnetIds` argument. However, note that this approach disables Terraform's ability to track and manage updates to `vpcSubnetIds`, so use it carefully to avoid unintended drift between your configuration and the actual infrastructure.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbProxy } from "./.gen/providers/aws/db-proxy";
interface MyConfig {
  auth: any;
  engineFamily: any;
  roleArn: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new DbProxy(this, "example", {
      lifecycle: {
        ignoreChanges: [vpcSubnetIds],
      },
      name: "example",
      vpcSubnetIds: [Token.asString(awsSubnetExample.id)],
      auth: config.auth,
      engineFamily: config.engineFamily,
      roleArn: config.roleArn,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
* `auth` - (Required) Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
* `debugLogging` - (Optional) Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
* `engineFamily` - (Required, Forces new resource) The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
* `idleClientTimeout` - (Optional) The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
* `requireTls` - (Optional) A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
* `roleArn` - (Required) The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
* `vpcSecurityGroupIds` - (Optional) One or more VPC security group IDs to associate with the new proxy.
* `vpcSubnetIds` - (Required) One or more VPC subnet IDs to associate with the new proxy.
* `tags` - (Optional) A mapping of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

`auth` blocks support the following:

* `authScheme` - (Optional) The type of authentication that the proxy uses for connections from the proxy to the underlying database. One of `SECRETS`.
* `clientPasswordAuthType` - (Optional) The type of authentication the proxy uses for connections from clients. Valid values are `MYSQL_CACHING_SHA2_PASSWORD`, `MYSQL_NATIVE_PASSWORD`, `POSTGRES_SCRAM_SHA_256`, `POSTGRES_MD5`, and `SQL_SERVER_AUTHENTICATION`.
* `description` - (Optional) A user-specified description about the authentication used by a proxy to log in as a specific database user.
* `iamAuth` - (Optional) Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of `DISABLED`, `REQUIRED`.
* `secretArn` - (Optional) The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.
* `username` - (Optional) The name of the database user to which the proxy connects.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) for the proxy.
* `arn` - The Amazon Resource Name (ARN) for the proxy.
* `endpoint` - The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `30m`)
- `update` - (Default `30m`)
- `delete` - (Default `60m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DB proxies using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbProxy } from "./.gen/providers/aws/db-proxy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DbProxy.generateConfigForImport(this, "example", "example");
  }
}

```

Using `terraform import`, import DB proxies using the `name`. For example:

```console
% terraform import aws_db_proxy.example example
```

<!-- cache-key: cdktf-0.20.8 input-2d1741eae1f95335d771f94d7ec8de417810561a472b6d0f5621e6c7cc32a883 -->