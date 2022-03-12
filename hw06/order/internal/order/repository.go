package order

import (
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	GetByID(id ID) *Order
	Store(o Order) error
}

// NewPsqlOrderRepository returns instance of psql repository for orders
func NewPsqlOrderRepository(db *pgx.Conn) *PsqlOrderRepository {
	return &PsqlOrderRepository{
		db: db,
	}
}

// PsqlOrderRepository is a structure that implements orders repository
type PsqlOrderRepository struct {
	db *pgx.Conn
}

func (r *PsqlOrderRepository) GetByID(id ID) *Order {
	return nil
}

func (r *PsqlOrderRepository) Store(o Order) error {
	return nil
}
