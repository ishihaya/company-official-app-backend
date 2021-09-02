package env

import (
	"os"
	"strconv"
)

func PORT() int {
	portStr, isExist := os.LookupEnv("PORT")
	if !isExist {
		return 8080
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	return port
}
