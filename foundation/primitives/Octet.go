package primitives

type Octet struct {
	value int8
}

func (p *Octet) Value() int8 {
	return p.value
}

func (p *Octet) SetValue(value int8) {
	p.value = value
}

func (p *Octet) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Octet).value)
}

func (p *Octet) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Octet)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Octet) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Octet) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Octet) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
