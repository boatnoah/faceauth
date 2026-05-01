package store

import (
	"context"
	"database/sql"
)

type AuthStorage struct {
	db *sql.DB
}

func (as *AuthStorage) Create(ctx context.Context) error {
	return nil
}
