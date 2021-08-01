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
	// TODO
	return "", nil
}
