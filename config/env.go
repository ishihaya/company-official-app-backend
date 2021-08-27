package config

import (
	"fmt"
	"os"
)

type Environment int

const (
	Production Environment = iota
	Development
	Local
)

func ENV() Environment {
	env := os.Getenv("APP_ENV")
	switch env {
	case "prd":
		return Production
	case "dev":
		return Development
	case "local":
		return Local
	default:
		panic(fmt.Sprintf("does not match environment %s", env))
	}
}

func IsProduction() bool {
	return ENV() == Production
}
