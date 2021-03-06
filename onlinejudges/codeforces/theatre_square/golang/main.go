package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, a, r float64
	fmt.Scanf("%f%f%f", &n, &m, &a)

	var w, h float64
	w = math.Ceil(float64(n) / float64(a))
	h = math.Ceil(float64(m) / float64(a))
	r = w * h
	fmt.Printf("%.0f\n", r)
}
