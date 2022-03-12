package account

import "github.com/google/uuid"

// NewID returns new identifier
func NewID() ID {
	return ID{
		Value: uuid.NewString(),
	}
}

type ID struct {
	Value string
}

type OwnerID struct {
	Value string
}

type Account struct {
	ID ID
	OwnerID OwnerID
	Balance int
}

// NewAccount creates new account entity
func NewAccount(ownerID string) Account {
	return Account{
		ID: NewID(),
		OwnerID: OwnerID{
			Value: ownerID,
		},
		Balance: 0,
	}
}