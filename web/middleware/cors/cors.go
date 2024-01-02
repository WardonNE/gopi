package cors

import (
	"fmt"
	"strings"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type CORSOptions struct {
	AllowCredentials bool
	AllowHeaders     []string
	AllowOrigin      []string
	AllowMethods     []string
	ExposeHeaders    []string
}

func New(options CORSOptions) middleware.IMiddleware {
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		resp := next(request)
		resp.Headers().Set("Access-Control-Allow-Credentials", fmt.Sprintf("%v", options.AllowCredentials))
		if len(options.AllowHeaders) > 0 {
			resp.Headers().Set("Access-Control-Allow-Headers", strings.Join(options.AllowHeaders, ","))
		}
		if len(options.AllowOrigin) > 0 {
			resp.Headers().Set("Access-Control-Allow-Origin", strings.Join(options.AllowOrigin, ","))
		} else {
			resp.Headers().Set("Access-Control-Allow-Origin", string(*request.Header("Origin")))
		}
		if len(options.AllowMethods) > 0 {
			resp.Headers().Set("Access-Control-Request-Method", strings.Join(options.AllowMethods, ","))
		} else {
			resp.Headers().Set("Access-Control-Request-Method", resp.Headers().Get("Allow"))
		}
		if len(options.ExposeHeaders) > 0 {
			resp.Headers().Set("Access-Control-Expose-Headers", strings.Join(options.ExposeHeaders, ","))
		}
		return resp
	}
}
