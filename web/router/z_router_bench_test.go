package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
)

func BenchmarkRouter_Handler(b *testing.B) {
	r := New()
	r.GET("/get", func(request *context.Request) context.IResponse {
		return context.NewResponse(200, "Hello World")
	})
	rr := r.HTTPRouter
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			req, _ := http.NewRequest("GET", "/get", nil)
			var resp = httptest.NewRecorder()
			rr.ServeHTTP(resp, req)
		}
	})
}

func BenchmarkRouter_HandlerWithValidate(b *testing.B) {
	r := New()
	r.SetValidateEngine(new(testvalidationengine))
	r.POST("/post", func(request *context.Request) context.IResponse {
		return context.NewResponse(200, request.Validated().(*testform).Name)
	}).Validate(new(testform), binding.JSON)
	rr := r.HTTPRouter
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			req, _ := http.NewRequest("POST", "/post", bytes.NewReader([]byte(`{"name": "user"}`)))
			var resp = httptest.NewRecorder()
			rr.ServeHTTP(resp, req)
		}
	})
}

func BenchmarkRouter_Action(b *testing.B) {
	r := New()
	r.Controller("api", new(testcontroller), func(group *RouteController) {
		group.GET("get", "Index")
	})
	rr := r.HTTPRouter
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			req, _ := http.NewRequest("GET", "/get", nil)
			var resp = httptest.NewRecorder()
			rr.ServeHTTP(resp, req)
		}
	})
}

func BenchmarkRouter_ActionWithValidate(b *testing.B) {
	r := New()
	r.SetValidateEngine(new(testvalidationengine))
	r.Controller("api", new(testcontroller), func(group *RouteController) {
		group.POST("post", "Valid").Validate(new(testform), binding.JSON)
	})
	rr := r.HTTPRouter
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			req, _ := http.NewRequest("POST", "/api/post", bytes.NewReader([]byte(`{"name": "user"}`)))
			var resp = httptest.NewRecorder()
			rr.ServeHTTP(resp, req)
		}
	})
}
