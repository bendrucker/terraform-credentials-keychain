package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var help bool

	flags := flag.NewFlagSet("helper", flag.ContinueOnError)
	flags.BoolVar(&help, "help", false, "print the help text")

	if err := flags.Parse(os.Args[1:]); err != nil {
		fail(err.Error())
	}

	if help {
		fmt.Println(helpText())
		return
	}

	if flags.NArg() != 2 {
		fmt.Fprintln(os.Stderr, helpText())
		os.Exit(1)
		return
	}

	command := flags.Arg(0)
	hostname := flags.Arg(1)

	switch command {
	case "get":
		Get(hostname)
	case "store":
		Store(hostname, "")
	case "forget":
		Forget(hostname)
	}
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
