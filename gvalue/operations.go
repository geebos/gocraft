package gvalue

func IfElse[T any](condition bool, t, f T) T {
	if condition {
		return t
	}
	return f
}
