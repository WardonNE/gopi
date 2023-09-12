package container

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

type Collection[E comparable] interface {
	serializer.JSONSerializer
	serializer.ArraySerializer[E]
	support.Stringable
	support.Clonable[Collection[E]]
	support.Countable
	support.Rangable[E]

	IsEmpty() bool
	IsNotEmpty() bool

	Contains(values ...E) bool
	ContainsAny(values ...E) bool

	Add(value E)
	AddAll(value ...E)

	Remove(value E)
	Clear()
}
