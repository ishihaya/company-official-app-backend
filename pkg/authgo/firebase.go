package authgo

import (
	"context"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"golang.org/x/xerrors"
)

type Client struct {
	*auth.Client
}

var sharedInstance *Client
var once sync.Once

func New() *Client {
	once.Do(func() {
		sharedInstance = newInstance()
	})
	return sharedInstance
}

func newInstance() *Client {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}
	firebaseCli, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	cli := &Client{firebaseCli}
	return cli
}

func (cl *Client) Verify(ctx context.Context, token string) (*auth.Token, error) {
	idToken, err := cl.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, xerrors.Errorf("failed to verify firebase idToken: %w", err)
	}
	return idToken, nil
}
