package operator

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
)

type appIDOperator struct{}

func NewAppIDOperator() operator.AppIDOperator {
	return &appIDOperator{}
}

func (a *appIDOperator) Generate(t time.Time) (entity.AppID, error) {
	// _, err := ulidgo.Generate(t)
	// if err != nil {
	// 	return "", nil
	// }
	// var id entity.AppID = ulid
	return "", nil
	// timeStamp := ulid.Timestamp(t)
	// entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// ulid, err := ulid.New(timeStamp, entropy)
	// if err != nil {
	// 	return "", err
	// }
	// return ulid.String(), nil
}
