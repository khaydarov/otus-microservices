package user

const (
	advertiser   = 1
	websiteOwner = 2
)

func NewType(v int) Type {
	return Type{
		v,
	}
}

type Type struct {
	value int
}

func (t *Type) GetValue() int {
	return t.value
}
