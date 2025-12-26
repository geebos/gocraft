# gocraft

[![Go](https://github.com/geebos/gocraft/actions/workflows/go.yml/badge.svg)](https://github.com/geebos/gocraft/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/geebos/gocraft.svg)](https://pkg.go.dev/github.com/geebos/gocraft)
[![Go Report Card](https://goreportcard.com/badge/github.com/geebos/gocraft)](https://goreportcard.com/report/github.com/geebos/gocraft)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Generic wrappers for popular Go libraries — write cleaner, more concise code.

📖 **[Documentation](https://geebos.github.io/gocraft/)** | 📦 **[pkg.go.dev](https://pkg.go.dev/github.com/geebos/gocraft)**

## Goals

- **Generic Wrappers**: Provide type-safe generic wrappers for common open-source libraries
- **Cleaner Code**: Reduce boilerplate and make your code more readable
- **Zero Cost**: Minimal overhead with compile-time type safety

## Quick Start

### Installation

```bash
go get github.com/geebos/gocraft
```

### gjson - Enhanced JSON Operations

```go
import "github.com/geebos/gocraft/gjson"

// Unmarshal JSON to a typed value
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
user, err := gjson.Unmarshal[User](`{"name":"John","age":30}`)

// Extract nested values using path expressions
data := `{"user": {"name": "John", "emails": ["a@b.com"]}}`
name, err := gjson.UnmarshalFromPath[string](data, "user.name")

// Quick dump to JSON string
jsonStr := gjson.Dumps(user)
```

### gvalue - Generic Value Utilities

```go
import "github.com/geebos/gocraft/gvalue"

// Create pointers to literal values
config := &Config{
    Timeout: gvalue.Ptr(30),
    Name:    gvalue.Ptr("default"),
}

// Ternary operator replacement
max := gvalue.IfElse(a > b, a, b)

// Get zero value of any type
zero := gvalue.Zero[int]() // 0
```

## Packages

| Package | Description |
|---------|-------------|
| [gjson](https://pkg.go.dev/github.com/geebos/gocraft/gjson) | Generic JSON encoding/decoding with path extraction support |
| [gvalue](https://pkg.go.dev/github.com/geebos/gocraft/gvalue) | Generic value utilities, type constraints, and helper functions |

## Requirements

- Go 1.18+ (generics support required)

## License

[MIT License](LICENSE)
