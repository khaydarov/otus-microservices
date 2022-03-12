package user

type Repository interface {
	FindUserByID(id ID)  *User
	FindUserByLoginAndPassword(login, password string) *User
	Store(user User) error
}
