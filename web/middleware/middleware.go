package middleware

import (
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
)

type IMiddleware = pipeline.IPipe[*context.Request, context.IResponse]
