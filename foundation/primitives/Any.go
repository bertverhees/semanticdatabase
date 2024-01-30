package primitives

import (
	"fmt"
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
	   equal alias "=", "==" (other: Any[1]): Boolean
	   Reference equality for reference types, value equality for value types.
	   Parameters	other	Other object for comparison.
	*/
	Equal(any IAny) *Boolean
}

type Any struct {
}

/*
1..1
instance_of (a_type: String[1]): Any
Create new instance of a type.
*/
func InstanceOf(aType String) IAny {
	var r IAny
	switch strings.ToLower(aType.Value()) {
	case "boolean, bool":
		r = new(Boolean)
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

/*
1..1
not_equal alias "!=", "≠" (other: Ordered[1]): Boolean
True if current object not equal to other. Returns not equal().
*/
func (a *Any) NotEqual(other IOrdered) *Boolean {
	return nil
}
