package algorithms_test

import "testing"
import "os"

import "github.com/stretchr/testify/assert"

import "github.com/lead-scm/pb/lib"
import "github.com/lead-scm/pb/algorithms"

func TestCommitString(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	main := &lib.Blob{}
	readme := &lib.Blob{}
	commit := &lib.Blob{}
	util := &lib.Blob{}
	commitTest := &lib.Blob{}

	main.Put()
	readme.Put()
	commit.Put()
	util.Put()
	commitTest.Put()

	index := [][]string{}
	index = append(index, []string{"main.go", main.Hash()})
	index = append(index, []string{"README", readme.Hash()})
	index = append(index, []string{"lib", "commit.go", commit.Hash()})
	index = append(index, []string{"lib", "util", "util.go", util.Hash()})
	index = append(index, []string{"lib", "commit_test.go", commitTest.Hash()})
	_, err := algorithms.BuildTreeFromWorkingIndex(index)

	assert.Equal(t, nil, err)

	// Teardown
	os.RemoveAll("./.pb")
}
