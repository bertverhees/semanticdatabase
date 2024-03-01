package primitives

import (
	"github.com/shopspring/decimal"
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

func (p *Real) returnRealFromINumeric(ordered INumeric) decimal.Decimal {
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

func (p *Real) returnRealFromIOrdered(ordered IOrdered) *Real {
	var r float32
	switch ordered.(type) {
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
	d1 := decimal.NewFromFloat32(p.value)
	d2 := p.returnRealFromINumeric(other)
	return NewReal(float32(d1.Add(d2).InexactFloat64()))
}

func (p *Real) Subtract(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Real) Multiply(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Real) Divide(other INumeric) INumeric {
	d1 := decimal.NewFromFloat32(p.value)
	d2 := p.returnRealFromINumeric(other)
	return NewReal(float32(d1.Div(d2).InexactFloat64()))
}

func (p *Real) Exponent(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Real) Negative() INumeric {
	//TODO implement me
	panic("implement me")
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
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Real) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Real) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Real) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Real) ToFixedNumberOfDecimals(precision *Integer) IFloat {
	value := float64(p.value)
	output := math.Pow(10, float64(precision.value))
	return NewReal(float32(float64(p.round(value*output)) / output))
}

func (p *Real) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
