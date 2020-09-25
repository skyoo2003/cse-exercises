// https://www.hackerrank.com/challenges/no-prefix-set/problem

package reducedstring

import (
	"fmt"
	"os"
)

func ReducedString(in *os.File) string {
	var line string
	fmt.Fscanf(in, "%s", &line)

	idx := 0
	for idx < len(line)-1 {
		prev, curr := line[idx], line[idx+1]
		if prev == curr {
			line = line[:idx] + line[idx+2:]
			idx = 0
		} else {
			idx++
		}
	}
	if len(line) == 0 {
		line = "Empty String"
	}
	fmt.Println(line)
	return line
}
