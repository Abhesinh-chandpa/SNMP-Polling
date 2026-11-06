// Package app provides the CLI entry point. SNMP and dashboard commands will
// be added only after their respective interfaces and tests are implemented.
package app

import (
	"errors"
	"fmt"
	"io"
)

const Version = "0.0.0-dev"

func Run(args []string, out io.Writer) error {
	if len(args) == 0 {
		return usage(out)
	}
	switch args[0] {
	case "help", "-h", "--help":
		return usage(out)
	case "version", "--version":
		_, err := fmt.Fprintln(out, "nmp-snmp", Version)
		return err
	default:
		return fmt.Errorf("unknown command %q (run 'nmp-snmp help')", args[0])
	}
}

func usage(out io.Writer) error {
	if out == nil {
		return errors.New("output writer is nil")
	}
	_, err := fmt.Fprintln(out, "nmp-snmp — SNMP tester (foundation build)\n\nCommands:\n  help       Show usage\n  version    Show version\n\nPlanned: get, walk, bulk, discover, poll, monitor, traps, stats")
	return err
}
