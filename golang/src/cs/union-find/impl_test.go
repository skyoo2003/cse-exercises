package unionfind

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWorstUnionFind(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`10
		6
		0 1
		1 2
		2 3
		4 5
		5 6
		6 7
		2
		0 4
		4 7
		`

	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	WorstUnionFind(in)
}

func TestOptimizedUnionFind(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`10
		6
		0 1
		1 2
		2 3
		4 5
		5 6
		6 7
		2
		0 4
		4 7
		`

	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	OptimizedUnionFind(in)
}
