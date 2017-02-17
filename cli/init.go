package cli

import "fmt"
import "log"
import "os"

import "github.com/mitchellh/cli"

// InitCommand is the controller for initializing a new Lead repository.
// TODO: Gracefully handle if already in a Lead repository.
type InitCommand struct {
	UI cli.Ui
}

// Help displays explanitory text for the InitCommand.
func (c *InitCommand) Help() string {
	return "Initialize a new pb repository at working directory"
}

// Synopsis is aliased to Help.
func (c *InitCommand) Synopsis() string {
	return c.Help()
}

// Run performs the initialization.
func (c *InitCommand) Run(_ []string) int {
	Init()
	return 0
}

// Init initializes a new pb repository at the working directory.
func Init() {
	err := os.MkdirAll("./.pb/objects", 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create("./.pb/index")
	if err != nil {
		log.Fatal(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lead (pb) project has been initialized at %s\n", pwd)
}
