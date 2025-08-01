---
subcategory: "Lake Formation"
layout: "aws"
page_title: "AWS: aws_lakeformation_resource_lf_tag"
description: |-
  Terraform resource for managing an AWS Lake Formation Resource LF Tag.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lakeformation_resource_lf_tag

Terraform resource for managing an AWS Lake Formation Resource LF Tag.

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
import { LakeformationResourceLfTag } from "./.gen/providers/aws/lakeformation-resource-lf-tag";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LakeformationResourceLfTag(this, "example", {
      database: [
        {
          name: Token.asString(awsGlueCatalogDatabaseExample.name),
        },
      ],
      lfTag: [
        {
          key: Token.asString(awsLakeformationLfTagExample.key),
          value: "stowe",
        },
      ],
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `lfTag` - (Required) Set of LF-tags to attach to the resource. See [LF Tag](#lf-tag) for more details.

Exactly one of the following is required:

* `database` - (Optional) Configuration block for a database resource. See [Database](#database) for more details.
* `table` - (Optional) Configuration block for a table resource. See [Table](#table) for more details.
* `tableWithColumns` - (Optional) Configuration block for a table with columns resource. See [Table With Columns](#table-with-columns) for more details.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalogId` - (Optional) Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.

### LF Tag

The following arguments are required:

* `key` - (Required) Key name for an existing LF-tag.
* `value` - (Required) Value from the possible values for the LF-tag.

The following argument is optional:

* `catalogId` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Database

The following argument is required:

* `name` - (Required) Name of the database resource. Unique to the Data Catalog.

The following argument is optional:

* `catalogId` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Table

The following argument is required:

* `databaseName` - (Required) Name of the database for the table. Unique to a Data Catalog.
* `name` - (Required, at least one of `name` or `wildcard`) Name of the table.
* `wildcard` - (Required, at least one of `name` or `wildcard`) Whether to use a wildcard representing every table under a database. Defaults to `false`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalogId` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Table With Columns

The following arguments are required:

* `columnNames` - (Required, at least one of `columnNames` or `wildcard`) Set of column names for the table.
* `databaseName` - (Required) Name of the database for the table with columns resource. Unique to the Data Catalog.
* `name` - (Required) Name of the table resource.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalogId` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.
* `columnWildcard` - (Optional) Option to add column wildcard. See [Column Wildcard](#column-wildcard) for more details.

### Column Wildcard

* `excludedColumnNames` - (Optional) Set of column names for the table to exclude. If `excludedColumnNames` is included, `wildcard` must be set to `true` to avoid Terraform reporting a difference.

## Attribute Reference

This resource exports no additional attributes.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `delete` - (Default `20m`)

## Import

You cannot import this resource.

<!-- cache-key: cdktf-0.20.8 input-f4c2650a2c1463738a0622997b13a7c264b52b84a9966a910944e21b21d9935f -->