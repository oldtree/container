package container

import (
	"container/list"
	"sync"
)

// SyncedList represents a goroutine-safe list.
type SyncedList struct {
	*sync.RWMutex
	queue *list.List
}

// NewSyncedList returns a SyncedList pointer.
func NewSyncedList() *SyncedList {
	return &SyncedList{
		RWMutex: &sync.RWMutex{},
		queue:   list.New(),
	}
}

// Front returns the first element of slist.
func (slist *SyncedList) Front() *list.Element {
	slist.RLock()
	defer slist.RUnlock()

	return slist.queue.Front()
}

// Back returns the last element of slist.
func (slist *SyncedList) Back() *list.Element {
	slist.RLock()
	defer slist.RUnlock()

	return slist.queue.Back()
}

// PushFront pushs an element to the head of slist.
func (slist *SyncedList) PushFront(v interface{}) *list.Element {
	slist.Lock()
	defer slist.Unlock()

	return slist.queue.PushFront(v)
}

// PushBack pushs an element to the tail of slist.
func (slist *SyncedList) PushBack(v interface{}) *list.Element {
	slist.Lock()
	defer slist.Unlock()

	return slist.queue.PushBack(v)
}

// InsertBefore inserts v before mark.
func (slist *SyncedList) InsertBefore(
	v interface{}, mark *list.Element) *list.Element {

	slist.Lock()
	defer slist.Unlock()

	return slist.queue.InsertBefore(v, mark)
}

// InsertAfter inserts v after mark.
func (slist *SyncedList) InsertAfter(
	v interface{}, mark *list.Element) *list.Element {

	slist.Lock()
	defer slist.Unlock()

	return slist.queue.InsertAfter(v, mark)
}

// Remove removes e from the slist.
func (slist *SyncedList) Remove(e *list.Element) interface{} {
	slist.Lock()
	defer slist.Unlock()

	return slist.queue.Remove(e)
}

// Clear resets the list queue.
func (slist *SyncedList) Clear() {
	slist.Lock()
	defer slist.Unlock()

	slist.queue.Init()
}

// Len returns length of the slist.
func (slist *SyncedList) Len() int {
	slist.RLock()
	defer slist.RUnlock()

	return slist.queue.Len()
}

// Iter returns a chan which output all elements.
func (slist *SyncedList) Iter() <-chan *list.Element {
	ch := make(chan *list.Element)
	go func() {
		slist.RLock()
		for e := slist.queue.Front(); e != nil; e = e.Next() {
			ch <- e
		}
		slist.RUnlock()
		close(ch)
	}()
	return ch
}
