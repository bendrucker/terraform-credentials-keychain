package credentialhelper

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
)

// New creates a new credential helper CLI
func New(name string, version string, helper Helper, flags *flag.FlagSet) *CLI {
	c := cli.NewCLI(name, version)

	if flags == nil {
		flags = flag.NewFlagSet(name, flag.ContinueOnError)
	}

	m := &meta{
		flags:  flags,
		helper: helper,
	}

	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return &getCommand{m}, nil
		},
		"store": func() (cli.Command, error) {
			return &storeCommand{m}, nil
		},
		"forget": func() (cli.Command, error) {
			return &forgetCommand{m}, nil
		},
	}

	return &CLI{
		cli:  c,
		meta: m,
	}
}

// CLI is a Terraform credential helper CLI
type CLI struct {
	*meta
	cli *cli.CLI
}

// Run runs the CLI with the provided arguments
func (c *CLI) Run(args []string) (int, error) {
	if err := c.flags.Parse(args); err != nil {
		return 1, err
	}

	if n := c.flags.NArg() - 1; n > 1 {
		return 1, fmt.Errorf("expected a hostname, got %d args: %v", n, c.flags.Args())
	} else if n == 0 {
		c.Error(c.cli.HelpFunc(c.cli.Commands))
		return 1, nil
	}

	c.cli.Args = c.flags.Args()
	return c.cli.Run()
}
