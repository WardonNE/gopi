package web

import (
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/pipeline"
)

type Handler = pipeline.Handler[*context.Request, context.IResponse]
