package compare

type Sortable[T any] interface {
	Sort(comparator Comparator[T])
}
