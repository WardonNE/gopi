package router

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type testcontroller struct {
	web.Controller
}

func (c *testcontroller) Index() context.IResponse {
	return context.NewResponse(200).JSON(map[string]any{
		"code":   200,
		"method": c.Method(),
	})
}

func (c *testcontroller) Valid() context.IResponse {
	form := c.Validated().(*testform)
	if form.Fails() {
		return context.NewResponse(400).JSON(map[string]any{
			"code":   400,
			"method": c.Method(),
			"errs":   form.Errors(),
		})
	}
	return context.NewResponse(200).JSON(map[string]any{
		"code":   200,
		"method": c.Method(),
		"name":   form.Name,
	})
}

func TestRouter_Handler(t *testing.T) {
	r := New()
	handler := func(r *context.Request) context.IResponse {
		return context.NewResponse(200).JSON(map[string]any{
			"code":   200,
			"method": r.Method(),
		})
	}
	r.Group("api", func(group *RouteGroup) {
		group.GET("handler", handler).AS("GetHandler")
		group.POST("handler", handler).AS("PostHandler")
		group.PUT("handler", handler).AS("PutHandler")
		group.PATCH("handler", handler).AS("PatchHandler")
		group.DELETE("handler", handler).AS("DeleteHandler")
		group.OPTIONS("handler", handler).AS("OptionsHandler")
		group.HEAD("handler", handler).AS("HeadHandler")
		group.CONNECT("handler", handler).AS("ConnectHandler")
		group.TRACE("handler", handler).AS("TraceHandler")
	})

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)

			req, err := http.NewRequest(method, "/api/handler", nil)
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			assert.Equal(t, 200, recorder.Result().StatusCode)
			assert.JSONEq(t, `{"code":200, "method":"`+method+`"}`, recorder.Body.String())
		}(method)
	}
	wg.Wait()
}

func TestRouter_HandlerWithMiddleware(t *testing.T) {
	r := New()
	handler := func(r *context.Request) context.IResponse {
		return context.NewResponse(200).JSON(map[string]any{
			"code":   200,
			"method": r.Method(),
		})
	}
	var mw1 middleware.IMiddleware = func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		request.Set("group-method", request.Method())
		return next(request)
	}
	var mw2 middleware.IMiddleware = func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		method := request.Method()
		// only GET and POST is allowed
		if method == http.MethodGet || method == http.MethodPost {
			return next(request)
		}
		return context.NewResponse(400).JSON(map[string]any{
			"code":   400,
			"method": request.Method(),
		})
	}
	r.Group("api", func(group *RouteGroup) {
		group.GET("handler", handler).Use(mw2).AS("GetHandler")
		group.POST("handler", handler).Use(mw2).AS("PostHandler")
		group.PUT("handler", handler).Use(mw2).AS("PutHandler")
		group.PATCH("handler", handler).Use(mw2).AS("PatchHandler")
		group.DELETE("handler", handler).Use(mw2).AS("DeleteHandler")
		group.OPTIONS("handler", handler).Use(mw2).AS("OptionsHandler")
		group.HEAD("handler", handler).Use(mw2).AS("HeadHandler")
		group.CONNECT("handler", handler).Use(mw2).AS("ConnectHandler")
		group.TRACE("handler", handler).Use(mw2).AS("TraceHandler")
	}).Use(mw1)

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)

			req, err := http.NewRequest(method, "/api/handler", nil)
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			if method == http.MethodGet || method == http.MethodPost {
				assert.Equal(t, 200, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":200, "method":"`+method+`"}`, recorder.Body.String())
			} else {
				assert.Equal(t, 400, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":400, "method":"`+method+`"}`, recorder.Body.String())
			}
		}(method)
	}
	wg.Wait()
}

type testform struct {
	validation.ValidateForm
	Name *string `json:"name"`
}

type testfilederror struct{ validator.FieldError }

func (fe testfilederror) Tag() string   { return "required" }
func (fe testfilederror) Field() string { return "Name" }
func (fe testfilederror) Error() string { return "field required" }

type testvalidationengine struct{}

func (v *testvalidationengine) Struct(data any) error {
	if data.(*testform).Name == nil {
		err := make(validator.ValidationErrors, 0)
		return append(err, testfilederror{})
	}
	return nil
}

func (v *testvalidationengine) Translator(locale string) ut.Translator {
	return nil
}

func TestRouter_HandlerWithValidation(t *testing.T) {
	r := New()
	r.SetValidateEngine(new(testvalidationengine))
	handler := func(r *context.Request) context.IResponse {
		form := r.Validated().(*testform)
		if form.Fails() {
			return context.NewResponse(400).JSON(map[string]any{
				"code":   400,
				"method": r.Method(),
				"errs":   form.Errors(),
			})
		}
		return context.NewResponse(200).JSON(map[string]any{
			"code":   200,
			"method": r.Method(),
			"name":   form.Name,
		})
	}

	form := new(testform)
	r.Group("api", func(group *RouteGroup) {
		group.GET("handler", handler).AS("GetHandler").Validate(form)
		group.POST("handler", handler).AS("PostHandler").Validate(form)
		group.PUT("handler", handler).AS("PutHandler").Validate(form)
		group.PATCH("handler", handler).AS("PatchHandler").Validate(form)
		group.DELETE("handler", handler).AS("DeleteHandler").Validate(form)
		group.OPTIONS("handler", handler).AS("OptionsHandler").Validate(form)
		group.HEAD("handler", handler).AS("HeadHandler").Validate(form)
		group.CONNECT("handler", handler).AS("ConnectHandler").Validate(form)
		group.TRACE("handler", handler).AS("TraceHandler").Validate(form)
	})

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)

			var reader io.Reader
			if method == http.MethodGet || method == http.MethodPost {
				reader = bytes.NewReader([]byte(`{"name": "testuser"}`))
			} else {
				reader = bytes.NewReader([]byte(`{"name": null}`))
			}
			req, err := http.NewRequest(method, "/api/handler", reader)
			assert.Nil(t, err)
			req.Header.Add("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			if method == http.MethodGet || method == http.MethodPost {
				assert.Equal(t, 200, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":200, "method":"`+method+`", "name":"testuser"}`, recorder.Body.String())
			} else {
				assert.Equal(t, 400, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":400, "method":"`+method+`", "errs": {"Name": ["field required"]}}`, recorder.Body.String())
			}
		}(method)
	}
	wg.Wait()
}

