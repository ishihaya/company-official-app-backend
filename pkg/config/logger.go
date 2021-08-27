package config

import (
	"os"
)

func ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}
