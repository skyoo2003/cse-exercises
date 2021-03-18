// https://www.hackerrank.com/challenges/no-prefix-set/problem
// nolint
package reducedstring

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Camelcase My solution
func Camelcase(in *os.File) int {
	var line string
	fmt.Fscanf(in, "%s", &line)

	fields := strings.FieldsFunc(line, unicode.IsUpper)
	fmt.Println(len(fields))
	return len(fields)
}
