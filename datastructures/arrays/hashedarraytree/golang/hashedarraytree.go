package hashedarraytree

import (
	"errors"
	"math"
)

const (
	DefaultPower    = 1 // log2(top)
	DefaultTop      = 2 // pow(2, power), log2(capacity)
	DefaultCapacity = 4 // pow(2, top)
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNoData          = errors.New("no data")
)

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

func (hat *HashedArrayTree) Size() int {
	return hat.size
}

func (hat *HashedArrayTree) Capacity() int {
	return hat.capacity
}

func (hat *HashedArrayTree) Get(index int) (interface{}, error) {
	if !hat.validIndex(index) {
		return nil, ErrIndexOutOfRange
	}
	ti, li := hat.topIndex(index), hat.leafIndex(index)
	return hat.top[ti][li], nil
}

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

func (hat *HashedArrayTree) Set(index int, value interface{}) error {
	if !hat.validIndex(index) {
		return ErrIndexOutOfRange
	}
	ti, li := hat.topIndex(index), hat.leafIndex(index)
	hat.top[ti][li] = value
	return nil
}

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

func (hat *HashedArrayTree) Pop() (interface{}, error) {
	return hat.Remove(hat.size - 1)
}

func (hat *HashedArrayTree) Shift() (interface{}, error) {
	return hat.Remove(0)
}
