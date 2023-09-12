package list

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/container"
	"github.com/wardonne/gopi/support/sort"
)

type List[E comparable] interface {
	container.Collection[E]
	sort.Sortable[E]
	support.ReverseRangable[E]

	Get(index int) E
	Pop() E
	Shift() E
	IndexOf(value E) int
	LastIndexOf(value E) int
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
