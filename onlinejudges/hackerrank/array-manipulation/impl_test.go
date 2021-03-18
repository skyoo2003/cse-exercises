package arraymanipulation

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`5 3
		1 2 100
		2 5 100
		3 4 100`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	ArrayManipulation(in)
}
