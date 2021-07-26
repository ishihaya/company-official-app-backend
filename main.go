package main

import (
	"fmt"
	"log"
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
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
}
