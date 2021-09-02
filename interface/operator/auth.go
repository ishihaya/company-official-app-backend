package operator

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/ishihaya/company-official-app-backend/infrastructure/auth"
	"golang.org/x/xerrors"
)

type authImpl struct {
	authClient *auth.Client
}

func NewAuth(authClient *auth.Client) operator.Auth {
	return &authImpl{
		authClient: authClient,
	}
}

// FindByToken - トークンから認証情報を見つける
func (a *authImpl) FindByToken(ctx context.Context, token string) (*entity.Auth, error) {
	authToken, err := a.authClient.Verify(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	auth := &entity.Auth{
		ID: authToken.UID,
	}
	return auth, nil
}
