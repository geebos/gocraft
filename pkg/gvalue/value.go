package gvalue

// Zero returns the zero value of any type T.
//
// This is equivalent to declaring a variable with var and returning it:
//
//	var v T
//	return v
//
// Zero is useful when you need the zero value of a type parameter
// in a generic function, especially for returning default values.
func Zero[T any]() T {
	var v T
	return v
}
