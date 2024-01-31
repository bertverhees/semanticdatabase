package primitives

type String struct {
	value string
}

func (p *String) Value() string {
	return p.value
}

func (p *String) SetValue(value string) {
	p.value = value
}

func (p *String) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*String).value)
}

func (p *String) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*String)
	//if ok {
	//	return
	//}
	return nil
}

func (p *String) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *String) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *String) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
