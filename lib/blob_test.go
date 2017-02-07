package lib_test

import "os"
import "testing"

import "github.com/stretchr/testify/assert"

import "gitlab.com/pab/pb/lib"

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

	assert.Equal(t, b0.String(), b1.String())

	// Teardown
	os.RemoveAll("./.pb")
}
