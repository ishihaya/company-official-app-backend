package main

import (
	"fmt"
	"net/http"

	"github.com/ishihaya/company-official-app-backend/config"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
	}))
	port := config.PORT()
	fmt.Println("Listening :", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
