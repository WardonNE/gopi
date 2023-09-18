package logger

import (
	"github.com/sirupsen/logrus"
)

var logger = Must(New(
	WithFormatter(jsonFormatter),
	WithLevel(logrus.InfoLevel),
))

func SetDefault(instance *Logger) {
	logger = instance
}

func Trace(message string, context map[string]any) {
	logger.Trace(message, context)
}

func Debug(message string, context map[string]any) {
	logger.Debug(message, context)
}

func Info(message string, context map[string]any) {
	logger.Info(message, context)
}

func Warn(message string, context map[string]any) {
	logger.Warn(message, context)
}

func Error(message string, context map[string]any) {
	logger.Error(message, context)
}

func Fatal(message string, context map[string]any) {
	logger.Fatal(message, context)
}

func Panic(message string, context map[string]any) {
	logger.Panic(message, context)
}
