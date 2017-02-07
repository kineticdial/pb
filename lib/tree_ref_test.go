package lib_test

import "testing"

import "gitlab.com/pab/pb/lib"
import "gitlab.com/pab/pb/testutil"

func TestTreeRefString(t *testing.T) {
	tr := &lib.TreeRef{
		Perms:   0100644,
		RefType: "blob",
		Name:    "README.md",
		Hash:    "abc123",
	}

	result := tr.String()
	expect := "100644\tblob\tREADME.md\tabc123"

	testutil.AssertString(result, expect, t)
}

func TestDecodeTreeRef(t *testing.T) {
	tr, err := lib.DecodeTreeRef("100644\tblob\tREADME.md\tabc123")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if tr.Perms != 0100644 {
		t.Logf("tr.Perms: '%o', expected: '%o'", tr.Perms, 0100644)
		t.Fail()
	}
}
