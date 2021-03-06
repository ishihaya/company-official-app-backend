package middleware

import (
	"net/http"

	"github.com/ishihaya/company-official-app-backend/interface/pkg/contextgo"
)

func CurrentTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(contextgo.SetCurrentTime(r.Context()))
		next.ServeHTTP(w, r)
	})
}
