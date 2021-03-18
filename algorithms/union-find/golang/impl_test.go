package unionfind

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWorstUnionFind(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

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
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	WorstUnionFind(in)
}

func TestOptimizedUnionFind(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

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
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	OptimizedUnionFind(in)
}
