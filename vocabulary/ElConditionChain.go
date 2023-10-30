package vocabulary

/**
Compound expression consisting of a chain of condition-gated expressions, and an
ungated else member that as a whole, represents an if/then/elseif/else chains.
Evaluated by iterating through items and for each one, evaluating its condition
, which if True, causes the evaluation result of the chain to be that item’s
result evaluation result. If no member of items has a True-returning condition ,
the evaluation result is the result of evaluating the else expression.
*/

// Interface definition
type IElConditionChain[T IBmmSimpleType] interface {
	// From: EL_DECISION_TABLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElConditionChain[T IBmmSimpleType] struct {
	// embedded for Inheritance
	ElDecisionTable[T]
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	Items []IElConditionalExpression[T] `yaml:"items" json:"items" xml:"items"`
}

// CONSTRUCTOR
func NewElConditionChain[T IBmmSimpleType]() *ElConditionChain[T] {
	elconditionchain := new(ElConditionChain[T])
	// Constants
	return elconditionchain
}

// BUILDER
type ElConditionChainBuilder[T IBmmSimpleType] struct {
	elconditionchain *ElConditionChain[T]
}

func NewElConditionChainBuilder[T IBmmSimpleType]() *ElConditionChainBuilder[T] {
	return &ElConditionChainBuilder[T]{
		elconditionchain: NewElConditionChain[T](),
	}
}

//BUILDER ATTRIBUTES
/**
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElConditionChainBuilder[T]) SetItems(v []IElConditionalExpression[T]) *ElConditionChainBuilder[T] {
	i.elconditionchain.Items = v
	return i
}

// From: ElDecisionTable
// Result expression of conditional, if its condition evaluates to True.
func (i *ElConditionChainBuilder[T]) SetElse(v T) *ElConditionChainBuilder[T] {
	i.elconditionchain.Else = v
	return i
}

func (i *ElConditionChainBuilder[T]) Build() *ElConditionChain[T] {
	return i.elconditionchain
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElConditionChain[T]) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElConditionChain[T]) IsBoolean() bool {
	return false
}
