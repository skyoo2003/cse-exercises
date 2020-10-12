package arraymanipulation

import (
	"fmt"
	"os"
)

// ArrayManipulation My solution
func ArrayManipulation(in *os.File) {
	var N, M int
	fmt.Fscanf(in, "%d %d", &N, &M)

	list := make([]int, N)
	for m := 0; m < M; m++ {
		var a, b, k int
		fmt.Fscanf(in, "%d %d %d", &a, &b, &k)
		list[a-1] += k
		if b < len(list) {
			list[b] -= k
		}
	}

	max, sum := 0, 0
	for n := 0; n < N; n++ {
		sum += list[n]
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
