package primitives

type Octet struct {
	value uint8
}

func NewOctet(value uint8) *Octet {
	d := new(Octet)
	d.value = value
	return d
}

func (p *Octet) Value() uint8 {
	return p.value
}

func (p *Octet) SetValue(value uint8) {
	p.value = value
}

func (p *Octet) IsEqual(b IAny) *Boolean {
	v := ConvertToOctetFromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *Octet) NotEqual(b IAny) *Boolean {
	v := ConvertToOctetFromIAny(b)
	return NewBoolean(p.value != v.value)
}

func (p *Octet) LessThan(other IOrdered) *Boolean {
	return NewBoolean(p.value < other.(*Octet).value)
}

func (p *Octet) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= other.(*Octet).value)
}

func (p *Octet) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > other.(*Octet).value)
}

func (p *Octet) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= other.(*Octet).value)
}
