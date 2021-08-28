package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/factory"
	"github.com/ishihaya/company-official-app-backend/pkg/logging"
)

type AuthMiddleware interface {
	AuthAPI(next http.Handler) http.Handler
}

type authMiddleware struct {
	authUsecase usecase.AuthUsecase
	logging     logging.Log
}

func NewAuthMiddleware(
	authUsecase usecase.AuthUsecase,
	logging logging.Log,
) AuthMiddleware {
	return &authMiddleware{
		authUsecase: authUsecase,
		logging:     logging,
	}
}

// AuthAPI - 認証API
func (a *authMiddleware) AuthAPI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// TODO validate
		typ := "Bearer"
		typPrefix := fmt.Sprintf("%s ", typ)
		idToken := strings.Replace(authHeader, typPrefix, "", 1)
		ctx := r.Context()
		auth, err := a.authUsecase.Get(ctx, idToken)
		if err != nil {
			a.logging.Errorf("failed to get auth: %+v", err)
			factory.JSON(w, http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
			return
		}
		r = r.WithContext(contextgo.SetAuthID(ctx, auth.ID))
		next.ServeHTTP(w, r)
	})
}
