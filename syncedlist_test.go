package container

import "testing"

func TestSyncedList(t *testing.T) {
	sl := NewSyncedList()

	insert := func() {
		for i := 0; i < 10; i++ {
			sl.PushBack(i)
		}
	}

	isEmpty := func() {
		if sl.Len() != 0 {
			t.Fail()
		}
	}

	// PushBack
	insert()

	// Len
	if sl.Len() != 10 {
		t.Fail()
	}

	// Iter
	i := 0
	for item := range sl.Iter() {
		if item.Value.(int) != i {
			t.Fail()
		}
		i++
	}

	// Front
	if sl.Front().Value.(int) != 0 {
		t.Fail()
	}

	// Back
	if sl.Back().Value.(int) != 9 {
		t.Fail()
	}

	// Remove
	for i := 0; i < 10; i++ {
		if sl.Remove(sl.Front()).(int) != i {
			t.Fail()
		}
	}
	isEmpty()

	// Clear
	insert()
	sl.Clear()
	isEmpty()
}
