package arraymanipulation

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`5 3
		1 2 100
		2 5 100
		3 4 100`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	ArrayManipulation(in)
}
