---
page_title: "magnolia_subscription Resource - terraform-provider-magnolia"
subcategory: ""
description: |-
 Magnolia subscription resource in the Terraform provider.
---

# Resource `magnolia_subscription`

Creates a user, permission and others whit a Magnolia SaaS subscription.

## Example Usage

```terraform
resource "magnolia_subscription" "my_hello_saas" {
  email    = "foo@example.org"
  password = "my_password"
}
```

## Schema

### Required

- **email** (String) The email address of the user for the Magnolia SaaS Subscription.
- **password** (String) The password for the Magnolia SaaS Subscription login.

## Ouput

- **id** The Magnolia SaaS Subscription identifier.
- **alias** The Magnolia SaaS Subscription alias.



