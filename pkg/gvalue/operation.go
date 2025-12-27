package gvalue

// IfElse returns t if condition is true, otherwise returns f.
//
// IfElse provides a generic ternary operator replacement for Go,
// which does not have a built-in ternary operator like ? : in other languages.
//
// Note that unlike a true ternary operator, both t and f are always evaluated
// before the function is called. If you need lazy evaluation, use if-else instead.
//
// Example:
//
//	max := IfElse(a > b, a, b)
//	status := IfElse(enabled, "on", "off")
func IfElse[T any](condition bool, t, f T) T {
	if condition {
		return t
	}
	return f
}

// Equal reports whether l and r are equal.
//
// Equal uses the == operator for comparison, so it only works with
// comparable types. For deep equality of complex types like slices
// and maps, use [reflect.DeepEqual] instead.
func Equal[T comparable](l, r T) bool {
	return l == r
}

// Less reports whether l is less than r.
//
// Less works with any ordered type, including integers, floats, and strings.
// For strings, the comparison is lexicographic based on Unicode code points.
func Less[T Ordered](l, r T) bool {
	return l < r
}

// LT reports whether l is less than r.
//
// LT is an alias for Less, provided for consistency with other comparison functions.
func LT[T Ordered](l, r T) bool {
	return l < r
}

// LTE reports whether l is less than or equal to r.
//
// LTE works with any ordered type, including integers, floats, and strings.
func LTE[T Ordered](l, r T) bool {
	return l <= r
}

// GT reports whether l is greater than r.
//
// GT works with any ordered type, including integers, floats, and strings.
func GT[T Ordered](l, r T) bool {
	return l > r
}

// GTE reports whether l is greater than or equal to r.
//
// GTE works with any ordered type, including integers, floats, and strings.
func GTE[T Ordered](l, r T) bool {
	return l >= r
}

// EQ reports whether l is equal to r.
//
// EQ is an alias for Equal, provided for consistency with other comparison functions.
// EQ uses the == operator for comparison, so it only works with comparable types.
func EQ[T comparable](l, r T) bool {
	return l == r
}
