---
page_title: "magnolia_git_config Resource - terraform-provider-magnolia"
subcategory: ""
description: |-
 Magnolia git configuration resource in the Terraform provider.
---

# Resource: magnolia_git_config

Creates and returns a git secret, and a webhook URL to inject later into your Git repository to interact whit a Magnolia SaaS subscription.

## Example Usage

```terraform
resource "magnolia_subscription" "my_hello_saas" {
  email    = "foo@example.org"
  password = "my_password"
}

resource "magnolia_git_config" "my_hello_saas" {
 subscription_id = magnolia_subscription.my_hello_saas.id
 git_provider    = "GITHUB"
 git_clone_url   = "my_ssh_clone_url"
}
```

## Argument Reference

- `subscription_id` - (Required) The Magnolia SaaS Subscription ID created.
- `git_provider` - (Required) The Git provider to store the Hello-SaaS Magnolia application (GITHUB, BITBUCKET or BITBUCKET_CLOUD).
- `git_clone_url` - (Required) The Git URL where your Hello-SaaS project is.

## Attributes Reference

- `public_key` - The public key for Magnolia SaaS Subscription instances to interact whit.
- `git_secret` - The Git secret generated to configure the webhook in the Hello-SaaS repository settings.
- `webhook_url` - The Webhook URL generated to configure the webhook endpoint in the Hello-SaaS repository settings.
