package floyd

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFloydAlgorithm(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

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
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	FloydAlgorithm(in)
}
