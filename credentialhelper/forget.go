package credentialhelper

import (
	"fmt"
	"strings"
)

type forgetCommand struct {
	*meta
}

func (c *forgetCommand) Run(args []string) int {
	if err := c.helper.Forget(args[0]); err != nil {
		c.Error(fmt.Sprintf("error forgetting credentials: %v", err))
		return 1
	}

	return 0
}

func (c *forgetCommand) Synopsis() string {
	return "Forget the credentials for the given hostname"
}

func (c *forgetCommand) Help() string {
	return strings.TrimSpace(forgetCommandHelp)
}

const forgetCommandHelp = `
To forget existing credentials, Terraform will run the "forget" command with any configured arguments,
plus the hostname for which credentials should be forgotten.
`
