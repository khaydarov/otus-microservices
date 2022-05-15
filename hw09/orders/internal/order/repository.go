package order

import "github.com/jackc/pgx/v4"

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

func (r *Repository) Store(order Order) error {
	return nil
}

func (r *Repository) Delete(order Order) error {
	return nil
}