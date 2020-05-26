# terraform-credentials-keychain

> A Terraform [credentials helper](https://www.terraform.io/docs/commands/cli-config.html#credentials-helpers) that stores your credentials in the system keychain

By default, `terraform login` writes your Terraform Cloud credentials (i.e. API token) as a plain text file in your home directory. Any program you run can read this file, potentially stealing your credentials. 

With this credential helper installed, your credentials will instead be stored in the system keychain. 

## Installing

Download an appropriate [release binary](https://github.com/bendrucker/terraform-credentials-keychain/releases) for your operating system/architecture. [Install it](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) into the appropriate Terraform plugin directory. Credentials helpers are stored in the same directory as provider plugins.

For example, on macOS, you should install the binary as:

```
~/.terraform.d/plugins/darwin_amd64/terraform-credentials-keychain
```

Releases for macOS are [signed and notarized](https://developer.apple.com/developer-id/) so that the system will trust the application.

## Usage

Run `terraform logout` for each Terraform host you connect to. For Terraform Cloud, you can run `terraform logout` directly. For Terraform Enterprise, supply the hostname of the Terraform Enterprise server. This will remove all plain text credentials stored in the `credentials.tfrc.json` or print an error if credential blocks are defined in `.terraformrc`. These credentials will bypass the credential helper if they are not removed. You should also revoke these API tokens from your Terraform Cloud user settings.

Add the credential helper to `~/.terraformrc` file:

```
credentials_helper "keychain" {}
```

Now when you use `terraform login` and `terraform logout`, they will use your system keychain rather than persisting credentials directly to disk!

[![asciicast](https://asciinema.org/a/334212.svg)](https://asciinema.org/a/334212)

Each time you run a `terraform` command that uses your credentials (e.g. `init`, `plan`, `apply`, etc.), the credential helper will read your credentials from the keychain, prompting for a password if needed.
