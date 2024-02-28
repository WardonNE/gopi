package support

// Clonable clonable interface
type Clonable[V any] interface {
	Clone() V
}
