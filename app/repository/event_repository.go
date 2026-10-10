package repository

import (
	"context"
	"errors"
	"fmt"

	"eticket-go/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("data tidak ditemukan")

type EventRepository interface {
	FindAll(ctx context.Context) ([]model.Event, error)
	FindByID(ctx context.Context, id int) (model.Event, error)
}

type eventPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) EventRepository {
	return &eventPostgresRepository{pool: pool}
}

func (r *eventPostgresRepository) FindAll(ctx context.Context) ([]model.Event, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, venue, event_date, price, quota
		FROM events
		ORDER BY id`,
	)

	if err != nil {
		return nil, fmt.Errorf("mengambil daftar event: %w", err)
	}

	defer rows.Close()

	hasil := []model.Event{}

	for rows.Next() {
		var e model.Event

		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Venue,
			&e.EventDate,
			&e.Price,
			&e.Quota,
		); err != nil {
			return nil, fmt.Errorf("membaca baris event: %w", err)
		}

		hasil = append(hasil, e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("membaca hasil query: %w", err)
	}

	return hasil, nil
}

func (r *eventPostgresRepository) FindByID(ctx context.Context, id int) (model.Event, error) {
	var e model.Event

	err := r.pool.QueryRow(ctx,
		`SELECT id, name, venue, event_date, price, quota
		FROM events
		WHERE id = $1`,
		id,
	).Scan(
		&e.ID,
		&e.Name,
		&e.Venue,
		&e.EventDate,
		&e.Price,
		&e.Quota,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Event{}, ErrNotFound
		}

		return model.Event{}, fmt.Errorf("mengambil event: %w", err)
	}

	return e, nil
}
