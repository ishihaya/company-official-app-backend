package logging

import (
	"sync"

	"github.com/ishihaya/cloudlog"
	"github.com/ishihaya/company-official-app-backend/pkg/env"
)

// You can define methods that you only use.
type Log interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

type log struct {
	*cloudlog.Logger
}

var sharedInstance Log
var once sync.Once

// You should call this if you use logger.
func New() Log {
	once.Do(func() {
		sharedInstance = newInstance()
	})
	return sharedInstance
}

func newInstance() Log {
	var logger *cloudlog.Logger
	var err error
	// serviceName is displayed in Error Reporting.
	serviceName := env.ServiceName()
	switch env.ENV() {
	// List runnning environments on cloud, such as GCP.
	case env.Production, env.Development:
		logger, err = cloudlog.NewCloudLogger(
			cloudlog.NeedErrorReporting(true),
			cloudlog.ServiceName(serviceName),
		)
	default:
		logger, err = cloudlog.NewLocalLogger()
	}
	if err != nil {
		panic(err)
	}
	return &log{logger}
}
