package gvalue

func IfElse[T any](condition bool, t, f T) T {
	if condition {
		return t
	}
	return f
}

func Equal[T comparable](l, r T) bool {
	return l == r
}

func Less[T Orderable](l, r T) bool {
	return l < r
}
