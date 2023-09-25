package router

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type RouteController struct {
	Prefix         string
	Routes         []*RouteAction
	Middlewares    *list.ArrayList[middleware.IMiddleware]
	Controller     web.IController
	ControllerType reflect.Type
}

func (group *RouteController) List() []IRoute {
	routes := make([]IRoute, 0)
	for _, route := range group.Routes {
		routes = append(routes, route)
	}
	return routes
}

func (group *RouteController) Use(middlewares ...middleware.IMiddleware) {
	group.Middlewares.AddAll(middlewares...)
}

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
			name:   "",
			method: method,
			path: strings.Join([]string{
				strings.TrimRight(group.Prefix, "/"),
				strings.TrimLeft(path, "/"),
			}, "/"),
			middlewares: group.Middlewares,
		},
		handler:        handler,
		controller:     group.Controller,
		controllerType: group.ControllerType,
	}
	group.Routes = append(group.Routes, action)
	return action
}

func (group *RouteController) HEAD(path, handler string) *RouteAction {
	return group.Route(http.MethodHead, path, handler)
}

func (group *RouteController) CONNECT(path, handler string) *RouteAction {
	return group.Route(http.MethodConnect, path, handler)
}

func (group *RouteController) OPTIONS(path, handler string) *RouteAction {
	return group.Route(http.MethodOptions, path, handler)
}

func (group *RouteController) TRACE(path, handler string) *RouteAction {
	return group.Route(http.MethodTrace, path, handler)
}

func (group *RouteController) GET(path, handler string) *RouteAction {
	return group.Route(http.MethodGet, path, handler)
}

func (group *RouteController) POST(path, handler string) *RouteAction {
	return group.Route(http.MethodPost, path, handler)
}

func (group *RouteController) PUT(path, handler string) *RouteAction {
	return group.Route(http.MethodPut, path, handler)
}

func (group *RouteController) PATCH(path, handler string) *RouteAction {
	return group.Route(http.MethodPatch, path, handler)
}

func (group *RouteController) DELETE(path, handler string) *RouteAction {
	return group.Route(http.MethodDelete, path, handler)
}
