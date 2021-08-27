package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/pkg/gincontext"
)

func CurrentTime(c *gin.Context) {
	gincontext.SetNow(c)
	c.Next()
}
