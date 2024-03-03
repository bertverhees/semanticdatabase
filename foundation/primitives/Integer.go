package primitives

type Integer struct {
	value int32
}

func NewInteger(value int32) *Integer {
	d := new(Integer)
	d.value = value
	return d
}

func (p *Integer) ConvertFromINumeric(ordered INumeric) INumeric {
	return nil
}

func (p *Integer) ConvertFromIOrdered(ordered IOrdered) IOrdered {
	return nil
}

func (p *Integer) Add(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Subtract(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Multiply(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Divide(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Exponent(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Negative() INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer) Value() int32 {
	return p.value
}

func (p *Integer) SetValue(value int32) {
	p.value = value
}

func (p *Integer) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Integer).value)
}

func (p *Integer) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Integer) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Integer) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Integer) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
