package vocabulary

/**
Meta-type of writable variables, including routine locals and the special
variable 'result'.
*/

// Interface definition
type IElWritableVariable interface {
	IElVariable
	//EL_WRITEABLE_VARIABLE
	Definition() IBmmWritableVariable
	SetDefinition(definition IBmmWritableVariable) error
}

// Struct definition
type ElWritableVariable struct {
	ElVariable
	// Attributes
	// Variable definition to which this reference refers.
	definition IBmmWritableVariable `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElWritableVariable) Definition() IBmmWritableVariable {
	return e.definition
}

func (e *ElWritableVariable) SetDefinition(definition IBmmWritableVariable) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElWritableVariable() *ElWritableVariable {
	elwritablevariable := new(ElWritableVariable)
	elwritablevariable.isWritable = true
	// Constants
	return elwritablevariable
}

// BUILDER
type ElWritableVariableBuilder struct {
	elwritablevariable *ElWritableVariable
	errors             []error
}

func NewElWritableVariableBuilder() *ElWritableVariableBuilder {
	return &ElWritableVariableBuilder{
		elwritablevariable: NewElWritableVariable(),
		errors:             make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Variable definition to which this reference refers.
func (i *ElWritableVariableBuilder) SetDefinition(v IBmmWritableVariable) *ElWritableVariableBuilder {
	i.AddError(i.elwritablevariable.SetDefinition(v))
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElWritableVariableBuilder) SetName(v string) *ElWritableVariableBuilder {
	i.AddError(i.elwritablevariable.SetName(v))
	return i
}

func (i *ElWritableVariableBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElWritableVariableBuilder) Build() *ElWritableVariable {
	return i.elwritablevariable
}

//FUNCTIONS
