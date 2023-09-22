package router

import "github.com/wardonne/gopi/web/middleware"

type IRouteGroup interface {
	List() []IRoute
	Use(middlewares ...middleware.IMiddleware)
}
