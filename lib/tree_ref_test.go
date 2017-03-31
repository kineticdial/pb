package lib_test

import "fmt"
import "os"
import "testing"

import "github.com/stretchr/testify/assert"

import "github.com/lead-scm/pb/lib"

func TestTreeRefString(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	blob := &lib.Blob{Contents: "abc123"}
	blob.Put()

	tr := &lib.TreeRef{
		Perms:   0100644,
		RefType: "blob",
		Name:    "README.md",
		Ref:     blob,
	}

	result := tr.String()
	expect := fmt.Sprintf("100644\tblob\tREADME.md\t%s", blob.Hash())

	assert.Equal(t, expect, result)

	// Teardown
	os.RemoveAll("./.pb")
}

func TestDecodeTreeRef(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	blob := &lib.Blob{Contents: "abc123"}
	blob.Put()

	rawEntry := fmt.Sprintf("100644\tblob\tREADME.md\t%s", blob.Hash())

	tr, err := lib.DecodeTreeRef(rawEntry)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if tr.Perms != 0100644 {
		t.Logf("tr.Perms: '%o', expected: '%o'", tr.Perms, 0100644)
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
