package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// to connect mysql db
	_ "github.com/go-sql-driver/mysql"
	"github.com/ishihaya/company-official-app-backend/config"
)

func main() {
	dsn := config.DSN()
	_, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v", err)
	}
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
