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
	// //From: ElAgent
// Closed arguments of a routine call as a tuple of objects.
func (i *ElFunctionAgentBuilder) SetClosedArgs ( v IElTuple ) *ElFunctionAgentBuilder{
	i.elfunctionagent.ClosedArgs = v
	return i
}
/**
	Optional list of names of open arguments of the call. If not provided, and the
	name refers to a routine with more arguments than supplied in closed_args , the
	missing arguments are inferred from the definition .
*/
func (i *ElFunctionAgentBuilder) SetOpenArgs ( v []string ) *ElFunctionAgentBuilder{
	i.elfunctionagent.OpenArgs = v
	return i
}
// Reference to definition of a routine for which this is an agent, if one exists.
func (i *ElFunctionAgentBuilder) SetDefinition ( v IBmmRoutine ) *ElFunctionAgentBuilder{
	i.elfunctionagent.Definition = v
	return i
}
// Name of the routine being called.
func (i *ElFunctionAgentBuilder) SetName ( v string ) *ElFunctionAgentBuilder{
	i.elfunctionagent.Name = v
	return i
}
func (i *ElFunctionAgentBuilder) SetIsWritable ( v bool ) *ElFunctionAgentBuilder{
	i.elfunctionagent.IsWritable = v
	return i
}
	// //From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFunctionAgentBuilder) SetScoper ( v IElValueGenerator ) *ElFunctionAgentBuilder{
	i.elfunctionagent.Scoper = v
	return i
}
	// //From: ElValueGenerator
func (i *ElFunctionAgentBuilder) SetIsWritable ( v bool ) *ElFunctionAgentBuilder{
	i.elfunctionagent.IsWritable = v
	return i
}
// Name used to represent the reference or other entity.
func (i *ElFunctionAgentBuilder) SetName ( v string ) *ElFunctionAgentBuilder{
	i.elfunctionagent.Name = v
	return i
}
	// //From: ElSimple
	// //From: ElTerminal
	// //From: ElExpression

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
