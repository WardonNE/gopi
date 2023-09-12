package sort

import "github.com/wardonne/gopi/support/compare"

type Sortable[T any] interface {
	Sort(comparer compare.Comparer[T])
}
