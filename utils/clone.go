package utils

import "github.com/wardonne/gopi/support"

func Clone[V any](src support.Clonable[V]) V {
	return src.Clone()
}
