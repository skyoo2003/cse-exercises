package contacts

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestContact(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer in.Close()

	testcase :=
		`4
		add hack
		add hackerrank
		find hac
		find hak`
	io.WriteString(in, testcase)
	in.Seek(0, os.SEEK_SET)
	Contact(in)
}
