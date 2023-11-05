package vocabulary

// Reference to a writable property.

// Interface definition
type IElPropertyRef interface {
	// From: EL_EXPRESSION
	IsBoolean() bool
	// From: EL_TERMINAL
	// From: EL_SIMPLE
	// From: EL_VALUE_GENERATOR
	//EL_FEATURE_REF
	Reference() string //redefined
	//EL_PROPERTY_REF
	EvalType() IBmmType //effected
}

// Struct definition
type ElPropertyRef struct {
	// embedded for Inheritance
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Property definition (within class).
	Definition IBmmProperty `yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True.
	IsWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// CONSTRUCTOR
func NewElPropertyRef() *ElPropertyRef {
	elpropertyref := new(ElPropertyRef)
	// Constants
	return elpropertyref
}

// BUILDER
type ElPropertyRefBuilder struct {
	elpropertyref *ElPropertyRef
}

func NewElPropertyRefBuilder() *ElPropertyRefBuilder {
	return &ElPropertyRefBuilder{
		elpropertyref: NewElPropertyRef(),
	}
}

// BUILDER ATTRIBUTES
// Property definition (within class).
func (i *ElPropertyRefBuilder) SetDefinition(v IBmmProperty) *ElPropertyRefBuilder {
	i.elpropertyref.Definition = v
	return i
}

// Defined to return True.
func (i *ElPropertyRefBuilder) SetIsWritable(v bool) *ElPropertyRefBuilder {
	i.elpropertyref.IsWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElPropertyRefBuilder) SetScoper(v IElValueGenerator) *ElPropertyRefBuilder {
	i.elpropertyref.Scoper = v
	return i
}

// From: ElValueGenerator
// Name used to represent the reference or other entity.
func (i *ElPropertyRefBuilder) SetName(v string) *ElPropertyRefBuilder {
	i.elpropertyref.Name = v
	return i
}

func (i *ElPropertyRefBuilder) Build() *ElPropertyRef {
	return i.elpropertyref
}

//FUNCTIONS
/**
Type definition (i.e. BMM meta-type definition object) of the constant, property
or variable, inferred by inspection of the current scoping instance. Return
definition.type .
*/
func (e *ElPropertyRef) EvalType() IBmmType {
	return nil
}

// From: EL_VALUE_GENERATOR
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElPropertyRef) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElPropertyRef) IsBoolean() bool {
	return false
}
