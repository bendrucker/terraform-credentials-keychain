package main

import (
	"flag"
	"fmt"
	"os"

	credentialhelper "github.com/bendrucker/terraform-credential-helper-sdk"
)

const (
	version = "dev"
)

func main() {
	flags := flag.NewFlagSet("terraform-credentials-keychain", flag.ContinueOnError)
	flags.String("keychain", "login", "The name of the macOS keychain where the credentials will be stored.")

	cli := credentialhelper.New("terraform-credentials-keychain", version, new(KeychainHelper), flags)

	status, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(status)

	return
}
