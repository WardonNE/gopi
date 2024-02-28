package compare

// Sortable sortable interface
type Sortable[T any] interface {
	Sort(comparator Comparator[T])
}
