package usecase

import (
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"golang.org/x/xerrors"
)

type AuthUsecase interface {
	Get(token string) (*entity.Auth, error)
}

type authUsecase struct {
	authOperator operator.AuthOperator
}

func NewAuthUsecase(authOperator operator.AuthOperator) AuthUsecase {
	return &authUsecase{}
}

func (a *authUsecase) Get(token string) (*entity.Auth, error) {
	auth, err := a.authOperator.FindByToken(token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return auth, nil
}
