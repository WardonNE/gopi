package serializer

// MapSerializer map serializer
type MapSerializer[K comparable, V any] interface {
	ToMap() map[K]V
	FromMap(values map[K]V)
}
