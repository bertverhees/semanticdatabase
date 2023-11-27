package vocabulary

// An agent whose signature is of a procedure, i.e. has no result type.

// Interface definition
type IElProcedureAgent interface {
	IElAgent
	//EL_PROCEDURE_AGENT
	EvalType() IBmmProcedureType
}

// Struct definition
type ElProcedureAgent struct {
	ElAgent
	// Attributes
	// Reference to definition of routine for which this is a call instance.
	definition IBmmProcedure `yaml:"definition" json:"definition" xml:"definition"`
}

// CONSTRUCTOR
func NewElProcedureAgent() *ElProcedureAgent {
	elprocedureagent := new(ElProcedureAgent)
	// Constants
	return elprocedureagent
}

// BUILDER
type ElProcedureAgentBuilder struct {
	elprocedureagent *ElProcedureAgent
}

func NewElProcedureAgentBuilder() *ElProcedureAgentBuilder {
	return &ElProcedureAgentBuilder{
		elprocedureagent: NewElProcedureAgent(),
	}
}

// BUILDER ATTRIBUTES
// Reference to definition of routine for which this is a call instance.
func (i *ElProcedureAgentBuilder) SetDefinition(v IBmmProcedure) *ElProcedureAgentBuilder {
	i.elprocedureagent.definition = v
	return i
}

// From: ElAgent
// Closed arguments of a routine call as a tuple of objects.
func (i *ElProcedureAgentBuilder) SetClosedArgs(v IElTuple) *ElProcedureAgentBuilder {
	i.elprocedureagent.closedArgs = v
	return i
}

// From: ElAgent
/**
Optional list of names of open arguments of the call. If not provided, and the
name refers to a routine with more arguments than supplied in closed_args , the
missing arguments are inferred from the definition .
*/
func (i *ElProcedureAgentBuilder) SetOpenArgs(v []string) *ElProcedureAgentBuilder {
	i.elprocedureagent.openArgs = v
	return i
}

// From: ElAgent
// name of the routine being called.
func (i *ElProcedureAgentBuilder) SetName(v string) *ElProcedureAgentBuilder {
	i.elprocedureagent.ElValueGenerator.name = v
	return i
}

// From: ElAgent
func (i *ElProcedureAgentBuilder) SetIsWritable(v bool) *ElProcedureAgentBuilder {
	i.elprocedureagent.ElValueGenerator.isWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElProcedureAgentBuilder) SetScoper(v IElValueGenerator) *ElProcedureAgentBuilder {
	i.elprocedureagent.scoper = v
	return i
}

func (i *ElProcedureAgentBuilder) Build() *ElProcedureAgent {
	return i.elprocedureagent
}

//FUNCTIONS
/**
Post_result : result = definition.signature. Eval type is the signature
corresponding to the (remaining) open arguments and return type, if any.
*/
func (e *ElProcedureAgent) EvalType() IBmmProcedureType {
	return nil
}

// From: EL_AGENT
//
//	Post_result_validity : result = open_arguments = Void
//
// True if there are no open arguments.
func (e *ElProcedureAgent) IsCallable() bool {
	return false
}

// From: EL_AGENT
// Generated full reference name, including scoping elements.
func (e *ElProcedureAgent) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElProcedureAgent) IsBoolean() bool {
	return false
}
