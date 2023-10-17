package router

import (
	libctx "context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
	"github.com/wardonne/gopi/web/middleware/validate"
)

// Router http router
type Router struct {
	*RouteGroup
	HTTPRouter     *httprouter.Router
	routes         []IRoute
	validateEngine validate.ValidationEngine
}

// New creates a new [Router] instance
func New() *Router {
	router := &Router{
		RouteGroup: &RouteGroup{
			Prefix:      "/",
			Middlewares: list.NewArrayList[middleware.IMiddleware](),
			RouteGroups: make([]IRouteGroup, 0),
			Routes:      make([]*RouteHandler, 0),
		},
		HTTPRouter: httprouter.New(),
	}
	router.router = router
	return router
}

// SetValidateEngine sets custom validate engine
func (router *Router) SetValidateEngine(ve validate.ValidationEngine) *Router {
	router.validateEngine = ve
	return router
}

// Prepare registers all routes into http server
func (router *Router) Prepare() {
	if router.routes == nil {
		router.routes = router.List()
	}
	for _, route := range router.routes {
		router.HTTPRouter.Handle(route.Method(), route.Path(), func(route IRoute) httprouter.Handle {
			return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				ctx := r.Context()
				ctx = libctx.WithValue(ctx, httprouter.ParamsKey, p)
				request := context.NewRequest(r, p)
				resp := route.HandleRequest(request)
				resp.Send(w, r)
			}
		}(route))
	}
}

// Run starts the http server
//
// # NOTICE: should call Prepare first before calling Run
func (router *Router) Run(addr string) error {
	return http.ListenAndServe(addr, router.HTTPRouter)
}
