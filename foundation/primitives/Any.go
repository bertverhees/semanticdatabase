package primitives

import (
	"fmt"
	"math"
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
}

type Any struct {
}

/*
function returns with INumeric true if ordered is greater than zero, else false
function returns with String true if ordered = true or false when ordered = false and an error if it is neither
function returns with Character true if  ordered = t or false when ordered = f and an error if it is neither
function returns with Boolean itself
*/
func ConvertToBooleanFromIAny(ordered IAny) *Boolean {
	var r bool
	switch ordered.(type) {
	case *Boolean:
		return ordered.(*Boolean)
	case *Double:
		r = ordered.(*Double).value > 0
	case *Real:
		r = ordered.(*Real).value > 0
	case *Integer:
		r = ordered.(*Integer).value > 0
	case *Integer64:
		r = ordered.(*Integer64).value > 0
	case *Octet:
		r = ordered.(*Octet).value > 0
	case *String:
		if ordered.(*String).value == "true" {
			r = true
		} else if ordered.(*String).value == "false" {
			r = false
		} else {
			panic("Cannot convert this String to Boolean:" + ordered.(*String).value)
		}
	case *Character:
		if ordered.(*Character).value == 't' {
			r = true
		} else if ordered.(*Character).value == 'f' {
			r = false
		} else {
			panic("Cannot convert this Character to Boolean:" + string(ordered.(*Character).value))
		}
	default:
		panic("Not valid type")
	}
	return NewBoolean(r)
}

/*
function returns with Float-like type the Round-value converted to rune
function returns with Integer-like and Octet type the value converted to rune
function returns with String the first character in a string
function returns with Character itself
function returns with Boolean 't' when true else 'f'
*/
func ConvertToCharacterFromIAny(ordered IAny) *Character {
	var r rune
	switch ordered.(type) {
	case *Boolean:
		if ordered.(*Boolean).Value() {
			r = 't'
		} else {
			r = 'f'
		}
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
	case *Octet:
		r = rune(ordered.(*Octet).Value())
	case *String:
		r = rune(ordered.(*String).Value()[0])
	case *Character:
		return ordered.(*Character)
	default:
		panic("Not valid type")
	}
	return NewCharacter(r)
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
