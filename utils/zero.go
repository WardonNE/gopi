package utils

import "reflect"

func Zero[E any]() E {
	var zero E
	return zero
}

func IsZero[E comparable](value E) bool {
	return value == Zero[E]()
}

func IsReflectZero[E any](value E) bool {
	return reflect.ValueOf(value).IsZero()
}
