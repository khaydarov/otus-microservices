package repository

import (
	"auth-app/model"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"os"
	"time"
)

var connection *pgx.Conn

func InitStorage() error {
	var err error
	connection, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	return nil
}

func StoreSession(session model.Session) error {
	sqlStmt := `INSERT INTO t_sessions (id, expires_in, user_id, email) VALUES ($1, to_timestamp($2), $3, $4)`
	_, err := connection.Exec(
		context.Background(),
		sqlStmt,
		session.Id,
		session.ExpiresIn.Unix(),
		session.UserId,
		session.UserEmail,
	)

	return err
}

func CreateSession(user model.User) model.Session {
	return model.Session{
		Id: uuid.New().String(),
		UserId: user.Id,
		UserEmail: user.Email,
		ExpiresIn: time.Now().Add(10 * 60 * time.Second),
	}
}

func GetSession(sessionId string) *model.Session {
	var (
		expiresIn pgtype.Timestamp
		userId int
		userEmail string
	)

	sqlStmt := `SELECT expires_in, user_id, email FROM t_sessions WHERE id = $1`
	err := connection.QueryRow(context.Background(), sqlStmt, sessionId).Scan(&expiresIn, &userId, &userEmail)
	if err != nil {
		return nil
	}

	return &model.Session{
		Id:        sessionId,
		UserId:    userId,
		UserEmail: userEmail,
		ExpiresIn: expiresIn.Time,
	}
}
