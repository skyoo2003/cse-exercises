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
		defer in.Close()
		io.WriteString(in, testcase.input)
		in.Seek(0, os.SEEK_SET)

		actual := ReducedString(in)
		if actual != testcase.expected {
			t.Errorf("actual: '%s' / expected: '%s'", actual, testcase.expected)
		}
	}
}
