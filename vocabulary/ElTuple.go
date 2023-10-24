package vocabulary

import (
	"vocabulary"
)

// Defines an array of optionally named items each of any type.

// Interface definition
type IElTuple interface {
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElTuple struct {
	// embedded for Inheritance
	ElExpression
	// Constants
	// Attributes
	/**
		Items in the tuple, potentially with names. Typical use is to represent an
		argument list to routine call.
	*/
	Items	List < EL_TUPLE_ITEM >	`yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	Type	IBmmTupleType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewElTuple() *ElTuple {
	eltuple := new(ElTuple)
	// Constants
	// From: ElExpression
	return eltuple
}
//BUILDER
type ElTupleBuilder struct {
	eltuple *ElTuple
}

func NewElTupleBuilder() *ElTupleBuilder {
	 return &ElTupleBuilder {
		eltuple : NewElTuple(),
	}
}

//BUILDER ATTRIBUTES
/**
	Items in the tuple, potentially with names. Typical use is to represent an
	argument list to routine call.
*/
func (i *ElTupleBuilder) SetItems ( v List < EL_TUPLE_ITEM > ) *ElTupleBuilder{
	i.eltuple.Items = v
	return i
}
// Static type inferred from literal value.
func (i *ElTupleBuilder) SetType ( v IBmmTupleType ) *ElTupleBuilder{
	i.eltuple.Type = v
	return i
}
	// //From: ElExpression

func (i *ElTupleBuilder) Build() *ElTuple {
	 return i.eltuple
}

//FUNCTIONS
// Return type .
func (e *ElTuple) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElTuple) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElTuple) IsBoolean (  )  bool {
	return nil
}
