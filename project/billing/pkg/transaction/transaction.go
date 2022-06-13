package transaction

import "billing/pkg/account"

func NewTransaction(description string) Transaction {
	var entries []Entry
	return Transaction{
		id:          NewID(),
		description: description,
		entries:     entries,
	}
}

type Transaction struct {
	id          ID
	description string
	entries     []Entry
	committed   bool
}

func (t *Transaction) AddEntry(account account.Account, entryType int, amount int) {
	if !t.committed {
		entry := NewEntry(entryType, t.id, account.ID, amount)
		t.entries = append(t.entries, entry)
	}
}

func (t *Transaction) Commit() {
	t.committed = true
}
