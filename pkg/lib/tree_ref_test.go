package lib_test

import (
	"testing"

	"github.com/Lead-SCM/pb/pkg/lib"
)

func TestTreeRefString(t *testing.T) {
	tr := &lib.TreeRef{
		Perms:   0100644,
		RefType: "blob",
		Name:    "README.md",
		Hash:    "abc123",
	}

	res := tr.String()
	expected := "100644\tblob\tREADME.md\tabc123"
	if res != expected {
		t.Logf("res: '%s', expected: '%s'", res, expected)
		t.Fail()
	}
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
