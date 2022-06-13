package session

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

func (r *Repository) Store(session Session) error {
	stmt := `INSERT INTO t_sessions (id, user_id, token, user_agent, ip_address, expires) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(
		context.Background(),
		stmt,
		session.ID,
		session.UserID.GetValue(),
		session.Token.Value,
		session.UserAgent,
		session.Ip,
		session.Expires,
	)

	if err != nil {
		return err
	}

	return nil
}
