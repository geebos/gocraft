package gweb

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestProcessor is a function type for processing requests.
// It receives context, gin context, and a request pointer (as any), then returns an error.
type RequestProcessor func(ctx context.Context, c *gin.Context, req any) error

// ResponseProcessor is a function type for processing responses.
// It receives context, gin context, response object (as any), and error.
type ResponseProcessor func(ctx context.Context, c *gin.Context, resp any, err error)

// config holds the configuration for the handler wrapper.
type config struct {
	requestProcessor  func(ctx context.Context, c *gin.Context, req any) error
	responseProcessor func(ctx context.Context, c *gin.Context, resp interface{}, err error)
}

// Option is a function type for configuring the handler wrapper.
type Option func(*config)

// WithRequestProcessor sets a custom request processor.
func WithRequestProcessor(processor RequestProcessor) Option {
	return func(cfg *config) {
		cfg.requestProcessor = processor
	}
}

// WithResponseProcessor sets a custom response processor.
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
		requestProcessor:  defaultRequestProcessorAny,
		responseProcessor: defaultResponseProcessorAny,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &HandlerWrapper{cfg: cfg}
}

// ProcessRequest processes a request using the configured request processor.
func (w *HandlerWrapper) ProcessRequest(ctx context.Context, c *gin.Context, req any) error {
	return w.cfg.requestProcessor(ctx, c, req)
}

// ProcessResponse processes a response using the configured response processor.
func (w *HandlerWrapper) ProcessResponse(ctx context.Context, c *gin.Context, resp interface{}, err error) {
	w.cfg.responseProcessor(ctx, c, resp, err)
}

// defaultRequestProcessorAny is the default request processor that uses binding to bind the request.
func defaultRequestProcessorAny(ctx context.Context, c *gin.Context, req any) error {
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

// defaultResponseProcessorAny is the default response processor that returns JSON response
// in the format {code, data, msg}. When err != nil, code is set to 503 and msg is set to err.Error().
func defaultResponseProcessorAny(ctx context.Context, c *gin.Context, resp interface{}, err error) {
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
