package router

import (
	"reflect"
	"runtime"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
	"github.com/wardonne/gopi/web/middleware/validate"
)

// RouteHandler route of handler function
type RouteHandler struct {
	Route
	handler Handler
}

// AS sets the name
func (route *RouteHandler) AS(name string) IRoute {
	route.name = name
	return route
}

// Use sets middlewares
func (route *RouteHandler) Use(middlewares ...middleware.IMiddleware) IRoute {
	route.middlewares.AddAll(middlewares...)
	return route
}

// Validate binds validation form to the route
func (route *RouteHandler) Validate(form validation.IValidateForm, bindings ...binding.Binding) IRoute {
	if route.router.validateEngine == nil {
		panic(ErrValidateEngineEmpty)
	}
	formType := reflect.TypeOf(form)
	if formType.Kind() != reflect.Ptr {
		panic("Non-ptr: " + formType.String())
	}
	form.SetEngine(route.router.validateEngine)
	route.validation = validate.New(form, bindings...)
	return route
}

// Handler returns the handler's name
func (route *RouteHandler) Handler() string {
	return runtime.FuncForPC(reflect.ValueOf(route.handler).Pointer()).Name()
}

// HandleRequest handles the http request
func (route *RouteHandler) HandleRequest(request *context.Request) context.IResponse {
	pl := new(pipeline.Pipeline[*context.Request, context.IResponse])
	pl = pl.Send(request)
	route.middlewares.Range(func(middleware middleware.IMiddleware) bool {
		pl.AppendThroughCallback(middleware)
		return true
	})
	if route.HasValidation() {
		pl = pl.AppendThroughCallback(route.validation)
	}
	return pl.Then(route.handler)
}
