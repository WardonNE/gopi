package compare

type Comparator[T any] interface {
	Compare(a, b T) int
}
