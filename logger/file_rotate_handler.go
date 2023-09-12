package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type FileRotateHandler struct {
	MaxAge        *time.Duration
	RotationTime  *time.Duration
	RotationSize  *int64
	RotationCount *uint
	LinkName      *string
	Clock         rotatelogs.Clock
	Location      *time.Location
}

func NewFileRotateHandler() *FileRotateHandler {
	return new(FileRotateHandler)
}

func (f *FileRotateHandler) WithMaxAge(maxAge time.Duration) *FileRotateHandler {
	f.MaxAge = &maxAge
	return f
}

func (f *FileRotateHandler) WithRotationTime(rotationTime time.Duration) *FileRotateHandler {
	f.RotationTime = &rotationTime
	return f
}

func (f *FileRotateHandler) WithRotationSize(rotationSize int64) *FileRotateHandler {
	f.RotationSize = &rotationSize
	return f
}

func (f *FileRotateHandler) WithLinkName(name string) *FileRotateHandler {
	f.LinkName = &name
	return f
}

func (f *FileRotateHandler) WithClock(clock rotatelogs.Clock) *FileRotateHandler {
	f.Clock = clock
	return f
}

func (f *FileRotateHandler) WithLocation(loc *time.Location) *FileRotateHandler {
	f.Location = loc
	return f
}

func (f *FileRotateHandler) Handle(logger *Logger) error {
	options := make([]rotatelogs.Option, 0)
	if f.MaxAge != nil {
		options = append(options, rotatelogs.WithMaxAge(*f.MaxAge))
	} else {
		options = append(options, rotatelogs.WithMaxAge(720*time.Hour))
	}
	if f.LinkName != nil {
		options = append(options, rotatelogs.WithLinkName(*f.LinkName))
	}
	if f.RotationTime != nil {
		options = append(options, rotatelogs.WithRotationTime(*f.RotationTime))
	} else {
		options = append(options, rotatelogs.WithRotationTime(24*time.Hour))
	}
	if f.RotationSize != nil {
		options = append(options, rotatelogs.WithRotationSize(*f.RotationSize))
	}
	if f.RotationCount != nil {
		options = append(options, rotatelogs.WithRotationCount(*f.RotationCount))
	}
	if f.Clock != nil {
		options = append(options, rotatelogs.WithClock(f.Clock))
	}
	if f.Location != nil {
		options = append(options, rotatelogs.WithLocation(f.Location))
	} else {
		options = append(options, rotatelogs.WithLocation(time.Local))
	}
	w, err := rotatelogs.New("", options...)
	if err != nil {
		return err
	}
	wmap := lfshook.WriterMap{
		logrus.TraceLevel: w,
		logrus.DebugLevel: w,
		logrus.InfoLevel:  w,
		logrus.WarnLevel:  w,
		logrus.ErrorLevel: w,
		logrus.FatalLevel: w,
		logrus.PanicLevel: w,
	}
	logger.AddHook(lfshook.NewHook(wmap, logger.Formatter))
	return nil
}
