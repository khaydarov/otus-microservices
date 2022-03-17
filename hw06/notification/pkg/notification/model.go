package notification

import "github.com/google/uuid"

func NewID() ID {
	return ID{
		Value: uuid.NewString(),
	}
}

type ID struct {
	Value string
}

type Notification struct {
	ID 		ID
	Text 	string `json:"text"`
	UserID  string `json:"userId"`
}