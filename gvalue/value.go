package gvalue

func Zero[T any]() T {
	var v T
	return v
}
