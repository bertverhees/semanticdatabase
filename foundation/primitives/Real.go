package primitives

type Real struct {
	value float32
}

func (p *Real) Value() float32 {
	return p.value
}

func (p *Real) SetValue(value float32) {
	p.value = value
}

func (p *Real) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Real).value)
}

func (p *Real) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Real) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Real) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Real) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
