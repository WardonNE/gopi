package list

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/collection"
	"github.com/wardonne/gopi/support/sort"
)

type List[E any] interface {
	collection.Collection[E]
	sort.Sortable[E]
	support.ReverseRangable[E]

	Get(index int) E
	Pop() E
	Shift() E
	IndexOf(matcher func(value E) bool) int
	LastIndexOf(matcher func(value E) bool) int
	SubList(from, to int) List[E]

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
	var _ List[E] = (*ArrayList[E])(nil)
	var _ List[E] = (*SyncArrayList[E])(nil)
	var _ List[E] = (*LinkedList[E])(nil)
	var _ List[E] = (*SyncLinkedList[E])(nil)
}
