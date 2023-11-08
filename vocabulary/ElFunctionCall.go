package vocabulary

/**
A call made on a closed function agent, returning a result. Equivalent to an
'application' of a function in Lambda calculus.
*/

// Interface definition
type IElFunctionCall interface {
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
	// From: EL_TERMINAL
	// From: EL_SIMPLE
	// From: EL_VALUE_GENERATOR
	//EL_FEATURE_REF
	// From: EL_AGENT_CALL
	Reference() string //redefined
}

// Struct definition
type ElFunctionCall struct {
	// embedded for Inheritance
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	ElAgentCall
	// Constants
	// Attributes
	// The function agent being called.
	Agent IElFunctionAgent `yaml:"agent" json:"agent" xml:"agent"`
	// Defined to return False.
	IsWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
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
}

func NewElFunctionCallBuilder() *ElFunctionCallBuilder {
	return &ElFunctionCallBuilder{
		elfunctioncall: NewElFunctionCall(),
	}
}

// BUILDER ATTRIBUTES
// The function agent being called.
func (i *ElFunctionCallBuilder) SetAgent(v IElFunctionAgent) *ElFunctionCallBuilder {
	i.elfunctioncall.Agent = v
	return i
}

// Defined to return False.
func (i *ElFunctionCallBuilder) SetIsWritable(v bool) *ElFunctionCallBuilder {
	i.elfunctioncall.IsWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFunctionCallBuilder) SetScoper(v IElValueGenerator) *ElFunctionCallBuilder {
	i.elfunctioncall.Scoper = v
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElFunctionCallBuilder) SetName(v string) *ElFunctionCallBuilder {
	i.elfunctioncall.Name = v
	return i
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
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFunctionCall) IsBoolean() bool {
	return false
}
