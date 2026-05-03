package auth

import "github.com/boatnoah/faceauth/internal/store"

type Authenticator interface {
	GenerateToken() (string, error)
	ValidateToken(token string) (string, error)
}

type SessionAuth struct {
	store store.Storage
}

func (sa *SessionAuth) GenerateToken() (string, error) {
	return "", nil
}

func (sa *SessionAuth) ValidateToken(token string) (string, error) {
	return "", nil
}
