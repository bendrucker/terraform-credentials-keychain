package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bendrucker/terraform-credentials-keychain/credentialhelper"
)

const (
	version = "dev"
)

func main() {
	helper := new(KeychainHelper)
	cli := credentialhelper.New("terraform-credentials-keychain", version, helper, nil)

	status, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(status)

	return
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func helpText() string {
	return strings.TrimSpace(`
terraform-credentials-keychain
	`)
}
