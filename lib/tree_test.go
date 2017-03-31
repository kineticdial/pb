package lib_test

import "os"
import "testing"
import "fmt"

import "github.com/stretchr/testify/assert"

import "github.com/lead-scm/pb/lib"

func TestTreeString(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	readme := &lib.Blob{}
	libDir := &lib.Tree{}
	rakefile := &lib.Blob{}

	readme.Put()
	libDir.Put()
	rakefile.Put()

	tree := &lib.Tree{
		Refs: []*lib.TreeRef{
			{0100644, "blob", "README.md", readme},
			{0040000, "tree", "lib", libDir},
			{0100644, "blob", "Rakefile", rakefile},
		},
	}

	result := tree.String()
	expect := fmt.Sprintf(
		"40000\ttree\tlib\t%s\n100644\tblob\tRakefile\t%s\n100644\tblob\tREADME.md\t%s\n",
		libDir.Hash(),
		rakefile.Hash(),
		readme.Hash(),
	)
	assert.Equal(t, expect, result)

	// Teardown
	os.RemoveAll("./.pb")
}

func TestTreePutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	readme := &lib.Blob{}
	libDir := &lib.Tree{}
	rakefile := &lib.Blob{}

	t0 := &lib.Tree{
		Refs: []*lib.TreeRef{
			{0100644, "blob", "README.md", readme},
			{0040000, "tree", "lib", libDir},
			{0100644, "blob", "Rakefile", rakefile},
		},
	}

	t0.Put()
	t1, err := lib.GetTree(t0.Hash())

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	assert.Equal(t, t0.String(), t1.String())

	// Teardown
	os.RemoveAll("./.pb")
}
