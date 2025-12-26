package gjson

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/geebos/gocraft/pkg/gvalue"
)

// ErrPathNotFound is returned when a JSON path expression does not match any value.
var ErrPathNotFound = fmt.Errorf("path not found")

// Unmarshal parses JSON-encoded data and returns a value of type T.
//
// The data parameter can be either []byte or string (or any type with
// an underlying type of []byte or string).
//
// Without options, Unmarshal uses the standard encoding/json decoder.
// With options, it creates a customized decoder based on the provided
// [DecodeOption] functions.
//
// Example:
//
//	type User struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//
//	user, err := Unmarshal[User](`{"name":"John","age":30}`)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(user.Name) // John
func Unmarshal[T any, D ~[]byte | ~string](d D, opts ...DecodeOption) (T, error) {
	var data = []byte(d)
	result := gvalue.Zero[T]()
	if len(opts) > 0 {
		return result, unmarshalWithOptions(data, &result, opts)
	}
	return result, json.Unmarshal(data, &result)
}

func unmarshalWithOptions(data []byte, ins any, opts []DecodeOption) error {
	var opt _option
	for _, fn := range opts {
		opt = fn(opt)
	}
	return opt.Decode(data, ins)
}

// Marshal returns the JSON encoding of v.
//
// The type parameter R specifies the return type, which can be either
// []byte or string (or any type with an underlying type of []byte or string).
//
// Without options, Marshal uses the standard encoding/json encoder.
// With options, it creates a customized encoder based on the provided
// [EncodeOption] functions.
//
// Example:
//
//	user := User{Name: "John", Age: 30}
//
//	// Get JSON as string
//	jsonStr, err := Marshal[string](user)
//
//	// Get JSON as bytes
//	jsonBytes, err := Marshal[[]byte](user)
func Marshal[R ~[]byte | ~string](v any, opts ...EncodeOption) (R, error) {
	var err error
	var data []byte
	if len(opts) == 0 {
		data, err = json.Marshal(v)
	} else {
		data, err = marshalWithOptions(v, opts)
	}
	return R(data), err
}

func marshalWithOptions(v any, opts []EncodeOption) ([]byte, error) {
	var opt _option
	for _, fn := range opts {
		opt = fn(opt)
	}
	return opt.Encode(v)
}

// MarshalIndent is like [Marshal] but applies indentation to format the output.
//
// Each JSON element begins on a new line that starts with prefix followed by
// one or more copies of indent according to the nesting depth.
//
// Example:
//
//	data := map[string]any{"name": "John", "age": 30}
//	json, err := MarshalIndent[string](data, "", "  ")
//	// Output:
//	// {
//	//   "age": 30,
//	//   "name": "John"
//	// }
func MarshalIndent[R ~[]byte | ~string](v any, prefix, indent string) (R, error) {
	data, err := json.MarshalIndent(v, prefix, indent)
	return R(data), err
}

// Cast converts a value to type T by marshaling to JSON and unmarshaling back.
//
// This is useful for converting between compatible struct types or for
// converting maps to structs and vice versa. Note that this involves
// JSON serialization overhead and may lose type information for some types.
//
// Example:
//
//	type UserInput struct {
//	    Name string `json:"name"`
//	}
//	type UserOutput struct {
//	    Name string `json:"name"`
//	}
//
//	input := UserInput{Name: "John"}
//	output, err := Cast[UserOutput](input)
func Cast[T any](from any) (T, error) {
	return Unmarshal[T](Dumps(from))
}

// Dumps returns the JSON string representation of v.
//
// Unlike [Marshal], Dumps ignores any encoding errors and returns
// an empty string if encoding fails. Use this for debugging or logging
// where error handling is not critical.
//
// Example:
//
//	user := User{Name: "John", Age: 30}
//	fmt.Println(Dumps(user)) // {"name":"John","age":30}
func Dumps[T any](v T) string {
	data, _ := Marshal[string](v)
	return data
}

// UnmarshalFromPath extracts a value from JSON data using a path expression
// and unmarshals it to type T.
//
// Path expressions follow the gjson syntax. Common patterns include:
//   - "name" - get a top-level field
//   - "user.name" - get a nested field
//   - "users.0" - get first array element
//   - "users.#" - get array length
//   - "users.#.name" - get all names from array
//
// Returns [ErrPathNotFound] if the path does not match any value.
//
// For the complete path syntax, see https://github.com/tidwall/gjson#path-syntax
//
// Example:
//
//	data := `{"user": {"name": "John", "emails": ["a@b.com", "c@d.com"]}}`
//
//	name, err := UnmarshalFromPath[string](data, "user.name")
//	// name = "John"
//
//	emails, err := UnmarshalFromPath[[]string](data, "user.emails")
//	// emails = ["a@b.com", "c@d.com"]
func UnmarshalFromPath[T any, D ~[]byte | ~string](data D, path string) (T, error) {
	result := gjson.GetBytes([]byte(data), path)
	if !result.Exists() {
		return gvalue.Zero[T](), fmt.Errorf("`%s` %w", path, ErrPathNotFound)
	}
	return Unmarshal[T](result.Raw)
}

// UnmarshalFromPathWithDefault extracts a value from JSON using a path expression,
// returning the provided default value if the path is not found or parsing fails.
//
// This is a convenience wrapper around [UnmarshalFromPath] that handles errors
// by returning a default value instead.
//
// Example:
//
//	data := `{"config": {"timeout": 30}}`
//
//	// Returns 30
//	timeout := UnmarshalFromPathWithDefault[int](data, "config.timeout", 60)
//
//	// Returns 60 (default) because path doesn't exist
//	retries := UnmarshalFromPathWithDefault[int](data, "config.retries", 60)
func UnmarshalFromPathWithDefault[T any, D ~[]byte | ~string](data D, path string, val T) T {
	result, err := UnmarshalFromPath[T](data, path)
	return gvalue.IfElse(err == nil, result, val)
}
