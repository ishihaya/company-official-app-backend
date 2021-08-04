package operator

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
)

type AuthOperator interface {
	FindByToken(ctx context.Context, token string) (*entity.Auth, error)
}
