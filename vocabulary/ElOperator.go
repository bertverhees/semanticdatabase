package vocabulary

// Abstract parent of operator types.

// Interface definition
type IElOperator interface {
	IElExpression
	//EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
	PrecedenceOverridden() bool
	SetPrecedenceOverridden(precedenceOverridden bool) error
	Symbol() string
	SetSymbol(symbol string) error
	Call() IElFunctionCall
	SetCall(call IElFunctionCall) error
}

// Struct definition
type ElOperator struct {
	ElExpression
	// Attributes
	/**
	True if the natural precedence of operators is overridden in the expression
	represented by this node of the expression tree. If True, parentheses should be
	introduced around the totality of the syntax expression corresponding to this
	operator node and its operands.
	*/
	precedenceOverridden bool `yaml:"precedenceoverridden" json:"precedenceoverridden" xml:"precedenceoverridden"`
	/**
	The symbol actually used in the expression, or intended to be used for
	serialisation. Must be a member of OPERATOR_DEF. symbols .
	*/
	symbol string `yaml:"symbol" json:"symbol" xml:"symbol"`
	/**
	Function call equivalent to this operator expression, inferred by matching
	operator against functions defined in interface of principal operand.
	*/
	call IElFunctionCall `yaml:"call" json:"call" xml:"call"`
}

func (e *ElOperator) PrecedenceOverridden() bool {
	return e.precedenceOverridden
}

func (e *ElOperator) SetPrecedenceOverridden(precedenceOverridden bool) error {
	e.precedenceOverridden = precedenceOverridden
	return nil
}

func (e *ElOperator) Symbol() string {
	return e.symbol
}

func (e *ElOperator) SetSymbol(symbol string) error {
	e.symbol = symbol
	return nil
}

func (e *ElOperator) Call() IElFunctionCall {
	return e.call
}

func (e *ElOperator) SetCall(call IElFunctionCall) error {
	e.call = call
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Operator definition derived from definition.operator_definition() .
func (e *ElOperator) OperatorDefinition() IBmmOperator {
	return nil
}

// Function call equivalent to this operator.
func (e *ElOperator) EquivalentCall() IElFunctionCall {
	return nil
}
