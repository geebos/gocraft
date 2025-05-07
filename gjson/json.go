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

func Unmarshal[T any, D ~[]byte | ~string](d D) (T, error) {
	var data = []byte(d)

	result := gvalue.Zero[T]()
	return result, json.Unmarshal(data, &result)
}

func Marshal[R ~[]byte | ~string](v any) (R, error) {
	data, err := json.Marshal(v)
	return R(data), err
}

func MarshalIndent[R ~[]byte | ~string](v any, prefix, indent string) (R, error) {
	data, err := json.MarshalIndent(v, prefix, indent)
	return R(data), err
}

func Cast[T any](from any) (T, error) {
	return Unmarshal[T](Dumps(from))
}

func Dumps[T any](v T) string {
	data, _ := Marshal[string](v)
	return data
}

// UnmarshalFromPath unmarshal result from data by path.
func UnmarshalFromPath[T any, D ~[]byte | ~string](data D, path string) (T, error) {
	result := gjson.GetBytes([]byte(data), path)
	if !result.Exists() {
		return gvalue.Zero[T](), fmt.Errorf("`%s` %w", path, ErrPathNotFound)
	}
	return Unmarshal[T](result.Raw)
}

// UnmarshalFromPathWithDefault unmarshal result from data by path.
func UnmarshalFromPathWithDefault[T any, D ~[]byte | ~string](data D, path string, val T) T {
	result, err := UnmarshalFromPath[T](data, path)
	return gvalue.IfElse(err == nil, result, val)
}
