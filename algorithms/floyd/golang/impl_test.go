package floyd

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFloydAlgorithm(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`5
		10
		1 2 1.0
		1 4 1.0
		1 5 1.0
		2 1 9.0
		2 3 3.0
		2 4 2.0
		3 4 4.0
		4 3 2.0
		4 5 3.0
		5 1 3.0
		2
		2 5
		3 2
		`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	Algorithm(in)
}
