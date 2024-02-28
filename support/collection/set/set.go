package set

import "github.com/wardonne/gopi/support/collection"

type Set[E comparable] interface {
	collection.Interface[E]
}

func implements[E comparable]() {
	var _ Set[E] = (*HashSet[E])(nil)
	var _ Set[E] = (*SyncHashSet[E])(nil)
	var _ Set[E] = (*LinkedHashSet[E])(nil)
	var _ Set[E] = (*SyncLinkedHashSet[E])(nil)
}
