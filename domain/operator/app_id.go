package operator

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
)

type AppIDOperator interface {
	Generate(t time.Time) (entity.AppID, error)
}
