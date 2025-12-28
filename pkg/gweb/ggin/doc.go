// Package ggin provides generic HTTP handler utilities for Gin framework.
//
// The package includes a generic Handler function and a NewHandlerWrapper function
// to provide type-safe HTTP handlers with configurable request and response processors.
//
// # Basic Usage
//
//	import (
//	    "context"
//	    "github.com/gin-gonic/gin"
//	    "github.com/geebos/gocraft/pkg/gweb/ggin"
//	)
//
//	// Create a wrapper with default processors for Gin
//	wrapper := ggin.NewHandlerWrapper()
//	router.POST("/users", ggin.Handler[CreateUserRequest, CreateUserResponse](wrapper, createUserHandler))
//
// # With Custom Processors
//
//	wrapper := ggin.NewHandlerWrapper(
//	    ggin.WithRequestProcessor(customReqProcessor),
//	    ggin.WithResponseProcessor(customRespProcessor),
//	)
//
// # Recommended Usage
//
// Define handlers in a separate package (e.g., user package):
//
//	package user
//
//	import (
//	    "context"
//	    "github.com/gin-gonic/gin"
//	)
//
//	type CreateUserRequest struct {
//	    Name  string `json:"name" binding:"required"`
//	    Email string `json:"email" binding:"required,email"`
//	}
//
//	type CreateUserResponse struct {
//	    ID    int    `json:"id"`
//	    Name  string `json:"name"`
//	    Email string `json:"email"`
//	}
//
//	func CreateUserHandler(ctx context.Context, c *gin.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
//	    // Business logic here
//	    return &CreateUserResponse{
//	        ID:    1,
//	        Name:  req.Name,
//	        Email: req.Email,
//	    }, nil
//	}
//
// For better code consistency, define a global wrapper and Handler function:
//
//	package main
//
//	import (
//	    "context"
//	    "github.com/gin-gonic/gin"
//	    "github.com/geebos/gocraft/pkg/gweb/ggin"
//	    "yourproject/user"
//	)
//
//	var wrapper = ggin.NewHandlerWrapper()
//
//	func Handler[Req any, Resp any](handler func(ctx context.Context, c *gin.Context, req *Req) (*Resp, error)) gin.HandlerFunc {
//	    return ggin.Handler[Req, Resp](wrapper, handler)
//	}
//
//	func setupRoutes(router *gin.Engine) {
//	    router.POST("/users", Handler[user.CreateUserRequest, user.CreateUserResponse](user.CreateUserHandler))
//	    router.GET("/users/:id", Handler[user.GetUserRequest, user.GetUserResponse](user.GetUserHandler))
//	}
package ggin
