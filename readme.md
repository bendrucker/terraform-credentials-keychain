# terraform-credentials-keychain [![tests](https://github.com/bendrucker/terraform-credentials-keychain/actions/workflows/test.yaml/badge.svg?branch=master)](https://github.com/bendrucker/terraform-credentials-keychain/actions/workflows/test.yaml)

> A Terraform [credentials helper](https://www.terraform.io/docs/commands/cli-config.html#credentials-helpers) that stores your credentials in the system keychain

By default, `terraform login` writes your Terraform Cloud credentials (i.e. API token) as a plain text file in your home directory. Any program you run can read this file, potentially stealing your credentials. 

With this credential helper installed, your credentials will instead be stored in the system keychain. This helper uses [99designs/keyring](https://github.com/99designs/keyring) and can use any credential storage backend it supports. Currently, only macOS is actively tested.

## Installing

Credentials helpers go in the same directory as Terraform provider plugins, and that directory names an architecture. The binary you put there has to match it. On Apple Silicon a mismatch used to work anyway, because Rosetta translated the `amd64` binary, but recent macOS versions drop Rosetta and the same mismatch now fails with `bad CPU type in executable`.

Deriving both the asset and the destination from `uname -m` keeps them in step. Set `version` to the [latest release](https://github.com/bendrucker/terraform-credentials-keychain/releases/latest):

```sh
version=0.2.2
arch=$([ "$(uname -m)" = arm64 ] && echo arm64 || echo amd64)
dir=~/.terraform.d/plugins/darwin_$arch

mkdir -p "$dir"
curl -fsSL "https://github.com/bendrucker/terraform-credentials-keychain/releases/download/v${version}/terraform-credentials-keychain_${version}_darwin_${arch}.tar.gz" \
  | tar -xzC "$dir" terraform-credentials-keychain
file "$dir/terraform-credentials-keychain"
```

The last line prints the architecture of what landed, which should match the directory it landed in.

On other platforms, download the [release asset](https://github.com/bendrucker/terraform-credentials-keychain/releases) matching your OS and architecture, then install the binary as `~/.terraform.d/plugins/<os>_<arch>/terraform-credentials-keychain`.

Releases for macOS are [signed and notarized](https://developer.apple.com/developer-id/) so that the system will trust the application.

## Usage

Run `terraform logout` for each Terraform host you connect to. For Terraform Cloud, you can run `terraform logout` directly. For Terraform Enterprise, supply the hostname of the Terraform Enterprise server. This will remove all plain text credentials stored in `credentials.tfrc.json` or print an error if credential blocks are defined in `.terraformrc`. These credentials will bypass the credential helper if they are not removed. You should also revoke these API tokens from your Terraform Cloud user settings.

Add the credential helper to `~/.terraformrc` file:

```
credentials_helper "keychain" {}
```

Now when you use `terraform login` and `terraform logout`, they will use your system keychain rather than persisting credentials directly to disk!

[![asciicast](https://asciinema.org/a/334212.svg)](https://asciinema.org/a/334212)

Each time you run a `terraform` command that uses your credentials (e.g. `init`, `plan`, `apply`, etc.), the credential helper will read your credentials from the keychain, prompting for a password if needed.

## Security

Any command that requires Terraform Cloud credentials, including most `terraform` commands, will prompt for the keychain password:

<img src="keychain.png" alt="macOS Keychain password prompt" width="434" />

For maximum security, click _Allow_ and enter your password every time it is required by Terraform or another program. If you run Terraform frequently, this may become tedious. If you click _Always Allow_, you will never be prompted for a password again. Your credentials will still be protected from a malicious program scanning your disk, but a program that calls `terraform-credentials-keychain get <host>` will still be able to obtain them. 

If you choose this option, consider [configuring your keychain to lock after a period of inactivity](https://support.apple.com/guide/keychain-access/mac-keychain-password-kyca1242/mac). When your keychain is locked, you will be prompted for the _keychain_ password before an application can access its contents, even if that application is trusted by the item. You can also use a dedicated keychain, instead of the default _login_ keychain:

```hcl
credentials_helper "keychain" {
  args = ["--keychain=terraform"]
}
```

After adding your credentials, you can open Keychain Access to edit the keychain's auto-lock settings:

```sh
open ~/Library/Keychains/terraform.keychain-db
```
