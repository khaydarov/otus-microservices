package account

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	GetByID(id ID) *Account
	GetByOwnerID(ownerId OwnerID) *Account
	Store(account *Account) error
	Update(account *Account) error
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

func (r *PsqlRepository) Store(account *Account) error {
	sqlStmt := `INSERT INTO t_accounts (id, owner_id, balance) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(
		context.Background(),
		sqlStmt,
		account.ID.Value,
		account.OwnerID.Value,
		account.Balance,
	)

	return err
}

func (r *PsqlRepository) GetByID(id ID) *Account {
	var (
		ownerId string
		balance int
	)
	sqlStmt := `SELECT owner_id, balance FROM t_accounts WHERE id = $1`
	err := r.db.QueryRow(context.Background(), sqlStmt, id.Value).Scan(&ownerId, &balance)

	if err != nil {
		return nil
	}

	return &Account{
		id,
		OwnerID{
			Value: ownerId,
		},
		balance,
	}
}

func (r *PsqlRepository) GetByOwnerID(ownerId OwnerID) *Account {
	var (
		id string
		balance int
	)
	sqlStmt := `SELECT id, balance FROM t_accounts WHERE owner_id = $1`
	err := r.db.QueryRow(context.Background(), sqlStmt, ownerId.Value).Scan(&id, &balance)

	if err != nil {
		return nil
	}

	return &Account{
		ID{
			Value: id,
		},
		ownerId,
		balance,
	}
}

func (r *PsqlRepository) Update(account *Account) error {
	sqlStmt := `UPDATE t_accounts SET balance = $2 WHERE id = $1`
	_, err := r.db.Exec(
		context.Background(),
		sqlStmt,
		account.ID.Value,
		account.Balance,
	)

	return err
}