package maps

import (
	"encoding/json"
	"fmt"

	"github.com/wardonne/gopi/support/builder"
)

type HashMap[K comparable, V any] struct {
	items map[K]*Entry[K, V]
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	hashMap := new(HashMap[K, V])
	hashMap.items = map[K]*Entry[K, V]{}
	return hashMap
}

func (m *HashMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.items)
}

func (m *HashMap[K, V]) UnmarshalJSON(data []byte) error {
	values := map[K]V{}
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	for key, value := range values {
		m.items[key] = &Entry[K, V]{key, value}
	}
	return nil
}

func (m *HashMap[K, V]) ToMap() map[K]V {
	values := make(map[K]V)
	for _, item := range m.items {
		values[item.Key] = item.Value
	}
	return values
}

func (m *HashMap[K, V]) FromMap(values map[K]V) {
	for key, value := range values {
		m.items[key] = &Entry[K, V]{key, value}
	}
}

func (m *HashMap[K, V]) String() string {
	if bytes, err := m.MarshalJSON(); err != nil {
		builder := builder.NewStringBuilder("{")
		for _, value := range m.items {
			builder.WriteString(fmt.Sprintf("%v: %v", value.Key, value.Value))
			builder.WriteRune(' ')
		}
		builder.TrimSpace()
		builder.WriteRune('}')
		return builder.String()
	} else {
		return string(bytes)
	}
}

func (m *HashMap[K, V]) Clone() Map[K, V] {
	hashMap := NewHashMap[K, V]()
	hashMap.items = m.items
	return hashMap
}

func (m *HashMap[K, V]) Copy() *HashMap[K, V] {
	return m.Clone().(*HashMap[K, V])
}

func (m *HashMap[K, V]) Count() int {
	return len(m.items)
}

func (m *HashMap[K, V]) Get(key K) (value V) {
	if v, ok := m.items[key]; ok {
		return v.Value
	} else {
		return
	}
}

func (m *HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, ok := m.items[key]; ok {
		return value.Value
	} else {
		return defaultValue
	}
}

func (m *HashMap[K, V]) Set(key K, value V) {
	if v, ok := m.items[key]; ok {
		v.Value = value
	} else {
		m.items[key] = &Entry[K, V]{key, value}
	}
}

func (m *HashMap[K, V]) Remove(key K) {
	delete(m.items, key)
}

func (m *HashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.items))
	for key := range m.items {
		keys = append(keys, key)
	}
	return keys
}

func (m *HashMap[K, V]) Values() []V {
	values := make([]V, len(m.items))
	for _, item := range m.items {
		values = append(values, item.Value)
	}
	return values
}

func (m *HashMap[K, V]) Entries() []*Entry[K, V] {
	entries := make([]*Entry[K, V], len(m.items))
	for _, item := range m.items {
		entries = append(entries, item)
	}
	return entries
}

func (m *HashMap[K, V]) Clear() {
	m.items = make(map[K]*Entry[K, V])
}

func (m *HashMap[K, V]) ContainsValue(matcher func(value V) bool) bool {
	for _, item := range m.items {
		if matcher(item.Value) {
			return true
		}
	}
	return false
}

func (m *HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.items[key]
	return ok
}

func (m *HashMap[K, V]) IsEmpty() bool {
	return len(m.items) == 0
}

func (m *HashMap[K, V]) IsNotEmpty() bool {
	return len(m.items) > 0
}

func (m *HashMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	for _, entry := range m.items {
		if !callback(entry) {
			break
		}
	}
}
