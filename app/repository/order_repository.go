package repository

import (
	"context"
	"errors"
	"fmt"

	"eticket-go/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	Create(ctx context.Context, o model.Order, items []model.OrderItem) (model.Order, error)
	FindAll(ctx context.Context) ([]model.Order, error)
	FindByID(ctx context.Context, id int) (model.Order, error)
	FindItemsByOrderID(ctx context.Context, orderID int) ([]model.OrderItem, error)
}

type orderPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) OrderRepository {
	return &orderPostgresRepository{pool: pool}
}

func (r *orderPostgresRepository) Create(ctx context.Context, o model.Order, items []model.OrderItem) (model.Order, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO orders (total_price)
		VALUES ($1)
		RETURNING id, status, created_at`,
		o.TotalPrice,
	).Scan(
		&o.ID,
		&o.Status,
		&o.CreatedAt,
	)

	if err != nil {
		return model.Order{}, fmt.Errorf("menyimpan pesanan: %w", err)
	}

	for i := range items {
		err := r.pool.QueryRow(ctx,
			`INSERT INTO order_items (order_id, event_id, quantity, price)
			VALUES ($1, $2, $3, $4)
			RETURNING id`,
			o.ID, items[i].EventID, items[i].Quantity, items[i].Price,
		).Scan(
			&items[i].ID,
		)

		if err != nil {
			return model.Order{}, fmt.Errorf("menyimpan item pesanan: %w", err)
		}

		items[i].OrderID = o.ID
	}

	o.Items = items

	return o, nil
}

func (r *orderPostgresRepository) FindAll(ctx context.Context) ([]model.Order, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, status, total_price, created_at
		FROM orders
		ORDER BY id`,
	)

	if err != nil {
		return nil, fmt.Errorf("mengambil daftar pesanan: %w", err)
	}

	defer rows.Close()

	hasil := []model.Order{}

	for rows.Next() {
		var o model.Order

		if err := rows.Scan(
			&o.ID,
			&o.Status,
			&o.TotalPrice,
			&o.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("membaca baris pesanan: %w", err)
		}

		hasil = append(hasil, o)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("membaca hasil query: %w", err)
	}

	return hasil, nil
}

func (r *orderPostgresRepository) FindByID(ctx context.Context, id int) (model.Order, error) {
	var o model.Order

	err := r.pool.QueryRow(ctx,
		`SELECT id, status, total_price, created_at
		FROM orders
		WHERE id = $1`,
		id,
	).Scan(
		&o.ID,
		&o.Status,
		&o.TotalPrice,
		&o.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, ErrNotFound
		}

		return model.Order{}, fmt.Errorf("mengambil pesanan: %w", err)
	}

	return o, nil
}

func (r *orderPostgresRepository) FindItemsByOrderID(ctx context.Context, orderID int) ([]model.OrderItem, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, order_id, event_id, quantity, price
		FROM order_items
		WHERE order_id = $1`,
		orderID,
	)

	if err != nil {
		return nil, fmt.Errorf("mengambil item pesanan: %w", err)
	}

	defer rows.Close()

	hasil := []model.OrderItem{}

	for rows.Next() {
		var item model.OrderItem

		if err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.EventID,
			&item.Quantity,
			&item.Price,
		); err != nil {
			return []model.OrderItem{}, fmt.Errorf("membaca baris item pesanan: %w", err)
		}

		hasil = append(hasil, item)
	}

	if err := rows.Err(); err != nil {
		return []model.OrderItem{}, fmt.Errorf("membaca hasil query: %w", err)
	}

	return hasil, nil
}
