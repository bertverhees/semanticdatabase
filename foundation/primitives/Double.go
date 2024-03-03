package primitives

import (
	"math"
	"strconv"
)

type Double struct {
	value float64
}

func NewDouble(value float64) *Double {
	d := new(Double)
	d.value = value
	return d
}

func (p *Double) ReturnDoubleFromINumeric(ordered INumeric) *Double {
	var r float64
	switch ordered.(type) {
	case *Double:
		r = ordered.(*Double).Value()
	case *Integer:
		r = float64(ordered.(*Integer).Value())
	case *Integer64:
		r = float64(ordered.(*Integer64).Value())
	case *Real:
		r = float64(ordered.(*Real).Value())
	default:
		panic("Not valid type")
	}
	return NewDouble(r)
}

func (p *Double) ReturnDoubleFromIOrdered(ordered IOrdered) *Double {
	var r float64
	switch ordered.(type) {
	case *Double:
		r = ordered.(*Double).Value()
	case *Real:
		r = float64(ordered.(*Real).Value())
	case *Integer:
		r = float64(ordered.(*Integer).Value())
	case *Integer64:
		r = float64(ordered.(*Integer64).Value())
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 64)
		if err != nil {
			panic("Cannot convert this string to float:" + ordered.(*String).value)
		}
		r = f
	case *Character:
		r = float64(ordered.(*Character).Value())
	case *Octet:
		r = float64(ordered.(*Octet).Value())
	default:
		panic("Not valid type")
	}
	return NewDouble(r)
}

func (p *Double) Add(other INumeric) INumeric {
	return NewDouble(p.value + p.ReturnDoubleFromINumeric(other).value)
}

func (p *Double) Subtract(other INumeric) INumeric {
	return NewDouble(p.value - p.ReturnDoubleFromINumeric(other).value)
}

func (p *Double) Multiply(other INumeric) INumeric {
	return NewDouble(p.value * p.ReturnDoubleFromINumeric(other).value)
}

func (p *Double) Divide(other INumeric) INumeric {
	return NewDouble(p.value / p.ReturnDoubleFromINumeric(other).value)
}

func (p *Double) Exponent(other INumeric) INumeric {
	return NewDouble(math.Pow(p.value, p.ReturnDoubleFromINumeric(other).value))
}

func (p *Double) Negative() INumeric {
	return NewDouble(-p.value)
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
	return NewBoolean(p.value < p.ReturnDoubleFromIOrdered(other).value)
}

func (p *Double) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.ReturnDoubleFromIOrdered(other).value)
}

func (p *Double) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.ReturnDoubleFromIOrdered(other).value)
}

func (p *Double) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.ReturnDoubleFromIOrdered(other).value)
}

func (p *Double) ToFixedNumberOfDecimals(precision *Integer) IFloat {
	output := math.Pow(10, float64(precision.value))
	return NewDouble(float64(p.round(p.value*output)) / output)
}

func (p *Double) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
