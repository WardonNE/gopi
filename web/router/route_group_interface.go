package router

import "github.com/wardonne/gopi/web/middleware"

// IRouteGroup interface of route group
type IRouteGroup interface {
	List() []IRoute
	Use(middlewares ...middleware.IMiddleware)
}
