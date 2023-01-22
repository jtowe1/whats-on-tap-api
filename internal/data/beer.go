package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Beer struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Version     int32        `json:"version"`
	CreatedAt   time.Time    `json:"-"`
	LastUpdated sql.NullTime `json:"-"`
}

type BeerModel struct {
	DB *sql.DB
}

func (b BeerModel) Get(id int64) (*Beer, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, name, version, created_at, last_updated
		FROM beer
		WHERE id = ?`

	var beer Beer

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := b.DB.QueryRowContext(ctx, query, id).Scan(
		&beer.Id,
		&beer.Name,
		&beer.Version,
		&beer.CreatedAt,
		&beer.LastUpdated,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &beer, nil
}
