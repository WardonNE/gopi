package set

import "github.com/wardonne/gopi/support/container"

type Set[E comparable] interface {
	container.Collection[E]
}
