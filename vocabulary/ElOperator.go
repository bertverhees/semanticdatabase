package vocabulary

// Abstract parent of operator types.

// Interface definition
type IElOperator interface {
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
	//EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
}

// Struct definition
type ElOperator struct {
	// embedded for Inheritance
	ElExpression
	// Constants
	// Attributes
	/**
	True if the natural precedence of operators is overridden in the expression
	represented by this node of the expression tree. If True, parentheses should be
	introduced around the totality of the syntax expression corresponding to this
	operator node and its operands.
	*/
	PrecedenceOverridden bool `yaml:"precedenceoverridden" json:"precedenceoverridden" xml:"precedenceoverridden"`
	/**
	The symbol actually used in the expression, or intended to be used for
	serialisation. Must be a member of OPERATOR_DEF. symbols .
	*/
	Symbol string `yaml:"symbol" json:"symbol" xml:"symbol"`
	/**
	Function call equivalent to this operator expression, inferred by matching
	operator against functions defined in interface of principal operand.
	*/
	Call IElFunctionCall `yaml:"call" json:"call" xml:"call"`
}

// CONSTRUCTOR
func NewElOperator() *ElOperator {
	eloperator := new(ElOperator)
	// Constants
	return eloperator
}

// BUILDER
type ElOperatorBuilder struct {
	eloperator *ElOperator
}

func NewElOperatorBuilder() *ElOperatorBuilder {
	return &ElOperatorBuilder{
		eloperator: NewElOperator(),
	}
}

//BUILDER ATTRIBUTES
/**
True if the natural precedence of operators is overridden in the expression
represented by this node of the expression tree. If True, parentheses should be
introduced around the totality of the syntax expression corresponding to this
operator node and its operands.
*/
func (i *ElOperatorBuilder) SetPrecedenceOverridden(v bool) *ElOperatorBuilder {
	i.eloperator.PrecedenceOverridden = v
	return i
}

/*
*
The symbol actually used in the expression, or intended to be used for
serialisation. Must be a member of OPERATOR_DEF. symbols .
*/
func (i *ElOperatorBuilder) SetSymbol(v string) *ElOperatorBuilder {
	i.eloperator.Symbol = v
	return i
}

/*
*
Function call equivalent to this operator expression, inferred by matching
operator against functions defined in interface of principal operand.
*/
func (i *ElOperatorBuilder) SetCall(v IElFunctionCall) *ElOperatorBuilder {
	i.eloperator.Call = v
	return i
}

func (i *ElOperatorBuilder) Build() *ElOperator {
	return i.eloperator
}

// FUNCTIONS
// Operator definition derived from definition.operator_definition() .
func (e *ElOperator) OperatorDefinition() IBmmOperator {
	return nil
}

// Function call equivalent to this operator.
func (e *ElOperator) EquivalentCall() IElFunctionCall {
	return nil
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElOperator) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElOperator) IsBoolean() bool {
	return false
}
