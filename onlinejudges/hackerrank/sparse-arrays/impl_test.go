package sparsearrays

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestSparseArrays(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`4
		aba
		baba
		aba
		xzxb
		3
		aba
		xzxb
		ab`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	SparseArrays(in)
}
