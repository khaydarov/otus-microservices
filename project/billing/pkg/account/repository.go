package account

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

func (r *Repository) Store(account Account) error {
	stmt := `INSERT INTO t_accounts (id, type) VALUES ($1, $2)`
	_, err := r.db.Exec(context.Background(), stmt, account.ID.GetValue(), account.Type)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO t_user_accounts (user_id, account_id) VALUES ($1, $2)`
	_, err = r.db.Exec(context.Background(), stmt, account.UserID, account.ID.GetValue())
	if err != nil {
		return err
	}

	return nil
}
