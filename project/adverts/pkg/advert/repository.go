package advert

import "github.com/jackc/pgx/v4"

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) FindByID(id ID) (Advert, error) {
	return Advert{}, nil
}

func (r *Repository) Store(advert Advert) error {
	return nil
}
