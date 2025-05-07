package gvalue

// Zero returns the zero value of type T
// Equivalent to var v T; return v
func Zero[T any]() T {
	var v T
	return v
}
