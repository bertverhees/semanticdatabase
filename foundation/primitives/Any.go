package primitives

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Abstract ancestor class for all other classes.
Usually maps to a type like Any or Object in an object-oriented technology.
Defined here to provide value and reference equality semantics.
*/
type IAny interface {
	/*
		1..1 (abstract)
		is_equal (other: Any[1]): Boolean
		Value equality: return True if this and other are attached to objects considered to be equal in value.
		Parameters other Other object for comparison.
	*/
	IsEqual(any IAny) *Boolean
	/*
	   1..1
	   not_equal alias "!=", "≠" (other: Ordered[1]): Boolean
	   True if current object not equal to other. Returns not equal().
	*/
	NotEqual(other IAny) *Boolean
	//IsComparable(other IAny) *Boolean
}

type Any struct {
}

/*
function returns with INumeric true if ordered is greater than zero, else false
function returns with String true if ordered = true or false when ordered = false and an error if it is neither
function returns with Character true if  ordered = t or false when ordered = f and an error if it is neither
function returns with Boolean itself
*/
func MakeIAnyComparableToBoolean(ordered IAny) (*Boolean, error) {
	switch ordered.(type) {
	case *Boolean:
		return ordered.(*Boolean), nil
	case *Double:
		return NewBoolean(ordered.(*Double).value > 0), nil
	case *Real:
		return NewBoolean(ordered.(*Real).value > 0), nil
	case *Integer:
		return NewBoolean(ordered.(*Integer).value > 0), nil
	case *Integer64:
		return NewBoolean(ordered.(*Integer64).value > 0), nil
	case *Octet:
		return NewBoolean(ordered.(*Octet).value > 0), nil
	case *String:
		if ordered.(*String).value == "true" {
			return NewBoolean(true), nil
		} else if ordered.(*String).value == "false" {
			return NewBoolean(false), nil
		} else {
			return nil, errors.New("Cannot convert this String to Boolean:" + ordered.(*String).value)
		}
	case *Character:
		if ordered.(*Character).value == 't' {
			return NewBoolean(true), nil
		} else if ordered.(*Character).value == 'f' {
			return NewBoolean(false), nil
		} else {
			return nil, errors.New("Cannot convert this Character to Boolean:" + string(ordered.(*Character).value))
		}
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float-like type the Round-value converted to rune
function returns with Integer-like and Octet type the value converted to rune
function returns with String the first character in a string
function returns with Character itself
function returns with Boolean 't' when true else 'f'
*/
func MakeIAnyComparableToCharacter(ordered IAny) (*Character, error) {
	switch ordered.(type) {
	case *Boolean:
		if ordered.(*Boolean).Value() {
			return NewCharacter('t'), nil
		} else {
			return NewCharacter('f'), nil
		}
	case *Double:
		i32 := math.Round(ordered.(*Double).Value())
		if i32 > math.MaxInt8 || i32 < math.MinInt8 {
			return nil, errors.New(fmt.Sprintf("Double: %v does not fit into int32 (Character)", i32))
		}
		return NewCharacter(int32(i32)), nil
	case *Real:
		i32 := math.Round(float64(ordered.(*Real).Value()))
		if i32 > math.MaxInt8 || i32 < math.MinInt8 {
			return nil, errors.New(fmt.Sprintf("Real: %v does not fit into int32 (Character)", i32))
		}
		return NewCharacter(int32(i32)), nil
	case *Integer:
		return NewCharacter(rune(ordered.(*Integer).Value())), nil
	case *Integer64:
		i32 := ordered.(*Integer64).Value()
		if i32 > math.MaxInt32 || i32 < math.MinInt32 {
			return nil, errors.New(fmt.Sprintf("Integer64: %v does not fit into int32 (Character)", i32))
		}
		return NewCharacter(int32(i32)), nil
	case *Octet:
		return NewCharacter(rune(ordered.(*Octet).Value())), nil
	case *String:
		return NewCharacter(rune(ordered.(*String).Value()[0])), nil
	case *Character:
		return ordered.(*Character), nil
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float- or Integer-like or Character type the value
function returns with String the parsed float
function returns with Character itself
function returns with Boolean an error
*/
func (p *Double) MakeIAnyComparableToDouble(ordered IAny) (*Double, error) {
	switch ordered.(type) {
	case *Double:
		return ordered.(*Double), nil
	case *Integer:
		return NewDouble(float64(ordered.(*Integer).Value())), nil
	case *Integer64:
		return NewDouble(float64(ordered.(*Integer64).Value())), nil
	case *Real:
		return NewDouble(float64(ordered.(*Real).Value())), nil
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 64)
		if err != nil {
			return nil, errors.New("Cannot convert this string to float:" + ordered.(*String).value)
		}
		return NewDouble(f), nil
	case *Character:
		return NewDouble(float64(ordered.(*Character).Value())), nil
	case *Octet:
		return NewDouble(float64(ordered.(*Octet).Value())), nil
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float- or Integer-like or Character type the value
function returns with String the parsed float
function returns with Boolean an error
*/
func MakeIAnyComparableToInteger(ordered IAny) (*Integer, error) {
	switch ordered.(type) {
	case *Double:
		i64 := math.Round(ordered.(*Double).Value())
		if i64 > math.MaxInt32 || i64 < math.MinInt32 {
			return nil, errors.New(fmt.Sprintf("Double: %v does not fit into int32 (Integer)", i64))
		}
		return NewInteger(int32(i64)), nil
	case *Integer:
		return ordered.(*Integer), nil
	case *Integer64:
		i64 := ordered.(*Integer64).Value()
		if i64 > math.MaxInt32 || i64 < math.MinInt32 {
			return nil, errors.New(fmt.Sprintf("Integer64: %v does not fit into int32 (Integer)", i64))
		}
		return NewInteger(int32(i64)), nil
	case *Real:
		i64 := math.Round(float64(ordered.(*Real).Value()))
		if i64 > math.MaxInt32 || i64 < math.MinInt32 {
			return nil, errors.New(fmt.Sprintf("Real: %v does not fit into int32 (Integer)", i64))
		}
		return NewInteger(int32(i64)), nil
	case *String:
		i, err := strconv.ParseInt(ordered.(*String).value, 10, 32)
		return NewInteger(int32(i)), err
	case *Character:
		return NewInteger(int32(float64(ordered.(*Character).Value()))), nil
	case *Octet:
		i, err := strconv.ParseInt(ordered.(*String).value, 10, 8)
		return NewInteger(int32(i)), err
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float- or Integer-like or Character type the value
function returns with String the parsed float
function returns with Boolean an error
*/
func MakeIAnyComparableToInteger64(ordered IAny) (*Integer64, error) {
	switch ordered.(type) {
	case *Double:
		return NewInteger64(int64(math.Round(ordered.(*Double).Value()))), nil
	case *Integer:
		return NewInteger64(int64(ordered.(*Integer).Value())), nil
	case *Integer64:
		return ordered.(*Integer64), nil
	case *Real:
		return NewInteger64(int64(math.Round(float64(ordered.(*Real).Value())))), nil
	case *String:
		i64, err := strconv.ParseInt(ordered.(*String).Value(), 10, 64)
		return NewInteger64(i64), err
	case *Character:
		return NewInteger64(int64(ordered.(*Character).Value())), nil
	case *Octet:
		return NewInteger64(int64(ordered.(*Octet).Value())), nil
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float- or Integer-like or Character type the value
function returns with String the parsed float
function returns with Boolean an error
*/
func MakeIAnyComparableToOctet(ordered IAny) (*Octet, error) {
	switch ordered.(type) {
	case *Double:
		i64 := math.Round(ordered.(*Double).Value())
		if i64 > math.MaxUint8 || i64 < 0 {
			return nil, errors.New(fmt.Sprintf("Double: %v does not fit into uint8 (Octet)", i64))
		}
		return NewOctet(uint8(i64)), nil
	case *Integer:
		i32 := math.Round(ordered.(*Double).Value())
		if i32 > math.MaxUint8 || i32 < 0 {
			return nil, errors.New(fmt.Sprintf("Double: %v does not fit into uint8 (Octet)", i32))
		}
		return NewOctet(uint8(i32)), nil
	case *Integer64:
		i64 := ordered.(*Integer64).Value()
		if i64 > math.MaxUint8 || i64 < 0 {
			return nil, errors.New(fmt.Sprintf("Integer64: %v does not fit into uint8 (Octet)", i64))
		}
		return NewOctet(uint8(i64)), nil
	case *Real:
		i64 := math.Round(float64(ordered.(*Real).Value()))
		if i64 > math.MaxUint8 || i64 < 0 {
			return nil, errors.New(fmt.Sprintf("Real: %v does not fit into intuint832 (Octet)", i64))
		}
		return NewOctet(uint8(i64)), nil
	case *String:
		i, err := strconv.ParseInt(ordered.(*String).value, 10, 8)
		return NewOctet(uint8(i)), err
	case *Character:
		return NewOctet(uint8(float64(ordered.(*Character).Value()))), nil
	case *Octet:
		i, err := strconv.ParseInt(ordered.(*String).value, 10, 8)
		return NewOctet(uint8(i)), err
	default:
		return nil, errors.New("Not a valid type")
	}
}

/*
function returns with Float- or Integer-like or Character type the value
function returns with String the parsed float
function returns with Character itself
function returns with Boolean an error
*/
func MakeIAnyComparableToReal(ordered IAny) (*Real, error) {
	switch ordered.(type) {
	case *Double:
		f64 := float64(ordered.(*Real).Value())
		if f64 > math.MaxFloat32 {
			return nil, errors.New(fmt.Sprintf("Double: %v does not fit into float32 (Real)", f64))
		}
		return NewReal(float32(f64)), nil
	case *Integer:
		return NewReal(float32(ordered.(*Integer).Value())), nil
	case *Integer64:
		f64 := float64(ordered.(*Real).Value())
		if f64 > math.MaxFloat32 {
			return nil, errors.New(fmt.Sprintf("Integer64: %v does not fit into float32 (Real)", f64))
		}
		return NewReal(float32(f64)), nil
	case *Real:
		return ordered.(*Real), nil
	case *String:
		f, err := strconv.ParseFloat(ordered.(*String).value, 32)
		if err != nil {
			return nil, errors.New("Cannot convert this string to float:" + ordered.(*String).value)
		}
		return NewReal(float32(f)), nil
	case *Character:
		return NewReal(float32(ordered.(*Character).Value())), nil
	case *Octet:
		return NewReal(float32(ordered.(*Character).Value())), nil
	default:
		return nil, errors.New("Not a valid type")
	}
}

func MakeIAnyComparableToString(ordered IAny) (*String, error) {
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
	return NewString(r), nil
}

/*
To be worked out
*/
func MakeIAnyComparableToUri(ordered IAny) (*String, error) {
	return MakeIAnyComparableToString(ordered)
}

/*
1..1
instance_of (a_type: String[1]): Any
Create new instance of a type.
*/
func InstanceOf(aType String) IAny {
	var r IAny
	switch strings.ToLower(aType.Value()) {
	case "boolean", "bool":
		r = new(Boolean)
	case "character", "rune", "char":
		r = new(Character)
	case "double", "float64":
		r = new(Double)
	case "real", "float32":
		r = new(Real)
	case "int64", "int":
		r = new(Integer64)
	case "int32":
		r = new(Integer64)
	case "byte", "octet":
		r = new(Octet)
	case "string":
		r = new(String)
	}
	return r
}

/*
1..1
type_of (an_object: Any[1]): String
Type name of an object as a string. May include generic parameters, as in "Interval<Time>".
*/
func TypeOf(anObject IAny) *String {
	var r string
	switch v := anObject.(type) {
	default:
		r = fmt.Sprintf("%v", v)
	}
	return NewString(r)
}
