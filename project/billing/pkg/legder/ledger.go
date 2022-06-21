package legder

import (
	"billing/pkg/account"
	"context"
	"github.com/jackc/pgx/v4"
)

func NewLedger(db *pgx.Conn) Ledger {
	return Ledger{
		db,
	}
}

type Ledger struct {
	db *pgx.Conn
}

func (l *Ledger) NewTransaction(description string) Transaction {
	return NewTransaction(description)
}

func (l *Ledger) Commit(transaction Transaction) error {
	ctx := context.Background()
	tx, err := l.db.Begin(ctx)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO t_transactions (id, description) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, stmt, transaction.id.value, transaction.description)
	if err != nil {
		_ = tx.Rollback(ctx)

		return err
	}

	for _, entry := range transaction.entries {
		stmt = `INSERT INTO t_postings (type, transaction_id, account_id, amount) VALUES ($1, $2, $3, $4)`
		_, err = tx.Exec(ctx, stmt, entry.Type.value, entry.TransactionID.value, entry.AccountID.GetValue(), entry.Amount)
	}

	if err != nil {
		_ = tx.Rollback(ctx)

		return err
	}

	return tx.Commit(ctx)
}

func (l *Ledger) GetAccountBalance(account account.Account) int {
	var (
		debits  int
		credits int
	)
	ctx := context.Background()
	_ = l.db.QueryRow(
		ctx,
		"SELECT sum(amount) FROM t_postings WHERE account_id = $1 and type = 1",
		account.ID.GetValue(),
	).Scan(&debits)

	_ = l.db.QueryRow(
		ctx,
		"SELECT sum(amount) FROM t_postings WHERE account_id = $1 and type = 2",
		account.ID.GetValue(),
	).Scan(&credits)

	if account.Type == 1 {
		return debits - credits
	}

	return credits - debits
}
