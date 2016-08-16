package container

import "testing"

func TestKeyedDeque(t *testing.T) {
	cases := []MapItem{
		{"a", 0},
		{"b", 1},
		{"c", 2},
	}

	deque := NewKeyedDeque()

	insert := func() {
		for _, item := range cases {
			deque.Push(item.Key, item.Val)
		}
	}

	isEmpty := func() {
		if deque.Len() != 0 {
			t.Fail()
		}
	}

	// Push
	insert()

	// Len
	if deque.Len() != 3 {
		t.Fail()
	}

	// Iter
	i := 0
	for e := range deque.Iter() {
		if e.Value.(int) != i {
			t.Fail()
		}
		i++
	}

	// HasKey, Get, Delete
	for _, item := range cases {
		if !deque.HasKey(item.Key) {
			t.Fail()
		}

		e, ok := deque.Get(item.Key)
		if !ok || e.Value.(int) != item.Val {
			t.Fail()
		}

		if deque.Delete(item.Key) != item.Val {
			t.Fail()
		}

		if deque.HasKey(item.Key) {
			t.Fail()
		}
	}
	isEmpty()

	// Clear
	insert()
	deque.Clear()
	isEmpty()
}
