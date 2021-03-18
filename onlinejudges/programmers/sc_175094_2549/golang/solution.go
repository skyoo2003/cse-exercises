// https://programmers.co.kr/skill_checks/175094?challenge_id=2549
// nolint
package golang

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solution(citations []int) int {
	r := 0
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, c := range citations {
		r = max(r, min(c, i+1))
	}
	return r
}
