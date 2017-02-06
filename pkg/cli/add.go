package cli

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Lead-SCM/pb/pkg/lib"
)

func Add(path string) {
	// TODO: Check if in repository.
	// Check if file exists
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
	b := &lib.Blob{string(buf.Bytes())}
	err = b.Put()
	if err != nil {
		log.Fatal(err)
	}

	// Write blob's key in index
	index, err := os.OpenFile("./.pb/index", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	s := fmt.Sprintf("%s %s\n", path, b.Hash())
	_, err = index.WriteString(s)
	if err != nil {
		log.Println("got here")
		log.Fatal(err)
	}
	index.Close()
}
