package utils

func If[T any](condition bool, a T, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}

func IfNull[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}
