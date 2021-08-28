package factory

import (
	"encoding/json"
	"net/http"

	"github.com/ishihaya/company-official-app-backend/pkg/logging"
)

func JSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logging.GetInstance().Errorf("failed to encode json: %+v", err)
	}
}
