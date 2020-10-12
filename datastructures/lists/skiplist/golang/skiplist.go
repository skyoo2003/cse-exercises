package skiplist

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	// ErrNotFound not found in skip list
	ErrNotFound = errors.New("not found")
	// ErrOutOfLevel out of level in skip list
	ErrOutOfLevel = errors.New("out of level")
)

const (
	// MaxLevel maximum level
	MaxLevel = 32
	// Probability probability to increase level randomly
	Probability = 0.5
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randomLevel() int {
	level := 1
	for rand.Float64() < Probability && level < MaxLevel {
		level++
	}
	return level
}

func getNode(nodes []*Node, level int) *Node {
	return nodes[level-1]
}

func setNode(nodes []*Node, level int, node *Node) {
	nodes[level-1] = node
}

// Node skip list's node data structure
type Node struct {
	key      int
	value    interface{}
	forwards []*Node
}

// NewNode create a node
func NewNode(key int, value interface{}) *Node {
	return &Node{
		key:      key,
		value:    value,
		forwards: make([]*Node, MaxLevel),
	}
}

// Forward get a forward node by level
func (n *Node) Forward(level int) *Node {
	return getNode(n.forwards, level)
}

// SkipList skip list data structure
type SkipList struct {
	header *Node
	size   int
	level  int
}

// NewSkipList create a skip list
func NewSkipList() *SkipList {
	header := NewNode(math.MaxInt64, nil)
	for i := 0; i < MaxLevel; i++ {
		header.forwards[i] = header
	}
	return &SkipList{
		header: header,
		size:   0,
		level:  1,
	}
}

// Size get size of skip list
func (sl *SkipList) Size() int {
	return sl.size
}

// Level get level of skip list
func (sl *SkipList) Level() int {
	return sl.level
}

// Front get header of skip list
func (sl *SkipList) Front() *Node {
	return sl.header
}

// Search find key in skip list
func (sl *SkipList) Search(key int) (*Node, error) {
	x := sl.header
	for lvl := sl.level; lvl >= 1; lvl-- {
		for x.Forward(lvl).key < key {
			x = x.Forward(lvl)
		}
	}
	x = x.Forward(1)
	if x.key == key {
		return x, nil
	}
	return nil, ErrNotFound
}

func (sl *SkipList) getUpdateAndCursor(key int) (update []*Node, x *Node) {
	update = make([]*Node, MaxLevel)
	x = sl.header
	for lvl := sl.level; lvl >= 1; lvl-- {
		for x.Forward(lvl).key < key {
			x = x.Forward(lvl)
		}
		setNode(update, lvl, x)
	}
	x = x.Forward(1)
	return
}

// Insert add an item into skip list
func (sl *SkipList) Insert(key int, value interface{}) error {
	update, x := sl.getUpdateAndCursor(key)
	if x.key == key {
		x.value = value
	} else {
		newLevel := randomLevel()
		if newLevel > sl.level {
			for lvl := sl.level + 1; lvl <= newLevel; lvl++ {
				setNode(update, lvl, sl.header)
			}
			sl.level = newLevel
		}
		x = NewNode(key, value)
		for lvl := 1; lvl <= newLevel; lvl++ {
			setNode(x.forwards, lvl, getNode(update, lvl).Forward(lvl))
			setNode(getNode(update, lvl).forwards, lvl, x)
		}
	}
	sl.size++
	return nil
}

// Delete delete an item by key from skip list
func (sl *SkipList) Delete(key int) error {
	update, x := sl.getUpdateAndCursor(key)
	if x.key == key {
		for lvl := 1; lvl <= sl.level; lvl++ {
			if getNode(update, lvl).Forward(lvl) != x {
				break
			}
			setNode(getNode(update, lvl).forwards, lvl, x.Forward(lvl))
		}
		x = nil
		for sl.level > 1 && sl.header.Forward(sl.level) == nil {
			sl.level--
		}
		sl.size--
	}
	return ErrNotFound
}

// String print information of skip list for debugging
func (sl *SkipList) String() string {
	var buf bytes.Buffer
	for lvl := sl.level; lvl >= 1; lvl-- {
		x := sl.header
		buf.WriteString(fmt.Sprintf("L(%d): ", lvl))
		for x.Forward(lvl) != sl.header {
			buf.WriteString(fmt.Sprintf("%d[%v]->", x.Forward(lvl).key, x.Forward(lvl).value))
			x = x.Forward(lvl)
		}
		buf.WriteString("NIL\n")
	}
	return buf.String()
}
