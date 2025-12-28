package ggin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestProcessor is a function type for processing requests with Gin context.
type RequestProcessor func(ctx context.Context, c *gin.Context, req any) error

// ResponseProcessor is a function type for processing responses with Gin context.
type ResponseProcessor func(ctx context.Context, c *gin.Context, resp any, err error)

// config holds the configuration for the handler wrapper.
type config struct {
	requestProcessor  RequestProcessor
	responseProcessor ResponseProcessor
}

// Option is a function type for configuring the handler wrapper.
type Option func(*config)

// defaultRequestProcessor is the default request processor that uses binding to bind the request.
func defaultRequestProcessor(ctx context.Context, c *gin.Context, req any) error {
	// Check if req is interface{}, if so skip binding
	if _, isInterface := req.(*interface{}); isInterface {
		return nil
	}

	// Use gin's ShouldBind to bind the request
	// req is already a pointer, so we can use type assertion to get the concrete type
	// and call ShouldBind on it
	if bindErr := c.ShouldBind(req); bindErr != nil {
		return bindErr
	}

	return nil
}

// defaultResponseProcessor is the default response processor that returns JSON response
// in the format {code, data, msg}. When err != nil, code is set to 503 and msg is set to err.Error().
func defaultResponseProcessor(ctx context.Context, c *gin.Context, resp interface{}, err error) {
	if err != nil {
		// Error response: {code: 503, data: null, msg: err.Error()}
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code": 503,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	// Success response: {code: 200, data: resp, msg: ""}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resp,
		"msg":  "",
	})
}

// WithRequestProcessor sets a custom request processor for Gin.
func WithRequestProcessor(processor RequestProcessor) Option {
	return func(cfg *config) {
		cfg.requestProcessor = processor
	}
}

// WithResponseProcessor sets a custom response processor for Gin.
func WithResponseProcessor(processor ResponseProcessor) Option {
	return func(cfg *config) {
		cfg.responseProcessor = processor
	}
}

// HandlerWrapper is a wrapper that can create handlers with configured processors.
type HandlerWrapper struct {
	cfg *config
}

// NewHandlerWrapper creates a new handler wrapper with the given options.
// If no options are provided, default processors will be used.
// This allows all routers to share the same wrapper instance.
//
// Example:
//
//	wrapper := NewHandlerWrapper(
//		WithRequestProcessor(customReqProcessor),
//		WithResponseProcessor(customRespProcessor),
//	)
func NewHandlerWrapper(opts ...Option) *HandlerWrapper {
	cfg := &config{
		requestProcessor:  defaultRequestProcessor,
		responseProcessor: defaultResponseProcessor,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &HandlerWrapper{cfg: cfg}
}

// ProcessRequest processes a request using the configured request processor.
func (w *HandlerWrapper) ProcessRequest(ctx context.Context, c *gin.Context, req any) error {
	if w.cfg.requestProcessor != nil {
		return w.cfg.requestProcessor(ctx, c, req)
	}
	return nil
}

// ProcessResponse processes a response using the configured response processor.
func (w *HandlerWrapper) ProcessResponse(ctx context.Context, c *gin.Context, resp interface{}, err error) {
	if w.cfg.responseProcessor != nil {
		w.cfg.responseProcessor(ctx, c, resp, err)
	}
}

// Handler returns a gin.HandlerFunc that processes requests using the configured processors.
// This is a generic function that can be used with any request and response types.
func Handler[Req any, Resp any](wrapper *HandlerWrapper, handler func(ctx context.Context, c *gin.Context, req *Req) (*Resp, error)) gin.HandlerFunc {
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
