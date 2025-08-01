---
subcategory: "GuardDuty"
layout: "aws"
page_title: "AWS: aws_guardduty_member"
description: |-
  Provides a resource to manage a GuardDuty member
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_guardduty_member

Provides a resource to manage a GuardDuty member. To accept invitations in member accounts, see the [`aws_guardduty_invite_accepter` resource](/docs/providers/aws/r/guardduty_invite_accepter.html).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.guardduty_detector import GuarddutyDetector
from imports.aws.guardduty_member import GuarddutyMember
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        member = GuarddutyDetector(self, "member",
            enable=True,
            provider=dev
        )
        primary = GuarddutyDetector(self, "primary",
            enable=True
        )
        aws_guardduty_member_member = GuarddutyMember(self, "member_2",
            account_id=member.account_id,
            detector_id=primary.id,
            email="required@example.com",
            invitation_message="please accept guardduty invitation",
            invite=True
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_guardduty_member_member.override_logical_id("member")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `account_id` - (Required) AWS account ID for member account.
* `detector_id` - (Required) The detector ID of the GuardDuty account where you want to create member accounts.
* `email` - (Required) Email address for member account.
* `invite` - (Optional) Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the Terraform state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
* `invitation_message` - (Optional) Message for invitation.
* `disable_email_notification` - (Optional) Boolean whether an email notification is sent to the accounts. Defaults to `false`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `relationship_status` - The status of the relationship between the member account and its primary account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `1m`)
- `update` - (Default `1m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import GuardDuty members using the primary GuardDuty detector ID and member AWS account ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.guardduty_member import GuarddutyMember
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GuarddutyMember.generate_config_for_import(self, "myMember", "00b00fd5aecc0ab60a708659477e9617:123456789012")
```

Using `terraform import`, import GuardDuty members using the primary GuardDuty detector ID and member AWS account ID. For example:

```console
% terraform import aws_guardduty_member.MyMember 00b00fd5aecc0ab60a708659477e9617:123456789012
```

<!-- cache-key: cdktf-0.20.8 input-c5e336d57f5bf72822b3e2253d69356130f927f1f97be4373acd113750a9d9aa -->