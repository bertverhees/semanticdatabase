package vocabulary

// Abstract parent of all typed expression meta-types.

// Interface definition
type IElExpression interface {
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElExpression struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewElExpression() *ElExpression {
	elexpression := new(ElExpression)
	// Constants
	return elexpression
}

// BUILDER
type ElExpressionBuilder struct {
	elexpression *ElExpression
}

func NewElExpressionBuilder() *ElExpressionBuilder {
	return &ElExpressionBuilder{
		elexpression: NewElExpression(),
	}
}

//BUILDER ATTRIBUTES

func (i *ElExpressionBuilder) Build() *ElExpression {
	return i.elexpression
}

//FUNCTIONS
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElExpression) EvalType() IBmmType {
	return nil
}

/*
*
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElExpression) IsBoolean() bool {
	return false
}
