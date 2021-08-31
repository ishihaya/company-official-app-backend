package operator

import (
	"context"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/ishihaya/company-official-app-backend/pkg/authgo"
	"golang.org/x/xerrors"
)

type auth struct {
	authClient *authgo.Client
}

func NewAuth(authClient *authgo.Client) operator.Auth {
	return &auth{
		authClient: authClient,
	}
}

// FindByToken - トークンから認証情報を見つける
func (a *auth) FindByToken(ctx context.Context, token string) (*entity.Auth, error) {
	authToken, err := a.authClient.Verify(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	auth := &entity.Auth{
		ID: authToken.UID,
	}
	return auth, nil
}
