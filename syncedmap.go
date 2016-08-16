package container

import (
	"sync"
)

type MapItem struct {
	Key interface{}
	Val interface{}
}

// SyncedMap represents a goroutine-safe map.
type SyncedMap struct {
	*sync.RWMutex
	data map[interface{}]interface{}
}

// NewSyncedMap returns a SyncedMap pointer.
func NewSyncedMap() *SyncedMap {
	return &SyncedMap{
		RWMutex: &sync.RWMutex{},
		data:    make(map[interface{}]interface{}),
	}
}

// Get returns the value mapped to key.
func (smap *SyncedMap) Get(key interface{}) (val interface{}, ok bool) {
	smap.RLock()
	defer smap.RUnlock()

	val, ok = smap.data[key]
	return
}

// Has returns whether the SyncedMap contains the key.
func (smap *SyncedMap) Has(key interface{}) bool {
	_, ok := smap.Get(key)
	return ok
}

// Set sets pair {key: val}.
func (smap *SyncedMap) Set(key interface{}, val interface{}) {
	smap.Lock()
	defer smap.Unlock()

	smap.data[key] = val
}

// Delete deletes the key in the map.
func (smap *SyncedMap) Delete(key interface{}) {
	smap.Lock()
	defer smap.Unlock()

	delete(smap.data, key)
}

// DeleteMulti deletes keys in batch.
func (smap *SyncedMap) DeleteMulti(keys []interface{}) {
	smap.Lock()
	defer smap.Unlock()

	for _, key := range keys {
		delete(smap.data, key)
	}
}

// Clear resets the data.
func (smap *SyncedMap) Clear() {
	smap.Lock()
	defer smap.Unlock()

	smap.data = make(map[interface{}]interface{})
}

// Iter returns a chan which output all items.
func (smap *SyncedMap) Iter() <-chan MapItem {
	ch := make(chan MapItem)
	go func() {
		smap.RLock()
		for key, val := range smap.data {
			ch <- MapItem{
				Key: key,
				Val: val,
			}
		}
		smap.RUnlock()
		close(ch)
	}()
	return ch
}

// Len returns the length of SyncedMap.
func (smap *SyncedMap) Len() int {
	smap.RLock()
	defer smap.RUnlock()

	return len(smap.data)
}
