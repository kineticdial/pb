package cli_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Lead-SCM/pb/pkg/cli"
	"github.com/Lead-SCM/pb/pkg/lib"
)

func TestAdd(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)
	f, _ := os.Create("./.pb/index")
	f.Close()
	f, _ = os.Create("./test_file.txt")
	f.WriteString("contents\n")
	f.Close()

	// Test
	cli.Add("./test_file.txt")

	f, _ = os.Open("./.pb/index")
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	f.Close()

	if string(buf.Bytes()) != "./test_file.txt 572c291421cd821a5e821e28766d0bdb719c379d\n" {
		t.Fail()
	}

	b, _ := lib.GetBlob("572c291421cd821a5e821e28766d0bdb719c379d")

	if b.Contents != "contents\n" {
		t.Fail()
	}

	// Teardown
	os.RemoveAll("./.pb")
	os.Remove("./test_file.txt")
}
