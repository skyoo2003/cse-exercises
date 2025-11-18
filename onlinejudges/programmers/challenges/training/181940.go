// https://school.programmers.co.kr/learn/courses/30/lessons/181940?language=go
package main

func solution(my_string string, k int) string {
	var buf []byte = make([]byte, len(my_string)*k)
	for i := 0; i < k; i++ {
		w := i * len(my_string)
		copy(buf[w:], my_string)
	}
	return string(buf)
}
