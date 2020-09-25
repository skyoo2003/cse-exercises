package main

import (
	"fmt"
	"sort"
)

func merge(a, b []int) []int {
	r := make([]int, 0)
	for len(a) > 0 && len(b) > 0 {
		if a[0] <= b[0] {
			r = append(r, a[0])
			a = a[1:]
		} else {
			r = append(r, b[0])
			b = b[1:]
		}
	}
	for len(a) > 0 {
		r = append(r, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		r = append(r, b[0])
		b = b[1:]
	}
	return r
}

func main() {
	a := []int{4, 10, 7, 1, 9}
	b := []int{8, 3, 5, 6, 2}
	sort.Ints(a)
	sort.Ints(b)
	r := merge(a, b)
	fmt.Println(r)
}
