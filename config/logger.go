package config

import (
	"os"

	"github.com/ishihaya/company-official-app-backend/infra/logger"
)

func Log() (logEnv string, logLevel string, logEncoding string) {
	logLevel = os.Getenv("LOG_LEVEL")
	logEncoding = os.Getenv("LOG_ENCODING")
	if IsProduction() {
		logEnv = logger.LogEnvProduction
		return
	}
	logEnv = logger.LogEnvDevelopment
	return
}
