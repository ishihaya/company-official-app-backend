package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
	}))
	fmt.Println("Listening 8080")
	http.ListenAndServe(":8080", nil)
}
