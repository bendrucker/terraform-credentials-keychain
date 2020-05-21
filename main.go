package main

import (
	"fmt"
	"os"

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
