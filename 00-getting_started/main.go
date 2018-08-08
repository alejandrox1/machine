package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	mVersion = flag.Bool("v", false, "Print version information")
	mDebu    = flag.Bool("d", false, "Enable debug mode")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: machine [OPTIONS] COMMAND [arg...]\n\n")
		fmt.Fprintf(os.Stderr, "Create and manage machines running Docker.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")

		flag.PrintDefaults()

		help := "\nCommands:\n"
		for _, command := range [][]string{
			{"active", "Get or set the active machine"},
			{"create", "Create a machine"},
			{"inspect", "Inspect information about a machine"},
			{"ip", "Get the IP address of a machine"},
			{"kill", "Kill a machine"},
			{"ls", "List machines"},
			{"restart", "Restart a machine"},
			{"rm", "Remove a machine"},
			{"ssh", "Log into or run a command on a machine with SSH"},
			{"start", "Start a machine"},
			{"stop", "Stop a machine"},
			{"upgrade", "Upgrade a machine to the latest version of Docker"},
			{"url", "Get the URL of a machine"},
		} {
			help += fmt.Sprintf("    %-10.10s%s\n", command[0], command[1])
		}
		help += "\nRun 'docker COMMAND --help' for more information on a command."
		fmt.Fprintf(os.Stderr, "%s\n", help)
	}

	flag.Parse()

	cli := DockerCli{}

	if err := cli.Cmd(flag.Args()...); err != nil {
		log.Fatal(err)
	}
}
