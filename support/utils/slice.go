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

// Diff return the elements in a that are not in b
func Diff[T constraints.Ordered](a []T, b []T) []T {
	diff := []T{}
	bSet := map[T]struct{}{}
	for _, item := range b {
		bSet[item] = struct{}{}
	}
	for _, x := range a {
		if _, ok := bSet[x]; !ok {
			diff = append(diff, x)
		}
	}
	return diff
}
