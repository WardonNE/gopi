package utils

import "golang.org/x/exp/constraints"

func Max[T constraints.Float | constraints.Integer](a, b T) T {
	return If[T](a >= b, a, b)
}

func Min[T constraints.Float | constraints.Integer](a, b T) T {
	return If[T](a < b, a, b)
}
