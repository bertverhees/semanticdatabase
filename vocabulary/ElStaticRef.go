package vocabulary

// Reference to a writable property, either a constant or computed.

// Interface definition
type IElStaticRef interface {
	IElFeatureRef
	EvalType() IBmmType
	Definition() IBmmStatic
	SetDefinition(definition IBmmStatic) error
}

// Struct definition
type ElStaticRef struct {
	ElFeatureRef
	// Attributes
	// Constant definition (within class).
	definition IBmmStatic `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElStaticRef) Definition() IBmmStatic {
	return e.definition
}

func (e *ElStaticRef) SetDefinition(definition IBmmStatic) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElStaticRef() *ElStaticRef {
	elstaticref := new(ElStaticRef)
	elstaticref.isWritable = false
	// Constants
	return elstaticref
}

// BUILDER
type ElStaticRefBuilder struct {
	elstaticref *ElStaticRef
	errors      []error
}

func NewElStaticRefBuilder() *ElStaticRefBuilder {
	return &ElStaticRefBuilder{
		elstaticref: NewElStaticRef(),
		errors:      make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Constant definition (within class).
func (i *ElStaticRefBuilder) SetDefinition(v IBmmProperty) *ElStaticRefBuilder {
	i.AddError(i.elstaticref.SetDefinition(v))
	return i
}

// Defined to return True.
//func (i *ElPropertyRefBuilder) SetIsWritable(v bool) *ElPropertyRefBuilder {

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElStaticRefBuilder) SetScoper(v IElValueGenerator) *ElStaticRefBuilder {
	i.AddError(i.elstaticref.SetScoper(v))
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElStaticRefBuilder) SetName(v string) *ElStaticRefBuilder {
	i.AddError(i.elstaticref.SetName(v))
	return i
}

func (i *ElStaticRefBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElStaticRefBuilder) Build() *ElStaticRef {
	return i.elstaticref
}

//FUNCTIONS
// From: EL_FEATURE_REF
/**
Generated full reference name, consisting of scoping elements and name
concatenated using dot notation.
*/
func (e *ElStaticRef) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElStaticRef) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElStaticRef) IsBoolean() bool {
	return false
}
