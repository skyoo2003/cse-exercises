package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(10, 1)
	sl.Insert(4, 2)
	sl.Insert(7, 3)
	sl.Insert(3, 4)
	sl.Insert(9, 5)
	sl.Insert(5, 6)
	sl.Insert(1, 7)
	sl.Insert(11, 8)
	sl.Insert(17, 9)
	sl.Insert(2, 10)
	sl.Insert(12, 11)
	sl.Insert(15, 12)
	sl.Insert(13, 13)
	sl.Insert(19, 14)

	node, _ := sl.Search(11)
	t.Log(node.key, node.value)

	t.Log("\n" + sl.String())
	sl.Delete(11)
	t.Log("\n" + sl.String())
}
