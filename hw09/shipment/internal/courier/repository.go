package courier

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

// NewCourierRepository returns CourierRepository instance
func NewCourierRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) GetFreeCourier() (Courier, error) {
	var (
		id int
		name string
	)

	sqlStmt := `SELECT * FROM t_couriers WHERE id NOT IN (SELECT courier_id FROM t_courier_reservations) LIMIT 1`
	err := r.db.QueryRow(context.Background(), sqlStmt).Scan(&id, &name)

	if err != nil {
		return Courier{}, errors.New("selection error")
	}

	return Courier{
		ID: id,
		Name: name,
	}, nil
}

func (r *Repository) Reserve(courier Courier, orderId string) error {
	sqlStmt := `INSERT INTO t_courier_reservations (courier_id, order_id) VALUES ($1, $2)`
	_, err := r.db.Exec(context.Background(), sqlStmt, courier.ID, orderId)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CancelReservation(orderId string) error {
	sqlStmt := `DELETE FROM t_courier_reservations WHERE order_id = $1`
	_, err := r.db.Exec(context.Background(), sqlStmt, orderId)

	if err != nil {
		return err
	}

	return nil
}