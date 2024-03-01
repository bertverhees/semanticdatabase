package primitives

import (
	"github.com/shopspring/decimal"
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

func (p *Double) returnDoubleFromINumeric(ordered INumeric) decimal.Decimal {
	var d decimal.Decimal
	switch ordered.(type) {
	case *Double:
		d = decimal.NewFromFloat(ordered.(*Double).Value())
	case *Integer:
		d = decimal.NewFromInt32(ordered.(*Integer).Value())
	case *Integer64:
		d = decimal.NewFromInt(ordered.(*Integer64).Value())
	case *Real:
		d = decimal.NewFromFloat32(ordered.(*Real).Value())
	default:
		panic("Not valid type")
	}
	return d
}

func (p *Double) returnDoubleFromIOrdered(ordered IOrdered) *Double {
	var r float64
	switch ordered.(type) {
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
	d1 := decimal.NewFromFloat(p.value)
	d2 := p.returnDoubleFromINumeric(other)
	return NewDouble(d1.Add(d2).InexactFloat64())
}

func (p *Double) Subtract(other INumeric) INumeric {
	d1 := decimal.NewFromFloat(p.value)
	d2 := p.returnDoubleFromINumeric(other)
	return NewDouble(d1.Sub(d2).InexactFloat64())
}

func (p *Double) Multiply(other INumeric) INumeric {
	d1 := decimal.NewFromFloat(p.value)
	d2 := p.returnDoubleFromINumeric(other)
	return NewDouble(d1.Mul(d2).InexactFloat64())
}

func (p *Double) Divide(other INumeric) INumeric {
	d1 := decimal.NewFromFloat(p.value)
	d2 := p.returnDoubleFromINumeric(other)
	return NewDouble(d1.Div(d2).InexactFloat64())
}

func (p *Double) Exponent(other INumeric) INumeric {
	d1 := decimal.NewFromFloat(p.value)
	d2 := p.returnDoubleFromINumeric(other)
	return NewDouble(d1.Pow(d2).InexactFloat64())
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
	return NewBoolean(p.value < p.returnDoubleFromIOrdered(other).value)
}

func (p *Double) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.returnDoubleFromIOrdered(other).value)
}

func (p *Double) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.returnDoubleFromIOrdered(other).value)
}

func (p *Double) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.returnDoubleFromIOrdered(other).value)
}

func (p *Double) ToFixedNumberOfDecimals(precision *Integer) IFloat {
	output := math.Pow(10, float64(precision.value))
	return NewDouble(float64(p.round(p.value*output)) / output)
}

func (p *Double) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
