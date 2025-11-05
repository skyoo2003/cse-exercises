// https://school.programmers.co.kr/learn/courses/30/lessons/181832
package main

import "fmt"

func solution(n int) [][]int {
	// Initialize table
	o := make([][]int, n)
	for i := 0; i < n; i++ {
		o[i] = make([]int, n)
	}

	// Fill table
	x, y := 0, 0
	dx, dy := 1, 0
	step, limit := 0, n
	flip := 0

	for i := 0; i < n*n; i++ {
		step += 1

		o[y][x] = i + 1

		// Change x, y direction (flipping)
		if step >= limit {
			flip += 1

			step = 0

			if flip%2 == 1 {
				limit -= 1
			}

			switch dx {
			case 0:
				switch flip % 4 {
				case 0:
					dx = 1
				case 2:
					dx = -1
				}
			case 1:
				dx = 0
			case -1:
				dx = 0
			}

			switch dy {
			case 0:
				switch flip % 4 {
				case 1:
					dy = 1
				case 3:
					dy = -1
				}
			case 1:
				dy = 0
			case -1:
				dy = 0
			}
		}

		x += dx
		y += dy
	}

	return o
}

func main() {
	fmt.Println(solution(4))
	fmt.Println(solution(5))
}
