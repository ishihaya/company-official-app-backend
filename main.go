package main

import (
	"fmt"

	"github.com/ishihaya/company-official-app-backend/interface/presentation/router"
	"github.com/ishihaya/company-official-app-backend/pkg/env"
)

func main() {
	fmt.Println("------main 0:50----------")
	r := router.New()
	r.HealthCheck()
	if !env.IsProduction() {
		r.Swagger()
	}
	r.Routes()
	r.RunServer(env.PORT())
}
