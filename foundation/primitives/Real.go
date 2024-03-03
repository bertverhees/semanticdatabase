package primitives

import (
	"math"
	"strconv"
)

type Real struct {
	value float32
}

func NewReal(value float32) *Real {
	d := new(Real)
	d.value = value
	return d
}

func (p *Real) ConvertFromINumeric(ordered INumeric) INumeric {
	var r float32
	switch ordered.(type) {
	case *Double:
		r = float32(ordered.(*Double).Value())
	case *Integer:
		r = float32(ordered.(*Integer).Value())
	case *Integer64:
		r = float32(ordered.(*Integer64).Value())
	case *Real:
		r = float32(ordered.(*Real).Value())
	default:
		panic("Not valid type")
	}
	return NewReal(r)
}

func (p *Real) ConvertFromIOrdered(ordered IOrdered) IOrdered {
	var r float32
	switch ordered.(type) {
	case *Double:
		r = float32(ordered.(*Double).Value())
	case *Real:
		r = float32(ordered.(*Real).Value())
	case *Integer:
		r = float32(ordered.(*Integer).Value())
	case *Integer64:
		r = float32(ordered.(*Integer64).Value())
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 32)
		if err != nil {
			panic("Cannot convert this string to float:" + ordered.(*String).value)
		}
		r = float32(f)
	case *Character:
		r = float32(ordered.(*Character).Value())
	case *Octet:
		r = float32(ordered.(*Octet).Value())
	default:
		panic("Not valid type")
	}
	return NewReal(r)
}

func (p *Real) Add(other INumeric) INumeric {
	return NewReal(p.value + p.ConvertFromINumeric(other).(*Real).value)
}

func (p *Real) Subtract(other INumeric) INumeric {
	return NewReal(p.value - p.ConvertFromINumeric(other).(*Real).value)
}

func (p *Real) Multiply(other INumeric) INumeric {
	return NewReal(p.value * p.ConvertFromINumeric(other).(*Real).value)
}

func (p *Real) Divide(other INumeric) INumeric {
	return NewReal(p.value / p.ConvertFromINumeric(other).(*Real).value)
}

func (p *Real) Exponent(other INumeric) INumeric {
	return NewReal(float32(math.Pow(float64(p.value), float64(p.ConvertFromINumeric(other).(*Real).value))))
}

func (p *Real) Negative() INumeric {
	return NewReal(-p.value)
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
	return NewBoolean(p.value < p.ConvertFromIOrdered(other).(*Real).value)
}

func (p *Real) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.ConvertFromIOrdered(other).(*Real).value)
}

func (p *Real) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.ConvertFromIOrdered(other).(*Real).value)
}

func (p *Real) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.ConvertFromIOrdered(other).(*Real).value)
}

func (p *Real) ToFixedNumberOfDecimals(precision *Integer) IFloat {
	value := float64(p.value)
	output := math.Pow(10, float64(precision.value))
	return NewReal(float32(float64(p.round(value*output)) / output))
}

func (p *Real) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
