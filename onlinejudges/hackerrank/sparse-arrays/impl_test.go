package sparsearrays

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestSparseArrays(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

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
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	SparseArrays(in)
}
