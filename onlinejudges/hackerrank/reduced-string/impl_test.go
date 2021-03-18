package reducedstring

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestReducedString(t *testing.T) {
	testcase := []struct {
		input, expected string
	}{
		{input: "aaabccddd", expected: "abd"},
		{input: "aa", expected: "Empty String"},
		{input: "baab", expected: "Empty String"},
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

		actual := ReducedString(in)
		if actual != testcase.expected {
			t.Errorf("actual: '%s' / expected: '%s'", actual, testcase.expected)
		}
	}
}
