package middleware

import (
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/pipeline"
)

type IMiddleware = pipeline.IPipe[*context.Request, context.IResponse]
