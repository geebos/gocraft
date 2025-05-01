package gjson

import (
	"encoding/json"

	"github.com/geebos/gocraft/gvalue"
)

func Unmarshal[T any, D []byte | string](d D) (T, error) {
	var data []byte
	switch v := any(d).(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}

	result := gvalue.Zero[T]()
	return result, json.Unmarshal(data, &result)
}

func Marshal[R []byte | string](v any) (R, error) {
	data, err := json.Marshal(v)
	return R(data), err
}

func MarshalIndent[R []byte | string](v any, prefix, indent string) (R, error) {
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
