package account

import (
	"errors"
	"github.com/google/uuid"
)

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

func (a *Account) Deposit(amount int) error {
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance < amount {
		return errors.New("not enough money")
	}

	a.Balance -= amount
	return nil
}

// NewAccount creates new account entity
func NewAccount(ownerID string) *Account {
	return &Account{
		ID: NewID(),
		OwnerID: OwnerID{
			Value: ownerID,
		},
		Balance: 0,
	}
}