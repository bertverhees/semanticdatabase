package vocabulary

/**
Meta-type for reference to a non-abstract type as an object. Assumed to be
accessible at run-time. Typically represented syntactically as TypeName or
{TypeName} . May be used as a value, or as the qualifier for a function or
constant access.
*/

// Interface definition
type IElTypeRef interface {
	IElValueGenerator
	EvalType() IBmmType
	Type() IBmmType
	SetType(_type IBmmType) error
	IsMutable() bool
	SetIsMutable(isMutable bool) error //effected
}

// Struct definition
type ElTypeRef struct {
	ElValueGenerator
	// Attributes
	// _type, directly from the name of the reference, e.g. {SOME_TYPE} .
	_type     IBmmType `yaml:"type" json:"type" xml:"type"`
	isMutable bool     `yaml:"ismutable" json:"ismutable" xml:"ismutable"`
}

func (e *ElTypeRef) Type() IBmmType {
	return e._type
}

func (e *ElTypeRef) SetType(_type IBmmType) error {
	e._type = _type
	return nil
}

func (e *ElTypeRef) IsMutable() bool {
	return e.isMutable
}

func (e *ElTypeRef) SetIsMutable(isMutable bool) error {
	e.isMutable = isMutable
	return nil
}

// CONSTRUCTOR
func NewElTypeRef() *ElTypeRef {
	eltyperef := new(ElTypeRef)
	// Constants
	return eltyperef
}

// BUILDER
type ElTypeRefBuilder struct {
	eltyperef *ElTypeRef
	errors    []error
}

func NewElTypeRefBuilder() *ElTypeRefBuilder {
	return &ElTypeRefBuilder{
		eltyperef: NewElTypeRef(),
		errors:    make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// _type, directly from the name of the reference, e.g. {SOME_TYPE} .
func (i *ElTypeRefBuilder) SetType(v IBmmType) *ElTypeRefBuilder {
	i.AddError(i.eltyperef.SetType(v))
	return i
}
func (i *ElTypeRefBuilder) SetIsMutable(v bool) *ElTypeRefBuilder {
	i.AddError(i.eltyperef.SetIsMutable(v))
	return i
}

// From: ElValueGenerator
func (i *ElTypeRefBuilder) SetIsWritable(v bool) *ElTypeRefBuilder {
	i.AddError(i.eltyperef.SetIsWritable(v))
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElTypeRefBuilder) SetName(v string) *ElTypeRefBuilder {
	i.AddError(i.eltyperef.SetName(v))
	return i
}

func (i *ElTypeRefBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElTypeRefBuilder) Build() *ElTypeRef {
	return i.eltyperef
}

// FUNCTIONS
// Return type .
func (e *ElTypeRef) EvalType() IBmmType {
	return nil
}

// From: EL_VALUE_GENERATOR
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElTypeRef) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElTypeRef) IsBoolean() bool {
	return false
}
