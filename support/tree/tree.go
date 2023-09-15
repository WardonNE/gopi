package tree

import (
	"github.com/wardonne/gopi/support"
	"github.com/wardonne/gopi/support/compare"
	"github.com/wardonne/gopi/support/serializer"
)

type Tree[E any] interface {
	serializer.JSONSerializer
	serializer.ArraySerializer[E]
	support.Stringable
	support.Clonable[Tree[E]]
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
	Comparator() compare.Comparator[E]
	First() (E, bool)
	Last() (E, bool)
}

func implements[E any]() {
	var _ Tree[E] = (*AVLTree[E])(nil)
	var _ Tree[E] = (*RBTree[E])(nil)
	var _ Tree[E] = (*SyncAVLTree[E])(nil)
	var _ Tree[E] = (*SyncRBTree[E])(nil)
}
