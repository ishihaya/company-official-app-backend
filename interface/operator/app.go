package operator

import (
	"math/rand"
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/oklog/ulid/v2"
	"golang.org/x/xerrors"
)

type appID struct{}

func NewAppID() operator.AppID {
	return &appID{}
}

// Generate - アプリケーションで使用するIDを生成する関数
func (a *appID) Generate(t time.Time) (id entity.AppID, err error) {
	timeStamp := ulid.Timestamp(t)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid, err := ulid.New(timeStamp, entropy)
	if err != nil {
		return "", xerrors.Errorf(": %w", err)
	}
	id = entity.AppID(ulid.String())
	return
}
