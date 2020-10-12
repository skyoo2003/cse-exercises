package arrayleftrotation

import (
	"fmt"
	"os"
)

// ArrayLeftRotation My solution
func ArrayLeftRotation(in *os.File) {
	var N, D int
	fmt.Fscanf(in, "%d %d", &N, &D)
	list := make([]int, 0)
	for n := 0; n < N; n++ {
		var value int
		fmt.Fscanf(in, "%d", &value)
		list = append(list, value)
	}
	for n := 0; n < N; n++ {
		where := (n + D) % N
		fmt.Printf("%d ", list[where])
	}
	fmt.Println()
}
