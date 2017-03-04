package cli_test

import "bytes"
import "io"
import "os"
import "testing"

import "github.com/stretchr/testify/assert"

import "github.com/lead-scm/pb/cli"
import "github.com/lead-scm/pb/lib"

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

	expect := "./test_file.txt bfe5ed57e6e323555b379c660aa8d35b70c2f8f07cf03ad6747266495ac13be0"
	assert.Equal(t, expect, string(buf.Bytes()))

	b, _ := lib.GetBlob("bfe5ed57e6e323555b379c660aa8d35b70c2f8f07cf03ad6747266495ac13be0")

	assert.Equal(t, "contents\n", b.Contents)

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

	expect := "./test_file1.txt 8f88da056e2ed130ee23b3b61245d2e0948fe335236dcb23a100a087f92130f2\n./test_file2.txt e2c722b4fd19f306c7c5967bb9b53cb54e01384cf7ad0bdedd42c8bf3525359f"

	assert.Equal(t, expect, string(buf.Bytes()))

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

	expect := "./test_file.txt 89ff0296725a405977e9d6b1fd7384b79c1982b22b1a926a8f93e483db60fefa"
	assert.Equal(t, expect, string(buf.Bytes()))

	// Teardown
	os.RemoveAll("./.pb")
	os.Remove("./test_file.txt")
}
