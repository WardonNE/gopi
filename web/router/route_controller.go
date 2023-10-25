package router

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// IController interface of controller
type IController interface {
	// Init inits controller
	Init(request *context.Request)
}

// RouteController a route group of [web.IController]
type RouteController struct {
	router *Router

	Prefix             string
	RouteGroups        []IRouteGroup
	Routes             []*RouteAction
	Middlewares        *list.ArrayList[middleware.IMiddleware]
	ControllerInstance IController
	ControllerType     reflect.Type
}

// List lists all routes in current group
func (group *RouteController) List() []IRoute {
	routes := make([]IRoute, 0)
	for _, route := range group.Routes {
		routes = append(routes, route)
	}
	for _, routeGroup := range group.RouteGroups {
		routes = append(routes, routeGroup.List()...)
	}
	return routes
}

// Use sets middlewares to current group
func (group *RouteController) Use(middlewares ...middleware.IMiddleware) {
	group.Middlewares.AddAll(middlewares...)
}

// Group registes a sub group of routes to current route
func (group *RouteController) Group(prefix string, callback func(group *RouteController)) *RouteController {
	routeGroup := &RouteController{
		router: group.router,
		Prefix: strings.Join([]string{
			strings.TrimRight(group.Prefix, "/"),
			strings.TrimLeft(prefix, "/"),
		}, "/"),
		Routes:             make([]*RouteAction, 0),
		Middlewares:        group.Middlewares,
		ControllerInstance: group.ControllerInstance,
		ControllerType:     group.ControllerType,
	}
	callback(routeGroup)
	group.RouteGroups = append(group.RouteGroups, routeGroup)
	return routeGroup
}

// Controller registes a sub route group with specific controller instance and returns an instance of [RouteController]
func (group *RouteController) Controller(prefix string, controller IController, callback func(group *RouteController)) *RouteController {
	routeGroup := &RouteController{
		router: group.router,
		Prefix: strings.Join([]string{
			strings.TrimRight(group.Prefix, "/"),
			strings.TrimLeft(prefix, "/"),
		}, "/"),
		Routes:             make([]*RouteAction, 0),
		Middlewares:        group.Middlewares,
		ControllerInstance: controller,
		ControllerType:     reflect.TypeOf(controller),
	}
	callback(routeGroup)
	group.RouteGroups = append(group.RouteGroups, routeGroup)
	return routeGroup
}

// Route registers controller's action to current route group and returns an instance of [RouteAction]
//
// the action should only receive only the receiver
//
// the action should only return an implemention of [context.IResponse]
//
// # NOTICE: The handler CAN'T be an empty string
//
// # NOTICE: The handler CAN'T be UNEXPORTED
//
// Example:
//
//	type Controller struct{}
//
//	func (c *Controller) Login() context.IResponse
func (group *RouteController) Route(method, path, handler string) *RouteAction {
	handler = strings.TrimSpace(handler)
	if len(handler) == 0 {
		panic("handler name is empty")
	}
	if handler[0] > 96 || handler[0] < 65 {
		panic("handler is not a public method")
	}
	handlerType, exists := group.ControllerType.MethodByName(handler)
	if !exists {
		panic(group.ControllerType.Name() + "." + handler + " not found")
	}

	if numIn := handlerType.Type.NumIn(); numIn != 1 {
		panic("invalid number of input")
	}

	if numOut := handlerType.Type.NumOut(); numOut != 1 {
		panic("invalid number of output")
	}

	if outputType := handlerType.Type.Out(0); !outputType.Implements(reflect.TypeOf((*context.IResponse)(nil)).Elem()) {
		panic("invalid return type, should be an implemention of context.IResponse")
	}

	action := &RouteAction{
		Route: Route{
			router: group.router,
			name:   "",
			method: method,
			path: strings.Join([]string{
				strings.TrimRight(group.Prefix, "/"),
				strings.TrimLeft(path, "/"),
			}, "/"),
			middlewares: group.Middlewares,
		},
		handler:        handler,
		controller:     group.ControllerInstance,
		controllerType: group.ControllerType,
	}
	group.Routes = append(group.Routes, action)
	return action
}

// HEAD registers an action with method [http.MethodHead]
func (group *RouteController) HEAD(path, handler string) *RouteAction {
	return group.Route(http.MethodHead, path, handler)
}

// CONNECT registers an action with method [http.MethodConnect]
func (group *RouteController) CONNECT(path, handler string) *RouteAction {
	return group.Route(http.MethodConnect, path, handler)
}

// OPTIONS registers an action with method [http.MethodOptions]
func (group *RouteController) OPTIONS(path, handler string) *RouteAction {
	return group.Route(http.MethodOptions, path, handler)
}

// TRACE registers an action with method [http.MethodTrace]
func (group *RouteController) TRACE(path, handler string) *RouteAction {
	return group.Route(http.MethodTrace, path, handler)
}

// GET registers an action with method [http.MethodGet]
func (group *RouteController) GET(path, handler string) *RouteAction {
	return group.Route(http.MethodGet, path, handler)
}

// POST registers an action with method [http.MethodPost]
func (group *RouteController) POST(path, handler string) *RouteAction {
	return group.Route(http.MethodPost, path, handler)
}

// PUT registers an action with method [http.MethodPut]
func (group *RouteController) PUT(path, handler string) *RouteAction {
	return group.Route(http.MethodPut, path, handler)
}

// PATCH registers an action with method [http.MethodPatch]
func (group *RouteController) PATCH(path, handler string) *RouteAction {
	return group.Route(http.MethodPatch, path, handler)
}

// DELETE registers an action with method [http.MethodDelete]
func (group *RouteController) DELETE(path, handler string) *RouteAction {
	return group.Route(http.MethodDelete, path, handler)
}
