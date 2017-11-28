package arrayleftrotation

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestArrayLeftRotation(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`5 4
		1 2 3 4 5`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	ArrayLeftRotation(in)
}
