package primitives

import (
	"math"
)

type Real struct {
	Any
	Float
	value float32
}

func NewReal(value float32) *Real {
	d := new(Real)
	d.value = value
	return d
}

func (p *Real) Add(other INumeric) (INumeric, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewReal(p.value + d.value), nil
}

func (p *Real) Subtract(other INumeric) (INumeric, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewReal(p.value - d.value), nil
}

func (p *Real) Multiply(other INumeric) (INumeric, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewReal(p.value * d.value), nil
}

func (p *Real) Divide(other INumeric) (INumeric, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewReal(p.value / d.value), nil
}

func (p *Real) Exponent(other INumeric) (INumeric, error) {
	d, e := other.AsInteger()
	if e == nil {
		result, err := NewDouble(math.Pow(float64(p.value), float64(d.value))).AsReal()
		return result, err
	} else {
		return nil, e
	}
}

func (p *Real) Negative() (INumeric, error) {
	return NewReal(-p.value), nil
}

func (p *Real) Value() float32 {
	return p.value
}

func (p *Real) SetValue(value float32) {
	p.value = value
}

func (p *Real) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Real) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Real) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Real) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsReal()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}
