package primitives

import "math"

type Integer struct {
	Any
	value int32
}

func NewInteger(value int32) *Integer {
	d := new(Integer)
	d.value = value
	return d
}

func (p *Integer) Add(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewInteger(p.value + d.value), nil
}

func (p *Integer) Subtract(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewInteger(p.value - d.value), nil
}

func (p *Integer) Multiply(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewInteger(p.value * d.value), nil
}

func (p *Integer) Divide(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewInteger(p.value / d.value), nil
}

func (p *Integer) Exponent(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e == nil {
		result, err := NewDouble(math.Pow(float64(p.value), float64(d.value))).AsInteger()
		return result, err
	} else {
		return nil, e
	}
}

func (p *Integer) Negative() (INumeric, error) {
	return NewInteger(-p.value), nil
}

func (p *Integer) Value() int32 {
	return p.value
}

func (p *Integer) SetValue(value int32) {
	p.value = value
}

func (p *Integer) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Integer) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Integer) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Integer) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}
