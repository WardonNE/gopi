package router

import (
	libctx "context"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/contract"
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

var ErrValidateEngineEmpty = errors.New("validate engine is nil, please call SetValidateEngine to set it first")

// Router http router
type Router struct {
	*RouteGroup
	HTTPRouter     *httprouter.Router
	routes         []IRoute
	validateEngine validation.Engine
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
	router.HTTPRouter.PanicHandler = defaultErrorHandler
	return router
}

// SetValidateEngine sets custom validate engine
func (router *Router) SetValidateEngine(ve validation.Engine) *Router {
	router.validateEngine = ve
	return router
}

// SetErrorHandler sets custom error handler
func (router *Router) SetErrorHandler(handler contract.ErrorHandler) *Router {
	router.HTTPRouter.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		if e, ok := i.(error); ok {
			resp := handler.Render(r, e)
			resp.Send(w, r)
		} else {
			defaultErrorHandler(w, r, i)
		}
	}
	return router
}

// Registe used to registe routes by custom callback
func (router *Router) Registe(register Register) {
	register(router)
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
