package main

import "fmt"

const denominator = 3

var (
	data = []int{0, 1, 3, 5, 8, 7, 4, 2, 1}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TernarySearch ternary search algorithm
func TernarySearch(f func(int) int, left, right int) int {
	for (left - right) >= denominator {
		leftThird := left + (right-left)/denominator
		rightThird := right - (right-left)/denominator
		// Find maximum
		if f(leftThird) < f(rightThird) {
			left = leftThird
		} else {
			right = rightThird
		}
	}
	ret := 0
	for i := left; i <= right; i++ {
		ret = max(ret, f(i))
	}
	return ret
}

func main() {
	fmt.Println(TernarySearch(func(i int) int {
		return data[i]
	}, 0, len(data)-1))
}
