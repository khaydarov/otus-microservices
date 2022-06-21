package legder

import "github.com/google/uuid"

func NewID() ID {
	return ID{
		uuid.NewString(),
	}
}

func WithValue(value string) ID {
	return ID{
		value,
	}
}

type ID struct {
	value string
}

func (id *ID) GetValue() string {
	return id.value
}
