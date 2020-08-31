package main

import "fmt"

func main() {
	var n int
	var f, s, t int
	solves := 0

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&f, &s, &t)
		v := f<<2 + s<<1 + t
		if v >= 5 || v == 3 { // 1 1 1 / 1 1 0 / 1 0 1 / 0 1 1
			solves += 1
		}
	}
	fmt.Println(solves)
}
