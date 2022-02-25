package repository

import (
	"auth-app/entity"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"time"
)

// SessionRepository is responsible for Session persistence
type SessionRepository interface {
	StoreSession(session entity.Session) error
	CreateSession(user entity.User) entity.Session
	GetSession(sessionId string) *entity.Session
}

// NewSessionRepository creates new Psql implementation of session repo
func NewSessionRepository(conn pgx.Conn) SessionRepository {
	return &psqlSessionRepository{
		conn,
	}
}

type psqlSessionRepository struct {
	db pgx.Conn
}

// StoreSession persists Session entity
func (u *psqlSessionRepository) StoreSession(session entity.Session) error {
	sqlStmt := `INSERT INTO t_sessions (id, expires_in, user_id, user_name) VALUES ($1, to_timestamp($2), $3, $4)`
	_, err := u.db.Exec(
		context.Background(),
		sqlStmt,
		session.Id,
		session.ExpiresIn.Unix(),
		session.UserId,
		session.UserName,
	)

	return err
}

// CreateSession creates new Session entity
func (u *psqlSessionRepository) CreateSession(user entity.User) entity.Session {
	return entity.Session{
		Id: uuid.New().String(),
		UserId: user.Id,
		UserName: user.Name,
		ExpiresIn: time.Now().Add(10 * 60 * time.Second),
	}
}

// GetSession returns Session entity from storage
func (u *psqlSessionRepository) GetSession(sessionId string) *entity.Session {
	var (
		expiresIn pgtype.Timestamp
		userId int
		userName string
	)

	sqlStmt := `SELECT expires_in, user_id, user_name FROM t_sessions WHERE id = $1`
	err := u.db.QueryRow(context.Background(), sqlStmt, sessionId).Scan(&expiresIn, &userId, &userName)
	if err != nil {
		return nil
	}

	return &entity.Session{
		Id:        sessionId,
		UserId:    userId,
		UserName:  userName,
		ExpiresIn: expiresIn.Time,
	}
}
