package credentialhelper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type storeCommand struct {
	*meta
}

func (c *storeCommand) Run(args []string) int {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		c.Error(fmt.Sprintf("error reading credentials from stdin: %v", err))
		return 1
	}

	if err := c.helper.Store(args[0], string(bytes)); err != nil {
		c.Error(fmt.Sprintf("error storing credentials: %v", err))
		return 1
	}

	return 0
}

func (c *storeCommand) Synopsis() string {
	return "Store the credentials for the given hostname"
}

func (c *storeCommand) Help() string {
	return strings.TrimSpace(storeCommandHelp)
}

const storeCommandHelp = `
To store new credentials, Terraform will run the "store" command with any configured arguments,
plus the hostname for which credentials should be retrieved. It will write the credentials to be
stored to stdin.
`
