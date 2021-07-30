package context

import (
	"time"

	"github.com/gin-gonic/gin"
)

const currentTimeKey = "currentTime"

// context.Valueから現在時刻を取得する関数
func Now(c *gin.Context) (time.Time, bool) {
	ct, isExist := c.Get(currentTimeKey)
	return ct.(time.Time), isExist
}

// context.Valueに現在時刻を格納する関数
func SetNow(c *gin.Context) {
	c.Set(currentTimeKey, time.Now().UTC())
}

// テスト用にmockしたいtime.Timeをcontext.Valueに格納する関数
func MockNow(c *gin.Context, mockTime time.Time) {
	c.Set(currentTimeKey, mockTime)
}
