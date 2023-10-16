package reqid

import (
	"github.com/google/uuid"
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

const defaultHeaderKey = "X-Request-ID"

func defaultGenerator() string {
	return uuid.NewString()
}

// Default returns default request id middleware
func Default() middleware.IMiddleware {
	return New(defaultHeaderKey, defaultGenerator)
}

// New creates a new request id middleware
func New(key string, generator func() string) middleware.IMiddleware {
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		id := generator()
		resp := next(request)
		if resp != nil && !resp.HasHeader(key) {
			resp.SetHeader(key, id)
		}
		return resp
	}
}
