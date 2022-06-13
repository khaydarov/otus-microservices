package transaction

import "github.com/jackc/pgx/v4"

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

type Repository struct {
	db *pgx.Conn
}

func (l *Repository) Commit(transaction Transaction) error {
	return nil
}
