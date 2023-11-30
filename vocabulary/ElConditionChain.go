package vocabulary

import "errors"

/**
Compound expression consisting of a chain of condition-gated expressions, and an
ungated else member that as a whole, represents an if/then/elseif/else chains.
Evaluated by iterating through items and for each one, evaluating its condition
, which if True, causes the evaluation result of the chain to be that item’s
result evaluation result. If no member of items has a True-returning condition ,
the evaluation result is the result of evaluating the else expression.
*/

// Interface definition
type IElConditionChain[T IElTerminal] interface {
	IElDecisionTable[T]
}

// Struct definition
type ElConditionChain[T IElTerminal] struct {
	ElDecisionTable[T]
	// Constants
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	items []IElConditionalExpression[T] `yaml:"items" json:"items" xml:"items"`
}

func (e *ElConditionChain[T]) SetItems(items []IElTerminal) error {
	e.items = make([]IElConditionalExpression[T], 0)
	for _, s := range items {
		s, ok := s.(IElConditionalExpression[T])
		if !ok {
			return errors.New("_type-assertion in ElConditionChain[T]->SetItems went wrong")
		} else {
			e.items = append(e.items, s)
		}
	}
	return nil
}

// CONSTRUCTOR
func NewElConditionChain[T IElTerminal]() *ElConditionChain[T] {
	elconditionchain := new(ElConditionChain[T])
	// Constants
	return elconditionchain
}

// BUILDER
type ElConditionChainBuilder[T IElTerminal] struct {
	elconditionchain *ElConditionChain[T]
	errors           []error
}

func NewElConditionChainBuilder[T IElTerminal]() *ElConditionChainBuilder[T] {
	return &ElConditionChainBuilder[T]{
		elconditionchain: NewElConditionChain[T](),
		errors:           make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElConditionChainBuilder[T]) SetItems(v []IElTerminal) *ElConditionChainBuilder[T] {
	i.AddError(i.elconditionchain.SetItems(v))
	return i
}

// From: ElDecisionTable
// result expression of conditional, if its condition evaluates to True.
func (i *ElConditionChainBuilder[T]) SetElse(v T) *ElConditionChainBuilder[T] {
	i.elconditionchain._else = v
	return i
}

func (i *ElConditionChainBuilder[T]) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElConditionChain[T]) IsBoolean() bool {
	return false
}
