package notification

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	Store(notification Notification) error
	GetByUserID(userId string) []*Notification
}

func NewPsqlRepository(db *pgx.Conn) *PsqlRepository {
	return &PsqlRepository{
		db,
	}
}

type PsqlRepository struct {
	db *pgx.Conn
}

func (r *PsqlRepository) Store(notification Notification) error {
	sqlStmt := `INSERT INTO t_notifications (id, user_id, text) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(
		context.Background(),
		sqlStmt,
		notification.ID.Value,
		notification.UserID,
		notification.Text,
	)

	return err
}

func (r *PsqlRepository) GetByUserID(userId string) []*Notification {
	var notifications []*Notification

	sqlStmt := `SELECT id, text FROM t_notifications WHERE user_id = $1`
	rows, _ := r.db.Query(context.Background(), sqlStmt, userId)

	for rows.Next() {
		var id string
		var text string

		err := rows.Scan(&id, &text)
		if err != nil {
			continue
		}

		notifications = append(notifications, &Notification{
			ID: ID{
				Value: id,
			},
			UserID: userId,
			Text: text,
		})

	}

	return notifications
}
