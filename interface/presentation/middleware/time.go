package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
)

func CurrentTime(c *gin.Context) {
	contextgo.SetNow(c)
	c.Next()
}
