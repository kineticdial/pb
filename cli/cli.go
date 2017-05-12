package cli

import (
	"flag"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

const appName = "pb"
const version = "0.0.0"

// Main parses command from CLI arguments.
func Main() int {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI(appName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{UI: ui}, nil
		},
		"init": func() (cli.Command, error) {
			return &InitCommand{UI: ui}, nil
		},
		"commit": func() (cli.Command, error) {
			msgPtr := flag.String("message", "", "commit message")
			authorPtr := flag.String("author", "", "commit Author")
			flag.Parse()

			log.Printf("Msg: %s\n", *msgPtr)
			log.Printf("Author: %s\n", *authorPtr)

			return &CommitCommand{UI: ui, Msg: *msgPtr, Author: *authorPtr}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
