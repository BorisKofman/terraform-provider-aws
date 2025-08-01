---
subcategory: "CodeStar Connections"
layout: "aws"
page_title: "AWS: aws_codestarconnections_connection"
description: |-
  Provides a CodeStar Connection
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_codestarconnections_connection

Provides a CodeStar Connection.

~> **NOTE:** The `aws_codestarconnections_connection` resource is created in the state `PENDING`. Authentication with the connection provider must be completed in the AWS Console. See the [AWS documentation](https://docs.aws.amazon.com/dtconsole/latest/userguide/connections-update.html) for details.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.codepipeline import Codepipeline
from imports.aws.codestarconnections_connection import CodestarconnectionsConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, location, type, category, name, owner, provider, version, category1, name1, owner1, provider1, version1):
        super().__init__(scope, name)
        example = CodestarconnectionsConnection(self, "example",
            name="example-connection",
            provider_type="Bitbucket"
        )
        aws_codepipeline_example = Codepipeline(self, "example_1",
            artifact_store=[CodepipelineArtifactStore(
                location=location,
                type=type
            )
            ],
            name="tf-test-pipeline",
            role_arn=codepipeline_role.arn,
            stage=[CodepipelineStage(
                action=[CodepipelineStageAction(
                    category="Source",
                    configuration={
                        "BranchName": "main",
                        "ConnectionArn": example.arn,
                        "FullRepositoryId": "my-organization/test"
                    },
                    name="Source",
                    output_artifacts=["source_output"],
                    owner="AWS",
                    provider="CodeStarSourceConnection",
                    version="1"
                )
                ],
                name="Source"
            ), CodepipelineStage(
                action=[CodepipelineStageAction(
                    category=category,
                    name=name,
                    owner=owner,
                    provider=provider,
                    version=version
                )
                ],
                name="Build"
            ), CodepipelineStage(
                action=[CodepipelineStageAction(
                    category=category1,
                    name=name1,
                    owner=owner1,
                    provider=provider1,
                    version=version1
                )
                ],
                name="Deploy"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_codepipeline_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
* `provider_type` - (Optional) The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub`, `GitHubEnterpriseServer`, `GitLab` or `GitLabSelfManaged`. Changing `provider_type` will create a new resource. Conflicts with `host_arn`
* `host_arn` - (Optional) The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `provider_type`
* `tags` - (Optional) Map of key-value resource tags to associate with the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The codestar connection ARN.
* `arn` - The codestar connection ARN.
* `connection_status` - The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CodeStar connections using the ARN. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.codestarconnections_connection import CodestarconnectionsConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CodestarconnectionsConnection.generate_config_for_import(self, "testConnection", "arn:aws:codestar-connections:us-west-1:0123456789:connection/79d4d357-a2ee-41e4-b350-2fe39ae59448")
```

Using `terraform import`, import CodeStar connections using the ARN. For example:

```console
% terraform import aws_codestarconnections_connection.test-connection arn:aws:codestar-connections:us-west-1:0123456789:connection/79d4d357-a2ee-41e4-b350-2fe39ae59448
```

<!-- cache-key: cdktf-0.20.8 input-c05e929b2e681c4892cf9898e41bee582859ab20eb2c610d55a91e57ae50caf7 -->