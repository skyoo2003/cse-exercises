// https://www.hackerrank.com/challenges/qheap1/problem
// nolint
package qheap1

import (
	"errors"
	"fmt"
	"os"
)

const numChildren = 2

const (
	leftChild  = 1
	rightChild = 2
)

const (
	cmdAdd = iota + 1
	cmdDelete
	cmdPrint
)

// Value value type
type Value int

// Heap heap data structure
type Heap struct {
	Data []Value
}

// NewHeap create a heap
func NewHeap() *Heap {
	return &Heap{
		Data: make([]Value, 0),
	}
}

func (t *Heap) parent(index int) (Value, int) {
	if index > 0 {
		where := index / numChildren
		return t.Data[where], where
	}
	return 0, -1
}

func (t *Heap) leftChild(index int) (Value, int) {
	where := index*numChildren + leftChild
	if where < len(t.Data) {
		return t.Data[where], where
	}
	return 0, -1
}

func (t *Heap) rightChild(index int) (Value, int) {
	where := index*numChildren + rightChild
	if where < len(t.Data) {
		return t.Data[where], where
	}
	return 0, -1
}

// Insert insert a value into heap
func (t *Heap) Insert(value Value) error {
	t.Data = append(t.Data, value)
	idx := len(t.Data) - 1
	parentValue, parentIdx := t.parent(idx)
	for parentIdx != -1 {
		if parentValue < value {
			break
		}
		t.Data[idx], t.Data[parentIdx] = t.Data[parentIdx], t.Data[idx]
		idx = parentIdx
		parentValue, parentIdx = t.parent(idx)
	}
	return nil
}

// Delete a value from heap
func (t *Heap) Delete(value Value) error {
	if len(t.Data) > 0 {
		where := -1
		for idx, item := range t.Data {
			if item == value {
				where = idx
				break
			}
		}

		last := len(t.Data) - 1
		t.Data[where], t.Data[last] = t.Data[last], t.Data[where]
		t.Data = t.Data[:last]

		for where != -1 && where < len(t.Data) {
			minValue := t.Data[where]
			minIdx := where

			leftValue, leftIdx := t.leftChild(where)
			if leftIdx != -1 && leftValue < minValue {
				minValue = leftValue
				minIdx = leftIdx
			}

			rightValue, rightIdx := t.rightChild(where)
			if rightIdx != -1 && rightValue < minValue {
				minValue = rightValue
				minIdx = rightIdx
			}

			if minValue < t.Data[where] {
				t.Data[where], t.Data[minIdx] = t.Data[minIdx], t.Data[where]
				where = minIdx
			} else {
				break
			}
		}

		return nil
	}
	return errors.New("no data")
}

// Top find maximum value in heap
func (t *Heap) Top() (Value, error) {
	if len(t.Data) > 0 {
		return t.Data[0], nil
	}
	return 0, errors.New("no data")
}

// QHeap1 My solution
func QHeap1(in *os.File) {
	var Q int
	heap := NewHeap()
	fmt.Fscanf(in, "%d", &Q)
	for q := 0; q < Q; q++ {
		var cmd int
		var value Value
		fmt.Fscanf(in, "%d", &cmd)
		switch cmd {
		case cmdAdd:
			fmt.Fscanf(in, "%d", &value)
			if err := heap.Insert(value); err != nil {
				panic(err)
			}
		case cmdDelete:
			fmt.Fscanf(in, "%d", &value)
			if err := heap.Delete(value); err != nil {
				panic(err)
			}
		case cmdPrint:
			value, _ = heap.Top()
			fmt.Println(value)
		}
	}
}
