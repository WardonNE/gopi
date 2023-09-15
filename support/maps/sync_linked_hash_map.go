package maps

import (
	"errors"
	"sync"
)

type SyncLinkedHashMap[K comparable, V any] struct {
	mu      *sync.RWMutex
	hashmap *LinkedHashMap[K, V]
}

func NewSyncLinkedHashMap[K comparable, V any]() *SyncLinkedHashMap[K, V] {
	hashMap := new(SyncLinkedHashMap[K, V])
	hashMap.mu = new(sync.RWMutex)
	hashMap.hashmap = NewLinkedHashMap[K, V]()
	return hashMap
}

func (m *SyncLinkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.MarshalJSON()
}

func (m *SyncLinkedHashMap[K, V]) UnmarshalJSON(data []byte) error {
	return errors.New("not implements")
}

func (m *SyncLinkedHashMap[K, V]) ToMap() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ToMap()
}

func (m *SyncLinkedHashMap[K, V]) FromMap(values map[K]V) {
	panic("not implements")
}

func (m *SyncLinkedHashMap[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.String()
}

func (m *SyncLinkedHashMap[K, V]) Clone() Map[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	hashMap := new(SyncLinkedHashMap[K, V])
	hashMap.mu = new(sync.RWMutex)
	hashMap.hashmap = m.hashmap.Copy()
	return hashMap
}

func (m *SyncLinkedHashMap[K, V]) Copy() *SyncLinkedHashMap[K, V] {
	return m.Clone().(*SyncLinkedHashMap[K, V])
}

func (m *SyncLinkedHashMap[K, V]) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Count()
}

func (m *SyncLinkedHashMap[K, V]) Get(key K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Get(key)
}

func (m *SyncLinkedHashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.GetOrDefault(key, defaultValue)
}

func (m *SyncLinkedHashMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Set(key, value)
}

func (m *SyncLinkedHashMap[K, V]) Remove(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Remove(key)
}

func (m *SyncLinkedHashMap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Keys()
}

func (m *SyncLinkedHashMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Values()
}

func (m *SyncLinkedHashMap[K, V]) Entries() []*Entry[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Entries()
}

func (m *SyncLinkedHashMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Clear()
}

func (m *SyncLinkedHashMap[K, V]) ContainsValue(matcher func(value V) bool) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ContainsValue(matcher)
}

func (m *SyncLinkedHashMap[K, V]) ContainsKey(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ContainsKey(key)
}

func (m *SyncLinkedHashMap[K, V]) IsEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.IsEmpty()
}

func (m *SyncLinkedHashMap[K, V]) IsNotEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.IsNotEmpty()
}

func (m *SyncLinkedHashMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.hashmap.Range(callback)
}

func (m *SyncLinkedHashMap[K, V]) ReverseRange(callback func(entry *Entry[K, V]) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.hashmap.ReverseRange(callback)
}
