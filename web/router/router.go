package router

import (
	libctx "context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/contract"
	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
	"github.com/wardonne/gopi/web/middleware/cors"
)

var ErrValidateEngineEmpty = errors.New("validate engine is nil, please call SetValidateEngine to set it first")

// Router http router
type Router struct {
	*RouteGroup
	HTTPRouter     *httprouter.Router
	routes         []IRoute
	validateEngine validation.Engine
	corsOptions    *cors.CORSOptions
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

// SetCORS sets cors configs
func (router *Router) SetCORS(options *cors.CORSOptions) *Router {
	router.corsOptions = options
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
		switch v := i.(type) {
		case error:
			resp := handler.Render(r, v)
			resp.Send(w, r)
		case string:
			resp := handler.Render(r, errors.New(v))
			resp.Send(w, r)
		case fmt.Stringer:
			resp := handler.Render(r, errors.New(v.String()))
			resp.Send(w, r)
		default:
			defaultErrorHandler(w, r, i)
		}
	}
	return router
}

// Registe used to registe routes by custom callback
func (router *Router) Registe(register Register) {
	register(router)
}

// Run starts the http server
//
// # NOTICE: should call Prepare first before calling Run
func (router *Router) Run(addr string) error {
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
	if router.corsOptions != nil {
		router.HTTPRouter.HandleOPTIONS = true
		router.HTTPRouter.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Credentials", fmt.Sprintf("%v", router.corsOptions.AllowCredentials))
			if len(router.corsOptions.AllowHeaders) > 0 {
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(router.corsOptions.AllowHeaders, ","))
			}
			if len(router.corsOptions.AllowMethods) > 0 {
				w.Header().Set("Access-Control-Request-Method", strings.Join(router.corsOptions.AllowMethods, ","))
			} else {
				w.Header().Set("Access-Control-Request-Method", w.Header().Get("Allow"))
			}
			if len(router.corsOptions.AllowOrigin) > 0 {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			} else {
				w.Header().Set("Access-Control-Allow-Origin", strings.Join(router.corsOptions.AllowOrigin, ","))
			}
			if len(router.corsOptions.ExposeHeaders) > 0 {
				w.Header().Set("Access-Control-Expose-Headers", strings.Join(router.corsOptions.ExposeHeaders, ","))
			}
			w.WriteHeader(http.StatusNoContent)
		})
	}
	return http.ListenAndServe(addr, router.HTTPRouter)
}
