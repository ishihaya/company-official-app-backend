package operator

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
)

type AppID interface {
	Generate(t time.Time) (entity.AppID, error)
}
