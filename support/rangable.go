package support

// Rangable rangable interface
type Rangable[E any] interface {
	Range(callback func(value E) bool)
}

// ReverseRangable reverse rangable interface
type ReverseRangable[E any] interface {
	ReverseRange(callback func(value E) bool)
}
