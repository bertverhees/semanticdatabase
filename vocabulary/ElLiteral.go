package vocabulary

/**
Literal value of any type known in the model, including primitive types. Defined
via a BMM_LITERAL_VALUE .
*/

// Interface definition
type IElLiteral interface {
	IElSimple
	//EL_LITERAL
	EvalType() IBmmType
	//effected
	Value() IBmmLiteralValue[IBmmType]
	SetValue(value IBmmLiteralValue[IBmmType]) error
}

// Struct definition
type ElLiteral struct {
	ElSimple
	// Attributes
	// The reference item from which the value of this node can be computed.
	value IBmmLiteralValue[IBmmType] `yaml:"value" json:"value" xml:"value"`
}

func (e *ElLiteral) Value() IBmmLiteralValue[IBmmType] {
	return e.value
}

func (e *ElLiteral) SetValue(value IBmmLiteralValue[IBmmType]) error {
	e.value = value
	return nil
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
	errors    []error
}

func NewElLiteralBuilder() *ElLiteralBuilder {
	return &ElLiteralBuilder{
		elliteral: NewElLiteral(),
		errors:    make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// The reference item from which the value of this node can be computed.
func (i *ElLiteralBuilder) SetValue(v IBmmLiteralValue[IBmmType]) *ElLiteralBuilder {
	i.AddError(i.elliteral.SetValue(v))
	return i
}

func (i *ElLiteralBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElLiteralBuilder) Build() *ElLiteral {
	return i.elliteral
}

// FUNCTIONS
// Return value.type .
func (e *ElLiteral) EvalType() IBmmType {
	return nil
}
