package primitives

import (
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

func (p *Double) returnDoubleFromIOrdered(ordered IOrdered) *Double {
	var r float64
	switch ordered.(type) {
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 64)
		if err != nil {
			panic("Cannot convert this string to float:" + ordered.(*String).value)
		}
		r = f
	case *Real:
		r = float64(ordered.(*Real).Value())
	case *Integer:
		r = float64(ordered.(*Integer).Value())
	case *Integer64:
		r = float64(ordered.(*Integer64).Value())
	case *Character:
		r = float64(ordered.(*Character).Value())
	case *Octet:
		r = float64(ordered.(*Octet).Value())
	default:
		panic("Unknown type")
	}
	return NewDouble(r)
}

func (p *Double) Add(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Subtract(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Multiply(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Divide(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Exponent(other INumeric) INumeric {
	//TODO implement me
	panic("implement me")
}

func (p *Double) Negative() INumeric {
	//TODO implement me
	panic("implement me")
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
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Double) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Double) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Double) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
