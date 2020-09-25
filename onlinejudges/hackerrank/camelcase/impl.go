// https://www.hackerrank.com/challenges/no-prefix-set/problem

package reducedstring

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Camelcase(in *os.File) int {
	var line string
	fmt.Fscanf(in, "%s", &line)

	fields := strings.FieldsFunc(line, func(r rune) bool {
		return unicode.IsUpper(r)
	})
	fmt.Println(len(fields))
	return len(fields)
}
