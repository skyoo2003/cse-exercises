package dynamicarray

import "testing"

func TestDynamicArray(t *testing.T) {
	arr := NewDynamicArrayWithSize(2)
	if arr.Size() != 0 {
		t.Error("Invalid initial size")
	}
	if arr.Capacity() != 2 {
		t.Error("Invalid initial capacity")
	}
	arr.Append(1, 2, 3)
	if arr.Size() != 3 {
		t.Error("Invalid size after appending")
	}
	if arr.Capacity() == 2 {
		t.Error("Invalid capacity after appending")
	}
	arr.Append(4, 5, 6)
	arr.Append(4, 5, 6)
	arr.Insert(2, 9, 8, 7)
	arr.Prepend(10, 11)

	expected := []interface{}{10, 11, 1, 2, 9, 8, 7, 3, 4, 5, 6, 4, 5, 6}
	actual, _ := arr.Range(0, arr.Size())
	if len(actual) != len(expected) {
		t.Error("Invalid buffer")
	}
	for i, value := range actual {
		if expected[i] != value {
			t.Error("Invalid buffer")
			t.FailNow()
		}
	}
	arr.Remove(2)
	t.Log(arr.Range(0, arr.Size()))
	t.Log(arr.Size(), arr.Capacity())
}
