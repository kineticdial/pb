package db_test

import (
	"os"
	"testing"

	"github.com/Lead-SCM/pb/db"
)

type Obj struct {
	Foo string
}

func TestPutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	o0 := Obj{"bar"}
	err := db.Put("key", o0)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var o1 Obj
	err = db.Get("key", &o1)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if o0 != o1 {
		t.Log("o0 != o1")
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
