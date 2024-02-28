package collection

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/serializer"
)

// Interface is the base interface of [List] and [Set]
type Interface[E any] interface {
	serializer.JSONSerializer
	serializer.ArraySerializer[E]
	support.Stringable
	support.Clonable[Interface[E]]
	support.Countable
	support.Rangable[E]
	// IsEmpty returns wheather the collection is empty
	IsEmpty() bool
	// IsNotEmpty returns wheather the collection is not empty
	IsNotEmpty() bool
	// Contains returns wheather the collection contains element which matches the matcher
	Contains(matcher Matcher[E]) bool
	// Add add a new element
	Add(value E)
	// AddAll all new elements
	AddAll(value ...E)
	// Remove remove all elements which matches the matcher
	Remove(matcher Matcher[E])
	// Clear clears the collection
	Clear()
}
