package contextgo

import (
	"context"

	"golang.org/x/xerrors"
)

var authIDKey = "authID"

func AuthID(ctx context.Context) (string, error) {
	if authID := ctx.Value(&authIDKey); authID != nil {
		return authID.(string), nil
	}
	return "", xerrors.New("authID not set")
}

func SetAuthID(ctx context.Context, authID string) context.Context {
	return context.WithValue(ctx, &authIDKey, authID)
}
