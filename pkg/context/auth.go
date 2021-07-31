package context

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

const authIDKey = "authID"

// SetAuthID - AuthIDをcontextにセット
func SetAuthID(c *gin.Context, authID string) {
	c.Set(authIDKey, authID)
}

// GetAuthID - contextからAuthIDを取得
func GetAuthID(c *gin.Context) (string, error) {
	authID, isExist := c.Get(authIDKey)
	if !isExist {
		return "", xerrors.New("authID not set")
	}
	return authID.(string), nil
}
