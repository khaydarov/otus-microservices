package user

type Email struct {
	value string
}

func (email *Email) GetValue() string {
	return email.value
}

func NewEmail(email string) Email {
	return Email{
		email,
	}
}
