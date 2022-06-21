package site

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

// Store TODO Make atomic
func (r *Repository) Store(site Site) error {
	ctx := context.Background()
	stmt := `INSERT INTO t_sites (id, title, code, domains) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, stmt, site.ID.GetValue(), site.Title, site.Code.GetValue(), site.Domains)

	if err != nil {
		return err
	}

	stmt = `INSERT INTO t_user_sites (user_id, site_id) VALUES ($1, $2)`
	_, err = r.db.Exec(ctx, stmt, site.UserID, site.ID.GetValue())
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByUserID(userID string) []Site {
	ctx := context.Background()
	stmt := `SELECT t1.id, t1.title, t1.code, t1.domains FROM t_sites t1 JOIN t_user_sites t2 ON t1.id = t2.site_id WHERE t2.user_id = $1`

	var sites []Site
	rows, _ := r.db.Query(ctx, stmt, userID)
	for rows.Next() {
		var (
			id      string
			title   string
			code    string
			domains []string
		)

		_ = rows.Scan(&id, &title, &code, &domains)
		sites = append(sites, Site{
			ID: ID{
				id,
			},
			UserID: userID,
			Title:  title,
			Code: Code{
				code,
			},
			Domains: domains,
		})
	}

	return sites
}
