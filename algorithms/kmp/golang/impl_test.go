package kmp

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestKMPSearch(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	// testcase :=
	// 	`ABC ABCDAB ABCDABCDABDE
	// 	ABCDABD`
	testcase :=
		`ABCDABCFAABBACABC
		ABC`

	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	KMPSearch(in)
}
