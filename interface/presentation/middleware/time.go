package middleware

import (
	"net/http"

	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
)

func CurrentTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextgo.SetCurrentTime(r.Context())
		next.ServeHTTP(w, r)
	})
}
