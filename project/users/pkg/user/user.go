package user

type User struct {
	ID       ID
	Email    Email
	Password string
	Type     Type
}

func NewUser(email, password string, typeValue int) User {
	return User{
		NewID(),
		NewEmail(email),
		password,
		NewType(typeValue),
	}
}
