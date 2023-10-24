package vocabulary

import (
	"vocabulary"
)

// An agent whose signature is of a function, i.e. has a result type.

// Interface definition
type IElFunctionAgent interface {
	EvalType (  )  IBmmFunctionType
	// From: EL_AGENT
	EvalType (  )  IBmmRoutineType
	IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void
	Reference (  )  string
	// From: EL_FEATURE_REF
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
type ElFunctionAgent struct {
	// embedded for Inheritance
	ElAgent
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	/**
		Reference to definition of a routine for which this is a direct call instance,
		if one exists.
	*/
	Definition	IBmmFunction	`yaml:"definition" json:"definition" xml:"definition"`
}

//CONSTRUCTOR
func NewElFunctionAgent() *ElFunctionAgent {
	elfunctionagent := new(ElFunctionAgent)
	// Constants
	// From: ElAgent
	// From: ElFeatureRef
	// From: ElValueGenerator
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elfunctionagent
}
//BUILDER
type ElFunctionAgentBuilder struct {
	elfunctionagent *ElFunctionAgent
}

func NewElFunctionAgentBuilder() *ElFunctionAgentBuilder {
	 return &ElFunctionAgentBuilder {
		elfunctionagent : NewElFunctionAgent(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Reference to definition of a routine for which this is a direct call instance,
		if one exists.
	*/
func (i *ElFunctionAgentBuilder) SetDefinition ( v IBmmFunction ) *ElFunctionAgentBuilder{
	i.elfunctionagent.Definition = v
	return i
}

func (i *ElFunctionAgentBuilder) Build() *ElFunctionAgent {
	 return i.elfunctionagent
}

//FUNCTIONS
/**
	Post_result : Result = definition.signature. Eval type is the signature
	corresponding to the (remaining) open arguments and return type, if any.
*/
func (e *ElFunctionAgent) EvalType (  )  IBmmFunctionType {
	return nil
}
// From: EL_AGENT
/**
	Post_result : Result = definition.signature. Eval type is the signature
	corresponding to the (remaining) open arguments and return type, if any.
*/
func (e *ElFunctionAgent) EvalType (  )  IBmmRoutineType {
	return nil
}
// From: EL_AGENT
// True if there are no open arguments.
func (e *ElFunctionAgent) IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void {
	return nil
}
// From: EL_AGENT
// Generated full reference name, including scoping elements.
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElFunctionAgent) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFunctionAgent) IsBoolean (  )  bool {
	return nil
}
