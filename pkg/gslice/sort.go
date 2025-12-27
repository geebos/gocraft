package gslice

import "sort"

// Sort sorts a copy of the slice and returns the sorted copy.
//
// Sort uses the provided less function to determine the order.
// The original slice is not modified. If the slice is nil, Sort returns nil.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	sorted := Sort(numbers, gvalue.Less[int])
//	// sorted is []int{1, 1, 3, 4, 5}
//	// numbers is still []int{3, 1, 4, 1, 5}
func Sort[T any](s []T, less func(T, T) bool) []T {
	if s == nil {
		return nil
	}
	result := make([]T, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// StealSort sorts the slice in place and returns it.
//
// StealSort uses the provided less function to determine the order.
// The original slice is modified. If the slice is nil, StealSort returns nil.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	sorted := StealSort(numbers, gvalue.Less[int])
//	// sorted is []int{1, 1, 3, 4, 5}
//	// numbers is also []int{1, 1, 3, 4, 5} (same underlying array)
func StealSort[T any](s []T, less func(T, T) bool) []T {
	if s == nil {
		return nil
	}
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}
