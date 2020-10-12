package dynamicarray

import (
	"errors"
)

var (
	// ErrIndexOutOfRange index out of range in dynamic array
	ErrIndexOutOfRange = errors.New("index out of range")
	// ErrNoData can not found any data in dynamic array
	ErrNoData = errors.New("no data")
)

const (
	// GrowthFactor ratio to increase/decrease capacity of dynamic array
	GrowthFactor = 2
)

// DynamicArray dynamic array data structure
type DynamicArray struct {
	buffer   []interface{}
	capacity int
	size     int
}

// NewDynamicArray create a dynamic array
func NewDynamicArray() DynamicArray {
	return NewDynamicArrayWithSize(0)
}

// NewDynamicArrayWithSize create a dynmaic array with initial size
func NewDynamicArrayWithSize(size int) DynamicArray {
	return DynamicArray{
		buffer:   make([]interface{}, size),
		capacity: size,
		size:     0,
	}
}

func (a *DynamicArray) validIndex(index int) bool {
	if index < 0 || index > a.size {
		return false
	}
	return true
}

func (a *DynamicArray) resize(claimSize int) error {
	if claimSize >= a.capacity { // Need to expand
		for claimSize >= a.capacity {
			if err := a.expand(); err != nil {
				return err
			}
		}
	} else if claimSize <= a.capacity/4 {
		for claimSize <= a.capacity/4 {
			if err := a.shrink(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DynamicArray) expand() error {
	if a.capacity == 0 {
		a.realloc(1)
	} else {
		a.realloc(a.capacity * 2)
	}
	return nil
}

func (a *DynamicArray) shrink() error {
	if a.capacity >= 2 {
		a.realloc(a.capacity / 2)
	} else {
		a.realloc(1)
	}
	return nil
}

func (a *DynamicArray) realloc(capacity int) error {
	newBuffer := make([]interface{}, capacity)
	copy(newBuffer, a.buffer)
	a.buffer = newBuffer
	a.capacity = capacity
	return nil
}

// Capacity get capacity of dynamic array
func (a *DynamicArray) Capacity() int {
	return a.capacity
}

// Size get size of dynamic array
func (a *DynamicArray) Size() int {
	return a.size
}

// Get get a value by index in dynamic array
func (a *DynamicArray) Get(index int) (interface{}, error) {
	if !a.validIndex(index) {
		return nil, ErrIndexOutOfRange
	}
	return a.buffer[index], nil
}

// Range get values between begin and end in dynamic array
func (a *DynamicArray) Range(begin, end int) ([]interface{}, error) {
	if !a.validIndex(begin) || !a.validIndex(end) {
		return nil, ErrIndexOutOfRange
	}
	buf := make([]interface{}, 0)
	for i := begin; i < end; i++ {
		if value, err := a.Get(i); err == nil {
			buf = append(buf, value)
		} else {
			return nil, err
		}
	}
	return buf, nil
}

// Set set a value by index in dynamic array
func (a *DynamicArray) Set(index int, value interface{}) error {
	if !a.validIndex(index) {
		return ErrIndexOutOfRange
	}
	a.buffer[index] = value
	return nil
}

// Append insert values in the end of dynamic array
func (a *DynamicArray) Append(values ...interface{}) error {
	newSize := a.size + len(values)
	if err := a.resize(newSize); err != nil {
		return err
	}
	for i, j := a.size, 0; i < newSize; i, j = i+1, j+1 {
		a.buffer[i] = values[j]
	}
	a.size = newSize
	return nil
}

// Prepend insert values in front of dynamic array
func (a *DynamicArray) Prepend(values ...interface{}) error {
	lenValues := len(values)
	newSize := a.size + lenValues
	if err := a.resize(newSize); err != nil {
		return err
	}
	// Move items in the buffer by the length of values
	for i := a.size - 1; i >= 0; i-- {
		a.buffer[i+lenValues] = a.buffer[i]
	}
	for i := 0; i < lenValues; i++ {
		a.buffer[i] = values[i]
	}
	a.size = newSize
	return nil
}

// Insert insert values by index in dynamic array
func (a *DynamicArray) Insert(index int, values ...interface{}) error {
	lenValues := len(values)
	newSize := a.size + lenValues
	if err := a.resize(newSize); err != nil {
		return err
	}
	// Move items in the middle of the buffer by the length of values
	for i := a.size - 1; i >= index; i-- {
		a.buffer[i+lenValues] = a.buffer[i]
	}
	for i, j := index, 0; i < index+lenValues; i, j = i+1, j+1 {
		a.buffer[i] = values[j]
	}
	a.size = newSize
	return nil
}

// Remove remove a value by index from dynamic array
func (a *DynamicArray) Remove(index int) (interface{}, error) {
	if a.size == 0 {
		return nil, ErrNoData
	}
	if !a.validIndex(index) {
		return nil, ErrIndexOutOfRange
	}
	newSize := a.size - 1
	if err := a.resize(newSize); err != nil {
		return nil, err
	}
	value, _ := a.Get(index)
	for i := index + 1; i < a.size; i++ {
		a.buffer[i-1] = a.buffer[i]
	}
	a.size = newSize
	return value, nil
}

// Pop remove a value in the end of dynamic array and get the value
func (a *DynamicArray) Pop() (interface{}, error) {
	return a.Remove(a.size - 1)
}

// Shift remove a value in front of dynamic array and get the value
func (a *DynamicArray) Shift() (interface{}, error) {
	return a.Remove(0)
}
