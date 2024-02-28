package serializer

// ArraySerializer array serializer
type ArraySerializer[E any] interface {
	ToArray() []E
	FromArray(values []E)
}
