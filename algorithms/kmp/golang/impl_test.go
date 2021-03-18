package kmp

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestKMPSearch1(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`ABC ABCDAB ABCDABCDABDE
		ABCDABD`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	Search(in)
}

func TestKMPSearch2(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`ABCDABCFAABBACABC
		ABC`

	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	Search(in)
}
