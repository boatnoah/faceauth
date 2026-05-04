package auth

import (
	"context"

	"github.com/boatnoah/faceauth/internal/store"
)

type Authenticator interface {
	GenerateToken() (string, error)
	ValidateToken(token string) (string, error)
}

type SessionAuth struct {
	store store.Storage
}

func (sa *SessionAuth) GenerateToken(ctx context.Context) (string, error) {
	sessionToken, err := sa.store.Session.Create(ctx)

	if err != nil {
		return "", err
	}

	return sessionToken, err
}

func (sa *SessionAuth) ValidateToken(token string) (string, error) {
	return "", nil
}
