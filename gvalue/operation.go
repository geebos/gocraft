package gvalue

// IfElse returns t if condition is true, otherwise returns f
// Acts as a generic ternary operator replacement
// @param condition - boolean evaluation condition
// @param t - value returned when condition is true
// @param f - value returned when condition is false
func IfElse[T any](condition bool, t, f T) T {
	if condition {
		return t
	}
	return f
}

// Equal checks strict equality of two comparable values
// Uses == operator under the hood
// @param l - left operand
// @param r - right operand
// @return true only if l and r are exactly equal
func Equal[T comparable](l, r T) bool {
	return l == r
}

// Less compares two orderable values (numeric types)
// @param l - left operand
// @param r - right operand
// @return true if l is less than r
func Less[T Ordered](l, r T) bool {
	return l < r
}
