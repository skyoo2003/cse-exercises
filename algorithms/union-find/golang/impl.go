package unionfind

import (
	"fmt"
	"os"
)

type WorstDisjointSet struct {
	parents []int
}

func NewWorstDisjointSet(N int) *WorstDisjointSet {
	parents := make([]int, 0)
	for n := 0; n < N; n++ {
		parents = append(parents, n)
	}
	return &WorstDisjointSet{
		parents: parents,
	}
}

func (s *WorstDisjointSet) find(i int) int {
	if s.parents[i] == i {
		return i
	}
	return s.find(s.parents[i])
}

func (s *WorstDisjointSet) union(f, t int) {
	pf, pt := s.parents[f], s.parents[t]
	if pf != pt {
		s.parents[pf] = pt
	}
}

func WorstUnionFind(in *os.File) {
	var N int
	fmt.Fscanf(in, "%d", &N)
	// Initialize the value with a self-index
	ds := NewWorstDisjointSet(N)
	var P int
	fmt.Fscanf(in, "%d", &P)
	// Making unions according to direction (f <-> t)
	for p := 0; p < P; p++ {
		var f, t int
		fmt.Fscanf(in, "%d %d", &f, &t)
		fmt.Printf("Union(%d, %d)\n", f, t)
		ds.union(f, t)
	}
	var F int
	fmt.Fscanf(in, "%d", &F)
	// Prints the same index for two given indexes
	for f := 0; f < F; f++ {
		var f, t int
		fmt.Fscanf(in, "%d %d", &f, &t)
		fmt.Printf("Find(%d) == Find(%d) ? %v\n", f, t, ds.find(f) == ds.find(t))
	}
}

type OptimizedDisjointSet struct {
	parents []int
	ranks   []int
}

func NewOptimizedDisjointSet(N int) *OptimizedDisjointSet {
	parents := make([]int, 0)
	ranks := make([]int, 0)
	for n := 0; n < N; n++ {
		parents = append(parents, n)
		ranks = append(ranks, 0)
	}
	return &OptimizedDisjointSet{
		parents: parents,
		ranks:   ranks,
	}
}

func (s *OptimizedDisjointSet) find(i int) int {
	if s.parents[i] == i {
		return i
	}
	s.parents[i] = s.find(s.parents[i]) // memoization
	return s.parents[i]
}

func (s *OptimizedDisjointSet) union(f, t int) {
	pf, pt := s.parents[f], s.parents[t]
	if pf != pt {
		if s.ranks[pf] > s.ranks[pt] {
			pf, pt = pt, pf
		}
		s.parents[pf] = pt
		if s.ranks[pf] == s.ranks[pt] {
			s.ranks[pt] = s.ranks[pt] + 1
		}
	}
}

func OptimizedUnionFind(in *os.File) {
	var N int
	fmt.Fscanf(in, "%d", &N)
	// Initialize the value with a self-index
	ds := NewOptimizedDisjointSet(N)
	var P int
	fmt.Fscanf(in, "%d", &P)
	// Making unions according to direction (f <-> t)
	for p := 0; p < P; p++ {
		var f, t int
		fmt.Fscanf(in, "%d %d", &f, &t)
		fmt.Printf("Union(%d, %d)\n", f, t)
		ds.union(f, t)
	}
	var F int
	fmt.Fscanf(in, "%d", &F)
	// Prints the same index for two given indexes
	for f := 0; f < F; f++ {
		var f, t int
		fmt.Fscanf(in, "%d %d", &f, &t)
		fmt.Printf("Find(%d) == Find(%d) ? %v\n", f, t, ds.find(f) == ds.find(t))
	}
}
