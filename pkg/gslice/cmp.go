package gslice

// CmpWith creates a comparison function that compares elements against a fixed value.
//
// CmpWith takes a comparison function cmp and a fixed value b, and returns
// a new function that can be used with operations like Filter or Find.
// The returned function takes a single value a and returns the result of cmp(a, b).
//
// This is useful for creating predicates from comparison functions, especially
// when used with functions from the gvalue package like LTE, GT, etc.
//
// Example:
//
//	import (
//		"github.com/yourorg/gocraft/pkg/gvalue"
//		"github.com/yourorg/gocraft/pkg/gslice"
//	)
//
//	numbers := []int{1, 2, 3, 4, 5}
//	// Find all numbers greater than 3
//	greaterThan3 := Filter(numbers, CmpWith(gvalue.GT[int], 3))
//	// greaterThan3 is []int{4, 5}
//
//	// Find the first number less than or equal to 2
//	value, found := Find(numbers, CmpWith(gvalue.LTE[int], 2))
//	// value is 1, found is true
func CmpWith[T any](cmp func(T, T) bool, b T) func(T) bool {
	return func(a T) bool {
		return cmp(a, b)
	}
}
