package pb_test

import (
	"os"
	"strings"
	"testing"

	"github.com/Lead-SCM/pb"
)

func TestBlobPutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	b0 := &pb.Blob{"foo"}
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

	s0 := b0.ToString()
	s1 := b1.ToString()
	if strings.Compare(s0, s1) != 0 {
		t.Logf("b0: '%s'; b1: '%s'", s0, s1)
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
