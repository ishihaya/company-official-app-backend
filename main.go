package main

import (
	"github.com/ishihaya/company-official-app-backend/config"
	"github.com/ishihaya/company-official-app-backend/infra/logger"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/router"
)

func main() {
	logger.New(config.Log())
	r := router.New()
	r.Routes()
	if config.IsLocal() {
		r.LoadSwagger()
	}
	r.RunServer()
}
