package cli

import "fmt"
import "log"
import "os"

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
