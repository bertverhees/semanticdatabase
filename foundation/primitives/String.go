package primitives

type String struct {
	Any
	value string
}

func NewString(value string) *String {
	v := new(String)
	v.value = value
	return v
}

func (p *String) Value() string {
	return p.value
}

func (p *String) SetValue(value string) {
	p.value = value
}

func (p *String) IsEqual(b IAny) *Boolean {
	return NewBoolean(p.value == b.(*String).value)
}

func (p *String) Equal(any IAny) *Boolean {
	return p.IsEqual(any)
}
