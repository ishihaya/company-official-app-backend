package main

import (
	"github.com/ishihaya/company-official-app-backend/interface/presentation/router"
	"github.com/ishihaya/company-official-app-backend/pkg/env"
)

func main() {
	r := router.New()
	r.HealthCheck()
	if !env.IsProduction() {
		r.Swagger()
	}
	r.Routes()
	r.RunServer(env.PORT())
}
