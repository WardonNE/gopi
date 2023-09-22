package router

import (
	"github.com/wardonne/gopi/binding"
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/middleware"
)

type IRoute interface {
	AS(name string) IRoute
	Use(middlewares ...middleware.IMiddleware) IRoute
	Validate(form validation.IValidateForm, bindings ...binding.Binding) IRoute
	Name() string
	Method() string
	Path() string
	Middlewares() []middleware.IMiddleware
	Handler() string
	HasValidation() bool
	HandleRequest(request *context.Request) context.IResponse
}

type Route struct {
	name        string
	method      string
	path        string
	middlewares []middleware.IMiddleware
	validation  *middleware.ValidateMiddleware
}

func (route *Route) Name() string {
	return route.name
}

func (route *Route) Method() string {
	return route.method
}

func (route *Route) Path() string {
	return route.path
}

func (route *Route) Middlewares() []middleware.IMiddleware {
	return route.middlewares
}

func (route *Route) HasValidation() bool {
	return route.validation != nil
}
