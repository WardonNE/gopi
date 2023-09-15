package iterator

type Iterator[E any] interface {
	HasNext() bool
	Next() E
}
