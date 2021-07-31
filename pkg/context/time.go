package context

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

const currentTimeKey = "currentTime"

func Now(c *gin.Context) (time.Time, error) {
	now, isExist := c.Get(currentTimeKey)
	if !isExist {
		return time.Time{}, xerrors.New("authID not set")
	}
	return now.(time.Time), nil
}

func SetNow(c *gin.Context) {
	c.Set(currentTimeKey, time.Now().UTC())
}

// テスト用にmockしたいtime.Timeをcontext.Valueに格納する関数
func MockNow(c *gin.Context, mockTime time.Time) {
	c.Set(currentTimeKey, mockTime)
}
