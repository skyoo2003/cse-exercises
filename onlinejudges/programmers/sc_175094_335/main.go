// https://programmers.co.kr/skill_checks/175094?challenge_id=335
// Ref: https://m.blog.naver.com/orbis1020/220664563768

package main

// gcd 파라미터는 항상 a > b
func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

func solution(w int, h int) int64 {
	var answer int64 = 0

	var c int
	if w < h {
		c = gcd(h, w)
	} else {
		c = gcd(w, h)
	}

	noPoints := c == 1
	if noPoints {
		answer = int64(w*h) - int64(w+h-1)
	} else {
		answer = int64(w*h) - int64(w+h-c)
	}

	return answer
}
