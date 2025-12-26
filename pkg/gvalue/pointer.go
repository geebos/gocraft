package gvalue

// Ptr returns a pointer to a copy of the given value v.
//
// This function is useful for creating pointers to literal values
// or expressions where taking the address directly is not allowed:
//
//	// Instead of:
//	s := "hello"
//	config := &Config{Name: &s}
//
//	// You can write:
//	config := &Config{Name: Ptr("hello")}
//
// Note that Ptr creates a new variable internally, so the returned
// pointer points to a copy of v, not to the original value.
func Ptr[T any](v T) *T {
	return &v
}

// Of dereferences a pointer and returns its underlying value.
//
// Of panics if v is nil. Use this function only when you are certain
// that the pointer is not nil.
//
// For safe dereferencing with a default value, consider using IfElse:
//
//	value := IfElse(ptr != nil, *ptr, defaultValue)
func Of[T any](v *T) T {
	return *v
}
