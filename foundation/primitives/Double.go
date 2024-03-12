package primitives

import (
	"math"
)

type Double struct {
	Any
	value float64
}

func NewDouble(value float64) *Double {
	d := new(Double)
	d.value = value
	return d
}

func (p *Double) Add(other INumeric) (INumeric, error) {
	d, e := other.AsDouble()
	if e != nil {
		return nil, e
	}
	return NewDouble(p.value + d.value), nil
}

func (p *Double) Subtract(other INumeric) (INumeric, error) {
	d, e := other.AsDouble()
	if e != nil {
		return nil, e
	}
	return NewDouble(p.value - d.value), nil
}

func (p *Double) Multiply(other INumeric) (INumeric, error) {
	d, e := other.AsDouble()
	if e != nil {
		return nil, e
	}
	return NewDouble(p.value * d.value), nil
}

func (p *Double) Divide(other INumeric) (INumeric, error) {
	d, e := other.AsDouble()
	if e != nil {
		return nil, e
	}
	return NewDouble(p.value / d.value), nil
}

func (p *Double) Exponent(other INumeric) (INumeric, error) {
	d, e := other.AsDouble()
	if e != nil {
		return nil, e
	}
	return NewDouble(math.Pow(p.value, d.value)), nil
}

func (p *Double) Negative() (INumeric, error) {
	return NewDouble(-p.value), nil
}

func (p *Double) Value() float64 {
	return p.value
}

func (p *Double) SetValue(value float64) {
	p.value = value
}

func (p *Double) LessThan(other IOrdered) *Boolean {
	return NewBoolean(p.value < p.ConvertFromIOrdered(other).(*Double).value)
}

func (p *Double) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.ConvertFromIOrdered(other).(*Double).value)
}

func (p *Double) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.ConvertFromIOrdered(other).(*Double).value)
}

func (p *Double) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.ConvertFromIOrdered(other).(*Double).value)
}

func (p *Double) ToFixedNumberOfDecimals(precision *Integer) IFloat {
	output := math.Pow(10, float64(precision.value))
	return NewDouble(float64(p.round(p.value*output)) / output)
}

func (p *Double) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
