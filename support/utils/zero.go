package utils

import "reflect"

// Zero returns a zero value of specific type
func Zero[E any]() E {
	var zero E
	return zero
}

// IsZero returns whether the specific value is zero value of its type using "==" operation
func IsZero[E comparable](value E) bool {
	return value == Zero[E]()
}

// IsReflectZero returns whether the specific value is zero value of its type using [reflect.IsZero]
func IsReflectZero[E any](value E) bool {
	return reflect.ValueOf(value).IsZero()
}
