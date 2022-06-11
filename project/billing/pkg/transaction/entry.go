package transaction

import "billing/pkg/account"

const (
	EntryDebit  = 1
	EntryCredit = 2
)

func NewEntry(entryType int, transaction Transaction, account account.Account, amount int) Entry {
	return Entry{
		NewID(),
		NewEntryType(entryType),
		transaction.ID,
		account.ID,
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
