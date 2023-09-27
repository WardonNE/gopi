package collection

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

type Collection[E any] interface {
	serializer.JSONSerializer
	serializer.ArraySerializer[E]
	support.Stringable
	support.Clonable[Collection[E]]
	support.Countable
	support.Rangable[E]

	IsEmpty() bool
	IsNotEmpty() bool
	Contains(matcher func(value E) bool) bool
	Add(value E)
	AddAll(value ...E)
	Remove(matcher func(value E) bool)
	Clear()
}
