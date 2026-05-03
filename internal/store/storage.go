package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Token interface {
		Create(context.Context) error
	}
	Profile interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Token:   &TokenStorage{db},
		Profile: &ProfileStorage{db},
	}
}
