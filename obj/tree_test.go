package obj_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/Lead-SCM/pb/obj"
)

func TestTreePutGet(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)
	m0 := make(map[string]string)
	m1 := make(map[string]string)

	// Test
	t0 := obj.NewTree(m0, m1)
	err := t0.Put()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t1, err := obj.GetTree(t0.Hash())
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	eq := reflect.DeepEqual(t0.Trees, t1.Trees)
	if !eq {
		t.Log("Expected t0 and t1 to have identical sub-trees")
		t.Fail()
	}

	eq = reflect.DeepEqual(t0.Blobs, t1.Blobs)
	if !eq {
		t.Log("Expected t0 and t1 to have identical blobs")
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
}
