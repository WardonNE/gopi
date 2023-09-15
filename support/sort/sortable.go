package sort

import "github.com/wardonne/gopi/support/compare"

type Sortable[T any] interface {
	Sort(comparator compare.Comparator[T])
}
