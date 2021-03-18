package dynamicarray

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`2 5
		1 0 5
		1 1 7
		1 0 3
		2 1 0
		2 1 1`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	DynamicArray(in)
}
