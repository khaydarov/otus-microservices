package account

const (
	Debit  = 1
	Credit = 2
)

func NewAccount(userID string) Account {
	return Account{
		NewID(),
		userID,
		Credit,
	}
}

type Account struct {
	ID     ID
	UserID string
	Type   int
}
