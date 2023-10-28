package vocabulary

// An agent whose signature is of a function, i.e. has a result type.

// Interface definition
type IElFunctionAgent interface {
	EvalType() IBmmFunctionType
	// From: EL_AGENT
	IsCallable() bool
	Reference() string
	// From: EL_EXPRESSION
	IsBoolean() bool
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
	Definition IBmmFunction `yaml:"definition" json:"definition" xml:"definition"`
}

// CONSTRUCTOR
func NewElFunctionAgent() *ElFunctionAgent {
	elfunctionagent := new(ElFunctionAgent)
	// Constants
	return elfunctionagent
}

// BUILDER
type ElFunctionAgentBuilder struct {
	elfunctionagent *ElFunctionAgent
}

func NewElFunctionAgentBuilder() *ElFunctionAgentBuilder {
	return &ElFunctionAgentBuilder{
		elfunctionagent: NewElFunctionAgent(),
	}
}

//BUILDER ATTRIBUTES
/**
Reference to definition of a routine for which this is a direct call instance,
if one exists.
*/
func (i *ElFunctionAgentBuilder) SetDefinition(v IBmmFunction) *ElFunctionAgentBuilder {
	i.elfunctionagent.Definition = v
	return i
}

// From: ElAgent
// Closed arguments of a routine call as a tuple of objects.
func (i *ElFunctionAgentBuilder) SetClosedArgs(v IElTuple) *ElFunctionAgentBuilder {
	i.elfunctionagent.ClosedArgs = v
	return i
}

// From: ElAgent
/**
Optional list of names of open arguments of the call. If not provided, and the
name refers to a routine with more arguments than supplied in closed_args , the
missing arguments are inferred from the definition .
*/
func (i *ElFunctionAgentBuilder) SetOpenArgs(v []string) *ElFunctionAgentBuilder {
	i.elfunctionagent.OpenArgs = v
	return i
}

// From: ElAgent
// Name of the routine being called.
func (i *ElFunctionAgentBuilder) SetName(v string) *ElFunctionAgentBuilder {
	i.elfunctionagent.Name = v
	return i
}

// From: ElAgent
func (i *ElFunctionAgentBuilder) SetIsWritable(v bool) *ElFunctionAgentBuilder {
	i.elfunctionagent.IsWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFunctionAgentBuilder) SetScoper(v IElValueGenerator) *ElFunctionAgentBuilder {
	i.elfunctionagent.Scoper = v
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
func (e *ElFunctionAgent) EvalType() IBmmFunctionType {
	return nil
}

// From: EL_AGENT
// Post_result_validity : Result = open_arguments = Void
// True if there are no open arguments.
func (e *ElFunctionAgent) IsCallable() bool {
	return false
}

// From: EL_AGENT
// Generated full reference name, including scoping elements.
func (e *ElFunctionAgent) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFunctionAgent) IsBoolean() bool {
	return false
}
