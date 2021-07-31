package contextgo

import (
	"net/http/httptest"
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
func SetMockTime(c *gin.Context, mockTime time.Time) {
	c.Set(currentTimeKey, mockTime)
}

// テストで同時刻を扱いたい時にcontext.Valueに現在時刻を格納して取得する関数
func GetMockNow() (time.Time, error) {
	gin.SetMode(gin.ReleaseMode)
	timeContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	SetNow(timeContext)
	mockTime, err := Now(timeContext)
	if err != nil {
		return time.Time{}, err
	}
	return mockTime, nil
}
