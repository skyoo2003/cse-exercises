// https://www.hackerrank.com/challenges/qheap1/problem

package qheap1

import (
	"errors"
	"fmt"
	"os"
)

type Value int

type Heap struct {
	Data []Value
}

func NewHeap() *Heap {
	return &Heap{
		Data: make([]Value, 0),
	}
}

func (t *Heap) parent(index int) (Value, int) {
	if index > 0 {
		where := index / 2
		return t.Data[where], where
	}
	return 0, -1
}

func (t *Heap) leftChild(index int) (Value, int) {
	where := index*2 + 1
	if where < len(t.Data) {
		return t.Data[where], where
	}
	return 0, -1
}

func (t *Heap) rightChild(index int) (Value, int) {
	where := index*2 + 2
	if where < len(t.Data) {
		return t.Data[where], where
	}
	return 0, -1
}

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
		t.Data = append(t.Data[:last])

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
	return errors.New("No data")
}

func (t *Heap) Top() (Value, error) {
	if len(t.Data) > 0 {
		return t.Data[0], nil
	}
	return 0, errors.New("No data")
}

func QHeap1(in *os.File) {
	var Q int
	heap := NewHeap()
	fmt.Fscanf(in, "%d", &Q)
	for q := 0; q < Q; q++ {
		var cmd int
		var value Value
		fmt.Fscanf(in, "%d", &cmd)
		switch cmd {
		case 1:
			fmt.Fscanf(in, "%d", &value)
			heap.Insert(value)
		case 2:
			fmt.Fscanf(in, "%d", &value)
			heap.Delete(value)
		case 3:
			value, _ = heap.Top()
			fmt.Println(value)
		}
	}
}
