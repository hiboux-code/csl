package csl

import (
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	// CSL_ERR = GetConsole cannot get information about the console
	CSL_ERR Mode = iota
	// CSL_NULL = No console attached
	CSL_NULL
	// CSL_TERM = console is a terminal
	CSL_TERM
	// CSL_REDIRECT = output is redirected to a file
	CSL_REDIRECT
)

// Mode represents the console type
type Mode int

// GetConsole returns the type of console that is attached
// to the program. See package constants for more information.
func GetConsole() Mode {
	h, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return CSL_ERR
	}

	if h == 0 {
		return CSL_NULL
	}

	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		return CSL_TERM
	}

	return CSL_REDIRECT
}
