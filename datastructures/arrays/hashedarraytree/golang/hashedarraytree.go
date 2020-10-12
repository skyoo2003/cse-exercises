package hashedarraytree

import (
	"errors"
	"math"
)

const (
	// DefaultPower log2(top)
	DefaultPower = 1
	// DefaultTop pow(2, power), log2(capacity)
	DefaultTop = 2
	// DefaultCapacity pow(2, top)
	DefaultCapacity = 4
)

var (
	// ErrIndexOutOfRange index out of rnage in hat
	ErrIndexOutOfRange = errors.New("index out of range")
	// ErrNoData can not found any data in hat
	ErrNoData = errors.New("no data")
)

// HashedArrayTree hashed array tree data structure
type HashedArrayTree struct {
	top      [][]interface{}
	capacity int
	power    int
	size     int
}

func makeTop(size int) [][]interface{} {
	top := make([][]interface{}, size)
	for i := 0; i < size; i++ {
		top[i] = make([]interface{}, size)
	}
	return top
}

func pow2(n int) int {
	return int(math.Pow(2.0, float64(n)))
}

func log2(n int) int {
	return int(math.Log2(float64(n)))
}

// NewHashedArrayTree create a hashed array tree
func NewHashedArrayTree() HashedArrayTree {
	return HashedArrayTree{
		top:      makeTop(DefaultTop),
		capacity: DefaultCapacity,
		power:    DefaultPower,
		size:     0,
	}
}

func (hat *HashedArrayTree) validIndex(index int) bool {
	if index < 0 || index > hat.size {
		return false
	}
	return true
}

