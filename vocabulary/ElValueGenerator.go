package vocabulary

import (
	"vocabulary"
)

// Meta-type representing a value-generating simple expression.

// Interface definition
type IElValueGenerator interface {
	Reference (  )  string
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElValueGenerator struct {
	// embedded for Inheritance
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// Name used to represent the reference or other entity.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

//CONSTRUCTOR
func NewElValueGenerator() *ElValueGenerator {
	elvaluegenerator := new(ElValueGenerator)
	// Constants
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elvaluegenerator
}
//BUILDER
type ElValueGeneratorBuilder struct {
	elvaluegenerator *ElValueGenerator
}

func NewElValueGeneratorBuilder() *ElValueGeneratorBuilder {
	 return &ElValueGeneratorBuilder {
		elvaluegenerator : NewElValueGenerator(),
	}
}

//BUILDER ATTRIBUTES
func (i *ElValueGeneratorBuilder) SetIsWritable ( v bool ) *ElValueGeneratorBuilder{
	i.elvaluegenerator.IsWritable = v
	return i
}
// Name used to represent the reference or other entity.
func (i *ElValueGeneratorBuilder) SetName ( v string ) *ElValueGeneratorBuilder{
	i.elvaluegenerator.Name = v
	return i
}
	// //From: ElSimple
	// //From: ElTerminal
	// //From: ElExpression

func (i *ElValueGeneratorBuilder) Build() *ElValueGenerator {
	 return i.elvaluegenerator
}

//FUNCTIONS
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElValueGenerator) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElValueGenerator) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElValueGenerator) IsBoolean (  )  bool {
	return nil
}
