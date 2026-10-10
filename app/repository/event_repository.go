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
	Create(ctx context.Context, e model.Event) (model.Event, error)
	Update(ctx context.Context, e model.Event) (model.Event, error)
	Delete(ctx context.Context, id int) error
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

func (r *eventPostgresRepository) Create(ctx context.Context, e model.Event) (model.Event, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO events (name, venue, event_date, price, quota)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		e.Name, e.Venue, e.EventDate, e.Price, e.Quota,
	).Scan(
		&e.ID,
	)

	if err != nil {
		return model.Event{}, fmt.Errorf("menyimpan event: %w", err)
	}

	return e, nil
}

func (r *eventPostgresRepository) Update(ctx context.Context, e model.Event) (model.Event, error) {
	err := r.pool.QueryRow(ctx,
		`UPDATE events
		SET name = $1, venue = $2, event_date = $3, price = $4, quota = $5
		WHERE id = $6
		RETURNING id`,
		e.Name, e.Venue, e.EventDate, e.Price, e.Quota, e.ID,
	).Scan(
		&e.ID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Event{}, ErrNotFound
		}

		return model.Event{}, fmt.Errorf("memperbarui event: %w", err)
	}

	return e, nil
}

func (r *eventPostgresRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.pool.Exec(ctx,
		`DELETE FROM events
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("menghapus event: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
