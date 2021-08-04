package main

import (
	"github.com/ishihaya/company-official-app-backend/config"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/router"
	"github.com/ishihaya/company-official-app-backend/pkg/logger"
)

func main() {
	logger.New(config.Log())
	r := router.New()
	r.Routes()
	r.RunServer()
}
