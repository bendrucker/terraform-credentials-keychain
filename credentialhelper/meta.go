package credentialhelper

import (
	"flag"
	"fmt"
	"os"
)

type meta struct {
	flags  *flag.FlagSet
	helper Helper
}

func (m *meta) Error(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}
