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
		defer func() { _ = in.Close() }()
		if _, err := io.WriteString(in, testcase.input); err != nil {
			t.Error(err)
			t.FailNow()
		}
		if _, err := in.Seek(0, io.SeekStart); err != nil {
			t.Error(err)
			t.FailNow()
		}

		actual := Camelcase(in)
		if actual != testcase.expected {
			t.Errorf("actual: '%d' / expected: '%d'", actual, testcase.expected)
		}
	}
}
