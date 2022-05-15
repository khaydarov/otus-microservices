package order

import "github.com/google/uuid"

type ID struct {
	value string
}

func (v *ID) GetValue() string {
	return v.value
}

func createID() ID {
	value := uuid.NewString()

	return ID{
		value,
	}
}

type Order struct {
	ID
}

// CreateOrder returns new Order
func CreateOrder() Order {
	return Order{
		createID(),
	}
}
