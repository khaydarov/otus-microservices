package account

import "github.com/jackc/pgx/v4"

type Repository interface {
	GetByID(id ID) *Account
	Store(account Account) error
}

// NewRepository returns instance of psql repository
func NewRepository(db *pgx.Conn) *PsqlRepository {
	return &PsqlRepository{
		db: db,
	}
}

// PsqlRepository implementation of repository
type PsqlRepository struct {
	db *pgx.Conn
}
