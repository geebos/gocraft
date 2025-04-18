package gvalue

func Ptr[T any](v T) *T {
	return &v
}

func Of[T any](v *T) T {
	return *v
}
