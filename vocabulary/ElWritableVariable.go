package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of writable variables, including routine locals and the special
	variable 'Result'.
*/

// Interface definition
type IElWritableVariable interface {
	// From: EL_VARIABLE
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElWritableVariable struct {
	// embedded for Inheritance
	ElVariable
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Variable definition to which this reference refers.
	Definition	IBmmWritableVariable	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True in all cases.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

//CONSTRUCTOR
func NewElWritableVariable() *ElWritableVariable {
	elwritablevariable := new(ElWritableVariable)
	// Constants
	// From: ElVariable
	// From: ElValueGenerator
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elwritablevariable
}
//BUILDER
type ElWritableVariableBuilder struct {
	elwritablevariable *ElWritableVariable
}

func NewElWritableVariableBuilder() *ElWritableVariableBuilder {
	 return &ElWritableVariableBuilder {
		elwritablevariable : NewElWritableVariable(),
	}
}

//BUILDER ATTRIBUTES
	// Variable definition to which this reference refers.
func (i *ElWritableVariableBuilder) SetDefinition ( v IBmmWritableVariable ) *ElWritableVariableBuilder{
	i.elwritablevariable.Definition = v
	return i
}
	// Defined to return True in all cases.
func (i *ElWritableVariableBuilder) SetIsWritable ( v bool ) *ElWritableVariableBuilder{
	i.elwritablevariable.IsWritable = v
	return i
}

func (i *ElWritableVariableBuilder) Build() *ElWritableVariable {
	 return i.elwritablevariable
}

//FUNCTIONS
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElWritableVariable) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElWritableVariable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElWritableVariable) IsBoolean (  )  bool {
	return nil
}
