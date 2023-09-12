package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

var jsonFormatter = NewJSONFormatter().
	TimeFormat(time.DateTime).
	DisableHTMLEscape().
	FieldKeyMsg("message").
	FieldKeyTime("timestamps")

type JSONFormatter struct {
	logrus.JSONFormatter
}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{
		JSONFormatter: logrus.JSONFormatter{
			FieldMap: make(logrus.FieldMap),
		},
	}
}

func (formatter *JSONFormatter) TimeFormat(layout string) *JSONFormatter {
	formatter.TimestampFormat = layout
	return formatter
}

func (formatter *JSONFormatter) DisableTimestamp() *JSONFormatter {
	formatter.JSONFormatter.DisableTimestamp = true
	return formatter
}

func (formatter *JSONFormatter) DisableHTMLEscape() *JSONFormatter {
	formatter.JSONFormatter.DisableHTMLEscape = true
	return formatter
}

func (formatter *JSONFormatter) DataKey(key string) *JSONFormatter {
	formatter.JSONFormatter.DataKey = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyMsg(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyMsg] = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyLevel(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyLevel] = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyTime(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyTime] = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyLogrusError(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyLogrusError] = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyFunc(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyFunc] = key
	return formatter
}

func (formatter *JSONFormatter) FieldKeyFile(key string) *JSONFormatter {
	formatter.JSONFormatter.FieldMap[logrus.FieldKeyFile] = key
	return formatter
}
