package vocabulary

// Reference to a writable property.

// Interface definition
type IElPropertyRef interface {
	IElFeatureRef
	//EL_PROPERTY_REF
	EvalType() IBmmType //effected
	Definition() IBmmProperty
	SetDefinition(definition IBmmProperty) error
}

// Struct definition
type ElPropertyRef struct {
	ElFeatureRef
	// Attributes
	// Property definition (within class).
	definition IBmmProperty `yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True.
}

func (e *ElPropertyRef) Definition() IBmmProperty {
	return e.definition
}

func (e *ElPropertyRef) SetDefinition(definition IBmmProperty) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElPropertyRef() *ElPropertyRef {
	elpropertyref := new(ElPropertyRef)
	elpropertyref.isWritable = true
	// Constants
	return elpropertyref
}

// BUILDER
type ElPropertyRefBuilder struct {
	elpropertyref *ElPropertyRef
	errors        []error
}

func NewElPropertyRefBuilder() *ElPropertyRefBuilder {
	return &ElPropertyRefBuilder{
		elpropertyref: NewElPropertyRef(),
		errors:        make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Property definition (within class).
func (i *ElPropertyRefBuilder) SetDefinition(v IBmmProperty) *ElPropertyRefBuilder {
	i.AddError(i.elpropertyref.SetDefinition(v))
	return i
}

// Defined to return True.
//func (i *ElPropertyRefBuilder) SetIsWritable(v bool) *ElPropertyRefBuilder {

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElPropertyRefBuilder) SetScoper(v IElValueGenerator) *ElPropertyRefBuilder {
	i.AddError(i.elpropertyref.SetScoper(v))
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElPropertyRefBuilder) SetName(v string) *ElPropertyRefBuilder {
	i.AddError(i.elpropertyref.SetName(v))
	return i
}

func (i *ElPropertyRefBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElPropertyRefBuilder) Build() *ElPropertyRef {
	return i.elpropertyref
}

//FUNCTIONS
/**
_type definition (i.e. BMM meta-type definition object) of the constant, property
or variable, inferred by inspection of the current scoping instance. Return
definition.type .
*/
func (e *ElPropertyRef) EvalType() IBmmType {
	return nil
}
