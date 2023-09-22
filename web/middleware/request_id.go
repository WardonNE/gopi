package middleware

import (
	"github.com/google/uuid"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/context"
)

type RequestID struct {
	headerKey string
	generator func() string
}

func NewRequestID() *RequestID {
	return &RequestID{
		headerKey: "X-Request-ID",
		generator: func() string {
			return uuid.New().String()
		},
	}
}

func (r *RequestID) WithGenerator(generator func() string) *RequestID {
	r.generator = generator
	return r
}

func (r *RequestID) WithCustomHeaderKey(key string) *RequestID {
	r.headerKey = key
	return r
}

func (r *RequestID) Handle(request *context.Request, next web.Handler) context.IResponse {
	id := r.generator()
	resp := next(request)
	if resp != nil && !resp.HasHeader(r.headerKey) {
		resp.SetHeader(r.headerKey, id)
	}
	return resp
}
