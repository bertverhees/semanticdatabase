package primitives

import (
	"fmt"
	"strconv"
)

type String struct {
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (p *String) ConvertFromIOrdered(ordered IOrdered) *String {
	var r string
	switch ordered.(type) {
	case *Double:
		r = strconv.FormatFloat(ordered.(*Double).Value(), 'f', -1, 64)
	case *Real:
		r = strconv.FormatFloat(float64(ordered.(*Real).Value()), 'f', -1, 32)
	case *Integer:
		r = strconv.FormatInt(int64(ordered.(*Integer).Value()), 10)
	case *Integer64:
		r = strconv.FormatInt(ordered.(*Integer64).Value(), 10)
	case *Character:
		r = fmt.Sprintf("%c", ordered.(*Character).Value())
	case *Octet:
		r = strconv.FormatInt(int64(ordered.(*Octet).Value()), 10)
	default:
		r = ordered.(*String).value
	}
	return NewString(r)
}

func (p *String) Value() string {
	return p.value
}

func (p *String) SetValue(value string) {
	p.value = value
}

func (p *String) IsEqual(b IAny) *Boolean {
	v := ConvertToStringFromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *String) NotEqual(b IAny) *Boolean {
	v := ConvertToStringFromIAny(b)
	return NewBoolean(p.value != v.value)
}

func (p *String) LessThan(other IOrdered) *Boolean {
	return NewBoolean(p.value < p.ConvertFromIOrdered(other).value)
}

func (p *String) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.ConvertFromIOrdered(other).value)
}

func (p *String) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.ConvertFromIOrdered(other).value)
}

func (p *String) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.ConvertFromIOrdered(other).value)
}
