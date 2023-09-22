package middleware

import (
	"time"

	"github.com/wardonne/gopi/logger"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/context"
)

type Logger struct {
	logger *logger.Logger
}

func NewLogger(options *logger.Logger) *Logger {
	log := &Logger{
		logger: options,
	}
	return log
}

func (l *Logger) Handle(request *context.Request, next web.Handler) context.IResponse {
	startTime := time.Now()
	resp := next(request)
	l.logger.WithFields(map[string]any{
		"status":     resp.StatusCode(),
		"method":     request.Method(),
		"path":       request.Path(),
		"query":      request.Request.URL.Query(),
		"ip":         request.ClientIP(),
		"user-agent": request.Header("User-Agent", "").ToString(),
		"latency":    time.Since(startTime).String(),
	}).Info()
	return resp
}
