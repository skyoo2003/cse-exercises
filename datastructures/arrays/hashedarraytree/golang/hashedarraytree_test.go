package hashedarraytree

import "testing"

func TestHashedArrayTree(t *testing.T) {
	hat := NewHashedArrayTree()
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	hat.Append(1, 2, 3)
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	hat.Append(4, 5, 6)
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	hat.Insert(2, 9, 8, 7)
	hat.Prepend(10, 11)
	hat.Remove(2)
	// 10, 11, 2, 9, 8, 7, 3, 4, 5, 6
	t.Log(hat.top, hat.size, hat.capacity, hat.power)

	expected := []interface{}{10, 11, 2, 9, 8, 7, 3, 4, 5, 6}
	actual, _ := hat.Range(0, hat.Size())
	t.Log(actual)
	if len(actual) != len(expected) {
		t.Error("invalid range size")
	}
	for i, value := range actual {
		if expected[i] != value {
			t.Error("Invalid range")
			t.FailNow()
		}
	}
}
