// nolint
package golang

// always a bigger than b
func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

// always a bigger than b
func lcd(a, b int) int {
	return (a * b) / gcd(a, b)
}

func solution(n int, m int) []int {
	var g, l int
	if n < m {
		g, l = gcd(m, n), lcd(m, n)
	} else {
		g, l = gcd(n, m), lcd(n, m)
	}
	return []int{g, l}
}
