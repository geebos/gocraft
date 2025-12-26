// Package gjson provides enhanced JSON encoding/decoding utilities with generics support.
//
// The package wraps the standard encoding/json package and tidwall/gjson library,
// providing a more ergonomic API using Go generics introduced in Go 1.18.
//
// # Basic Usage
//
// For simple JSON operations:
//
//	// Unmarshal JSON to a typed value
//	user, err := gjson.Unmarshal[User](jsonData)
//
//	// Marshal a value to JSON
//	jsonStr, err := gjson.Marshal[string](user)
//
//	// Quick dump to JSON string (ignores errors)
//	str := gjson.Dumps(user)
//
// # Path-based Extraction
//
// Extract values from JSON using path expressions (powered by tidwall/gjson):
//
//	data := `{"user": {"name": "John", "age": 30}}`
//
//	// Extract nested value
//	name, err := gjson.UnmarshalFromPath[string](data, "user.name")
//
//	// Extract with default value
//	age := gjson.UnmarshalFromPathWithDefault[int](data, "user.age", 0)
//
// # Encoding/Decoding Options
//
// Customize JSON handling with options:
//
//	// Decode with strict number handling
//	data, err := gjson.Unmarshal[map[string]any](json, gjson.WithUseNumber())
//
//	// Encode with custom formatting
//	json, err := gjson.Marshal[string](v, gjson.WithIndent("", "  "))
//
// For path expression syntax, refer to https://github.com/tidwall/gjson#path-syntax
package gjson

