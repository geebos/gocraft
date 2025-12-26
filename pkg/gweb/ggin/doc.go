// Package ggin provides generic HTTP handler utilities for Gin framework.
//
// The package includes a generic Handler function that works with gweb.HandlerWrapper
// to provide type-safe HTTP handlers with configurable request and response processors.
//
// # Basic Usage
//
//	import (
//	    "github.com/geebos/gocraft/pkg/gweb"
//	    "github.com/geebos/gocraft/pkg/gweb/ggin"
//	)
//
//	wrapper := gweb.NewHandlerWrapper()
//	router.POST("/users", ggin.Handler[CreateUserRequest, CreateUserResponse](wrapper, createUserHandler))
//
// For more examples, see the gweb package documentation.
package ggin
