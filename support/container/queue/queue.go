package queue

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

type Queue[E comparable] interface {
	serializer.JSONSerializer
	support.Countable
	support.Stringable

	IsEmpty() bool
	IsNotEmpty() bool

	Peek() (E, bool)
	Enqueue()
	Dequeue()
}
