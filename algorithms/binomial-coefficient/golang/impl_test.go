package binomialcoefficient

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestBinRecursive(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase := `50 7`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)

	BinRecursive(in)
}

func TestBinDP(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase := `50 7`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)

	BinDP(in)
}
