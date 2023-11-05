package vocabulary

// Boolean-returning expression.

// Interface definition
type IElBooleanExpression interface {
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
	//EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
	//EL_CONSTRAINED
}

// Struct definition
type ElBooleanExpression struct {
	// embedded for Inheritance
	ElConstrained
	ElExpression
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewElBooleanExpression() *ElBooleanExpression {
	elbooleanexpression := new(ElBooleanExpression)
	// Constants
	return elbooleanexpression
}

// BUILDER
type ElBooleanExpressionBuilder struct {
	elbooleanexpression *ElBooleanExpression
}

func NewElBooleanExpressionBuilder() *ElBooleanExpressionBuilder {
	return &ElBooleanExpressionBuilder{
		elbooleanexpression: NewElBooleanExpression(),
	}
}

// BUILDER ATTRIBUTES
// From: ElConstrained
// The base expression of this constrained form.
func (i *ElBooleanExpressionBuilder) SetBaseExpression(v IElExpression) *ElBooleanExpressionBuilder {
	i.elbooleanexpression.BaseExpression = v
	return i
}

func (i *ElBooleanExpressionBuilder) Build() *ElBooleanExpression {
	return i.elbooleanexpression
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElBooleanExpression) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElBooleanExpression) IsBoolean() bool {
	return false
}
