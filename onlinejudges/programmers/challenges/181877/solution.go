// https://school.programmers.co.kr/learn/courses/30/lessons/181877
package main

func solution(myString string) string {
	var buf []byte = make([]byte, len(myString))
	for i := 0; i < len(myString); i++ {
		c := myString[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		buf[i] = c
	}
	return string(buf)
}
