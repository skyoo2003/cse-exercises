package kmp

// ref. https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type KMP struct {
	S string
}

func NewKMP(S string) *KMP {
	return &KMP{
		S: S,
	}
}

func (t *KMP) table(W string) []int {
	T := make([]int, 0)
	for i := 0; i < len(W); i++ {
		T = append(T, 0)
	}
	T = append(T, 0)
	pos, cnd := 1, 0

	T[0] = -1

	for pos < len(W) {
		if W[pos] == W[cnd] {
			T[pos] = T[cnd]
		} else {
			T[pos] = cnd
			cnd = T[cnd]
			for cnd >= 0 && W[pos] != W[cnd] {
				cnd = T[cnd]
			}
		}
		pos, cnd = pos+1, cnd+1
	}
	T[pos] = cnd

	return T
}

func (t *KMP) search(W string) ([]int, int) {
	S := t.S
	T := t.table(W)
	j, k := 0, 0

	P := make([]int, 0)
	nP := 0
	for j < len(S) {
		if W[k] == S[j] {
			j, k = j+1, k+1
			if k == len(W) {
				P = append(P, j-k)
				nP = nP + 1
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

func KMPSearch(in *os.File) {
	r := bufio.NewReader(in)
	sl, _, _ := r.ReadLine()
	S := strings.TrimSpace(string(sl))
	wl, _, _ := r.ReadLine()
	W := strings.TrimSpace(string(wl))

	kmp := NewKMP(S)
	P, nP := kmp.search(W)
	fmt.Printf("Positions(%v), Count(%d)", P, nP)
}
