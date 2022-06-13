package site

import "github.com/jackc/pgx/v4"

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) Store(site Site) error {
	return nil
}

//func (r *Repository) FindByCode()
