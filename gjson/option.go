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

type EncodeOption func(opt _option) _option
type DecodeOption func(opt _option) _option

func WithUseNumber() DecodeOption {
	return func(opt _option) _option {
		opt.UseNumber = gvalue.Ptr(true)
		return opt
	}
}

func WithDisableUnknownFields() DecodeOption {
	return func(opt _option) _option {
		opt.DisableUnknownFields = gvalue.Ptr(true)
		return opt
	}
}

func WithEscapeHtml(escape bool) EncodeOption {
	return func(opt _option) _option {
		opt.EscapeHtml = gvalue.Ptr(escape)
		return opt
	}
}

func WithIndent(prefix, indent string) EncodeOption {
	return func(opt _option) _option {
		opt.IndentPrefix = gvalue.Ptr(prefix)
		opt.Indent = gvalue.Ptr(indent)
		return opt
	}
}

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
