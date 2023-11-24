package vocabulary

/**
Literal value of any type known in the model, including primitive types. Defined
via a BMM_LITERAL_VALUE .
*/

// Interface definition
type IElLiteral interface {
	IElSimple
	//EL_LITERAL
	EvalType() IBmmType //effected
}

// Struct definition
type ElLiteral struct {
	ElSimple
	// Attributes
	// The reference item from which the value of this node can be computed.
	Value IBmmLiteralValue[IBmmType] `yaml:"value" json:"value" xml:"value"`
}

// CONSTRUCTOR
func NewElLiteral() *ElLiteral {
	elliteral := new(ElLiteral)
	// Constants
	return elliteral
}

// BUILDER
type ElLiteralBuilder struct {
	elliteral *ElLiteral
}

func NewElLiteralBuilder() *ElLiteralBuilder {
	return &ElLiteralBuilder{
		elliteral: NewElLiteral(),
	}
}

// BUILDER ATTRIBUTES
// The reference item from which the value of this node can be computed.
func (i *ElLiteralBuilder) SetValue(v IBmmLiteralValue[IBmmType]) *ElLiteralBuilder {
	i.elliteral.Value = v
	return i
}

func (i *ElLiteralBuilder) Build() *ElLiteral {
	return i.elliteral
}

// FUNCTIONS
// Return value.type .
func (e *ElLiteral) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElLiteral) IsBoolean() bool {
	return false
}
