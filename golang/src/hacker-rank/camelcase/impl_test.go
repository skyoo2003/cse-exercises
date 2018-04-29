package reducedstring

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCamelcase(t *testing.T) {
	testcase := []struct {
		input    string
		expected int
	}{
		{input: "saveChangesInTheEditor", expected: 5},
		{input: "helloWorld", expected: 2},
	}
	for _, testcase := range testcase {
		in, _ := ioutil.TempFile(os.TempDir(), "")
		defer in.Close()
		io.WriteString(in, testcase.input)
		in.Seek(0, os.SEEK_SET)

		actual := Camelcase(in)
		if actual != testcase.expected {
			t.Errorf("actual: '%d' / expected: '%d'", actual, testcase.expected)
		}
	}
}
