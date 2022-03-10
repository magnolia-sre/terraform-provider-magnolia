---
page_title: "magnolia_subscription Resource - terraform-provider-magnolia"
subcategory: ""
description: |-
 Magnolia subscription resource in the Terraform provider.
---

# Resource: magnolia_subscription

Creates a user, permission and others whit a Magnolia SaaS subscription.

## Example Usage

```terraform
resource "magnolia_subscription" "my_hello_saas" {
  email    = "foo@example.org"
  password = "my_password"
}
```

## Argument Reference

- `email` - (Required) The email address of the user for the Magnolia SaaS Subscription.
- `password` - (Required) The password for the Magnolia SaaS Subscription login.

## Attributes Reference

- `id` - The Magnolia SaaS Subscription identifier.
- `alias` - The Magnolia SaaS Subscription alias.
