package lib_test

import (
	"os"
	"testing"

	"github.com/Lead-SCM/pb/pkg/lib"
)

func TestTreeString(t *testing.T) {
	tree := &lib.Tree{
		Refs: []*lib.TreeRef{
			{0100644, "blob", "README.md", "abc123"},
			{0040000, "tree", "lib", "bcd234"},
			{0100644, "blob", "Rakefile", "cde345"},
		},
	}

	res := tree.String()
	expected := "40000\ttree\tlib\tbcd234\n100644\tblob\tRakefile\tcde345\n100644\tblob\tREADME.md\tabc123\n"
	if res != expected {
		t.Logf("res: '%s', expected: '%s'", res, expected)
		t.Fail()
	}
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

	if t0.String() != t1.String() {
		t.Logf("t0: '%s', t1: '%s'", t0.String(), t1.String())
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
