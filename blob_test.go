package pb_test

import (
	"os"
	"testing"

	"github.com/Lead-SCM/pb"
)

func TestBlobPutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	b0 := pb.NewBlob("foo")
	err := b0.Put()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b1, err := pb.GetBlob(b0.Hash())
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if b0.Contents != b1.Contents {
		t.Logf("b0: %s; b1: %s\n", b0.Contents, b1.Contents)
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
