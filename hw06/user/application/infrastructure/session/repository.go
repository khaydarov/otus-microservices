package session

import (
	"context"
	"github.com/jackc/pgx/v4"
	"hw06/user/domain/session"
)

// NewSessionRepository returns psql repository
func NewSessionRepository(db *pgx.Conn) *PsqlRepository {
	return &PsqlRepository{
		db,
	}
}

type PsqlRepository struct {
	db *pgx.Conn
}

func (r *PsqlRepository) GetByToken(token session.Token) *session.Session {
	return nil
}

func (r *PsqlRepository) Store(s session.Session) error {
	sqlStmt := `
		INSERT INTO t_sessions (user_id, token, user_agent, ip_address, expires, created_at)
		VALUES ($1, $2, $3, $4, to_timestamp($5), to_timestamp($6)) RETURNING id`

	_, err := r.db.Exec(
		context.Background(),
		sqlStmt,
		s.UserID.Value,
		s.Token.Value,
		s.UserAgent,
		s.Ip,
		s.Expires.Unix(),
		s.CreatedAt.Unix(),
	)

	return err
}