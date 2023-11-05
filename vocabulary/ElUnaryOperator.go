package vocabulary

// Unary operator expression node.

// Interface definition
type IElUnaryOperator interface {
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
	//EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
}

// Struct definition
type ElUnaryOperator struct {
	// embedded for Inheritance
	ElOperator
	ElExpression
	// Constants
	// Attributes
	// Operand node.
	Operand IElExpression `yaml:"operand" json:"operand" xml:"operand"`
}

// CONSTRUCTOR
func NewElUnaryOperator() *ElUnaryOperator {
	elunaryoperator := new(ElUnaryOperator)
	// Constants
	return elunaryoperator
}

// BUILDER
type ElUnaryOperatorBuilder struct {
	elunaryoperator *ElUnaryOperator
}

func NewElUnaryOperatorBuilder() *ElUnaryOperatorBuilder {
	return &ElUnaryOperatorBuilder{
		elunaryoperator: NewElUnaryOperator(),
	}
}

// BUILDER ATTRIBUTES
// Operand node.
func (i *ElUnaryOperatorBuilder) SetOperand(v IElExpression) *ElUnaryOperatorBuilder {
	i.elunaryoperator.Operand = v
	return i
}

// From: ElOperator
/**
True if the natural precedence of operators is overridden in the expression
represented by this node of the expression tree. If True, parentheses should be
introduced around the totality of the syntax expression corresponding to this
operator node and its operands.
*/
func (i *ElUnaryOperatorBuilder) SetPrecedenceOverridden(v bool) *ElUnaryOperatorBuilder {
	i.elunaryoperator.PrecedenceOverridden = v
	return i
}

// From: ElOperator
/**
The symbol actually used in the expression, or intended to be used for
serialisation. Must be a member of OPERATOR_DEF. symbols .
*/
func (i *ElUnaryOperatorBuilder) SetSymbol(v string) *ElUnaryOperatorBuilder {
	i.elunaryoperator.Symbol = v
	return i
}

// From: ElOperator
/**
Function call equivalent to this operator expression, inferred by matching
operator against functions defined in interface of principal operand.
*/
func (i *ElUnaryOperatorBuilder) SetCall(v IElFunctionCall) *ElUnaryOperatorBuilder {
	i.elunaryoperator.Call = v
	return i
}

func (i *ElUnaryOperatorBuilder) Build() *ElUnaryOperator {
	return i.elunaryoperator
}

// FUNCTIONS
// From: EL_OPERATOR
// Operator definition derived from definition.operator_definition() .
func (e *ElUnaryOperator) OperatorDefinition() IBmmOperator {
	return nil
}

// From: EL_OPERATOR
// Function call equivalent to this operator.
func (e *ElUnaryOperator) EquivalentCall() IElFunctionCall {
	return nil
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElUnaryOperator) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElUnaryOperator) IsBoolean() bool {
	return false
}
