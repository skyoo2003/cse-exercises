// https://www.hackerrank.com/challenges/sparse-arrays/problem
// nolint
package sparsearrays

import (
	"fmt"
	"os"
)

// SparseArrays My solution
func SparseArrays(in *os.File) {
	var N, Q int
	list := make([]string, 0)
	fmt.Fscanf(in, "%d", &N)
	for n := 0; n < N; n++ {
		var value string
		fmt.Fscanf(in, "%s", &value)
		list = append(list, value)
	}
	fmt.Fscanf(in, "%d", &Q)
	for q := 0; q < Q; q++ {
		var query string
		fmt.Fscanf(in, "%s", &query)
		cnt := 0
		for _, phrase := range list {
			if phrase == query {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}
