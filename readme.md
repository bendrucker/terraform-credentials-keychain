# terraform-credentials-keychain

> A Terraform [credentials helper](https://www.terraform.io/docs/commands/cli-config.html#credentials-helpers) that stores your credentials in the system keychain

By default, `terraform login` writes your Terraform Cloud credentials (i.e. API token) as a plain text file in your home directory. Any program you run can read this file, potentially stealing your credentials. 

With this credential helper installed, your credentials will instead be stored in the system keychain