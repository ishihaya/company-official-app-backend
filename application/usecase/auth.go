package usecase

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"golang.org/x/xerrors"
)

type AuthUsecase interface {
	Get(ctx context.Context, token string) (*entity.Auth, error)
}

type authUsecase struct {
	authOperator operator.AuthOperator
}

func NewAuthUsecase(authOperator operator.AuthOperator) AuthUsecase {
	return &authUsecase{
		authOperator: authOperator,
	}
}

func (a *authUsecase) Get(ctx context.Context, token string) (*entity.Auth, error) {
	auth, err := a.authOperator.FindByToken(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return auth, nil
}
