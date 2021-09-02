package main

import (
	"github.com/ishihaya/company-official-app-backend/common/env"
	"github.com/ishihaya/company-official-app-backend/interface/router"
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
