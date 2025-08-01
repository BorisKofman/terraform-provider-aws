---
subcategory: "SESv2 (Simple Email V2)"
layout: "aws"
page_title: "AWS: aws_sesv2_email_identity"
description: |-
  Terraform resource for managing an AWS SESv2 (Simple Email V2) Email Identity.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sesv2_email_identity

Terraform resource for managing an AWS SESv2 (Simple Email V2) Email Identity.

## Example Usage

### Basic Usage

#### Email Address Identity

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sesv2_email_identity import Sesv2EmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Sesv2EmailIdentity(self, "example",
            email_identity="testing@example.com"
        )
```

#### Domain Identity

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sesv2_email_identity import Sesv2EmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Sesv2EmailIdentity(self, "example",
            email_identity="example.com"
        )
```

#### Configuration Set

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sesv2_configuration_set import Sesv2ConfigurationSet
from imports.aws.sesv2_email_identity import Sesv2EmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Sesv2ConfigurationSet(self, "example",
            configuration_set_name="example"
        )
        aws_sesv2_email_identity_example = Sesv2EmailIdentity(self, "example_1",
            configuration_set_name=example.configuration_set_name,
            email_identity="example.com"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sesv2_email_identity_example.override_logical_id("example")
```

#### DKIM Signing Attributes (BYODKIM)

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sesv2_email_identity import Sesv2EmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Sesv2EmailIdentity(self, "example",
            dkim_signing_attributes=Sesv2EmailIdentityDkimSigningAttributes(
                domain_signing_private_key="MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...",
                domain_signing_selector="example"
            ),
            email_identity="example.com"
        )
```

## Argument Reference

The following arguments are required:

* `email_identity` - (Required) The email address or domain to verify.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `configuration_set_name` - (Optional) The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
* `dkim_signing_attributes` - (Optional) The configuration of the DKIM authentication settings for an email domain identity.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### dkim_signing_attributes

* `domain_signing_private_key` - (Optional, Sensitive) [Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.

-> **NOTE:** You have to delete the first and last lines ('-----BEGIN PRIVATE KEY-----' and '-----END PRIVATE KEY-----', respectively) of the generated private key. Additionally, you have to remove the line breaks in the generated private key. The resulting value is a string of characters with no spaces or line breaks.

* `domain_signing_selector` - (Optional) [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
* `next_signing_key_length` - (Optional) [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day. Valid values: `RSA_1024_BIT`, `RSA_2048_BIT`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Email Identity.
* `dkim_signing_attributes` - A list of objects that contains at most one element with information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for Easy DKIM.
    * `current_signing_key_length` - [Easy DKIM] The key length of the DKIM key pair in use.
    * `last_key_generation_timestamp` - [Easy DKIM] The last time a key pair was generated for this identity.
    * `next_signing_key_length` - [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.
    * `signing_attributes_origin` - A string that indicates how DKIM was configured for the identity. `AWS_SES` indicates that DKIM was configured for the identity by using Easy DKIM. `EXTERNAL` indicates that DKIM was configured for the identity by using Bring Your Own DKIM (BYODKIM).
    * `status` - Describes whether or not Amazon SES has successfully located the DKIM records in the DNS records for the domain. See the [AWS SES API v2 Reference](https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_DkimAttributes.html#SES-Type-DkimAttributes-Status) for supported statuses.
    * `tokens` - If you used Easy DKIM to configure DKIM authentication for the domain, then this object contains a set of unique strings that you use to create a set of CNAME records that you add to the DNS configuration for your domain. When Amazon SES detects these records in the DNS configuration for your domain, the DKIM authentication process is complete. If you configured DKIM authentication for the domain by providing your own public-private key pair, then this object contains the selector for the public key.
* `identity_type` - The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).
* `verified_for_sending_status` - Specifies whether or not the identity is verified.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SESv2 (Simple Email V2) Email Identity using the `email_identity`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sesv2_email_identity import Sesv2EmailIdentity
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Sesv2EmailIdentity.generate_config_for_import(self, "example", "example.com")
```

Using `terraform import`, import SESv2 (Simple Email V2) Email Identity using the `email_identity`. For example:

```console
% terraform import aws_sesv2_email_identity.example example.com
```

<!-- cache-key: cdktf-0.20.8 input-d126beb9f9a5eaed8432b97a8433209e089df6b5fd1db728654b464d8e34e7b8 -->