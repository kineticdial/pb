package lib

import "bytes"
import "fmt"
import "io"
import "io/ioutil"
import "os"
import "strings"

// UpsertToIndex updates/inserts blob and path combiation to the working index.
func UpsertToIndex(path string, b *Blob) error {
	index, err := os.Open("./.pb/index")
	if err != nil {
		return err
	}

	indexBuf := bytes.NewBuffer(nil)
	_, err = io.Copy(indexBuf, index)
	if err != nil {
		return err
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
	return err
}

func formatIndexLine(path string, b *Blob) string {
	return fmt.Sprintf("%s %s", path, b.Hash())
}
