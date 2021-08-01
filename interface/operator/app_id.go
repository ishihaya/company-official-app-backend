package operator

import (
	"math/rand"
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/oklog/ulid/v2"
)

type appIDOperator struct{}

func NewAppIDOperator() operator.AppIDOperator {
	return &appIDOperator{}
}

// Generate - アプリケーションの識別子生成関数
// NOTE: Unit Testは行わない
func (a *appIDOperator) Generate(t time.Time) (entity.AppID, error) {
	timeStamp := ulid.Timestamp(t)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid, err := ulid.New(timeStamp, entropy)
	if err != nil {
		return "", err
	}
	id := entity.AppID(ulid.String())
	return id, nil
}
