package auth

import (
	"context"

	"github.com/boatnoah/faceauth/internal/store"
)

type Authenticator interface {
	GenerateToken(context.Context) (string, error)
	ValidateToken(context.Context, string) error
}

type SessionAuth struct {
	store store.Storage
}

func New(store store.Storage) *SessionAuth {
	return &SessionAuth{
		store: store,
	}
}

func (sa *SessionAuth) GenerateToken(ctx context.Context) (string, error) {
	sessionToken, err := sa.store.Session.Create(ctx)

	if err != nil {
		return "", err
	}

	return sessionToken, err
}

func (sa *SessionAuth) ValidateToken(ctx context.Context, token string) error {
	_, err := sa.store.Session.FindBySessionToken(ctx, token)

	if err != nil {
		return err
	}

	return nil
}
