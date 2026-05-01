package store

import (
	"context"
	"database/sql"
)

type ProfileStorage struct {
	db *sql.DB
}

func (ps *ProfileStorage) Create(ctx context.Context) error {
	return nil
}
