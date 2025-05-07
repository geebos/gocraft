package gvalue

// Ptr returns a pointer to the given value
// Helpful for getting address of temporary values
// @param v - value to be wrapped in a pointer
// @return pointer to the value v
func Ptr[T any](v T) *T {
	return &v
}

// Of dereferences a pointer safely
// Panic when v is nil
// @param v - pointer to value
// @return value pointed to by v
func Of[T any](v *T) T {
	return *v
}
