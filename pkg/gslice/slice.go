package gslice

// Map transforms each element of the input slice using the provided function.
//
// Map returns a new slice containing the results of applying fn to each element
// of the input slice. The input slice is not modified.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	doubled := Map(numbers, func(n int) int { return n * 2 })
//	// doubled is []int{2, 4, 6, 8}
func Map[T, R any](s []T, fn func(T) R) []R {
	if s == nil {
		return nil
	}
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
//
// Filter creates a new slice with all elements from the input slice for which
// the predicate function returns true. The input slice is not modified.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
//	// evens is []int{2, 4}
func Filter[T any](s []T, fn func(T) bool) []T {
	if s == nil {
		return nil
	}
	result := make([]T, 0, len(s))
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces the slice to a single value by applying a function cumulatively.
//
// Reduce applies the function fn to each element of the slice, along with an
// accumulator value. The accumulator is initialized with the initial value,
// and the function is called for each element in order.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
//	// sum is 10
func Reduce[T, R any](s []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, v := range s {
		result = fn(result, v)
	}
	return result
}

// Find returns the first element that satisfies the predicate, along with a boolean
// indicating whether such an element was found.
//
// If no element satisfies the predicate, Find returns the zero value of T and false.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	value, found := Find(numbers, func(n int) bool { return n > 3 })
//	// value is 4, found is true
func Find[T any](s []T, fn func(T) bool) (T, bool) {
	var zero T
	if s == nil {
		return zero, false
	}
	for _, v := range s {
		if fn(v) {
			return v, true
		}
	}
	return zero, false
}

// Any returns true if at least one element in the slice satisfies the predicate.
//
// Any returns false if the slice is nil or empty, or if no element satisfies the predicate.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := Any(numbers, func(n int) bool { return n%2 == 0 })
//	// hasEven is true
func Any[T any](s []T, fn func(T) bool) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate.
//
// All returns true if the slice is nil or empty (vacuous truth).
//
// Example:
//
//	numbers := []int{2, 4, 6, 8}
//	allEven := All(numbers, func(n int) bool { return n%2 == 0 })
//	// allEven is true
func All[T any](s []T, fn func(T) bool) bool {
	if s == nil {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Concat concatenates multiple slices into a single slice.
//
// Concat returns a new slice containing all elements from all input slices
// in the order they are provided. The input slices are not modified.
//
// Example:
//
//	s1 := []int{1, 2}
//	s2 := []int{3, 4}
//	s3 := []int{5}
//	result := Concat(s1, s2, s3)
//	// result is []int{1, 2, 3, 4, 5}
func Concat[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Slice returns a sub-slice of the input slice.
//
// Slice extracts elements from index start (inclusive) to index end (exclusive).
// It is similar to the built-in slice expression s[start:end], but returns
// a new slice rather than sharing the underlying array.
//
// If start or end are out of bounds, Slice panics (same behavior as built-in slicing).
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sub := Slice(numbers, 1, 4)
//	// sub is []int{2, 3, 4}
func Slice[T any](s []T, start, end int) []T {
	if s == nil {
		return nil
	}
	sub := s[start:end]
	result := make([]T, len(sub))
	copy(result, sub)
	return result
}
