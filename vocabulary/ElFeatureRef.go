package vocabulary

import (
	"vocabulary"
)

/**
	A reference that is scoped by a containing entity and requires a context
	qualifier if it is not the currently scoping entity.
*/

// Interface definition
type IElFeatureRef interface {
	Reference (  )  string
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElFeatureRef struct {
	// embedded for Inheritance
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	Scoper	IElValueGenerator	`yaml:"scoper" json:"scoper" xml:"scoper"`
}

//CONSTRUCTOR
func NewElFeatureRef() *ElFeatureRef {
	elfeatureref := new(ElFeatureRef)
	// Constants
	// From: ElValueGenerator
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elfeatureref
}
//BUILDER
type ElFeatureRefBuilder struct {
	elfeatureref *ElFeatureRef
}

func NewElFeatureRefBuilder() *ElFeatureRefBuilder {
	 return &ElFeatureRefBuilder {
		elfeatureref : NewElFeatureRef(),
	}
}

//BUILDER ATTRIBUTES
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFeatureRefBuilder) SetScoper ( v IElValueGenerator ) *ElFeatureRefBuilder{
	i.elfeatureref.Scoper = v
	return i
}
	// //From: ElValueGenerator
func (i *ElFeatureRefBuilder) SetIsWritable ( v bool ) *ElFeatureRefBuilder{
	i.elfeatureref.IsWritable = v
	return i
}
// Name used to represent the reference or other entity.
func (i *ElFeatureRefBuilder) SetName ( v string ) *ElFeatureRefBuilder{
	i.elfeatureref.Name = v
	return i
}
	// //From: ElSimple
	// //From: ElTerminal
	// //From: ElExpression

func (i *ElFeatureRefBuilder) Build() *ElFeatureRef {
	 return i.elfeatureref
}

//FUNCTIONS
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFeatureRef) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElFeatureRef) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElFeatureRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFeatureRef) IsBoolean (  )  bool {
	return nil
}
