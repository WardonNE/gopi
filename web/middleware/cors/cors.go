package cors

import (
	"fmt"
	"strings"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/support/maps"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type CorsOptions struct {
	AllowCredentials bool
	AllowHeaders     []string
	AllowOrigin      string
	AllowMethods     []string
	ExposeHeaders    []string
}

func New(options CorsOptions) middleware.IMiddleware {
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		headers := maps.NewHashMap[string, string]()
		headers.Set("Access-Control-Allow-Credentials", fmt.Sprintf("%v", options.AllowCredentials))
		headers.Set("Access-Control-Allow-Headers", strings.Join(options.AllowHeaders, ","))
		headers.Set("Access-Control-Allow-Origin", options.AllowOrigin)
		headers.Set("Access-Control-Request-Method", strings.Join(options.AllowMethods, ","))
		if len(options.ExposeHeaders) > 0 {
			headers.Set("Access-Control-Expose-Headers", strings.Join(options.ExposeHeaders, ","))
		}
		resp := next(request)
		resp.SetHeaders(headers.ToMap())
		return resp
	}
}
