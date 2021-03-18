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
	if err := arr.Append(1, 2, 3); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if arr.Size() != 3 {
		t.Error("Invalid size after appending")
	}
	if arr.Capacity() == 2 {
		t.Error("Invalid capacity after appending")
	}
	if err := arr.Append(4, 5, 6); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := arr.Append(4, 5, 6); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := arr.Insert(2, 9, 8, 7); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := arr.Prepend(10, 11); err != nil {
		t.Error(err)
		t.FailNow()
	}

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
	if _, err := arr.Remove(2); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(arr.Range(0, arr.Size()))
	t.Log(arr.Size(), arr.Capacity())
}