func (hat *HashedArrayTree) resize(claimSize int) error {
	if claimSize > hat.capacity {
		for claimSize > hat.capacity {
			if err := hat.expand(); err != nil {
				return err
			}
		}
	} else if claimSize < log2(hat.capacity) {
		for claimSize < log2(hat.capacity) {
			if err := hat.shrink(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (hat *HashedArrayTree) expand() error {
	if hat.capacity < DefaultCapacity {
		hat.realloc(1)
	} else {
		hat.realloc(hat.power + 1)
	}
	return nil
}

func (hat *HashedArrayTree) shrink() error {
	if hat.capacity > DefaultCapacity {
		hat.realloc(hat.power - 1)
	} else {
		hat.realloc(1)
	}
	return nil
}

func (hat *HashedArrayTree) realloc(newPower int) error {
	if newPower == hat.power {
		return nil
	}
	newTop := makeTop(pow2(newPower))
	if newPower > hat.power {
		for i, j := 0, 0; i < len(newTop) && j < len(hat.top); i, j = i+1, j+2 {
			copy(newTop[i], append(hat.top[j], hat.top[j+1]...))
		}
	} else if newPower < hat.power {
		for i, j := 0, 0; i < len(newTop) && j < len(hat.top); i, j = i+2, j+1 {
			copy(newTop[i], hat.top[j][0:len(hat.top[j])/2])
			copy(newTop[i+1], hat.top[j][len(hat.top[j])/2:len(hat.top[j])])
		}
	}
	hat.top = newTop
	hat.power = newPower
	hat.capacity = pow2(len(newTop))
	return nil
}

func (hat *HashedArrayTree) topIndex(index int) int {
	return index >> hat.power
}

func (hat *HashedArrayTree) leafIndex(index int) int {
	return index & ((1 << hat.power) - 1)
}

// Size get size of hashed array tree
func (hat *HashedArrayTree) Size() int {
	return hat.size
}

// Capacity get capacity of hashed array tree
func (hat *HashedArrayTree) Capacity() int {
	return hat.capacity
}

// Get get a value by index in hashed array tree
func (hat *HashedArrayTree) Get(index int) (interface{}, error) {
	if !hat.validIndex(index) {
		return nil, ErrIndexOutOfRange
	}
	ti, li := hat.topIndex(index), hat.leafIndex(index)
	return hat.top[ti][li], nil
}

// Range get values between begin and end in hahsed array tree
func (hat *HashedArrayTree) Range(begin, end int) ([]interface{}, error) {
	if !hat.validIndex(begin) || !hat.validIndex(end) {
		return nil, ErrIndexOutOfRange
	}
	buf := make([]interface{}, 0)
	for i := begin; i < end; i++ {
		if value, err := hat.Get(i); err == nil {
			buf = append(buf, value)
		} else {
			return nil, err
		}
	}
	return buf, nil
}

// Set set a value by index in hashed array tree
func (hat *HashedArrayTree) Set(index int, value interface{}) error {
	if !hat.validIndex(index) {
		return ErrIndexOutOfRange
	}
	ti, li := hat.topIndex(index), hat.leafIndex(index)
	hat.top[ti][li] = value
	return nil
}

// Append insert values in the end of hashed array tree
func (hat *HashedArrayTree) Append(values ...interface{}) error {
	newSize := hat.size + len(values)
	if err := hat.resize(newSize); err != nil {
		return err
	}
	for i, j := hat.size, 0; i < newSize; i, j = i+1, j+1 {
		ti, li := hat.topIndex(i), hat.leafIndex(i)
		hat.top[ti][li] = values[j]
	}
	hat.size = newSize
	return nil
}

// Prepend insert values in front of hashed array tree
func (hat *HashedArrayTree) Prepend(values ...interface{}) error {
	lenValues := len(values)
	newSize := hat.size + lenValues
	if err := hat.resize(newSize); err != nil {
		return err
	}
	// Move items in the buffer by the length of values
	for i := hat.size - 1; i >= 0; i-- {
		bti, bli := hat.topIndex(i), hat.leafIndex(i)
		ati, ali := hat.topIndex(i+lenValues), hat.leafIndex(i+lenValues)
		hat.top[ati][ali] = hat.top[bti][bli]
	}
	for i := 0; i < lenValues; i++ {
		ti, li := hat.topIndex(i), hat.leafIndex(i)
		hat.top[ti][li] = values[i]
	}
	hat.size = newSize
	return nil
}

// Insert insert values by index in hashed array tree
func (hat *HashedArrayTree) Insert(index int, values ...interface{}) error {
	lenValues := len(values)
	newSize := hat.size + lenValues
	if err := hat.resize(newSize); err != nil {
		return err
	}
	// Move items in the middle of the buffer by the length of values
	for i := hat.size - 1; i >= index; i-- {
		bti, bli := hat.topIndex(i), hat.leafIndex(i)
		ati, ali := hat.topIndex(i+lenValues), hat.leafIndex(i+lenValues)
		hat.top[ati][ali] = hat.top[bti][bli]
	}
	for i, j := index, 0; i < index+lenValues; i, j = i+1, j+1 {
		ti, li := hat.topIndex(i), hat.leafIndex(i)
		hat.top[ti][li] = values[j]
	}
	hat.size = newSize
	return nil
}

// Remove remove a value by index from hashed array tree
func (hat *HashedArrayTree) Remove(index int) (interface{}, error) {
	if hat.size == 0 {
		return nil, ErrNoData
	}
	if !hat.validIndex(index) {
		return nil, ErrIndexOutOfRange
	}
	newSize := hat.size - 1
	if err := hat.resize(newSize); err != nil {
		return nil, err
	}

	value, _ := hat.Get(index)
	for i := index + 1; i < hat.size; i++ {
		bti, bli := hat.topIndex(i), hat.leafIndex(i)
		ati, ali := hat.topIndex(i-1), hat.leafIndex(i-1)

		hat.top[ati][ali] = hat.top[bti][bli]
	}
	hat.size = newSize
	return value, nil
}

// Pop remove a value in the end of hashed array tree and get the value
func (hat *HashedArrayTree) Pop() (interface{}, error) {
	return hat.Remove(hat.size - 1)
}

// Shift remove a value in front of hashed array tree and get the value
func (hat *HashedArrayTree) Shift() (interface{}, error) {
	return hat.Remove(0)
}
