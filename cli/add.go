package cli

import "bytes"
import "io"
import "log"
import "os"

import "github.com/lead-scm/pb/lib"
import "github.com/mitchellh/cli"

// AddCommand is the controller for staging files to the working area. Right now
// you have to add one file at a time, but in the future (TODO) you can specify
// whole directories.
type AddCommand struct {
	UI cli.Ui
}

// Help displays explanitory text for the AddCommand.
func (c *AddCommand) Help() string {
	return "Add a file to the working index"
}

// Synopsis is aliased to Help.
func (c *AddCommand) Synopsis() string {
	return c.Help()
}

// Run stages a given file to the repository.
func (c *AddCommand) Run(args []string) int {
	path := args[0]
	Add(path)
	return 0
}

// Add a file to the working index. If the file location is already in the
// working index, update the blob reference for that location.
func Add(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	// Persist file as blob in object store
	b := &lib.Blob{Contents: string(buf.Bytes())}
	err = b.Put()
	if err != nil {
		log.Fatal(err)
	}

	err = lib.UpsertToIndex(path, b)
	if err != nil {
		log.Fatal(err)
	}
}
