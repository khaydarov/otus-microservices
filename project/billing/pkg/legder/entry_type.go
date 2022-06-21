package legder

func NewEntryType(v int) EntryType {
	return EntryType{
		v,
	}
}

type EntryType struct {
	value int
}

func (t *EntryType) GetValue() int {
	return t.value
}
