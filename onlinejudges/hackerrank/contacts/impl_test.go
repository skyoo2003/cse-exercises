package contacts

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestContact(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`4
		add hack
		add hackerrank
		find hac
		find hak`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	Contact(in)
}
