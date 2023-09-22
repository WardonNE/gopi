package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type Router struct {
	*RouteGroup
	HTTPRouter *httprouter.Router
	routes     []IRoute
}

func New() *Router {
	router := &Router{
		RouteGroup: &RouteGroup{
			Prefix:      "/",
			Middlewares: make([]middleware.IMiddleware, 0),
			RouteGroups: make([]IRouteGroup, 0),
			Routes:      make([]*RouteHandler, 0),
		},
		HTTPRouter: httprouter.New(),
	}
	return router
}

func (router *Router) Run(addr ...string) error {
	if router.routes == nil {
		router.routes = router.List()
	}
	for _, route := range router.routes {
		router.HTTPRouter.Handle(route.Method(), route.Path(), func(route IRoute) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				request := context.NewRequest(r, p)
				resp := route.HandleRequest(request)
				resp.Send(w, r)
			}
		}(route))
	}
	return http.ListenAndServe(addr[0], router.HTTPRouter)
}
