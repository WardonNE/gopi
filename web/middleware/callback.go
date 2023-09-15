package middleware

import (
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web"
)

type callback struct {
	handler pipeline.PipeHandler[*context.Request, context.IResponse]
}

func (m *callback) Handle(request *context.Request, next web.Handler) context.IResponse {
	return m.handler(request, next)
}

func Callback(handler pipeline.PipeHandler[*context.Request, context.IResponse]) *callback {
	return &callback{
		handler: handler,
	}
}