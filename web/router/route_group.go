package router

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/middleware"
)

type RouteGroup struct {
	Prefix      string
	RouteGroups []IRouteGroup
	Routes      []*RouteHandler
	Middlewares []middleware.IMiddleware
}

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

func (group *RouteGroup) Use(middlewares ...middleware.IMiddleware) {
	group.Middlewares = append(group.Middlewares, middlewares...)
}

func (group *RouteGroup) Group(prefix string, callback func(group *RouteGroup)) *RouteGroup {
	routeGroup := &RouteGroup{
		Prefix: strings.Join([]string{
			strings.TrimRight(group.Prefix, "/"),
			strings.TrimLeft(prefix, "/"),
		}, "/"),
		RouteGroups: make([]IRouteGroup, 0),
		Routes:      make([]*RouteHandler, 0),
		Middlewares: group.Middlewares,
	}
	callback(routeGroup)
	return routeGroup
}

func (group *RouteGroup) Controller(prefix string, controller web.IController, callback func(group *RouteController)) *RouteController {
	routeGroup := &RouteController{
		Prefix: strings.Join([]string{
			strings.TrimRight(group.Prefix, "/"),
			strings.TrimLeft(prefix, "/"),
		}, "/"),
		Routes:         make([]*RouteAction, 0),
		Middlewares:    group.Middlewares,
		Controller:     controller,
		ControllerType: reflect.TypeOf(controller),
	}
	callback(routeGroup)
	group.RouteGroups = append(group.RouteGroups, routeGroup)
	return routeGroup
}

func (group *RouteGroup) Route(method, path string, handler web.Handler) *RouteHandler {
	route := &RouteHandler{
		Route: Route{
			name:   "",
			method: method,
			path: strings.Join([]string{
				strings.TrimRight(group.Prefix, "/"),
				strings.TrimLeft(path, "/"),
			}, "/"),
			middlewares: group.Middlewares,
		},
		handler: handler,
	}
	group.Routes = append(group.Routes, route)
	return route
}

func (group *RouteGroup) HEAD(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodHead, path, handler)
}

func (group *RouteGroup) CONNECT(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodConnect, path, handler)
}

func (group *RouteGroup) OPTIONS(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodOptions, path, handler)
}

func (group *RouteGroup) TRACE(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodTrace, path, handler)
}

func (group *RouteGroup) GET(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodGet, path, handler)
}

func (group *RouteGroup) POST(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodPost, path, handler)
}

func (group *RouteGroup) PUT(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodPut, path, handler)
}

func (group *RouteGroup) PATCH(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodPatch, path, handler)
}

func (group *RouteGroup) DELETE(path string, handler web.Handler) *RouteHandler {
	return group.Route(http.MethodDelete, path, handler)
}
