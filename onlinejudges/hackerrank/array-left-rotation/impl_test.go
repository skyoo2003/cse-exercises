package arrayleftrotation

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestArrayLeftRotation(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`5 4
		1 2 3 4 5`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	ArrayLeftRotation(in)
}
