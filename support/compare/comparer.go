package compare

type Comparer[T any] interface {
	Compare(a, b T) int
}
