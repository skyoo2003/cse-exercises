package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var curr int
	prev := -1
	order, cum := 1, 0
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&curr)
		if curr <= 0 {
			break
		}
		if prev != curr {
			order += cum
			cum = 1
		} else {
			cum++
		}
		if order <= k {
			cnt++
		}
		prev = curr
	}
	fmt.Println(cnt)
}
