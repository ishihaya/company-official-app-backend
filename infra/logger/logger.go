package logger

import (
	"sync"

	"go.uber.org/zap"
)

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

var Logging logging
var zapLogger *zap.Logger
var sugar *zap.SugaredLogger
var once sync.Once

func (log *loggingImpl) Debug(msg string, field ...zap.Field) {
	zapLogger.Debug(msg, field...)
}

func (log *loggingImpl) Error(msg string, field ...zap.Field) {
	zapLogger.Error(msg, field...)
}

func (log *loggingImpl) Info(msg string, field ...zap.Field) {
	zapLogger.Info(msg, field...)
}

func (log *loggingImpl) Warn(msg string, field ...zap.Field) {
	zapLogger.Warn(msg, field...)
}

func (log *loggingImpl) Fatal(msg string, field ...zap.Field) {
	zapLogger.Fatal(msg, field...)
}

func (log *loggingImpl) Debugf(template string, args ...interface{}) {
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Debugf(template, args)
}

func (log *loggingImpl) Errorf(template string, args ...interface{}) {
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Errorf(template, args)
}

func (log *loggingImpl) Infof(template string, args ...interface{}) {
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Infof(template, args)
}

func (log *loggingImpl) Warnf(template string, args ...interface{}) {
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Warnf(template, args)
}

func (log *loggingImpl) Fatalf(template string, args ...interface{}) {
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Fatalf(template, args)
}

func New() {}