package logger

import (
	"fmt"
	"log"
	"sync"

	stackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

const LogEnvDevelopment = "development"
const LogEnvProduction = "production"

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

func New(logEnv string, logLevel string, logEncoding string) {
	once.Do(func() {
		level := unmarshalLogLevel(logLevel)
		switch logEnv {
		case LogEnvDevelopment:
			zapLogger = newDevelopmentLogger(level, logEncoding)
		case LogEnvProduction:
			zapLogger = newProductionLogger(level, logEncoding)
		default:
			log.Fatal(fmt.Sprintf("does not match logEnv: [%s]", logEnv))
		}
		sugar = zapLogger.Sugar()
		Logging = new(loggingImpl)
	})
}

// see: https://github.com/uber-go/zap/blob/425214515ff452748375576b20c82524849177c6/zapcore/level.go#L126-L146
func unmarshalLogLevel(text string) zapcore.Level {
	switch text {
	case "debug", "DEBUG":
		return zap.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		return zap.InfoLevel
	case "warn", "WARN":
		return zap.WarnLevel
	case "error", "ERROR":
		return zap.ErrorLevel
	case "dpanic", "DPANIC":
		return zap.DPanicLevel
	case "panic", "PANIC":
		return zap.PanicLevel
	case "fatal", "FATAL":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func newDevelopmentLogger(logLevel zapcore.Level, logEncoding string) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.Encoding = logEncoding
	config.EncoderConfig = stackdriver.EncoderConfig
	logger, err := config.Build()
	if err != nil {
		log.Fatal(fmt.Sprintf("init development logging configs error: %v", err))
	}
	return logger
}

func newProductionLogger(logLevel zapcore.Level, logEncoding string) *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.Sampling.Initial = 1
	config.Sampling.Thereafter = 1
	config.Encoding = logEncoding
	config.EncoderConfig = stackdriver.EncoderConfig
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatal(fmt.Sprintf("init production logging configs error: %v", err))
	}
	return logger
}
