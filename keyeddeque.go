package container

import (
	"container/list"
	"sync"
)

// KeyedDeque represents a keyed deque.
type KeyedDeque struct {
	*sync.RWMutex
	*SyncedList
	index         map[interface{}]*list.Element
	invertedIndex map[*list.Element]interface{}
}

// NewKeyedDeque returns a KeyedDeque pointer.
func NewKeyedDeque() *KeyedDeque {
	return &KeyedDeque{
		RWMutex:       &sync.RWMutex{},
		SyncedList:    NewSyncedList(),
		index:         make(map[interface{}]*list.Element),
		invertedIndex: make(map[*list.Element]interface{}),
	}
}

// Push pushs a keyed-value to the end of deque.
func (deque *KeyedDeque) Push(key interface{}, val interface{}) {
	deque.Lock()
	defer deque.Unlock()

	if e, ok := deque.index[key]; ok {
		deque.SyncedList.Remove(e)
	}
	deque.index[key] = deque.SyncedList.PushBack(val)
	deque.invertedIndex[deque.index[key]] = key
}

// Get returns the keyed value.
func (deque *KeyedDeque) Get(key interface{}) (*list.Element, bool) {
	deque.RLock()
	defer deque.RUnlock()

	v, ok := deque.index[key]
	return v, ok
}

// Has returns whether key already exists.
func (deque *KeyedDeque) HasKey(key interface{}) bool {
	_, ok := deque.Get(key)
	return ok
}

// Delete deletes a value named key.
func (deque *KeyedDeque) Delete(key interface{}) (v interface{}) {
	deque.RLock()
	e, ok := deque.index[key]
	deque.RUnlock()

	deque.Lock()
	defer deque.Unlock()

	if ok {
		v = deque.SyncedList.Remove(e)
		delete(deque.index, key)
		delete(deque.invertedIndex, e)
	}

	return
}

// Removes overwrites list.List.Remove.
func (deque *KeyedDeque) Remove(e *list.Element) (v interface{}) {
	deque.RLock()
	key, ok := deque.invertedIndex[e]
	deque.RUnlock()

	if ok {
		v = deque.Delete(key)
	}

	return
}

// Clear resets the deque.
func (deque *KeyedDeque) Clear() {
	deque.Lock()
	defer deque.Unlock()

	deque.SyncedList.Clear()
	deque.index = make(map[interface{}]*list.Element)
	deque.invertedIndex = make(map[*list.Element]interface{})
}
