package store

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
)

const length = 16

type SessionStorage struct {
	db *sql.DB
}

type Sessions struct {
	ID        int32  `json:"id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expired_at"`
}

func (ts *SessionStorage) Create(ctx context.Context) (string, error) {

	query := `	
		INSERT INTO sessions (token, expired_at) 
		VALUES ($1, $2)
	`

	sessionToken, err := createSessionToken()
	if err != nil {
		return "", err
	}

	err = ts.db.QueryRowContext(ctx, query).Err()

	if err != nil {
		return "", err
	}

	return sessionToken, nil
}

func createSessionToken() (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
