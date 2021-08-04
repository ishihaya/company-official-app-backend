package operator

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/ishihaya/company-official-app-backend/pkg/authgo"
	"golang.org/x/xerrors"
)

// NOTE: operatorは複雑なロジックのみtestを導入する 現在は導入していないが柔軟に対応する

type authOperator struct {
	authClient *authgo.Client
}

func NewAuthOperator(authClient *authgo.Client) operator.AuthOperator {
	return &authOperator{
		authClient: authClient,
	}
}

// FindByToken - トークンから認証情報を見つける
func (a *authOperator) FindByToken(token string) (*entity.Auth, error) {
	ctx := context.Background()
	authToken, err := a.authClient.Verify(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	auth := &entity.Auth{
		ID: authToken.UID,
	}
	return auth, nil
}
