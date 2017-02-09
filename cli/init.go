package cli

import "fmt"
import "log"
import "os"

import "github.com/mitchellh/cli"

type InitCommand struct {
	Ui cli.Ui
}

func (c *InitCommand) Help() string {
	return "Initialize a new pb repository at working directory"
}

func (c *InitCommand) Synopsis() string {
	return c.Help()
}

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
