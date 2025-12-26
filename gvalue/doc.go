// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gvalue provides generic utility functions for value manipulation.
//
// The package includes type constraints for numeric types, pointer operations,
// comparison functions, and common utility functions that leverage Go generics.
//
// # Type Constraints
//
// The package provides several type constraints for use with generic functions:
//
//   - [Signed]: all signed integer types
//   - [Unsigned]: all unsigned integer types
//   - [Integer]: all integer types (signed and unsigned)
//   - [Float]: all floating-point types
//   - [Complex]: all complex number types
//   - [Ordered]: all types that support ordering operators
//
// # Utility Functions
//
// Common utility functions include:
//
//   - [Zero]: returns the zero value of any type
//   - [Ptr]: creates a pointer to a value
//   - [Of]: dereferences a pointer
//   - [IfElse]: ternary operator replacement
//   - [Equal]: equality comparison
//   - [Less]: less-than comparison
package gvalue
