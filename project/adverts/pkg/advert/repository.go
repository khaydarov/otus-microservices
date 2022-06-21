package advert

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func NewRepository(db *pgx.Conn) Repository {
	return Repository{
		db,
	}
}

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) FindByID(id ID) (Advert, error) {
	ctx := context.Background()
	stmt := `SELECT t1.title,
					t1.description,
					t1.link,
					t1.image,
					t2.dates,
					t2.devices,
					t2.hits,
					t2.cost
			FROM t_adverts t1 JOIN t_adverts_targeting t2 ON t1.id = t2.advert_id WHERE t1.id = $1`

	var (
		title       string
		description string
		link        string
		image       string
		dates       []string
		devices     []string
		hits        int
		cost        int
	)

	err := r.db.QueryRow(ctx, stmt, id.GetValue()).Scan(&title, &description, &link, &image, &dates, &devices, &hits, &cost)
	if err != nil {
		return Advert{}, err
	}

	return Advert{
		id,
		title,
		description,
		link,
		image,
		Targeting{
			devices,
			dates,
			hits,
			cost,
		},
	}, nil
}

// Store TODO make inserts atomic
func (r *Repository) Store(advert Advert) error {
	ctx := context.Background()
	stmt := `INSERT INTO t_adverts (id, title, description, link, image) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(
		ctx,
		stmt,
		advert.ID.GetValue(),
		advert.Title,
		advert.Description,
		advert.Link,
		advert.Image,
	)

	if err != nil {
		return err
	}

	stmt = `INSERT INTO t_adverts_targeting (advert_id, dates, devices, hits, cost) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(
		ctx,
		stmt,
		advert.ID.GetValue(),
		advert.Targeting.Dates,
		advert.Targeting.Devices,
		advert.Targeting.Hits,
		advert.Targeting.Cost,
	)

	return err
}
