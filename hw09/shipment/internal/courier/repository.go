package courier

import "github.com/jackc/pgx/v4"

type Repository struct {
	db *pgx.Conn
}

// NewCourierRepository returns CourierRepository instance
func NewCourierRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) Reserve(courier Courier) error {
	return nil
}

func (r *Repository) CancelReservation(courier Courier) error {
	return nil
}