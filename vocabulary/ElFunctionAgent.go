package vocabulary

import "errors"

// An agent whose signature is of a function, i.e. has a result type.

// Interface definition
type IElFunctionAgent interface {
	IElAgent
	//EL_FUNCTION_AGENT
}

// Struct definition
type ElFunctionAgent struct {
	ElAgent
	// Attributes
}

func (e *ElFunctionAgent) SetDefinition(definition IBmmRoutine) error {
	s, ok := definition.(IBmmFunction)
	if !ok {
		return errors.New("_type-assertion in ElFunctionAgent->SetDefinition went wrong")
	} else {
		e.definition = s
		return nil
	}
}

// CONSTRUCTOR
func NewElFunctionAgent() *ElFunctionAgent {
	elfunctionagent := new(ElFunctionAgent)
	elfunctionagent.openArgs = make([]string, 0)
	elfunctionagent.isWritable = false
	// Constants
	return elfunctionagent
}

// BUILDER
type ElFunctionAgentBuilder struct {
	elfunctionagent *ElFunctionAgent
	errors          []error
}

func NewElFunctionAgentBuilder() *ElFunctionAgentBuilder {
	return &ElFunctionAgentBuilder{
		elfunctionagent: NewElFunctionAgent(),
		errors:          make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Reference to definition of a routine for which this is a direct call instance,
if one exists.
*/
func (i *ElFunctionAgentBuilder) SetDefinition(v IBmmFunction) *ElFunctionAgentBuilder {
	i.AddError(i.elfunctionagent.SetDefinition(v))
	return i
}

// From: ElAgent
// Closed arguments of a routine call as a tuple of objects.
func (i *ElFunctionAgentBuilder) SetClosedArgs(v IElTuple) *ElFunctionAgentBuilder {
	i.AddError(i.elfunctionagent.SetClosedArgs(v))
	return i
}

// From: ElAgent
/**
Optional list of names of open arguments of the call. If not provided, and the
name refers to a routine with more arguments than supplied in closed_args , the
missing arguments are inferred from the definition .
*/
func (i *ElFunctionAgentBuilder) SetOpenArgs(v []string) *ElFunctionAgentBuilder {
	i.AddError(i.elfunctionagent.SetOpenArgs(v))
	return i
}

// From: ElAgent
// name of the routine being called.
func (i *ElFunctionAgentBuilder) SetName(v string) *ElFunctionAgentBuilder {
	i.AddError(i.elfunctionagent.SetName(v))
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElFunctionAgentBuilder) SetScoper(v IElValueGenerator) *ElFunctionAgentBuilder {
	i.AddError(i.elfunctionagent.SetScoper(v))
	return i
}

func (i *ElFunctionAgentBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElFunctionAgentBuilder) Build() *ElFunctionAgent {
	return i.elfunctionagent
}

//FUNCTIONS
