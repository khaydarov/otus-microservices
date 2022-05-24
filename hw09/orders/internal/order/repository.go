package order

import (
	"context"
	"github.com/jackc/pgx/v4"
	"hw09/orders/internal/tracer"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) Store(ctx context.Context, order Order) error {
	ctx, span := tracer.NewSpan(ctx, "INSERT INTO t_orders (id) VALUES ($1)")
	defer span.End()

	sqlStmt := `INSERT INTO t_orders (id) VALUES ($1)`
	_, err := r.db.Exec(ctx, sqlStmt, order.ID.GetValue())

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(order Order) error {
	sqlStmt := `DELETE FROM t_orders WHERE id = $1`
	_, err := r.db.Exec(context.Background(), sqlStmt, order.ID.GetValue())

	if err != nil {
		return err
	}

	return nil
}