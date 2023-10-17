package utils

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Max returns the max value in given values
func Max[T constraints.Ordered](values ...T) T {
	return slices.Max(values)
}

// Min returns the min value in given values
func Min[T constraints.Ordered](values ...T) T {
	return slices.Min(values)
}
