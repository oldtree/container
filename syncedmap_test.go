package container

import (
	"sync"
	"testing"
)

func TestSyncedMap(t *testing.T) {
	cases := []MapItem{
		{"a", 0},
		{"b", 1},
		{"c", 2},
	}

	sm := NewSyncedMap()

	set := func() {
		group := sync.WaitGroup{}
		for _, item := range cases {
			group.Add(1)
			go func(item MapItem) {
				sm.Set(item.Key, item.Val)
				group.Done()
			}(item)
		}
		group.Wait()
	}

	isEmpty := func() {
		if sm.Len() != 0 {
			t.Fail()
		}
	}

	// Set
	set()
	if sm.Len() != len(cases) {
		t.Fail()
	}

Loop:
	// Iter
	for item := range sm.Iter() {
		for _, c := range cases {
			if item.Key == c.Key && item.Val == c.Val {
				continue Loop
			}
		}
		t.Fail()
	}

	// Get, Delete, Has
	for _, item := range cases {
		val, ok := sm.Get(item.Key)
		if !ok || val != item.Val {
			t.Fail()
		}

		sm.Delete(item.Key)
		if sm.Has(item.Key) {
			t.Fail()
		}
	}
	isEmpty()

	// DeleteMulti
	set()
	sm.DeleteMulti([]interface{}{"a", "b", "c"})
	isEmpty()

	// Clear
	set()
	sm.Clear()
	isEmpty()
}
