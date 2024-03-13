package primitives

import "math"

type Integer64 struct {
	Any
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

func (p *Integer64) Add(other INumeric) (INumeric, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewInteger64(p.value + d.value), nil
}

func (p *Integer64) Subtract(other INumeric) (INumeric, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewInteger64(p.value - d.value), nil
}

func (p *Integer64) Multiply(other INumeric) (INumeric, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewInteger64(p.value * d.value), nil
}

func (p *Integer64) Divide(other INumeric) (INumeric, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewInteger64(p.value / d.value), nil
}

func (p *Integer64) Exponent(other INumeric) (INumeric, error) {
	d, e := other.AsInteger64()
	if e == nil {
		result, err := NewDouble(math.Pow(float64(p.value), float64(d.value))).AsInteger64()
		return result, err
	} else {
		return nil, e
	}
}

func (p *Integer64) Negative() (INumeric, error) {
	return NewInteger64(-p.value), nil
}

func (p *Integer64) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Integer64) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Integer64) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Integer64) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsInteger64()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}