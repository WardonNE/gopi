package maps

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/compare"
	"github.com/wardonne/gopi/support/serializer"
)

type Map[K comparable, V any] interface {
	serializer.JSONSerializer
	serializer.MapSerializer[K, V]
	support.Countable
	support.Rangable[*Entry[K, V]]
	support.Stringable
	support.Clonable[Map[K, V]]

	Get(key K) V
	GetOrDefault(key K, defaultValue V) V
	Set(key K, value V)
	Remove(key K)
	Keys() []K
	Values() []V
	Entries() []*Entry[K, V]
	Clear()
	ContainsValue(matcher func(value V) bool) bool
	ContainsKey(key K) bool
	IsEmpty() bool
	IsNotEmpty() bool
}

type OrderedMap[K comparable, V any] interface {
	Map[K, V]

	Comparator() compare.Comparator[K]
	FirstKey() K
	LastKey() K
}

func implements[K, V comparable]() {
	var _ Map[K, V] = (*HashMap[K, V])(nil)
	var _ Map[K, V] = (*SyncHashMap[K, V])(nil)
	var _ Map[K, V] = (*LinkedHashMap[K, V])(nil)
	var _ Map[K, V] = (*SyncLinkedHashMap[K, V])(nil)

	var _ OrderedMap[K, V] = (*TreeMap[K, V])(nil)
	var _ OrderedMap[K, V] = (*SyncTreeMap[K, V])(nil)
}
