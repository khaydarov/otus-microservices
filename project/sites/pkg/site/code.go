package site

import "github.com/google/uuid"

func NewCode() Code {
	return Code{
		uuid.NewString(),
	}
}

func CodeWithValue(value string) ID {
	return ID{
		value,
	}
}

type Code struct {
	value string
}

func (code *Code) GetValue() string {
	return code.value
}
