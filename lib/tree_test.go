package lib_test

import "os"
import "testing"

import "gitlab.com/pab/pb/lib"
import "gitlab.com/pab/pb/testutil"

func TestTreeString(t *testing.T) {
	tree := &lib.Tree{
		Refs: []*lib.TreeRef{
			{0100644, "blob", "README.md", "abc123"},
			{0040000, "tree", "lib", "bcd234"},
			{0100644, "blob", "Rakefile", "cde345"},
		},
	}

	result := tree.String()
	expect := "40000\ttree\tlib\tbcd234\n100644\tblob\tRakefile\tcde345\n100644\tblob\tREADME.md\tabc123\n"
	testutil.AssertString(result, expect, t)
}

func TestTreePutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	t0 := &lib.Tree{
		Refs: []*lib.TreeRef{
			{0100644, "blob", "README.md", "abc123"},
			{0040000, "tree", "lib", "bcd234"},
			{0100644, "blob", "Rakefile", "cde345"},
		},
	}

	t0.Put()
	t1, err := lib.GetTree(t0.Hash())

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	testutil.AssertString(t1.String(), t0.String(), t)

	// Teardown
	os.RemoveAll("./.pb")
}
