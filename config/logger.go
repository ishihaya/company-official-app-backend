package config

import (
	"os"

	"github.com/ishihaya/company-official-app-backend/infra/logger"
)

func Log() (logEnv string, logLevel string, logEncoding string) {
	appEnv := os.Getenv("APP_ENV")
	logLevel = os.Getenv("LOG_LEVEL")
	logEncoding = os.Getenv("LOG_ENCODING")
	switch appEnv {
	case "prd":
		logEnv = logger.LogEnvProduction
		return
	default:
		logEnv = logger.LogEnvDevelopment
		return
	}
}
