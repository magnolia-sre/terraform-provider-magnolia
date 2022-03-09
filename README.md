# Magnolia Terraform Provider

- Website: https://www.terraform.io
- Documentation: https://registry.terraform.io/providers/magnolia/magnolia/latest/docs
- Mailing list: http://groups.google.com/group/terraform-tool
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.16

## Building The Provider locally (Mac Sample)

1. You need to create a file name .terraformrc with the content:
https://www.terraform.io/cli/config/config-file#explicit-installation-method-configuration

```
provider_installation {
    filesystem_mirror {
        path    = "/var/tmp/terraform/providers/"
        include = ["registry.terraform.io/magnolia-sre/*"]
    }
    direct {
        exclude = ["registry.terraform.io/magnolia-sre/*"]
    }
}
```

This file will tell Terraform using local file system instead of direct download from registry

2. You need to build the provider and copy to /var/tmp/terraform/providers/
To ensure required sub directories created, please run

```
mkdir -p /var/tmp/terraform/providers/registry.terraform.io/magnolia-sre/magnolia/0.1.7/darwin_amd64/
```

then build and put provider

```
go build -o terraform-provider-magnolia_v0.1.7 && mv terraform-provider-magnolia_v0.1.7 /var/tmp/terraform/providers/registry.terraform.io/magnolia-sre/magnolia/0.1.7/darwin_amd64/
```

3. Delete .terraform and file .terraform.lock.hcl  in terraform-provider-magnolia-demo, and run terraform init  again, and you are using local provider

```
rm .terraform.lock.hcl && rm -rf .terraform/providers/registry.terraform.io/magnolia-sre/ && terraform init
```

