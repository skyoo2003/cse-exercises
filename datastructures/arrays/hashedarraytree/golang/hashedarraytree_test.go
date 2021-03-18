package hashedarraytree

import "testing"

func TestHashedArrayTree(t *testing.T) {
	hat := NewHashedArrayTree()
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	if err := hat.Append(1, 2, 3); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	if err := hat.Append(4, 5, 6); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(hat.top, hat.size, hat.capacity, hat.power)
	if err := hat.Insert(2, 9, 8, 7); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := hat.Prepend(10, 11); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := hat.Remove(2); err != nil {
		t.Error(err)
		t.FailNow()
	}
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
