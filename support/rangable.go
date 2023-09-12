package support

type Rangable[E any] interface {
	Range(callback func(value E) bool)
}

type ReverseRangable[E any] interface {
	ReverseRange(callback func(value E) bool)
}
