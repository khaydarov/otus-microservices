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

func (r *Repository) GetByUserID(userID string) (Account, error) {
	var (
		accountID string
		typeValue int
	)

	stmt := `SELECT t2.id, t2.type FROM t_user_accounts t1 JOIN t_accounts t2 ON t1.account_id = t2.id WHERE t1.user_id = $1`
	err := r.db.QueryRow(context.Background(), stmt, userID).Scan(&accountID, &typeValue)

	if err != nil {
		return Account{}, err
	}

	return Account{
		WithValue(accountID),
		userID,
		typeValue,
	}, nil
}

func (r *Repository) GetByID(id ID) (Account, error) {
	var typeValue int

	stmt := `SELECT type FROM t_accounts WHERE id = $1`
	err := r.db.QueryRow(context.Background(), stmt, id.GetValue()).Scan(&typeValue)

	if err != nil {
		return Account{}, err
	}

	return Account{
		id,
		"",
		typeValue,
	}, nil
}
