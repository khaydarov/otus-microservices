package legder

import "billing/pkg/account"

const (
	EntryDebit  = 1
	EntryCredit = 2
)

func NewEntry(entryType int, transactionID ID, accountID account.ID, amount int) Entry {
	return Entry{
		NewID(),
		NewEntryType(entryType),
		transactionID,
		accountID,
		amount,
	}
}

type Entry struct {
	ID            ID
	Type          EntryType
	TransactionID ID
	AccountID     account.ID
	Amount        int
}
