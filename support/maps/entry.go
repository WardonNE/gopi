package maps

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
