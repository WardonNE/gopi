package router

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/web/middleware"
)

// RouteGroup used to manage a group of [Route]
type RouteGroup struct {
	router *Router

	Prefix      string
	RouteGroups []IRouteGroup
	Routes      []*RouteHandler
	Middlewares *list.ArrayList[middleware.IMiddleware]
}

// List lists all routes in current group
func (group *RouteGroup) List() []IRoute {
	routes := make([]IRoute, 0)
	for _, route := range group.Routes {
		routes = append(routes, route)
	}
	for _, routeGroup := range group.RouteGroups {
		routes = append(routes, routeGroup.List()...)
	}
	return routes
}

// Use sets middlewares to the group
func (group *RouteGroup) Use(middlewares ...middleware.IMiddleware) {
	group.Middlewares.AddAll(middlewares...)
}

// Group registers sub route group and returns the sub group instance
func (group *RouteGroup) Group(prefix string, callback func(group *RouteGroup)) *RouteGroup {
	routeGroup := &RouteGroup{
		router: group.router,
		Prefix: strings.Join([]string{
			strings.TrimRight(group.Prefix, "/"),
			strings.TrimLeft(prefix, "/"),
		}, "/"),
		RouteGroups: make([]IRouteGroup, 0),
		Routes:      make([]*RouteHandler, 0),
		Middlewares: group.Middlewares,
	}
	group.RouteGroups = append(group.RouteGroups, routeGroup)
	callback(routeGroup)
	return routeGroup
}

// Controller registers a sub route group with specific controller instance and returns an instance of [RouteController]
func (group *RouteGroup) Controller(prefix string, controller IController, callback func(group *RouteController)) *RouteController {
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

// Route registers a handler route to current group and it returns an instance of [RouteHandler]
func (group *RouteGroup) Route(method, path string, handler Handler) *RouteHandler {
	pathWithPrefix := strings.Join([]string{
		strings.TrimRight(group.Prefix, "/"),
		strings.TrimLeft(path, "/"),
	}, "/")
	if path == "" {
		pathWithPrefix = pathWithPrefix[:len(pathWithPrefix)-1]
	}
	route := &RouteHandler{
		Route: Route{
			router:      group.router,
			name:        "",
			method:      method,
			path:        pathWithPrefix,
			middlewares: group.Middlewares,
		},
		handler: handler,
	}
	group.Routes = append(group.Routes, route)
	return route
}

// HEAD registers a handler route with method [http.MethodHead]
func (group *RouteGroup) HEAD(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodHead, path, handler)
}

// CONNECT registers a handler route with method [http.MethodConnect]
func (group *RouteGroup) CONNECT(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodConnect, path, handler)
}

// OPTIONS registers a handler route with method [http.MethodOptions]
func (group *RouteGroup) OPTIONS(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodOptions, path, handler)
}

// TRACE registers a handler route with method [http.MethodTrace]
func (group *RouteGroup) TRACE(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodTrace, path, handler)
}

// GET registers a handler route with method [http.MethodGet]
func (group *RouteGroup) GET(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodGet, path, handler)
}

// POST registers a handler route with method [http.MethodPost]
func (group *RouteGroup) POST(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodPost, path, handler)
}

// PUT registers a handler route with method [http.MethodPut]
func (group *RouteGroup) PUT(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodPut, path, handler)
}

// PATCH registers a handler route with method [http.MethodPatch]
func (group *RouteGroup) PATCH(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodPatch, path, handler)
}

// DELETE registers a handler route with method [http.MethodDelete]
func (group *RouteGroup) DELETE(path string, handler Handler) *RouteHandler {
	return group.Route(http.MethodDelete, path, handler)
}
