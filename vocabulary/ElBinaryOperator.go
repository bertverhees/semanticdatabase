package vocabulary

// Binary operator expression node.

// Interface definition
type IElBinaryOperator interface {
	IElOperator
	LeftOperand() IElExpression
	SetLeftOperand(leftOperand IElExpression) error
	RightOperand() IElExpression
	SetRightOperand(rightOperand IElExpression) error
}

// Struct definition
type ElBinaryOperator struct {
	ElOperator
	// Attributes
	// Left operand node.
	leftOperand IElExpression `yaml:"leftoperand" json:"leftoperand" xml:"leftoperand"`
	// Right operand node.
	rightOperand IElExpression `yaml:"rightoperand" json:"rightoperand" xml:"rightoperand"`
}

func (e *ElBinaryOperator) LeftOperand() IElExpression {
	return e.leftOperand
}

func (e *ElBinaryOperator) SetLeftOperand(leftOperand IElExpression) error {
	e.leftOperand = leftOperand
	return nil
}

func (e *ElBinaryOperator) RightOperand() IElExpression {
	return e.rightOperand
}

func (e *ElBinaryOperator) SetRightOperand(rightOperand IElExpression) error {
	e.rightOperand = rightOperand
	return nil
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
	errors           []error
}

func NewElBinaryOperatorBuilder() *ElBinaryOperatorBuilder {
	return &ElBinaryOperatorBuilder{
		elbinaryoperator: NewElBinaryOperator(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Left operand node.
func (i *ElBinaryOperatorBuilder) SetLeftOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.AddError(i.elbinaryoperator.SetLeftOperand(v))
	return i
}

// Right operand node.
func (i *ElBinaryOperatorBuilder) SetRightOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.AddError(i.elbinaryoperator.SetRightOperand(v))
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
	i.AddError(i.elbinaryoperator.SetPrecedenceOverridden(v))
	return i
}

// From: ElOperator
/**
The symbol actually used in the expression, or intended to be used for
serialisation. Must be a member of OPERATOR_DEF. symbols .
*/
func (i *ElBinaryOperatorBuilder) SetSymbol(v string) *ElBinaryOperatorBuilder {
	i.AddError(i.elbinaryoperator.SetSymbol(v))
	return i
}

// From: ElOperator
/**
Function call equivalent to this operator expression, inferred by matching
operator against functions defined in interface of principal operand.
*/
func (i *ElBinaryOperatorBuilder) SetCall(v IElFunctionCall) *ElBinaryOperatorBuilder {
	i.AddError(i.elbinaryoperator.SetCall(v))
	return i
}

func (i *ElBinaryOperatorBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElBinaryOperator) IsBoolean() bool {
	return false
}
