package maps

import (
	"errors"
	"sync"

	"github.com/wardonne/gopi/support/compare"
)

type SyncTreeMap[K comparable, V any] struct {
	mu    sync.RWMutex
	items *TreeMap[K, V]
}

func NewSyncTreeMap[K comparable, V any](comparator compare.Comparator[K]) *SyncTreeMap[K, V] {
	t := new(SyncTreeMap[K, V])
	t.items = NewTreeMap[K, V](comparator)
	return t
}

func (t *SyncTreeMap[K, V]) MarshalJSON() ([]byte, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.MarshalJSON()
}

func (t *SyncTreeMap[K, V]) UnmarshalJSON(data []byte) error {
	return errors.New("not implements")
}

func (t *SyncTreeMap[K, V]) ToMap() map[K]V {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.ToMap()
}

func (t *SyncTreeMap[K, V]) FromMap(values map[K]V) {
	panic(errors.New("not implements"))
}

func (t *SyncTreeMap[K, V]) String() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.String()
}

func (t *SyncTreeMap[K, V]) Clone() Map[K, V] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	m := NewSyncTreeMap[K, V](t.items.Comparator())
	m.items = t.items.Clone().(*TreeMap[K, V])
	return m
}

func (t *SyncTreeMap[K, V]) Copy() *SyncTreeMap[K, V] {
	return t.Clone().(*SyncTreeMap[K, V])
}

func (t *SyncTreeMap[K, V]) Get(key K) (value V) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.Get(key)
}

func (t *SyncTreeMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.GetOrDefault(key, defaultValue)
}

func (t *SyncTreeMap[K, V]) Set(key K, value V) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.items.Set(key, value)
}

func (t *SyncTreeMap[K, V]) Remove(key K) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.items.Remove(key)
}

func (t *SyncTreeMap[K, V]) Keys() []K {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.Keys()
}

func (t *SyncTreeMap[K, V]) Values() []V {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.Values()
}

func (t *SyncTreeMap[K, V]) Entries() []*Entry[K, V] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.Entries()
}

func (t *SyncTreeMap[K, V]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.items.Clear()
}

func (t *SyncTreeMap[K, V]) ContainsValue(matcher func(value V) bool) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.ContainsValue(matcher)
}

func (t *SyncTreeMap[K, V]) ContainsKey(key K) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.ContainsKey(key)
}

func (t *SyncTreeMap[K, V]) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.Count()
}

func (t *SyncTreeMap[K, V]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *SyncTreeMap[K, V]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *SyncTreeMap[K, V]) Comparator() compare.Comparator[K] {
	return t.items.Comparator()
}

func (t *SyncTreeMap[K, V]) FirstKey() K {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.FirstKey()
}

func (t *SyncTreeMap[K, V]) LastKey() K {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items.LastKey()
}

func (t *SyncTreeMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.items.Range(callback)
}
