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
type IElConditionChain interface {
	// From: EL_DECISION_TABLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElConditionChain struct {
	// embedded for Inheritance
	ElDecisionTable
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	Items []IElConditionalExpression `yaml:"items" json:"items" xml:"items"`
}

// CONSTRUCTOR
func NewElConditionChain() *ElConditionChain {
	elconditionchain := new(ElConditionChain)
	// Constants
	return elconditionchain
}

// BUILDER
type ElConditionChainBuilder struct {
	elconditionchain *ElConditionChain
}

func NewElConditionChainBuilder() *ElConditionChainBuilder {
	return &ElConditionChainBuilder{
		elconditionchain: NewElConditionChain(),
	}
}

//BUILDER ATTRIBUTES
/**
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElConditionChainBuilder) SetItems(v []IElConditionalExpression) *ElConditionChainBuilder {
	i.elconditionchain.Items = v
	return i
}

// From: ElDecisionTable
// Result expression of conditional, if its condition evaluates to True.
func (i *ElConditionChainBuilder) SetElse(v T) *ElConditionChainBuilder {
	i.elconditionchain.Else = v
	return i
}

func (i *ElConditionChainBuilder) Build() *ElConditionChain {
	return i.elconditionchain
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElConditionChain) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElConditionChain) IsBoolean() bool {
	return false
}
