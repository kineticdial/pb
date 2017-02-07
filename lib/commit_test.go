package lib_test

import "fmt"
import "os"
import "testing"
import "time"

import "github.com/stretchr/testify/assert"

import "gitlab.com/pab/pb/lib"

func TestCommitString(t *testing.T) {
	now := time.Now()
	c := &lib.Commit{
		Tree:   "treeHash",
		Parent: "",
		Msg:    "Initial commit",
		Author: "Patrick Arthur Brown <pat@pab.io>",
		Date:   now,
	}

	expect := fmt.Sprintf(
		"Commit:\t%s\nTree:\t%s\nParent:\t%s\nAuthor:\tPatrick Arthur Brown <pat@pab.io>\nDate:\t%s\n\n\tInitial commit\n",
		c.Hash(),
		"treeHash",
		c.Parent,
		now,
	)

	assert.Equal(t, expect, c.String())
}

func TestCommitPutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	now := time.Now()
	c0 := &lib.Commit{
		Tree:   "treeHash",
		Parent: "",
		Msg:    "Initial commit",
		Author: "Patrick Arthur Brown <pat@pab.io>",
		Date:   now,
	}
	c0.Put()

	c1, err := lib.GetCommit(c0.Hash())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c0.String(), c1.String())

	// Teardown
	os.RemoveAll("./.pb")
}
