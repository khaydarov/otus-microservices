package user

type Created struct {
	ID 			string
	FirstName 	string
	LastName 	string
}

// NewUserCreatedEvent returns user created event
func NewUserCreatedEvent(user User) Created {
	return Created{
		user.ID.Value,
		user.FirstName,
		user.LastName,
	}
}
