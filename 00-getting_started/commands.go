package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type DockerCli struct{}

func (cli *DockerCli) getMethod(args ...string) (func(...string) error, bool) {
	cameArgs := make([]string, len(args))
	for i, s := range args {
		if len(s) == 0 {
			return nil, false
		}
		camelArgs[i] = strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	}

	methodName := "Cmd" + strings.Join(camelArgs, "")
	method := reflect.ValueOf(cli).MethodByName(methodName)
	if !method.IsValid() {
		return nil, false
	}

	return method.Interface().(func(...string) error), true
}

func (cli *DockerCli) CmdHelp(args ...string) error {
	if len(args) > 0 {
		method, exists := cli.getMethod(args[0])
		if !exists {
			fmt.Fprintf(os.Stderr, "Error: command not found - %\n", args[0])
		} else {
			method("--help")
			return nil
		}
	}

	flag.Usage()
	return nil
}

// SubCmd implements a subcommand by creating a FlagSet struct.
func (cli *DockerCli) SubCmd(name, signature, description string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		options := ""
		fmt.Fprintf(os.Stderr, "\nUsage: docker %s %s%s\n\n%s\n\n", name, options, signature, description)
		flags.PrintDefaults()
		os.Exit(2)
	}
	return flags
}

// Cmd serves as the entrypoint for the DockerCli. It checks that command line
// arguments are properly passed and calls the corresponding method.
func (cli *DockerCli) Cmd(args ...string) error {
	if len(args) > 0 {
		method, exists := cli.getMethod(args[0])
		if !exists {
			fmt.Printf("Error: command doesn't exist - %s\n", args[0])
			return cli.CmdHelp()
		}
		return method(args[1:]...)
	}

	return cli.CmdHelp
}

func (cli *DockerCli) CmdCreate(args ...string) error {
	cmd := cli.SubCmd("machine create", "NAME", "create machine")

	driverDesc := fmt.Sprintf(
		"Driver to create machines with. Available drivers: %s",
		strings.Join(drivers.GetDriverNames(), ", "),
	)
	return nil
}
