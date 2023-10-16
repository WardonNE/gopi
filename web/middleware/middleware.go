package middleware

import (
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
)

// IMiddleware is an alias of [pipeline.Handler][*[context.Request], [context.IResponse]]
type IMiddleware = pipeline.Handler[*context.Request, context.IResponse]
