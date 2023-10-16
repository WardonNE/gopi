package router

import (
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// IRoute route interface
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

// Route basic route struct
type Route struct {
	router *Router

	name        string
	method      string
	path        string
	middlewares *list.ArrayList[middleware.IMiddleware]
	validation  middleware.IMiddleware
}

// Name returns the name of route
func (route *Route) Name() string {
	return route.name
}

// Method returns the method of route
func (route *Route) Method() string {
	return route.method
}

// Path returns the path of route
func (route *Route) Path() string {
	return route.path
}

// Middlewares returns the middlewares of the route
func (route *Route) Middlewares() []middleware.IMiddleware {
	return route.middlewares.ToArray()
}

// HasValidation returns whether the route has a binded validation
func (route *Route) HasValidation() bool {
	return route.validation != nil
}
