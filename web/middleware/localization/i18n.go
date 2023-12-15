package localization

import (
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// Default creates a localicazion middleware which get current locale from Accept-Language header
func Default() middleware.IMiddleware {
	return New(func(request *context.Request) string {
		return request.Header("Accept-Language", "en").String()
	})
}

// New creates a new localization middleware
func New(localeGetter func(request *context.Request) string) middleware.IMiddleware {
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		locale := localeGetter(request)
		request.Set("language", locale)
		return next(request)
	}
}
