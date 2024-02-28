package list

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/compare"
)

// Interface list interface
type Interface[E any] interface {
	collection.Interface[E]
	compare.Sortable[E]
	support.ReverseRangable[E]
	// Get get element by index
	Get(index int) E
	// First gets first element
	First() E
	// List gets last element
	Last() E
	// Pop removes the last element and returns it
	Pop() E
	// Shift removes the first element and returns it
	Shift() E
	// FirstWhere gets the first element which matches the matcher
	FirstWhere(matcher collection.Matcher[E]) (E, error)
	// LastWhere gets the last element which matches the matcher
	LastWhere(matcher collection.Matcher[E]) (E, error)
	// IndexOf gets the index of first element which matches the matcher
	IndexOf(matcher collection.Matcher[E]) int
	// LastIndexOf gets the index of last element which matches the matcher
	LastIndexOf(matcher collection.Matcher[E]) int
	// SubList creates a sub list
	SubList(from, to int) Interface[E]
	// Set sets value by index
	Set(index int, value E)
	// Push pushes a new element
	Push(value E)
	// PushAll pushes new elements
	PushAll(values ...E)
	// Unshift unshifts a new element
	Unshift(value E)
	// UnshiftAll unshifts new elements
	UnshiftAll(value ...E)
	// InsertBefore inserts a new element before index
	InsertBefore(index int, value E)
	// InsertAfter inserts a new element after index
	InsertAfter(index int, value E)
	// RemoveAt remove an element by specific index
	RemoveAt(index int)
	// Map proesses elements by callback
	Map(callback func(value E) E)
}

func implements[E any]() {
	var _ Interface[E] = (*ArrayList[E])(nil)
	var _ Interface[E] = (*SyncArrayList[E])(nil)
	var _ Interface[E] = (*LinkedList[E])(nil)
	var _ Interface[E] = (*SyncLinkedList[E])(nil)
}