func TestRouter_Action(t *testing.T) {
	r := New()
	r.Controller("/api", new(testcontroller), func(group *RouteController) {
		group.GET("handler", "Index").AS("GetAction")
		group.POST("handler", "Index").AS("PostAction")
		group.PUT("handler", "Index").AS("PutAction")
		group.PATCH("handler", "Index").AS("PatchAction")
		group.DELETE("handler", "Index").AS("DeleteAction")
		group.OPTIONS("handler", "Index").AS("OptionsAction")
		group.HEAD("handler", "Index").AS("HeadAction")
		group.CONNECT("handler", "Index").AS("ConnectAction")
		group.TRACE("handler", "Index").AS("TraceAction")
	})

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)

			req, err := http.NewRequest(method, "/api/handler", nil)
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			assert.Equal(t, 200, recorder.Result().StatusCode)
			assert.JSONEq(t, `{"code":200, "method":"`+method+`"}`, recorder.Body.String())
		}(method)
	}
	wg.Wait()
}

func TestRouter_ActionWithMiddleware(t *testing.T) {
	r := New()
	var mw1 middleware.IMiddleware = func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		request.Set("group-method", request.Method())
		return next(request)
	}
	var mw2 middleware.IMiddleware = func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		method := request.Method()
		// only GET and POST is allowed
		if method == http.MethodGet || method == http.MethodPost {
			return next(request)
		}
		return context.NewResponse(400).JSON(map[string]any{
			"code":   400,
			"method": request.Method(),
		})
	}
	r.Controller("/api", new(testcontroller), func(group *RouteController) {
		group.GET("handler", "Index").AS("GetAction").Use(mw2)
		group.POST("handler", "Index").AS("PostAction").Use(mw2)
		group.PUT("handler", "Index").AS("PutAction").Use(mw2)
		group.PATCH("handler", "Index").AS("PatchAction").Use(mw2)
		group.DELETE("handler", "Index").AS("DeleteAction").Use(mw2)
		group.OPTIONS("handler", "Index").AS("OptionsAction").Use(mw2)
		group.HEAD("handler", "Index").AS("HeadAction").Use(mw2)
		group.CONNECT("handler", "Index").AS("ConnectAction").Use(mw2)
		group.TRACE("handler", "Index").AS("TraceAction").Use(mw2)
	}).Use(mw1)

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)

			req, err := http.NewRequest(method, "/api/handler", nil)
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			if method == http.MethodGet || method == http.MethodPost {
				assert.Equal(t, 200, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":200, "method":"`+method+`"}`, recorder.Body.String())
			} else {
				assert.Equal(t, 400, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":400, "method":"`+method+`"}`, recorder.Body.String())
			}
		}(method)
	}
	wg.Wait()
}

func TestRouter_ActionWithValidation(t *testing.T) {
	r := New()
	r.SetValidateEngine(new(testvalidationengine))
	form := new(testform)
	r.Controller("/api", new(testcontroller), func(group *RouteController) {
		group.GET("handler", "Valid").AS("GetAction").Validate(form)
		group.POST("handler", "Valid").AS("PostAction").Validate(form)
		group.PUT("handler", "Valid").AS("PutAction").Validate(form)
		group.PATCH("handler", "Valid").AS("PatchAction").Validate(form)
		group.DELETE("handler", "Valid").AS("DeleteAction").Validate(form)
		group.OPTIONS("handler", "Valid").AS("OptionsAction").Validate(form)
		group.HEAD("handler", "Valid").AS("HeadAction").Validate(form)
		group.CONNECT("handler", "Valid").AS("ConnectAction").Validate(form)
		group.TRACE("handler", "Valid").AS("TraceAction").Validate(form)
	})

	r.Prepare()
	rr := r.HTTPRouter
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
		http.MethodConnect,
		http.MethodTrace,
	}
	wg := sync.WaitGroup{}
	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			handler, _, _ := rr.Lookup(method, "/api/handler")
			assert.NotNil(t, handler)
			var reader io.Reader
			if method == http.MethodGet || method == http.MethodPost {
				reader = bytes.NewReader([]byte(`{"name":"testuser"}`))
			} else {
				reader = bytes.NewReader([]byte(`{}`))
			}
			req, err := http.NewRequest(method, "/api/handler", reader)
			req.Header.Add("Content-Type", "application/json")
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			rr.ServeHTTP(recorder, req)
			if method == http.MethodGet || method == http.MethodPost {
				assert.Equal(t, 200, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":200, "method":"`+method+`", "name":"testuser"}`, recorder.Body.String())
			} else {
				assert.Equal(t, 400, recorder.Result().StatusCode)
				assert.JSONEq(t, `{"code":400, "method":"`+method+`", "errs":{"Name":["field required"]}}`, recorder.Body.String())
			}
		}(method)
	}
	wg.Wait()
}
