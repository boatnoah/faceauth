package store

import (
	"context"
	"database/sql"
)

type TokenStorage struct {
	db *sql.DB
}

func (ts *TokenStorage) Create(ctx context.Context) error {
	return nil
}
