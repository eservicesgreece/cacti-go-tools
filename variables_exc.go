// +build !solaris

package main

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var interactive = terminal.IsTerminal(int(os.Stdout.Fd())) // check if we are executed by a user
