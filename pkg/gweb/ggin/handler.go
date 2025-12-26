package ggin

import (
	"context"

	"github.com/geebos/gocraft/pkg/gweb"
	"github.com/gin-gonic/gin"
)

// Handler returns a gin.HandlerFunc that processes requests using the configured processors.
// This is a generic function that can be used with any request and response types.
func Handler[Req any, Resp any](wrapper *gweb.HandlerWrapper, handler func(ctx context.Context, c *gin.Context, req *Req) (*Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *Req
		var resp *Resp
		var err error

		// Get request context
		ctx := c.Request.Context()

		// Create request instance
		req = new(Req)

		// Process request using configured processor
		err = wrapper.ProcessRequest(ctx, c, req)
		if err != nil {
			wrapper.ProcessResponse(ctx, c, nil, err)
			return
		}

		// Call business handler function
		resp, err = handler(ctx, c, req)

		// Process response using configured processor
		wrapper.ProcessResponse(ctx, c, resp, err)
	}
}
