package maps

import (
	"errors"

	"github.com/wardonne/gopi/support/collection/list"
)

type LinkedHashMap[K comparable, V any] struct {
	hashmap *HashMap[K, V]
	list    *list.LinkedList[*Entry[K, V]]
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	m := new(LinkedHashMap[K, V])
	m.hashmap = NewHashMap[K, V]()
	m.list = list.NewLinkedList[*Entry[K, V]]()
	return m
}

func (m *LinkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
	return m.hashmap.MarshalJSON()
}

func (m *LinkedHashMap[K, V]) UnmarshalJSON(data []byte) error {
	return errors.New("not implements")
}

func (m *LinkedHashMap[K, V]) ToMap() map[K]V {
	return m.hashmap.ToMap()
}

func (m *LinkedHashMap[K, V]) FromMap(values map[K]V) {
	panic("not implements")
}

func (m *LinkedHashMap[K, V]) String() string {
	return m.hashmap.String()
}

func (m *LinkedHashMap[K, V]) Clone() Map[K, V] {
	hashMap := NewLinkedHashMap[K, V]()
	m.Range(func(entry *Entry[K, V]) bool {
		newEntry := &Entry[K, V]{entry.Key, entry.Value}
		hashMap.hashmap.items[entry.Key] = newEntry
		hashMap.list.Add(newEntry)
		return true
	})
	return hashMap
}

func (m *LinkedHashMap[K, V]) Copy() *LinkedHashMap[K, V] {
	return m.Clone().(*LinkedHashMap[K, V])
}

func (m *LinkedHashMap[K, V]) Count() int {
	return m.hashmap.Count()
}

func (m *LinkedHashMap[K, V]) Get(key K) V {
	return m.hashmap.Get(key)
}

func (m *LinkedHashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	return m.hashmap.GetOrDefault(key, defaultValue)
}

func (m *LinkedHashMap[K, V]) Set(key K, value V) {
	if entry, ok := m.hashmap.items[key]; ok {
		entry.Value = value
	} else {
		entry = &Entry[K, V]{key, value}
		m.hashmap.items[key] = entry
		m.list.Add(entry)
	}
}

func (m *LinkedHashMap[K, V]) Remove(key K) {
	if _, ok := m.hashmap.items[key]; ok {
		m.hashmap.Remove(key)
		m.list.Remove(func(value *Entry[K, V]) bool {
			return value.Key == key
		})
	}
}

func (m *LinkedHashMap[K, V]) Keys() []K {
	var keys = []K{}
	m.list.Range(func(value *Entry[K, V]) bool {
		keys = append(keys, value.Key)
		return true
	})
	return keys
}

func (m *LinkedHashMap[K, V]) Values() []V {
	var values = []V{}
	m.list.Range(func(value *Entry[K, V]) bool {
		values = append(values, value.Value)
		return true
	})
	return values
}

func (m *LinkedHashMap[K, V]) Entries() []*Entry[K, V] {
	return m.list.ToArray()
}

func (m *LinkedHashMap[K, V]) Clear() {
	m.hashmap.Clear()
	m.list.Clear()
}

func (m *LinkedHashMap[K, V]) ContainsValue(matcher func(value V) bool) bool {
	return m.hashmap.ContainsValue(matcher)
}

func (m *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	return m.hashmap.ContainsKey(key)
}

func (m *LinkedHashMap[K, V]) IsEmpty() bool {
	return m.hashmap.IsEmpty()
}

func (m *LinkedHashMap[K, V]) IsNotEmpty() bool {
	return m.hashmap.IsNotEmpty()
}

func (m *LinkedHashMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	m.list.Range(callback)
}

func (m *LinkedHashMap[K, V]) ReverseRange(callback func(entry *Entry[K, V]) bool) {
	m.list.ReverseRange(callback)
}
