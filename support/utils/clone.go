package utils

import "github.com/wardonne/gopi/support"

// Clone clones a new instance from src
func Clone[V any](src support.Clonable[V]) V {
	return src.Clone()
}
