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

	Get(index int) E
	First() E
	Last() E
	Pop() E
	Shift() E
	FirstWhere(matcher collection.Matcher[E]) (E, error)
	LastWhere(matcher collection.Matcher[E]) (E, error)
	IndexOf(matcher collection.Matcher[E]) int
	LastIndexOf(matcher collection.Matcher[E]) int
	SubList(from, to int) Interface[E]
	Where(matcher collection.Matcher[E]) Interface[E]

	Set(index int, value E)
	Push(value E)
	PushAll(values ...E)
	Unshift(value E)
	UnshiftAll(value ...E)
	InsertBefore(index int, value E)
	InsertAfter(index int, value E)

	RemoveAt(index int)

	Map(callback func(value E) E)
}

func implements[E any]() {
	var _ Interface[E] = (*ArrayList[E])(nil)
	var _ Interface[E] = (*SyncArrayList[E])(nil)
	var _ Interface[E] = (*LinkedList[E])(nil)
	var _ Interface[E] = (*SyncLinkedList[E])(nil)
}
