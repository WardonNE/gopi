package compare

import (
	"github.com/wardonne/gopi/support/utils"
	"golang.org/x/exp/constraints"
)

type NatureComparator[T constraints.Ordered] struct {
	Desc bool
}

func NewNatureComparator[T constraints.Ordered](desc bool) *NatureComparator[T] {
	return &NatureComparator[T]{desc}
}

func (c *NatureComparator[T]) Compare(a, b T) int {
	if a < b {
		return utils.If(c.Desc, 1, -1)
	} else if a > b {
		return utils.If(c.Desc, -1, 1)
	} else {
		return 0
	}
}
