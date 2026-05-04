package store

import (
	"context"
	"database/sql"
)

type ProfileStorage struct {
	db *sql.DB
}

type Profiles struct {
	UserID        int32  `json:"user_id"`
	FavoriteColor string `json:"favorite_color"`
	CreatedAt     string `json:"created_at"`
}

func (ps *ProfileStorage) Create(ctx context.Context, favoriteColor string) error {

	query := `
		INSERT INTO profiles (favorite_color)
		VALUES ($1)	
	`

	_, err := ps.db.ExecContext(ctx, query, favoriteColor)

	if err != nil {
		return err
	}

	return nil

}
