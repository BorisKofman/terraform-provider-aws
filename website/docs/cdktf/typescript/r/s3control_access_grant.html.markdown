---
subcategory: "S3 Control"
layout: "aws"
page_title: "AWS: aws_s3control_access_grant"
description: |-
  Provides a resource to manage an S3 Access Grant.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3control_access_grant

Provides a resource to manage an S3 Access Grant.
Each access grant has its own ID and gives an IAM user or role or a directory user, or group (the grantee) access to a registered location. You determine the level of access, such as `READ` or `READWRITE`.
Before you can create a grant, you must have an S3 Access Grants instance in the same Region as the S3 data.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3ControlAccessGrant } from "./.gen/providers/aws/s3-control-access-grant";
import { S3ControlAccessGrantsInstance } from "./.gen/providers/aws/s3-control-access-grants-instance";
import { S3ControlAccessGrantsLocation } from "./.gen/providers/aws/s3-control-access-grants-location";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3ControlAccessGrantsInstance(this, "example", {});
    const awsS3ControlAccessGrantsLocationExample =
      new S3ControlAccessGrantsLocation(this, "example_1", {
        dependsOn: [example],
        iamRoleArn: Token.asString(awsIamRoleExample.arn),
        locationScope: "s3://${" + awsS3BucketExample.bucket + "}/prefixA*",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3ControlAccessGrantsLocationExample.overrideLogicalId("example");
    const awsS3ControlAccessGrantExample = new S3ControlAccessGrant(
      this,
      "example_2",
      {
        accessGrantsLocationConfiguration: [
          {
            s3SubPrefix: "prefixB*",
          },
        ],
        accessGrantsLocationId: Token.asString(
          awsS3ControlAccessGrantsLocationExample.accessGrantsLocationId
        ),
        grantee: [
          {
            granteeIdentifier: Token.asString(awsIamUserExample.arn),
            granteeType: "IAM",
          },
        ],
        permission: "READ",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3ControlAccessGrantExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `accessGrantsLocationConfiguration` - (Optional) See [Location Configuration](#location-configuration) below for more details.
* `accessGrantsLocationId` - (Required) The ID of the S3 Access Grants location to with the access grant is giving access.
* `accountId` - (Optional) The AWS account ID for the S3 Access Grants location. Defaults to automatically determined account ID of the Terraform AWS provider.
* `grantee` - (Optional) See [Grantee](#grantee) below for more details.
* `permission` - (Required) The access grant's level of access. Valid values: `READ`, `WRITE`, `READWRITE`.
* `s3PrefixType` - (Optional) If you are creating an access grant that grants access to only one object, set this to `Object`. Valid values: `Object`.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### Location Configuration

The `accessGrantsLocationConfiguration` block supports the following:

* `s3SubPrefix` - (Optional) Sub-prefix.

### Grantee

The `grantee` block supports the following:

* `granteeIdentifier` - (Required) Grantee identifier.
* `granteeType` - (Required) Grantee types. Valid values: `DIRECTORY_USER`, `DIRECTORY_GROUP`, `IAM`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `accessGrantArn` - Amazon Resource Name (ARN) of the S3 Access Grant.
* `accessGrantId` - Unique ID of the S3 Access Grant.
* `grantScope` - The access grant's scope.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import S3 Access Grants using the `accountId` and `accessGrantId`, separated by a comma (`,`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3ControlAccessGrant } from "./.gen/providers/aws/s3-control-access-grant";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    S3ControlAccessGrant.generateConfigForImport(
      this,
      "example",
      "123456789012,04549c5e-2f3c-4a07-824d-2cafe720aa22"
    );
  }
}

```

Using `terraform import`, import S3 Access Grants using the `accountId` and `accessGrantId`, separated by a comma (`,`). For example:

```console
% terraform import aws_s3control_access_grants_location.example 123456789012,04549c5e-2f3c-4a07-824d-2cafe720aa22
```

<!-- cache-key: cdktf-0.20.8 input-5efb3eac5480790c668ba5342992abfc9ad39c578d69d5b78b3a8ae687607102 -->