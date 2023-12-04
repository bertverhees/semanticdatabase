package vocabulary

// Unary operator expression node.

// Interface definition
type IElUnaryOperator interface {
	IElOperator
	Operand() IElExpression
	SetOperand(operand IElExpression) error
}

// Struct definition
type ElUnaryOperator struct {
	ElOperator
	// Attributes
	// operand node.
	operand IElExpression `yaml:"operand" json:"operand" xml:"operand"`
}

func (e *ElUnaryOperator) Operand() IElExpression {
	return e.operand
}

func (e *ElUnaryOperator) SetOperand(operand IElExpression) error {
	e.operand = operand
	return nil
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
	errors          []error
}

func NewElUnaryOperatorBuilder() *ElUnaryOperatorBuilder {
	return &ElUnaryOperatorBuilder{
		elunaryoperator: NewElUnaryOperator(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// operand node.
func (i *ElUnaryOperatorBuilder) SetOperand(v IElExpression) *ElUnaryOperatorBuilder {
	i.AddError(i.elunaryoperator.SetOperand(v))
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
	i.AddError(i.elunaryoperator.SetPrecedenceOverridden(v))
	return i
}

// From: ElOperator
/**
The symbol actually used in the expression, or intended to be used for
serialisation. Must be a member of OPERATOR_DEF. symbols .
*/
func (i *ElUnaryOperatorBuilder) SetSymbol(v string) *ElUnaryOperatorBuilder {
	i.AddError(i.elunaryoperator.SetSymbol(v))
	return i
}

// From: ElOperator
/**
Function call equivalent to this operator expression, inferred by matching
operator against functions defined in interface of principal operand.
*/
func (i *ElUnaryOperatorBuilder) SetCall(v IElFunctionCall) *ElUnaryOperatorBuilder {
	i.AddError(i.elunaryoperator.SetCall(v))
	return i
}

func (i *ElUnaryOperatorBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElUnaryOperatorBuilder) Build() *ElUnaryOperator {
	return i.elunaryoperator
}

// FUNCTIONS
