package gjson_test

import (
	"encoding/json"
	"fmt"

	"github.com/geebos/gocraft/pkg/gjson"
)

func ExampleUnmarshal() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := `{"name":"John","age":30}`
	user, err := gjson.Unmarshal[User](data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(user.Name, user.Age)

	// Output:
	// John 30
}

func ExampleUnmarshal_withOptions() {
	// Using WithUseNumber to preserve numeric precision
	data := `{"id": 9007199254740993}`
	result, err := gjson.Unmarshal[map[string]any](data, gjson.WithUseNumber())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	num := result["id"].(json.Number)
	id, _ := num.Int64()
	fmt.Println(id)

	// Output:
	// 9007199254740993
}

func ExampleMarshal() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "John", Age: 30}

	// Get JSON as string
	jsonStr, err := gjson.Marshal[string](user)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(jsonStr)

	// Output:
	// {"name":"John","age":30}
}

func ExampleMarshal_withOptions() {
	data := map[string]string{"html": "<div>test</div>"}

	// Disable HTML escaping
	jsonStr, err := gjson.Marshal[string](data, gjson.WithEscapeHtml(false))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// Note: encoder adds trailing newline
	fmt.Print(jsonStr)

	// Output:
	// {"html":"<div>test</div>"}
}

func ExampleMarshalIndent() {
	data := map[string]any{"name": "John", "age": 30}

	jsonStr, err := gjson.MarshalIndent[string](data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(jsonStr)

	// Output:
	// {
	//   "age": 30,
	//   "name": "John"
	// }
}

func ExampleDumps() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "John", Age: 30}
	fmt.Println(gjson.Dumps(user))

	// Output:
	// {"name":"John","age":30}
}

func ExampleCast() {
	type UserInput struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type UserOutput struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	input := UserInput{Name: "John", Age: 30}
	output, err := gjson.Cast[UserOutput](input)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(output.Name, output.Age)

	// Output:
	// John 30
}

func ExampleUnmarshalFromPath() {
	data := `{
		"user": {
			"name": "John",
			"emails": ["john@example.com", "j@test.com"]
		}
	}`

	// Extract nested string
	name, err := gjson.UnmarshalFromPath[string](data, "user.name")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(name)

	// Extract array
	emails, err := gjson.UnmarshalFromPath[[]string](data, "user.emails")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(emails)

	// Output:
	// John
	// [john@example.com j@test.com]
}

func ExampleUnmarshalFromPathWithDefault() {
	data := `{"config": {"timeout": 30}}`

	// Existing path - returns actual value
	timeout := gjson.UnmarshalFromPathWithDefault[int](data, "config.timeout", 60)
	fmt.Println(timeout)

	// Non-existing path - returns default
	retries := gjson.UnmarshalFromPathWithDefault[int](data, "config.retries", 3)
	fmt.Println(retries)

	// Output:
	// 30
	// 3
}

func ExampleWithUseNumber() {
	// Large integer that would lose precision as float64
	data := `{"big_number": 9007199254740993}`

	result, _ := gjson.Unmarshal[map[string]any](data, gjson.WithUseNumber())
	num := result["big_number"].(json.Number)
	fmt.Println(num.String())

	// Output:
	// 9007199254740993
}

func ExampleWithDisableUnknownFields() {
	type User struct {
		Name string `json:"name"`
	}

	// JSON with unknown field
	data := `{"name": "John", "unknown": 123}`

	_, err := gjson.Unmarshal[User](data, gjson.WithDisableUnknownFields())
	if err != nil {
		fmt.Println("Error detected: unknown field")
	}

	// Output:
	// Error detected: unknown field
}

func ExampleWithEscapeHtml() {
	data := map[string]string{"content": "1 < 2 & 3 > 2"}

	// With HTML escaping disabled
	jsonStr, _ := gjson.Marshal[string](data, gjson.WithEscapeHtml(false))
	fmt.Print(jsonStr)

	// Output:
	// {"content":"1 < 2 & 3 > 2"}
}

func ExampleWithIndent() {
	data := map[string]int{"a": 1, "b": 2}

	jsonStr, _ := gjson.Marshal[string](data, gjson.WithIndent("", "    "))
	fmt.Print(jsonStr)

	// Output:
	// {
	//     "a": 1,
	//     "b": 2
	// }
}
