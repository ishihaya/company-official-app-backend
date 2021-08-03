package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	AuthAPI(c *gin.Context)
}

type authMiddleware struct {}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

// AuthAPI - 認証API
func (a *authMiddleware) AuthAPI(c *gin.Context) {
	// TODO
	// 先にentityから書く
}
