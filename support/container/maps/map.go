package maps

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

type Map[K, V comparable] interface {
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
	ContainsValue(value V) bool
	ContainsKey(key K) bool
	IsEmpty() bool
	IsNotEmpty() bool
}
