package order

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db *pgx.Conn
}

func (s *Service) Store(idempotencyKey string, o Order) error {
	ctx := context.Background()
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err1 := tx.Exec(
		ctx,
		"INSERT INTO t_orders (id, owner_id, price, title) VALUES ($1, $2, $3, $4)",
		o.ID.Value,
		o.CreatorID.Value,
		o.Price,
		o.Title,
	)

	_, err2 := tx.Exec(
		ctx,
		"INSERT INTO processed_orders (id, order_id) VALUES ($1, $2)",
		idempotencyKey,
		o.ID.Value,
	)

	if err1 != nil || err2 != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		if err1 == nil && err2 != nil {
			return err2
		}

		return err1
	}

	return tx.Commit(ctx)
}

func (s *Service) GetOrderIDByIdempotencyKey(idempotencyKey string) (string, error) {
	var orderID string
	err := s.db.QueryRow(
		context.Background(),
		"SELECT order_id FROM processed_orders WHERE id = $1",
		idempotencyKey,
	).Scan(&orderID)

	if err != nil {
		return "", err
	}

	return orderID, nil
}

func NewService(db *pgx.Conn) *Service {
	return &Service{
		db,
	}
}
