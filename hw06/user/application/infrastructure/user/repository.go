package user

import (
	"context"
	"github.com/jackc/pgx/v4"
	"hw06/user/domain/user"
)

// NewPsqlRepository returns instance of psql repository
func NewPsqlRepository(db *pgx.Conn) *PsqlRepository {
	return &PsqlRepository{db}
}

// PsqlRepository implements user domain repository
type PsqlRepository struct {
	db *pgx.Conn
}

func (r *PsqlRepository) FindUserByID(id user.ID)  *user.User {
	var (
		login 		string
		password 	string
		firstName   string
		lastName	string
	)

	sqlStmt := `SELECT login, password, firstname, lastname FROM t_users WHERE id = $1`
	err := r.db.QueryRow(context.Background(), sqlStmt, id.Value).Scan(&login, &password, &firstName, &lastName)
	if err != nil {
		return nil
	}

	return &user.User{
		ID:        id,
		Login:     login,
		Password:  user.Password{Value: password},
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (r *PsqlRepository) FindUserByLoginAndPassword(login, password string) *user.User {
	var (
		id			string
		firstName   string
		lastName	string
	)

	sqlStmt := `SELECT id, firstname, lastname FROM t_users WHERE login = $1 AND password = $2`
	err := r.db.QueryRow(context.Background(), sqlStmt, login, password).Scan(&id, &firstName, &lastName)
	if err != nil {
		return nil
	}

	return &user.User{
		ID:        user.ID{Value: id},
		Login:     login,
		Password:  user.Password{Value: password},
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (r *PsqlRepository) Store(user user.User) error {
	sqlStmt := `INSERT INTO t_users (id, login, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(
		context.Background(),
		sqlStmt,
		user.ID.Value,
		user.Login,
		user.Password.Value,
		user.FirstName,
		user.LastName,
	)

	return err
}