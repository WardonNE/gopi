package logger

import (
	"github.com/sirupsen/logrus"
)

type Option func(logger *Logger) error

func WithLevel(level logrus.Level) Option {
	return func(logger *Logger) error {
		logger.SetLevel(level)
		return nil
	}
}

func WithFormatter(formatter logrus.Formatter) Option {
	return func(logger *Logger) error {
		logger.SetFormatter(formatter)
		return nil
	}
}

func WithReportCaller(enable bool) Option {
	return func(logger *Logger) error {
		logger.SetReportCaller(enable)
		return nil
	}
}

func WithHandler(handler LoggerHandler) Option {
	return func(logger *Logger) error {
		return handler.Handle(logger)
	}
}
