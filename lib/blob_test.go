package lib_test

import "os"
import "testing"

import "gitlab.com/pab/pb/lib"
import "gitlab.com/pab/pb/testutil"

func TestBlobPutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)

	// Test
	b0 := &lib.Blob{"foo"}
	err := b0.Put()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b1, err := lib.GetBlob(b0.Hash())
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	testutil.AssertString(b0.String(), b1.String(), t)

	// Teardown
	os.RemoveAll("./.pb")
}
