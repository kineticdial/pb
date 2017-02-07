package cli

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gitlab.com/pab/pb/pkg/lib"
)

// Add stages a given file to the repository.
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
	b := &lib.Blob{string(buf.Bytes())}
	err = b.Put()
	if err != nil {
		log.Fatal(err)
	}

	index, err := os.Open("./.pb/index")
	if err != nil {
		log.Fatal(err)
	}

	indexBuf := bytes.NewBuffer(nil)
	_, err = io.Copy(indexBuf, index)
	if err != nil {
		log.Fatal(err)
	}
	indexContents := string(indexBuf.Bytes())
	index.Close()

	lines := strings.Split(indexContents, "\n")
	itemFound := false
	newLines := make([]string, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		if tokens[0] == path {
			newLines = append(newLines, formatIndexLine(path, b))
			itemFound = true
		} else {
			newLines = append(newLines, line)
		}
	}

	if !itemFound {
		newLines = append(newLines, formatIndexLine(path, b))
	}

	newContents := strings.TrimLeft(strings.Join(newLines, "\n"), "\n")
	ioutil.WriteFile("./.pb/index", []byte(newContents), 0666)
}

func formatIndexLine(path string, b *lib.Blob) string {
	return fmt.Sprintf("%s %s", path, b.Hash())
}
