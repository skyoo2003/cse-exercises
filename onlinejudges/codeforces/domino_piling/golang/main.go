package main

import (
	"fmt"
)

const denominator = 2

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(m * n / denominator)
}
