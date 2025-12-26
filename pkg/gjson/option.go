package gjson

import (
	"bytes"
	"encoding/json"

	"github.com/geebos/gocraft/pkg/gvalue"
)

type _option struct {
	// decode options
	UseNumber            *bool
	DisableUnknownFields *bool
	// encode options
	EscapeHtml   *bool
	IndentPrefix *string
	Indent       *string
}

// EncodeOption is a function that configures JSON encoding behavior.
// Use the With* functions to create EncodeOption values.
type EncodeOption func(opt _option) _option

// DecodeOption is a function that configures JSON decoding behavior.
// Use the With* functions to create DecodeOption values.
type DecodeOption func(opt _option) _option

// WithUseNumber configures the decoder to use [json.Number] instead of float64
// for JSON numbers.
//
// This is useful when you need to preserve the exact numeric representation,
// especially for large integers that would lose precision when converted to float64.
//
// When enabled, numeric values decoded into interface{} will be of type
// json.Number, which can be converted to int64 or float64 as needed.
//
// Example:
//
//	data := `{"id": 9007199254740993}`
//	result, err := Unmarshal[map[string]any](data, WithUseNumber())
//	num := result["id"].(json.Number)
//	id, _ := num.Int64() // preserves precision
func WithUseNumber() DecodeOption {
	return func(opt _option) _option {
		opt.UseNumber = gvalue.Ptr(true)
		return opt
	}
}

// WithDisableUnknownFields configures the decoder to return an error when
// the JSON contains fields that do not have corresponding struct fields.
//
// This is useful for strict API validation where you want to detect typos
// or unexpected fields in incoming JSON data.
//
// Example:
//
//	type User struct {
//	    Name string `json:"name"`
//	}
//
//	data := `{"name": "John", "unknown_field": 123}`
//	_, err := Unmarshal[User](data, WithDisableUnknownFields())
//	// err: json: unknown field "unknown_field"
func WithDisableUnknownFields() DecodeOption {
	return func(opt _option) _option {
		opt.DisableUnknownFields = gvalue.Ptr(true)
		return opt
	}
}

// WithEscapeHtml configures whether the encoder should escape
// HTML-sensitive characters (<, >, &) in JSON strings.
//
// By default, the standard library escapes these characters for safe
// embedding in HTML. Set escape to false to disable this behavior
// when the JSON is not intended for HTML contexts.
//
// Example:
//
//	data := map[string]string{"html": "<div>test</div>"}
//
//	// Default behavior: escapes HTML
//	json1, _ := Marshal[string](data)
//	// {"html":"\u003cdiv\u003etest\u003c/div\u003e"}
//
//	// Disabled: preserves original characters
//	json2, _ := Marshal[string](data, WithEscapeHtml(false))
//	// {"html":"<div>test</div>"}
func WithEscapeHtml(escape bool) EncodeOption {
	return func(opt _option) _option {
		opt.EscapeHtml = gvalue.Ptr(escape)
		return opt
	}
}

// WithIndent configures the encoder to output indented JSON.
//
// Each JSON element begins on a new line that starts with prefix followed
// by one or more copies of indent according to the nesting depth.
//
// Example:
//
//	data := map[string]any{"name": "John", "age": 30}
//	json, _ := Marshal[string](data, WithIndent("", "  "))
//	// {
//	//   "age": 30,
//	//   "name": "John"
//	// }
func WithIndent(prefix, indent string) EncodeOption {
	return func(opt _option) _option {
		opt.IndentPrefix = gvalue.Ptr(prefix)
		opt.Indent = gvalue.Ptr(indent)
		return opt
	}
}

// Decode parses JSON data into the target value with the configured options.
//
// The method applies any configured decoding options:
//   - DisableUnknownFields: returns error for JSON keys not matching struct fields
//   - UseNumber: preserves number precision with json.Number type
func (opt _option) Decode(data []byte, ins any) error {
	buf := bytes.NewReader(data)
	decoder := json.NewDecoder(buf)
	if opt.DisableUnknownFields != nil && *opt.DisableUnknownFields {
		decoder.DisallowUnknownFields()
	}
	if opt.UseNumber != nil && *opt.UseNumber {
		decoder.UseNumber()
	}
	return decoder.Decode(ins)
}

// Encode serializes a value to JSON with the configured formatting options.
//
// The method applies any configured encoding options:
//   - EscapeHtml: controls HTML-sensitive character escaping
//   - Indent: configures output indentation format
func (opt _option) Encode(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if opt.EscapeHtml != nil {
		encoder.SetEscapeHTML(*opt.EscapeHtml)
	}
	if opt.IndentPrefix != nil {
		encoder.SetIndent(*opt.IndentPrefix, *opt.Indent)
	}
	err := encoder.Encode(v)
	return buf.Bytes(), err
}
