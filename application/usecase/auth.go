package usecase

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"golang.org/x/xerrors"
)

type Auth interface {
	Get(ctx context.Context, token string) (*entity.Auth, error)
}

type auth struct {
	authOperator operator.Auth
}

func NewAuth(authOperator operator.Auth) Auth {
	return &auth{
		authOperator: authOperator,
	}
}

func (a *auth) Get(ctx context.Context, token string) (*entity.Auth, error) {
	auth, err := a.authOperator.FindByToken(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return auth, nil
}
