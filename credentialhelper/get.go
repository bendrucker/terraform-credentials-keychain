package credentialhelper

import (
	"fmt"
	"strings"
)

type getCommand struct {
	*meta
}

func (c *getCommand) Run(args []string) int {
	creds, err := c.helper.Get(args[0])
	if err != nil {
		c.Error(fmt.Sprintf("error getting credentials: %v", err))
	}

	fmt.Println(creds)
	return 0
}

func (c *getCommand) Synopsis() string {
	return "Retrieve the credentials for the given hostname"
}

func (c *getCommand) Help() string {
	return strings.TrimSpace(getCommandHelp)
}

const getCommandHelp = `
To retrieve credentials, Terraform will run the "get" command with any configured arguments,
plus the hostname for which credentials should be retrieved.
`
