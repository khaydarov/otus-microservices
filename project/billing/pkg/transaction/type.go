package transaction

func NewType(value int) Type {
	return Type{
		value,
	}
}

type Type struct {
	value int
}

func (t *Type) GetValue() int {
	return t.value
}
