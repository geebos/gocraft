// Package gweb provides generic HTTP handler wrapper utilities with customizable request
// and response processors.
//
// The package includes a factory function that creates a handler wrapper,
// allowing all routers to share the same wrapper configuration.
//
// # Recommended Usage
//
// Define handlers in a separate package (e.g., user package):
//
//		package user
//
//		import (
//		    "context"
//		    "github.com/gin-gonic/gin"
//		)
//
//		type CreateUserRequest struct {
//		    Name  string `json:"name" binding:"required"`
//		    Email string `json:"email" binding:"required,email"`
//		}
//
//		type CreateUserResponse struct {
//		    ID    int    `json:"id"`
//		    Name  string `json:"name"`
//		    Email string `json:"email"`
//		}
//
//		func CreateUserHandler(ctx context.Context, c *gin.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
//	     // The request should be validated by gin binding, so you don't need to validate it here
//		    // Business logic here
//		    // ... create user in database ...
//		    return &CreateUserResponse{
//		        ID:    1,
//		        Name:  req.Name,
//		        Email: req.Email,
//		    }, nil
//		}
//
// For better code consistency, it's recommended to define a global wrapper
// and a global Handler function in your application:
//
//	package main
//
//	import (
//	    "context"
//	    "github.com/gin-gonic/gin"
//	    "github.com/geebos/gocraft/pkg/gweb"
//	    "github.com/geebos/gocraft/pkg/gweb/ggin"
//	    "yourproject/user" // Import your user package
//	)
//
//	// Define a global wrapper (initialize in init() or main())
//	var wrapper = gweb.NewHandlerWrapper(
//	    // Add custom processors if needed
//	    gweb.WithRequestProcessor(customRequestProcessor),
//	    gweb.WithResponseProcessor(customResponseProcessor),
//	)
//
//	// Define a global Handler function for convenience
//	func Handler[Req any, Resp any](handler func(ctx context.Context, c *gin.Context, req *Req) (*Resp, error)) gin.HandlerFunc {
//	    return ggin.Handler[Req, Resp](wrapper, handler)
//	}
//
//	// Now use it in your routes
//	func setupRoutes(router *gin.Engine) {
//	    router.POST("/users", Handler[user.CreateUserRequest, user.CreateUserResponse](user.CreateUserHandler))
//	    router.GET("/users/:id", Handler[user.GetUserRequest, user.GetUserResponse](user.GetUserHandler))
//	}
//
// # Custom Processors
//
// Create a wrapper with custom request and response processors:
//
//	wrapper := gweb.NewHandlerWrapper(
//	    gweb.WithRequestProcessor(func(ctx context.Context, c *gin.Context, req any) error {
//	        // Custom request processing logic
//	        // req is already a pointer to the request type
//	        // ... custom binding logic ...
//	        return c.ShouldBind(req)
//	    }),
//	    gweb.WithResponseProcessor(func(ctx context.Context, c *gin.Context, resp any, err error) {
//	        if err != nil {
//	            c.JSON(500, gin.H{"error": err.Error()})
//	            return
//	        }
//	        c.JSON(200, gin.H{"code": 200, "data": resp, "msg": "ok"})
//	    }),
//	)
//
//	// Use the same wrapper for all routes
//	router.POST("/users", ggin.Handler[CreateUserRequest, CreateUserResponse](wrapper, createUserHandler))
//	router.GET("/users/:id", ggin.Handler[GetUserRequest, GetUserResponse](wrapper, getUserHandler))
package gweb
