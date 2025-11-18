// https://school.programmers.co.kr/learn/courses/30/lessons/181927
package main

func solution(num_list []int) []int {
	l_i := len(num_list) - 1
	l_p_i := l_i - 1
	v := 0
	if num_list[l_i] > num_list[l_p_i] {
		v = num_list[l_i] - num_list[l_p_i]
	} else {
		v = num_list[l_i] * 2
	}
	num_list = append(num_list, v)
	return num_list
}
