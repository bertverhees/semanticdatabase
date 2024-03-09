package primitives

import (
	"fmt"
	"math"
	"strconv"
)

type IOrderedNumeric interface {
	INumeric
	IOrdered
}

func (p *Double) ConvertToDoubleFromIOrderedNumeric(ordered IOrderedNumeric) *Double {
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

func ConvertToIntegerFromIOrderedNumeric(ordered IOrderedNumeric) *Integer {
	var r float64
	switch ordered.(type) {
	case *Double, *Real, *Integer, *Integer64:
		return p.ConvertFromINumeric(ordered.(INumeric)).(IOrdered)
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
	return NewInteger(r)
}

func ConvertToInteger64FromIOrderedNumeric(ordered IOrderedNumeric) *Integer64 {
	var r float64
	switch ordered.(type) {
	case *Double, *Real, *Integer, *Integer64:
		return p.ConvertFromINumeric(ordered.(INumeric)).(IOrdered)
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
	return NewInteger64(r)
}

func ConvertToOctetFromIOrderedNumeric(ordered IOrderedNumeric) *Octet {
	var r float64
	switch ordered.(type) {
	case *Double, *Real, *Integer, *Integer64:
		return p.ConvertFromINumeric(ordered.(INumeric)).(IOrdered)
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
	return NewOctet(r)
}

func ConvertToRealFromIOrderedNumeric(ordered IOrderedNumeric) *Real {
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
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 64)
		if err != nil {
			panic("Cannot convert this string to float:" + ordered.(*String).value)
		}
		r = f
	case *Character:
		r = float32(ordered.(*Character).Value())
	case *Octet:
		r = float32(ordered.(*Octet).Value())
	default:
		panic("Not valid type")
	}
	return NewReal(r)
}

func ConvertToStringFromIOrderedNumeric(ordered IOrderedNumeric) *String {
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
