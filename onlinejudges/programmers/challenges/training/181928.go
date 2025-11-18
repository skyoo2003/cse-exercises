// https://school.programmers.co.kr/learn/courses/30/lessons/181928
package main

// See https://cs.opensource.google/go/go/+/refs/tags/go1.25.3:src/math/pow10.go
var pow10rem = [...]float64{
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
	1e20, 1e21, 1e22, 1e23, 1e24, 1e25, 1e26, 1e27, 1e28, 1e29,
	1e30, 1e31,
}

var pow10quo = [...]float64{
	1e00, 1e32, 1e64, 1e96, 1e128, 1e160, 1e192, 1e224, 1e256, 1e288,
}

func pow10(n int) float64 {
	return pow10quo[uint(n)/32] * pow10rem[uint(n)%32]
}

func solution(num_list []int) int {
	even_n, odd_n := 0, 0
	even_v, odd_v := 0, 0
	for i := len(num_list) - 1; i >= 0; i-- {
		value := num_list[i]
		if value%2 == 0 {
			even_v += value * int(pow10(even_n))
			even_n += 1
		} else {
			odd_v += value * int(pow10(odd_n))
			odd_n += 1
		}
	}
	return even_v + odd_v
}
