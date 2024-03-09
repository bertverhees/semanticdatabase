package primitives

import (
	"fmt"
	"math"
	"strconv"
)

type INumeric interface {
	IAny
	//1..1 (abstract)
	//add alias "+"
	//Sum with other (commutative). Actual type of result depends on arithmetic balancing rules.
	Add(other INumeric) INumeric
	//1..1(abstract)
	//subtract alias "-" (
	//Result of subtracting other. Actual type of result depends on arithmetic balancing rules.
	Subtract(other INumeric) INumeric
	//1..1 (abstract)
	//multiply alias "*" (
	//Product by other. Actual type of result depends on arithmetic balancing rules.
	Multiply(other INumeric) INumeric
	//1..1 (abstract)
	//divide alias "/" (
	//Divide by`other`. Actual type of result depends on arithmetic balancing rules.
	Divide(other INumeric) INumeric
	//1..1 (abstract)
	//exponent alias "^" (
	//Expontiation of this by other.
	Exponent(other INumeric) INumeric
	//1..1 (abstract)
	//negative alias "-" (): Numeric
	//Generate negative of current value.
	Negative() INumeric
}

/*
function returns with INumeric true if ordered is greater than zero, else false
function returns with String true if ordered = true or false when ordered = false and an error if it is neither
function returns with Character true if  ordered = t or false when ordered = f and an error if it is neither
*/
func ConvertToBooleanFromINumeric(ordered INumeric) *Boolean {
	var r bool
	switch ordered.(type) {
	case *Double:
		r = ordered.(*Double).value > 0
	case *Real:
		r = ordered.(*Real).value > 0
	case *Integer:
		r = ordered.(*Integer).value > 0
	case *Integer64:
		r = ordered.(*Integer64).value > 0
	default:
		panic("Not valid type")
	}
	return NewBoolean(r)
}

func ConvertToCharacterFromINumeric(ordered INumeric) *Character {
	var r rune
	switch ordered.(type) {
	case *Double:
		f := ordered.(*Double).Value()
		r = rune(int(math.Round(f)))
	case *Real:
		f := ordered.(*Real).Value()
		r = rune(int(math.Round(float64(f))))
	case *Integer:
		r = rune(ordered.(*Integer).Value())
	case *Integer64:
		r = rune(ordered.(*Integer64).Value())
	default:
		panic("Not valid type")
	}
	return NewCharacter(r)
}

func (p *Double) ConvertToDoubleFromINumeric(ordered INumeric) *Double {
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

func ConvertToIntegerFromINumeric(ordered INumeric) *Integer {
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

func ConvertToInteger64FromINumeric(ordered INumeric) *Integer64 {
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

func ConvertToOctetFromINumeric(ordered INumeric) *Octet {
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

func ConvertToRealFromIOrdered(ordered IOrdered) *Real {
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
	return NewReal(r)
}

func ConvertToStringFromINumeric(ordered INumeric) *String {
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
