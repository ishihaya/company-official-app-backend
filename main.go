package main

import (
	"github.com/ishihaya/company-official-app-backend/interface/presentation/router"
)

func main() {
	r := router.New()
	r.Routes()
	r.RunServer()
}
