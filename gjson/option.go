package gjson

import (
	"bytes"
	"encoding/json"

	"github.com/geebos/gocraft/gvalue"
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

// EncodeOption defines the type for encoding configuration functions
type EncodeOption func(opt _option) _option

// DecodeOption defines the type for decoding configuration functions
type DecodeOption func(opt _option) _option

// WithUseNumber enables strict number parsing mode
// When enabled, decoder will use json.Number type instead of float64
// This preserves numeric precision but requires explicit type conversion
// Example usage:
//
//	Unmarshal(data, WithUseNumber())
func WithUseNumber() DecodeOption {
	return func(opt _option) _option {
		opt.UseNumber = gvalue.Ptr(true)
		return opt
	}
}

// WithDisableUnknownFields enables strict field validation
// When enabled, decoder returns error if JSON contains unknown fields
// Prevents silent ignoring of typos in field names
// Recommended for API request parsing
func WithDisableUnknownFields() DecodeOption {
	return func(opt _option) _option {
		opt.DisableUnknownFields = gvalue.Ptr(true)
		return opt
	}
}

// WithEscapeHtml controls HTML escaping in JSON encoding
// @param escape - true to enable HTML escaping (default), false to disable
func WithEscapeHtml(escape bool) EncodeOption {
	return func(opt _option) _option {
		opt.EscapeHtml = gvalue.Ptr(escape)
		return opt
	}
}

// WithIndent configures indentation for JSON output
// @param prefix - prefix for each indented line
// @param indent - indentation string (usually spaces)
func WithIndent(prefix, indent string) EncodeOption {
	return func(opt _option) _option {
		opt.IndentPrefix = gvalue.Ptr(prefix)
		opt.Indent = gvalue.Ptr(indent)
		return opt
	}
}

// Decode parses JSON data into the target interface with configured options
// Applies decoding settings:
// - DisableUnknownFields: rejects JSON keys not matching struct fields
// - UseNumber: preserves number precision with json.Number type
// @param data - raw JSON bytes to decode
// @param ins - pointer to target decoding structure
// @return error if decoding fails
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

// Encode serializes value to JSON with configured formatting options
// Applies encoding settings:
// - EscapeHtml: controls HTML-sensitive characters escaping
// - Indent: configures output indentation format
// @param v - value to be serialized
// @return formatted JSON bytes
// @return error if encoding fails
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
