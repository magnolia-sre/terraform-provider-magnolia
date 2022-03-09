---
page_title: "magnolia Provider"
subcategory: ""
description: |-
---

# Magnolia Provider

The Magnolia provider is used to interact with the `Magnolia SaaS Subscription Service.`

In order to use this Provider, you must have an active account with `Magnolia SaaS`.
More information can be found at https://www.magnolia-cms.com/

Use the navigation to the left to read about the available resources.

## Example Usage

```terraform
# Terraform 0.13+ requires providers to be declared in a "required_providers" block
terraform {
  required_providers {
    magnolia = {
      source = "magnolia-sre/magnolia"
    }
  }
}

# Configure the Magnolia Provider
provider "fastly" {}

# Create a Subscription and Configure you Hello SaaS
resource "magnolia_subscription" "my_saas" {
  email    = my_project_owner_email
  password = my_project_owner_email_password
}

resource "magnolia_git_config" "my_saas" {
  subscription_id = my_magnolia_subscription_id
  git_provider    = "GITHUB"
  git_clone_url   = my_github_repository_ssh_clone_url
}
```

## Authentication

Get an access token for Magnolia provider (the access token file will be located at `~/.mgnl/access_token`):

```
 npm install -g mgnlctl
 MGNL_PROFILE=hackathon mgnlctl login local-development -y
```
