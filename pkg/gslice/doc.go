// Package gslice provides generic utility functions for slice and array operations.
//
// The package includes common functional programming operations on slices and arrays,
// such as map, filter, reduce, find, sort, and concatenation.
//
// All operations are implemented as functions rather than methods, making them
// easy to use with any slice or array type.
//
// # Operations
//
// The package provides the following operations:
//
//   - [Map]: transforms each element of a slice
//   - [Filter]: filters elements based on a predicate
//   - [Reduce]: reduces a slice to a single value
//   - [Find]: finds the first element matching a predicate
//   - [Any]: checks if any element satisfies a predicate
//   - [All]: checks if all elements satisfy a predicate
//   - [Sort]: sorts a copy of the slice
//   - [StealSort]: sorts the slice in place
//   - [Concat]: concatenates multiple slices
//   - [Slice]: extracts a sub-slice
//   - [CmpWith]: creates a comparison function with a fixed value
//   - [Unique]: removes duplicate elements from a slice
//   - [Union]: merges multiple slices and removes duplicates
//   - [Intersection]: returns elements that exist in all input slices
//   - [Difference]: returns elements from the first slice not in other slices
package gslice
