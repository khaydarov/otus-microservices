package order

import "github.com/google/uuid"

const (
	StateNew = iota
	StatePaid = iota
	StateUnpaid = iota
)

func NewID() ID {
	return ID{
		Value: uuid.NewString(),
	}
}

// ID is a order identifier
type ID struct {
	Value string
}

func NewState() State {
	return State{
		Value: StateNew,
	}
}

type State struct {
	Value int
}

func (s *State) Paid() {
	s.Value = StatePaid
}

func (s *State) Unpaid() {
	s.Value = StateUnpaid
}

// Order is an base entity of order
type Order struct {
	ID        ID
	CreatorID ID
	Title     string
	Price     int
	State     State
}

// NewOrder returns instance of new order
func NewOrder(userID string, title string, price int) Order {
	return Order{
		ID: NewID(),
		CreatorID: ID{
			Value: userID,
		},
		Title: title,
		Price: price,
		State: NewState(),
	}
}
