package router

import (
	"reflect"
	"runtime"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type RouteHandler struct {
	Route
	handler web.Handler
}

func (route *RouteHandler) AS(name string) IRoute {
	route.name = name
	return route
}

func (route *RouteHandler) Use(middlewares ...middleware.IMiddleware) IRoute {
	route.middlewares.AddAll(middlewares...)
	return route
}

func (route *RouteHandler) Validate(form validation.IValidateForm, bindings ...binding.Binding) IRoute {
	formType := reflect.TypeOf(form)
	if formType.Kind() != reflect.Ptr {
		panic("Non-ptr: " + formType.String())
	}
	route.validation = middleware.Validation(form, bindings...)
	return route
}

func (route *RouteHandler) Handler() string {
	return runtime.FuncForPC(reflect.ValueOf(route.handler).Pointer()).Name()
}

func (route *RouteHandler) HandleRequest(request *context.Request) context.IResponse {
	pl := new(pipeline.Pipeline[*context.Request, context.IResponse])
	pl = pl.Send(request)
	route.middlewares.Range(func(middleware middleware.IMiddleware) bool {
		pl.AppendThroughCallback(middleware)
		return true
	})
	if route.HasValidation() {
		pl = pl.AppendThrough(route.validation)
	}
	return pl.Then(route.handler)
}
