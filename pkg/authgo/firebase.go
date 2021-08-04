package authgo

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"golang.org/x/xerrors"
)

type Client struct {
	*auth.Client
}

var client *Client

func init() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("failed to init firebase: %+v", err)
	}
	firebaseCli, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("failed to init firebase app auth: %+v", err)
	}
	client = &Client{firebaseCli}
}

func New() *Client {
	return client
}

func (cl *Client) Verify(ctx context.Context, token string) (*auth.Token, error) {
	idToken, err := cl.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf("failed to verify firebase idToken: %w", err)
	}
	return idToken, nil
}
