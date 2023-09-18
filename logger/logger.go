package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New(opts ...Option) (*Logger, error) {
	log := new(Logger)
	log.Logger = logrus.New()
	for _, option := range opts {
		if err := option(log); err != nil {
			return nil, err
		}
	}
	log.Out = os.Stdout
	return log, nil
}

func Must(logger *Logger, err error) *Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

func (l *Logger) Trace(message string, context map[string]any) {
	l.WithFields(context).Traceln(message)
}

func (l *Logger) Debug(message string, context map[string]any) {
	l.WithFields(context).Debugln(message)
}

func (l *Logger) Info(message string, context map[string]any) {
	l.WithFields(context).Infoln(message)
}

func (l *Logger) Warn(message string, context map[string]any) {
	l.WithFields(context).Warnln(message)
}

func (l *Logger) Error(message string, context map[string]any) {
	l.WithFields(context).Errorln(message)
}

func (l *Logger) Fatal(message string, context map[string]any) {
	l.WithFields(context).Fatalln(message)
}

func (l *Logger) Panic(message string, context map[string]any) {
	l.WithFields(context).Panicln(message)
}
