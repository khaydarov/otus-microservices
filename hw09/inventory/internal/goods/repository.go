package goods

import "github.com/jackc/pgx/v4"

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) Reserve(orderId string, goods []int) error {
	return nil
}

func (r *Repository) CancelReservation(orderId string) error {
	return nil
}
