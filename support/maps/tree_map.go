package maps

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wardonne/gopi/support/compare"
	"github.com/wardonne/gopi/support/tree"
)

type TreeMap[K comparable, V any] struct {
	tree  *tree.RBTree[K]
	items map[K]*Entry[K, V]
}

func NewTreeMap[K comparable, V any](comparator compare.Comparator[K]) *TreeMap[K, V] {
	t := new(TreeMap[K, V])
	t.tree = tree.NewRBTree[K](comparator)
	t.items = make(map[K]*Entry[K, V])
	return t
}

func (t *TreeMap[K, V]) MarshalJSON() ([]byte, error) {
	keys := t.tree.ToArray()
	values := make(map[K]V)
	for _, key := range keys {
		values[key] = t.items[key].Value
	}
	return json.Marshal(values)
}

func (t *TreeMap[K, V]) UnmarshalJSON(data []byte) error {
	return errors.New("not implements")
}

func (t *TreeMap[K, V]) ToMap() map[K]V {
	keys := t.tree.ToArray()
	values := make(map[K]V)
	for _, key := range keys {
		values[key] = t.items[key].Value
	}
	return values
}

func (t *TreeMap[K, V]) FromMap(values map[K]V) {
	panic(errors.New("not implements"))
}

func (t *TreeMap[K, V]) String() string {
	m := map[K]V{}
	t.tree.Range(func(value K) bool {
		m[value] = t.items[value].Value
		return true
	})
	return fmt.Sprintf("%v", m)
}

func (t *TreeMap[K, V]) Clone() Map[K, V] {
	m := NewTreeMap[K, V](t.tree.Comparator())
	keys := t.tree.ToArray()
	for _, key := range keys {
		m.Set(key, t.items[key].Value)
	}
	return m
}

func (t *TreeMap[K, V]) Get(key K) (value V) {
	if entry, ok := t.items[key]; ok {
		return entry.Value
	} else {
		return
	}
}

func (t *TreeMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if entry, ok := t.items[key]; ok {
		return entry.Value
	} else {
		return defaultValue
	}
}

func (t *TreeMap[K, V]) Set(key K, value V) {
	if _, ok := t.items[key]; ok {
		t.items[key].Value = value
	} else {
		t.items[key] = &Entry[K, V]{key, value}
		t.tree.Add(key)
	}
}

func (t *TreeMap[K, V]) Remove(key K) {
	if _, ok := t.items[key]; ok {
		delete(t.items, key)
		t.tree.Remove(key)
	}
}

func (t *TreeMap[K, V]) Keys() []K {
	return t.tree.ToArray()
}

func (t *TreeMap[K, V]) Values() (values []V) {
	keys := t.Keys()
	for _, key := range keys {
		values = append(values, t.items[key].Value)
	}
	return
}

func (t *TreeMap[K, V]) Entries() (entries []*Entry[K, V]) {
	keys := t.Keys()
	for _, key := range keys {
		entries = append(entries, t.items[key])
	}
	return
}

func (t *TreeMap[K, V]) Clear() {
	t.items = make(map[K]*Entry[K, V])
	t.tree.Clear()
}

func (t *TreeMap[K, V]) ContainsValue(matcher func(value V) bool) bool {
	for _, entry := range t.items {
		if matcher(entry.Value) {
			return true
		}
	}
	return false
}

func (t *TreeMap[K, V]) ContainsKey(key K) bool {
	_, ok := t.items[key]
	return ok
}

func (t *TreeMap[K, V]) Count() int {
	return len(t.items)
}

func (t *TreeMap[K, V]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *TreeMap[K, V]) IsNotEmpty() bool {
	return t.Count() > 0
}

func (t *TreeMap[K, V]) Comparator() compare.Comparator[K] {
	return t.tree.Comparator()
}

func (t *TreeMap[K, V]) FirstKey() (key K) {
	if k, ok := t.tree.First(); ok {
		return k
	} else {
		return
	}
}

func (t *TreeMap[K, V]) LastKey() (key K) {
	if k, ok := t.tree.Last(); ok {
		return k
	} else {
		return
	}
}

func (t *TreeMap[K, V]) Range(callback func(entry *Entry[K, V]) bool) {
	keys := t.tree.ToArray()
	for _, key := range keys {
		if !callback(t.items[key]) {
			break
		}
	}
}
