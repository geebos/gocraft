package gjson

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/geebos/gocraft/gvalue"
)

var (
	ErrPathNotFound = fmt.Errorf("path not found")
)

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by T.
// @param d - input data in []byte or string format
// @param opts - optional decode configuration
// @return parsed value of type T
// @return error if parsing fails
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
// @param v - value to be encoded
// @param opts - optional encode configuration
// @return JSON-encoded data in specified format R
// @return error if encoding fails
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

// MarshalIndent is like Marshal but applies Indent to format the output.
// @param v - value to be encoded
// @param prefix - prefix for each indented line
// @param indent - indentation string
// @return JSON-encoded data in specified format R
// @return error if encoding fails
func MarshalIndent[R ~[]byte | ~string](v any, prefix, indent string) (R, error) {
	data, err := json.MarshalIndent(v, prefix, indent)
	return R(data), err
}

// Cast converts between compatible types through JSON serialization.
// @param from - source value to be converted
// @return converted value of type T
// @return error if conversion fails
func Cast[T any](from any) (T, error) {
	return Unmarshal[T](Dumps(from))
}

// Dumps returns the JSON string representation of v.
// @param v - value to be serialized
// @return JSON string
func Dumps[T any](v T) string {
	data, _ := Marshal[string](v)
	return data
}

// UnmarshalFromPath extracts JSON value by path and unmarshals it to type T.
// @param data - source JSON data
// @param path - JSON path expression, refer to gjson documentation for supported syntax.
// @return parsed value of type T
// @return error if path not found or parsing fails
func UnmarshalFromPath[T any, D ~[]byte | ~string](data D, path string) (T, error) {
	result := gjson.GetBytes([]byte(data), path)
	if !result.Exists() {
		return gvalue.Zero[T](), fmt.Errorf("`%s` %w", path, ErrPathNotFound)
	}
	return Unmarshal[T](result.Raw)
}

// UnmarshalFromPathWithDefault provides fallback value when path not found or parse fail.
// @param data - source JSON data
// @param path - JSON path expression, refer to gjson documentation for supported syntax.
// @param val - default value to return
// @return parsed value of type T or default value if path not found or parsing fails
func UnmarshalFromPathWithDefault[T any, D ~[]byte | ~string](data D, path string, val T) T {
	result, err := UnmarshalFromPath[T](data, path)
	return gvalue.IfElse(err == nil, result, val)
}
