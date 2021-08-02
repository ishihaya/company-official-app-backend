package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Auth(c *gin.Context)
}

type authMiddleware struct {}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

// TODO 命名について考える Authだとentityと重複する
func (a *authMiddleware) Auth(c *gin.Context) {
	// TODO
	// 先にentityから書く
}
