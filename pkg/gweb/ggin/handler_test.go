package ggin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geebos/gocraft/pkg/gweb"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandler(t *testing.T) {
	Convey("TestHandler", t, func() {
		gin.SetMode(gin.TestMode)

		type TestRequest struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
		}

		type TestResponse struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		Convey("successful request and response", func() {
			wrapper := gweb.NewHandlerWrapper()
			handler := Handler(
				wrapper,
				func(ctx context.Context, c *gin.Context, req *TestRequest) (*TestResponse, error) {
					return &TestResponse{
						ID:    1,
						Name:  req.Name,
						Email: req.Email,
					}, nil
				},
			)

			body := `{"name":"John","email":"john@example.com"}`
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			handler(ctx)

			So(w.Code, ShouldEqual, http.StatusOK)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response["code"], ShouldEqual, 200)
			So(response["msg"], ShouldEqual, "")

			// Verify response data structure and content
			data, ok := response["data"].(map[string]interface{})
			So(ok, ShouldBeTrue)
			So(data["id"], ShouldEqual, float64(1)) // JSON numbers are unmarshaled as float64
			So(data["name"], ShouldEqual, "John")
			So(data["email"], ShouldEqual, "john@example.com")
		})

		Convey("request binding error", func() {
			wrapper := gweb.NewHandlerWrapper()
			handler := Handler(
				wrapper,
				func(ctx context.Context, c *gin.Context, req *TestRequest) (*TestResponse, error) {
					return &TestResponse{ID: 1}, nil
				},
			)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/test", nil)
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			handler(ctx)

			So(w.Code, ShouldEqual, http.StatusServiceUnavailable)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response["code"], ShouldEqual, 503)
			So(response["data"], ShouldBeNil)
			// Verify error message is present and not empty
			msg, ok := response["msg"].(string)
			So(ok, ShouldBeTrue)
			So(msg, ShouldNotBeEmpty)
		})

		Convey("business handler returns error", func() {
			wrapper := gweb.NewHandlerWrapper()
			businessErr := errors.New("business error")
			handler := Handler(
				wrapper,
				func(ctx context.Context, c *gin.Context, req *TestRequest) (*TestResponse, error) {
					return nil, businessErr
				},
			)

			body := `{"name":"John","email":"john@example.com"}`
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			handler(ctx)

			So(w.Code, ShouldEqual, http.StatusServiceUnavailable)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response["code"], ShouldEqual, 503)
			So(response["msg"], ShouldEqual, businessErr.Error())
			// Verify data is nil when error occurs
			So(response["data"], ShouldBeNil)
		})

		Convey("custom request processor", func() {
			customReqErr := errors.New("custom request error")
			wrapper := gweb.NewHandlerWrapper(
				gweb.WithRequestProcessor(func(ctx context.Context, c *gin.Context, req any) error {
					return customReqErr
				}),
			)
			handler := Handler(
				wrapper,
				func(ctx context.Context, c *gin.Context, req *TestRequest) (*TestResponse, error) {
					return &TestResponse{ID: 1}, nil
				},
			)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/test", nil)
			ctx.Request = req

			handler(ctx)

			So(w.Code, ShouldEqual, http.StatusServiceUnavailable)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response["code"], ShouldEqual, 503)
			So(response["msg"], ShouldEqual, customReqErr.Error())
			// Verify data is nil when request processor returns error
			So(response["data"], ShouldBeNil)
		})

		Convey("custom response processor", func() {
			wrapper := gweb.NewHandlerWrapper(
				gweb.WithResponseProcessor(func(ctx context.Context, c *gin.Context, resp any, err error) {
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"result": resp})
				}),
			)
			handler := Handler(
				wrapper,
				func(ctx context.Context, c *gin.Context, req *TestRequest) (*TestResponse, error) {
					return &TestResponse{
						ID:    1,
						Name:  req.Name,
						Email: req.Email,
					}, nil
				},
			)

			body := `{"name":"John","email":"john@example.com"}`
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			handler(ctx)

			// Should use custom response processor format
			So(w.Code, ShouldEqual, http.StatusOK)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)

			// Verify custom response format (not the default {code, data, msg})
			So(response["code"], ShouldBeNil)
			So(response["msg"], ShouldBeNil)

			// Verify result structure and content
			result, ok := response["result"].(map[string]interface{})
			So(ok, ShouldBeTrue)
			So(result["id"], ShouldEqual, float64(1))
			So(result["name"], ShouldEqual, "John")
			So(result["email"], ShouldEqual, "john@example.com")
		})
	})
}
