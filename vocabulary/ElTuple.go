package vocabulary

// Defines an array of optionally named items each of any type.

// Interface definition
type IElTuple interface {
	IElExpression
	EvalType() IBmmType
}

// Struct definition
type ElTuple struct {
	ElExpression
	// Attributes
	/**
	Items in the tuple, potentially with names. Typical use is to represent an
	argument list to routine call.
	*/
	Items []IElTupleItem `yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	Type IBmmTupleType `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewElTuple() *ElTuple {
	eltuple := new(ElTuple)
	// Constants
	return eltuple
}

// BUILDER
type ElTupleBuilder struct {
	eltuple *ElTuple
}

func NewElTupleBuilder() *ElTupleBuilder {
	return &ElTupleBuilder{
		eltuple: NewElTuple(),
	}
}

//BUILDER ATTRIBUTES
/**
Items in the tuple, potentially with names. Typical use is to represent an
argument list to routine call.
*/
func (i *ElTupleBuilder) SetItems(v []IElTupleItem) *ElTupleBuilder {
	i.eltuple.Items = v
	return i
}

// Static type inferred from literal value.
func (i *ElTupleBuilder) SetType(v IBmmTupleType) *ElTupleBuilder {
	i.eltuple.Type = v
	return i
}

func (i *ElTupleBuilder) Build() *ElTuple {
	return i.eltuple
}

// FUNCTIONS
// Return type .
func (e *ElTuple) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElTuple) IsBoolean() bool {
	return false
}
