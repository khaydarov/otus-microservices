package user

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

func (r *Repository) Save(user User) error {
	stmt := `INSERT INTO t_users (id, email, password, type) VALUES($1, $2, $3, $4)`
	_, err := r.db.Exec(
		context.Background(),
		stmt,
		user.ID.GetValue(),
		user.Email.GetValue(),
		user.Password,
		user.Type.GetValue(),
	)

	return err
}

func (r *Repository) FindByID(id ID) (User, error) {
	stmt := `SELECT * FROM t_users WHERE id = $1`

	var (
		email     string
		password  string
		typeValue int
	)

	row := r.db.QueryRow(context.Background(), stmt, id.GetValue())
	err := row.Scan(&email, &password, &typeValue)

	if err != nil {
		return User{}, err
	}

	return User{
		id,
		NewEmail(email),
		password,
		NewType(typeValue),
	}, nil
}

func (r *Repository) FindByEmailAndPassword(email Email, password string) (User, error) {
	stmt := `SELECT id, type FROM t_users WHERE email = $1 AND password = $2`
	var (
		id        string
		typeValue int
	)
	row := r.db.QueryRow(context.Background(), stmt, email.GetValue(), password)
	err := row.Scan(&id, &typeValue)

	if err != nil {
		return User{}, err
	}

	return User{
		WithValue(id),
		email,
		password,
		NewType(typeValue),
	}, nil
}

func (r *Repository) Delete(user User) error {
	stmt := `DELETE FROM t_users WHERE id = $1`
	_, err := r.db.Exec(context.Background(), stmt, user.ID.GetValue())
	if err != nil {
		return err
	}

	return nil
}
