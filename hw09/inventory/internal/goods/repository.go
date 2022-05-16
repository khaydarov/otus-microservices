package goods

import (
	"context"
	"errors"
	"fmt"
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

func (r *Repository) Reserve(orderId string, goods []int) error {
	sqlStmt := `INSERT INTO t_goods_reservations (order_id, good_id) VALUES ($1, $2)`
	for _, v := range goods {
		_, err := r.db.Exec(context.Background(), sqlStmt, orderId, v)

		if err != nil {
			return errors.New(fmt.Sprintf("could reserve good %d: %s", v, err))
		}
	}

	return nil
}

func (r *Repository) CancelReservation(orderId string) error {
	sqlStmt := `DELETE FROM t_goods_reservations WHERE order_id = $1`
	_, err := r.db.Exec(context.Background(), sqlStmt, orderId)

	if err != nil {
		return err
	}

	return nil
}