package vocabulary

import "errors"

/**
A call made on a closed function agent, returning a result. Equivalent to an
'application' of a function in Lambda calculus.
*/

// Interface definition
type IElFunctionCall interface {
	IElFeatureRef
	IElAgentCall
	EvalType() IBmmType
}

// Struct definition
type ElFunctionCall struct {
	ElAgentCall
	ElFeatureRef
	// Attributes
}

func (e *ElFunctionCall) SetAgent(agent IElFunctionAgent) error {
	s, ok := agent.(IElFunctionAgent)
	if !ok {
		return errors.New("_type-assertion in ElFunctionCall->SetAgent went wrong")
	} else {
		e.agent = s
		return nil
	}
}

// CONSTRUCTOR
func NewElFunctionCall() *ElFunctionCall {
	elfunctioncall := new(ElFunctionCall)
	// Constants
	return elfunctioncall
}

// BUILDER
type ElFunctionCallBuilder struct {
	elfunctioncall *ElFunctionCall
	errors           []error
}

func NewElFunctionCallBuilder() *ElFunctionCallBuilder {
	return &ElFunctionCallBuilder{
		elfunctioncall: NewElFunctionCall(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// The function agent being called.
func (i *ElFunctionCallBuilder) SetAgent(v IElFunctionAgent) *ElFunctionCallBuilder {
	i.elfunctioncall.agent = v
	return i
}

// Defined to return False.
func (i *ElFunctionCallBuilder) SetIsWritable(v bool) *ElFunctionCallBuilder {
	i.elfunctioncall.isWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFunctionCallBuilder) SetScoper(v IElValueGenerator) *ElFunctionCallBuilder {
	i.elfunctioncall.scoper = v
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElFunctionCallBuilder) SetName(v string) *ElFunctionCallBuilder {
	i.elfunctioncall.name = v
	return i
}

func (i *ElFunctionCallBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElFunctionCallBuilder) Build() *ElFunctionCall {
	return i.elfunctioncall
}

// FUNCTIONS
// Return agent.definition.type .
func (e *ElFunctionCall) EvalType() IBmmType {
	return nil
}

/*
*
Generated full reference name, consisting of any scoping elements, function name
and routine parameters enclosed in parentheses.
*/
func (e *ElFunctionCall) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFunctionCall) IsBoolean() bool {
	return false
}
