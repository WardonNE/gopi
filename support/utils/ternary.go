package utils

// If returns a when condition is true, returns b when condition is b
func If[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

// IfNull returns defaultValue when value is nil
func IfNull[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}
