package order

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) Store(order Order) error {
	sqlStmt := `INSERT INTO t_orders (id) VALUES ($1)`
	_, err := r.db.Exec(context.Background(), sqlStmt, order.ID.GetValue())

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