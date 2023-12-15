package logger

import (
	"net/url"
	"time"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// Information contains request and response informations
type Information struct {
	Status    int        `json:"status"`
	Method    string     `json:"method"`
	Path      string     `json:"path"`
	Query     url.Values `json:"query"`
	IP        string     `json:"ip"`
	UserAgent string     `json:"user_agent"`
	Latency   string     `json:"latency"`
}

// New creates a logger middleware from specific write function
func New(writer func(Information)) middleware.IMiddleware {
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		s := time.Now()
		resp := next(request)
		writer(Information{
			Status:    resp.StatusCode(),
			Method:    request.Method(),
			Path:      request.Path(),
			Query:     request.Request.URL.Query(),
			IP:        request.ClientIP(),
			UserAgent: request.Header("User-Agent", "").String(),
			Latency:   time.Since(s).String(),
		})
		return resp
	}
}
