package main

import (
	"container/heap"
	"fmt"
	"math"
)

// IntHeap int heap type alias
type IntHeap []int

// Len get size of heap
func (h IntHeap) Len() int {
	return len(h)
}

// Less whether i less then j
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swap each other
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push append a value into heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop pop a value into heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func merge(kheap ...*IntHeap) []int {
	allEmpty := false
	merged := make([]int, 0)
	for !allEmpty {
		minValue, minIndex := math.MaxInt64, -1
		for i, h := range kheap {
			n := len(*h)
			allEmpty = allEmpty && n == 0
			if n > 0 && (*h)[n-1] < minValue {
				minValue, minIndex = (*h)[n-1], i
			}
		}
		if minIndex >= 0 {
			merged = append(merged, minValue)
			heap.Pop(kheap[minIndex])
		}
	}
	return merged
}

func main() {
	a := &IntHeap{4, 10, 7, 1, 9}
	b := &IntHeap{8, 3, 5, 6, 2}
	c := &IntHeap{11, 14, 12, -1, -4}
	d := &IntHeap{-9, 1, 4, 5, 7}
	heap.Init(a)
	heap.Init(b)
	heap.Init(c)
	heap.Init(d)
	r := merge(a, b, c, d)
	fmt.Println(r)
}
