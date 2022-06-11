package user

import "github.com/google/uuid"

type ID struct {
	value string
}

func (id *ID) GetValue() string {
	return id.value
}

func WithValue(value string) ID {
	return ID{
		value,
	}
}

func NewID() ID {
	return ID{
		uuid.NewString(),
	}
}
