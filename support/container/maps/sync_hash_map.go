package maps

import "sync"

type SyncHashMap[K, V comparable] struct {
	mu      *sync.RWMutex
	hashmap *HashMap[K, V]
}

func NewSyncHashMap[K, V comparable]() *SyncHashMap[K, V] {
	hashMap := new(SyncHashMap[K, V])
	hashMap.mu = new(sync.RWMutex)
	hashMap.hashmap = NewHashMap[K, V]()
	return hashMap
}

func (m *SyncHashMap[K, V]) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.MarshalJSON()
}

func (m *SyncHashMap[K, V]) UnmarshalJSON(data []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.hashmap.UnmarshalJSON(data)
}

func (m *SyncHashMap[K, V]) ToMap() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ToMap()
}

func (m *SyncHashMap[K, V]) FromMap(values map[K]V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.FromMap(values)
}

func (m *SyncHashMap[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.String()
}

func (m *SyncHashMap[K, V]) Clone() Map[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	hashMap := NewSyncHashMap[K, V]()
	hashMap.hashmap = NewHashMap[K, V]()
	hashMap.hashmap.items = m.hashmap.items
	return hashMap
}

func (m *SyncHashMap[K, V]) Copy() *SyncHashMap[K, V] {
	return m.Clone().(*SyncHashMap[K, V])
}

func (m *SyncHashMap[K, V]) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Count()
}

func (m *SyncHashMap[K, V]) Get(key K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Get(key)
}

func (m *SyncHashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.GetOrDefault(key, defaultValue)
}

func (m *SyncHashMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Set(key, value)
}

func (m *SyncHashMap[K, V]) Remove(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Remove(key)
}

func (m *SyncHashMap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Keys()
}

func (m *SyncHashMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Values()
}

func (m *SyncHashMap[K, V]) Entries() []*Entry[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.Entries()
}

func (m *SyncHashMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hashmap.Clear()
}

func (m *SyncHashMap[K, V]) ContainsValue(value V) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ContainsValue(value)
}

func (m *SyncHashMap[K, V]) ContainsKey(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.ContainsKey(key)
}

func (m *SyncHashMap[K, V]) IsEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.IsEmpty()
}

func (m *SyncHashMap[K, V]) IsNotEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashmap.IsNotEmpty()
}

func (m *SyncHashMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.hashmap.Range(callback)
}
