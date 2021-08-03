package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/request"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/logger"
)

type AuthMiddleware interface {
	AuthAPI(c *gin.Context)
}

type authMiddleware struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthMiddleware(authUsecase usecase.AuthUsecase) AuthMiddleware {
	return &authMiddleware{
		authUsecase: authUsecase,
	}
}

// AuthAPI - 認証API
func (a *authMiddleware) AuthAPI(c *gin.Context) {
	req := &request.AuthAPI{
		IDToken: c.Request.Header.Get("Authorization"),
	}
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Logging.Warnf("request not valid: %+v", err)
		c.JSON(http.StatusBadRequest, apperror.ErrValidation.Error())
		return
	}

	auth, err := a.authUsecase.Get(req.IDToken)
	if err != nil {
		logger.Logging.Errorf("failed to get auth: %+v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
	}

	contextgo.SetAuthID(c, auth.ID)

	c.Next()
}
