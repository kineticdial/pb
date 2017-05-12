package cli

import "bytes"
import "io"
import "log"
import "os"
import "strings"
import "regexp"
import "time"

import "github.com/lead-scm/pb/lib"
import "github.com/lead-scm/pb/algorithms"
import "github.com/mitchellh/cli"

type CommitCommand struct {
	UI     cli.Ui
	Msg    string
	Author string
}

func (c *CommitCommand) Help() string {
	return "Commits working index"
}

// Synopsis is aliased to Help.
func (c *CommitCommand) Synopsis() string {
	return c.Help()
}

func (c *CommitCommand) Run(args []string) int {
	Commit(c.Msg, c.Author)
	return 0
}

func Commit(msg string, author string) {
	index := parseWorkingIndex()

	tree, err := algorithms.BuildTreeFromWorkingIndex(index)

	prevCommit := findPrevCommit()

	commit := lib.Commit{
		Tree:   tree.Hash(),
		Parent: prevCommit,
		Msg:    msg,
		Author: author,
		Date:   time.Now(),
	}

	commit.Put()

	if err != nil {
		log.Fatal(err)
	}
}

func parseWorkingIndex() [][]string {
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
	workingIndex := make([][]string, 0)
	for _, line := range lines {
		s := regexp.MustCompile("(/| )").Split(line, -1)
		workingIndex = append(workingIndex, s)
	}

	return workingIndex
}

func findPrevCommit() string {
	return "abc123"
}
