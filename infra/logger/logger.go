package logger

import "go.uber.org/zap"

type logging interface {
	Debug(msg string, field ...zap.Field)
	Error(msg string, field ...zap.Field)
	Info(msg string, field ...zap.Field)
	Warn(msg string, field ...zap.Field)
	Fatal(msg string, field ...zap.Field)
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

type loggingImpl struct{}
