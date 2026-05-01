package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Auth interface {
		Create(context.Context) error
	}
	Token interface {
		Create(context.Context) error
	}
	Profile interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Auth:    &AuthStorage{db},
		Token:   &TokenStorage{db},
		Profile: &ProfileStorage{db},
	}
}
