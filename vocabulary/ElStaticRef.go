package vocabulary

// Reference to a writable property, either a constant or computed.

// Interface definition
type IElStaticRef interface {
	IElFeatureRef
	EvalType() IBmmType
}

// Struct definition
type ElStaticRef struct {
	ElFeatureRef
	// Attributes
	// Constant definition (within class).
	Definition IBmmStatic `yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return False.
	IsWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// CONSTRUCTOR
func NewElStaticRef() *ElStaticRef {
	elstaticref := new(ElStaticRef)
	// Constants
	return elstaticref
}

// BUILDER
type ElStaticRefBuilder struct {
	elstaticref *ElStaticRef
}

func NewElStaticRefBuilder() *ElStaticRefBuilder {
	return &ElStaticRefBuilder{
		elstaticref: NewElStaticRef(),
	}
}

// BUILDER ATTRIBUTES
// Constant definition (within class).
func (i *ElStaticRefBuilder) SetDefinition(v IBmmStatic) *ElStaticRefBuilder {
	i.elstaticref.Definition = v
	return i
}

// Defined to return False.
func (i *ElStaticRefBuilder) SetIsWritable(v bool) *ElStaticRefBuilder {
	i.elstaticref.IsWritable = v
	return i
}

// From: ElFeatureRef
// Scoping expression, which must be a EL_VALUE_GENERATOR .
func (i *ElStaticRefBuilder) SetScoper(v IElValueGenerator) *ElStaticRefBuilder {
	i.elstaticref.Scoper = v
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElStaticRefBuilder) SetName(v string) *ElStaticRefBuilder {
	i.elstaticref.Name = v
	return i
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
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElStaticRef) IsBoolean() bool {
	return false
}
