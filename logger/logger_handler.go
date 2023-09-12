package logger

type LoggerHandler interface {
	Handle(logger *Logger) error
}
