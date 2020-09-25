package dynamicarray

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`2 5
		1 0 5
		1 1 7
		1 0 3
		2 1 0
		2 1 1`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	DynamicArray(in)
}
