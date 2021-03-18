// nolint
package kmp

// ref. https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// KMP kmp string matching algorithm data structure
type KMP struct {
	S string
}

// NewKMP create a kmp
func NewKMP(s string) *KMP {
	return &KMP{
		S: s,
	}
}

func (t *KMP) table(w string) []int {
	T := make([]int, 0)
	for i := 0; i < len(w); i++ {
		T = append(T, 0)
	}
	T = append(T, 0)
	pos, cnd := 1, 0

	T[0] = -1

	for pos < len(w) {
		if w[pos] == w[cnd] {
			T[pos] = T[cnd]
		} else {
			T[pos] = cnd
			cnd = T[cnd]
			for cnd >= 0 && w[pos] != w[cnd] {
				cnd = T[cnd]
			}
		}
		pos, cnd = pos+1, cnd+1
	}
	T[pos] = cnd

	return T
}

func (t *KMP) search(w string) ([]int, int) {
	S := t.S
	T := t.table(w)
	j, k := 0, 0

	P := make([]int, 0)
	nP := 0
	for j < len(S) {
		if w[k] == S[j] {
			j, k = j+1, k+1
			if k == len(w) {
				P = append(P, j-k)
				nP++
				k = T[k]
			}
		} else {
			k = T[k]
			if k < 0 {
				j, k = j+1, k+1
			}
		}
	}
	return P, nP
}

// Search My solution
func Search(in *os.File) {
	r := bufio.NewReader(in)
	sl, _, _ := r.ReadLine()
	S := strings.TrimSpace(string(sl))
	wl, _, _ := r.ReadLine()
	W := strings.TrimSpace(string(wl))

	kmp := NewKMP(S)
	P, nP := kmp.search(W)
	fmt.Printf("Positions(%v), Count(%d)", P, nP)
}
