package vocabulary

/**
Meta-type of read-only variables, including routine parameter and the special
variable 'Self'.
*/

// Interface definition
type IElReadonlyVariable interface {
	IElVariable
	//EL_READONLY_VARIABLE
	Definition() IBmmReadonlyVariable
	SetDefinition(definition IBmmReadonlyVariable) error
}

// Struct definition
type ElReadonlyVariable struct {
	ElVariable
	// Attributes
	// Variable definition to which this reference refers.
	definition IBmmReadonlyVariable `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElReadonlyVariable) Definition() IBmmReadonlyVariable {
	return e.definition
}

func (e *ElReadonlyVariable) SetDefinition(definition IBmmReadonlyVariable) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElReadonlyVariable() *ElReadonlyVariable {
	elreadonlyvariable := new(ElReadonlyVariable)
	elreadonlyvariable.isWritable = false
	// Constants
	return elreadonlyvariable
}

// BUILDER
type ElReadonlyVariableBuilder struct {
	elreadonlyvariable *ElReadonlyVariable
	errors             []error
}

func NewElReadonlyVariableBuilder() *ElReadonlyVariableBuilder {
	return &ElReadonlyVariableBuilder{
		elreadonlyvariable: NewElReadonlyVariable(),
		errors:             make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Variable definition to which this reference refers.
func (i *ElReadonlyVariableBuilder) SetDefinition(v IBmmWritableVariable) *ElReadonlyVariableBuilder {
	i.AddError(i.elreadonlyvariable.SetDefinition(v))
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElReadonlyVariableBuilder) SetName(v string) *ElReadonlyVariableBuilder {
	i.AddError(i.elreadonlyvariable.SetName(v))
	return i
}

func (i *ElReadonlyVariableBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElReadonlyVariableBuilder) Build() *ElReadonlyVariable {
	return i.elreadonlyvariable
}

//FUNCTIONS
