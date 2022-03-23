package user

import "github.com/google/uuid"

// ID is User unique identifier
type ID struct {
	Value string
}

// Password is User user password
type Password struct {
	Value string
}

// User a base entity of domain
type User struct {
	ID        ID
	Login     string
	Password  Password
	FirstName string
	LastName  string
}

// NewID Generates new identifier
func NewID() ID {
	return ID{
		Value: uuid.New().String(),
	}
}

// NewPassword Generates new password
func NewPassword() Password {
	return Password{
		Value: uuid.NewString(),
	}
}

// NewUser Creates new user entity
func NewUser(login, firstName, lastName string) User {
	return User{
		NewID(),
		login,
		NewPassword(),
		firstName,
		lastName,
	}
}