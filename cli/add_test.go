package cli_test

import "bytes"
import "io"
import "os"
import "testing"

import "github.com/stretchr/testify/assert"

import "gitlab.com/pab/pb/cli"
import "gitlab.com/pab/pb/lib"

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

	expect := "./test_file.txt 572c291421cd821a5e821e28766d0bdb719c379d"
	assert.Equal(t, string(buf.Bytes()), expect)

	b, _ := lib.GetBlob("572c291421cd821a5e821e28766d0bdb719c379d")

	assert.Equal(t, b.Contents, "contents\n")

	// Teardown
	os.RemoveAll("./.pb")
	os.Remove("./test_file.txt")
}

func TestAddAdditional(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)
	f, _ := os.Create("./.pb/index")
	f.Close()
	f, _ = os.Create("./test_file1.txt")
	f.WriteString("contents1\n")
	f.Close()
	f, _ = os.Create("./test_file2.txt")
	f.WriteString("contents2\n")
	f.Close()

	// Test
	cli.Add("./test_file1.txt")
	cli.Add("./test_file2.txt")

	f, _ = os.Open("./.pb/index")
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	f.Close()

	expect := "./test_file1.txt e53b356df5a76565a769e67ad656e581afc9d1b6\n./test_file2.txt b56fa959a159c3a3ee54824989711aab309805ba"

	assert.Equal(t, string(buf.Bytes()), expect)

	// Teardown
	os.RemoveAll("./.pb")
	os.Remove("./test_file1.txt")
	os.Remove("./test_file2.txt")
}

func TestAddMutate(t *testing.T) {
	// Setup
	os.MkdirAll("./.pb/objects", 0777)
	f, _ := os.Create("./.pb/index")
	f.Close()
	f, _ = os.Create("./test_file.txt")
	f.WriteString("contents\n")
	f.Close()

	// Test
	cli.Add("./test_file.txt")

	f, _ = os.OpenFile("./test_file.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	f.WriteString("additional contents\n")
	f.Close()

	cli.Add("./test_file.txt")

	f, _ = os.Open("./.pb/index")
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	f.Close()

	expect := "./test_file.txt 5b2489e7bf4366347c0921dcef135c5870a19757"
	assert.Equal(t, string(buf.Bytes()), expect)

	// Teardown
	os.RemoveAll("./.pb")
	os.Remove("./test_file.txt")
}
