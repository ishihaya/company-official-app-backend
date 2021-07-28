package config

import "os"

func PORT() string {
	port, isExist := os.LookupEnv("PORT")
	if !isExist {
		port = "8080"
	}
	return port
}
