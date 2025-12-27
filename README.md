<p align="center">
  <img src="assets/icon/512x512.png" alt="gocraft logo" width="180" height="180">
</p>

<h1 align="center">gocraft</h1>

<p align="center">
  Generic wrappers for popular Go libraries — write cleaner, more concise code.
</p>

<p align="center">
  <a href="https://github.com/geebos/gocraft/actions/workflows/ci.yml"><img src="https://github.com/geebos/gocraft/actions/workflows/ci.yml/badge.svg" alt="Go"></a>
  <a href="https://pkg.go.dev/github.com/geebos/gocraft"><img src="https://pkg.go.dev/badge/github.com/geebos/gocraft.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/geebos/gocraft"><img src="https://goreportcard.com/badge/github.com/geebos/gocraft" alt="Go Report Card"></a>
  <a href="https://codecov.io/gh/geebos/gocraft"><img src="https://codecov.io/gh/geebos/gocraft/branch/main/graph/badge.svg" alt="codecov"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"></a>
</p>

📖 **[Documentation](https://geebos.github.io/gocraft/)** | 📦 **[pkg.go.dev](https://pkg.go.dev/github.com/geebos/gocraft)**

## Goals

- **Generic Wrappers**: Provide type-safe generic wrappers for common open-source libraries
- **Cleaner Code**: Reduce boilerplate and make your code more readable

## Quick Start

### Installation

```bash
go get github.com/geebos/gocraft
```

### gjson - Enhanced JSON Operations

```go
import "github.com/geebos/gocraft/pkg/gjson"

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
import "github.com/geebos/gocraft/pkg/gvalue"

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

### gslice - Generic Slice Operations

```go
import (
    "github.com/geebos/gocraft/pkg/gslice"
    "github.com/geebos/gocraft/pkg/gvalue"
    "cmp"
)

// Map, Filter, Reduce
numbers := []int{1, 2, 3, 4, 5}
doubled := gslice.Map(numbers, func(n int) int { return n * 2 })
evens := gslice.Filter(numbers, func(n int) bool { return n%2 == 0 })
sum := gslice.Reduce(numbers, 0, func(acc, n int) int { return acc + n })

// Find, Any, All
value, found := gslice.Find(numbers, func(n int) bool { return n > 3 })
hasEven := gslice.Any(numbers, func(n int) bool { return n%2 == 0 })
allPositive := gslice.All(numbers, func(n int) bool { return n > 0 })

// Sort (clone) and StealSort (in-place)
sorted := gslice.Sort(numbers, cmp.Compare[int])
gslice.StealSort(numbers, cmp.Compare[int])

// Set operations
s1, s2, s3 := []int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}
union := gslice.Union(s1, s2, s3)           // [1, 2, 3, 4, 5]
intersection := gslice.Intersection(s1, s2, s3) // [3]
difference := gslice.Difference(s1, s2, s3)     // [1]

// Use with gvalue comparison functions
greaterThan3 := gslice.Filter(numbers, gslice.CmpWith(gvalue.GT[int], 3))
```

### gweb - Generic HTTP Handler Wrappers

```go
import (
    "github.com/geebos/gocraft/pkg/gweb"
    "github.com/geebos/gocraft/pkg/gweb/ggin"
)

// Create a handler wrapper with custom processors
wrapper := gweb.NewHandlerWrapper(
    gweb.WithRequestProcessor(customRequestProcessor),
    gweb.WithResponseProcessor(customResponseProcessor),
)

// Define type-safe handlers
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUserHandler(ctx context.Context, c *gin.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
    // Business logic here
    return &CreateUserResponse{ID: 1, Name: req.Name, Email: req.Email}, nil
}

// Use in routes
router.POST("/users", ggin.Handler[CreateUserRequest, CreateUserResponse](wrapper, createUserHandler))
```

## Packages

| Package | Description |
|---------|-------------|
| [gjson](https://pkg.go.dev/github.com/geebos/gocraft/pkg/gjson) | Generic JSON encoding/decoding with path extraction support |
| [gvalue](https://pkg.go.dev/github.com/geebos/gocraft/pkg/gvalue) | Generic value utilities, type constraints, and helper functions |
| [gslice](https://pkg.go.dev/github.com/geebos/gocraft/pkg/gslice) | Generic slice and array operations (map, filter, reduce, sort, set operations) |
| [gweb](https://pkg.go.dev/github.com/geebos/gocraft/pkg/gweb) | Generic HTTP handler wrappers with customizable request/response processors |

## Requirements

- Go 1.18+ (generics support required)

## License

[MIT License](LICENSE)
