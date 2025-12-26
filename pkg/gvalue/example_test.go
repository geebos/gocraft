package gvalue_test

import (
	"fmt"

	"github.com/geebos/gocraft/pkg/gvalue"
)

func ExampleZero() {
	// Get zero value of int
	intZero := gvalue.Zero[int]()
	fmt.Println(intZero)

	// Get zero value of string
	strZero := gvalue.Zero[string]()
	fmt.Printf("%q\n", strZero)

	// Get zero value of pointer
	ptrZero := gvalue.Zero[*int]()
	fmt.Println(ptrZero)

	// Output:
	// 0
	// ""
	// <nil>
}

func ExamplePtr() {
	// Create a pointer to a string literal
	name := gvalue.Ptr("John")
	fmt.Println(*name)

	// Useful for struct initialization with optional fields
	type Config struct {
		Timeout *int
	}
	config := Config{Timeout: gvalue.Ptr(30)}
	fmt.Println(*config.Timeout)

	// Output:
	// John
	// 30
}

func ExampleOf() {
	// Dereference a pointer
	value := 42
	ptr := &value
	result := gvalue.Of(ptr)
	fmt.Println(result)

	// Output:
	// 42
}

func ExampleIfElse() {
	// Basic ternary operation
	a, b := 10, 20
	max := gvalue.IfElse(a > b, a, b)
	fmt.Println(max)

	// String selection
	enabled := true
	status := gvalue.IfElse(enabled, "on", "off")
	fmt.Println(status)

	// Output:
	// 20
	// on
}

func ExampleEqual() {
	// Compare integers
	fmt.Println(gvalue.Equal(1, 1))
	fmt.Println(gvalue.Equal(1, 2))

	// Compare strings
	fmt.Println(gvalue.Equal("hello", "hello"))
	fmt.Println(gvalue.Equal("hello", "world"))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleLess() {
	// Compare integers
	fmt.Println(gvalue.Less(1, 2))
	fmt.Println(gvalue.Less(2, 1))

	// Compare floats
	fmt.Println(gvalue.Less(1.5, 2.5))

	// Compare strings (lexicographic order)
	fmt.Println(gvalue.Less("apple", "banana"))

	// Output:
	// true
	// false
	// true
	// true
}
