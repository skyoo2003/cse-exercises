// https://www.hackerrank.com/challenges/dynamic-array/problem
// nolint
package dynamicarray

import (
	"fmt"
	"os"
)

// DynamicArray My solution
func DynamicArray(in *os.File) {
	var N, Q int
	fmt.Fscanf(in, "%d %d", &N, &Q)

	fmt.Println(N, Q)

	lastAnswer := 0
	seqList := make([][]int, N)
	for n := 0; n < N; n++ {
		seqList[n] = make([]int, 0)
	}

	for q := 0; q < Q; q++ {
		var qt, x, y int
		fmt.Fscanf(in, "%d %d %d", &qt, &x, &y)

		where := ((x ^ lastAnswer) % N)
		switch qt {
		case 1:
			seqList[where] = append(seqList[where], y)
		case 2:
			seq := seqList[where]
			lastAnswer = seq[y%len(seq)]
			fmt.Println(lastAnswer)
		}
	}
}
