package web

import (
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
)

type Handler = pipeline.Handler[*context.Request, context.IResponse]
