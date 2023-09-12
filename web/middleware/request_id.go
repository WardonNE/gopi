package middleware

import (
	uuid "github.com/satori/go.uuid"
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/web"
)

type RequestID struct {
	headerKey string
	generator func() string
}

func NewRequestID() *RequestID {
	return &RequestID{
		headerKey: "X-Request-ID",
		generator: func() string {
			return uuid.NewV4().String()
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
