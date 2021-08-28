package factory

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
