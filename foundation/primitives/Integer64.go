package primitives

type Integer64 struct {
	value int64
}

func NewInteger64(value int64) *Integer64 {
	d := new(Integer64)
	d.value = value
	return d
}

func (p *Integer64) Value() int64 {
	return p.value
}

func (p *Integer64) SetValue(value int64) {
	p.value = value
}

func (Integer64) Add(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (Integer64) Subtract(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (Integer64) Multiply(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (Integer64) Divide(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (Integer64) Exponent(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (Integer64) Negative() INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Integer64) IsEqual(b IAny) *Boolean {
	v := ConvertToInteger64FromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *Integer64) NotEqual(b IAny) *Boolean {
	v := ConvertToInteger64FromIAny(b)
	return NewBoolean(p.value != v.value)
}

func (Integer64) LessThan(other IOrdered) *Boolean {
	//TODO implement me
	panic("implement me")
}

func (Integer64) LessThanOrEqual(other IOrdered) *Boolean {
	//TODO implement me
	panic("implement me")
}

func (Integer64) GreaterThan(other IOrdered) *Boolean {
	//TODO implement me
	panic("implement me")
}

func (Integer64) GreaterThanOrEqual(other IOrdered) *Boolean {
	//TODO implement me
	panic("implement me")
}
