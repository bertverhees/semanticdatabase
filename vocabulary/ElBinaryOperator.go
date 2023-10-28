package vocabulary

// Binary operator expression node.

// Interface definition
type IElBinaryOperator interface {
	// From: EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElBinaryOperator struct {
	// embedded for Inheritance
	ElOperator
	ElExpression
	// Constants
	// Attributes
	// Left operand node.
	LeftOperand IElExpression `yaml:"leftoperand" json:"leftoperand" xml:"leftoperand"`
	// Right operand node.
	RightOperand IElExpression `yaml:"rightoperand" json:"rightoperand" xml:"rightoperand"`
}

// CONSTRUCTOR
func NewElBinaryOperator() *ElBinaryOperator {
	elbinaryoperator := new(ElBinaryOperator)
	// Constants
	return elbinaryoperator
}

// BUILDER
type ElBinaryOperatorBuilder struct {
	elbinaryoperator *ElBinaryOperator
}

func NewElBinaryOperatorBuilder() *ElBinaryOperatorBuilder {
	return &ElBinaryOperatorBuilder{
		elbinaryoperator: NewElBinaryOperator(),
	}
}

// BUILDER ATTRIBUTES
// Left operand node.
func (i *ElBinaryOperatorBuilder) SetLeftOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.elbinaryoperator.LeftOperand = v
	return i
}

// Right operand node.
func (i *ElBinaryOperatorBuilder) SetRightOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.elbinaryoperator.RightOperand = v
	return i
}

// From: ElOperator
/**
True if the natural precedence of operators is overridden in the expression
represented by this node of the expression tree. If True, parentheses should be
introduced around the totality of the syntax expression corresponding to this
operator node and its operands.
*/
func (i *ElBinaryOperatorBuilder) SetPrecedenceOverridden(v bool) *ElBinaryOperatorBuilder {
	i.elbinaryoperator.PrecedenceOverridden = v
	return i
}

// From: ElOperator
/**
The symbol actually used in the expression, or intended to be used for
serialisation. Must be a member of OPERATOR_DEF. symbols .
*/
func (i *ElBinaryOperatorBuilder) SetSymbol(v string) *ElBinaryOperatorBuilder {
	i.elbinaryoperator.Symbol = v
	return i
}

// From: ElOperator
/**
Function call equivalent to this operator expression, inferred by matching
operator against functions defined in interface of principal operand.
*/
func (i *ElBinaryOperatorBuilder) SetCall(v IElFunctionCall) *ElBinaryOperatorBuilder {
	i.elbinaryoperator.Call = v
	return i
}

func (i *ElBinaryOperatorBuilder) Build() *ElBinaryOperator {
	return i.elbinaryoperator
}

// FUNCTIONS
// From: EL_OPERATOR
// Operator definition derived from definition.operator_definition() .
func (e *ElBinaryOperator) OperatorDefinition() IBmmOperator {
	return nil
}

// From: EL_OPERATOR
// Function call equivalent to this operator.
func (e *ElBinaryOperator) EquivalentCall() IElFunctionCall {
	return nil
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElBinaryOperator) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElBinaryOperator) IsBoolean() bool {
	return false
}
