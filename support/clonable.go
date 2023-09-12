package support

type Clonable[V any] interface {
	Clone() V
}
