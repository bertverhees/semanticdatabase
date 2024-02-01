package primitives

type Double struct {
	value float64
}

func (p *Double) Add(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Subtract(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Multiply(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Divide(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Exponent(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Negative() INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Value() float64 {
	return p.value
}

func (p *Double) SetValue(value float64) {
	p.value = value
}

func (p *Double) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Double).value)
}

func (p *Double) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Double) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Double) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Double) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
