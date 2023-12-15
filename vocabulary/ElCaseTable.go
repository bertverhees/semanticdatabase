package vocabulary

import "errors"

/**
Compound expression consisting of a list of value-range / expression pairs, and
an else member that as a whole, represents a case statement flavour of decision
table. Evaluated by iterating through items and for each one, comparing input to
the item value_range . If the input is in the range, the evaluation result of
the table is that item’s result evaluation result. If no member of items has a
True-returning condition , the evaluation result is the result of evaluating the
else expression.
*/

// Interface definition
type IElCaseTable[T IElTerminal] interface {
	IElDecisionTable[T]
	TestValue() IElValueGenerator
	SetTestValue(testValue IElValueGenerator) error
}

// Struct definition
type ElCaseTable[T IElTerminal] struct {
	ElDecisionTable[T]
	// Attributes
	// Expressing generating the input value for the case table.
	testValue IElValueGenerator `yaml:"testvalue" json:"testvalue" xml:"testvalue"`
}

func (e *ElCaseTable[T]) TestValue() IElValueGenerator {
	return e.testValue
}

func (e *ElCaseTable[T]) SetTestValue(testValue IElValueGenerator) error {
	e.testValue = testValue
	return nil
}

func (e *ElCaseTable[T]) SetItems(items []IElDecisionBranch[T]) error {
	e.items = make([]IElDecisionBranch[T], 0)
	for _, s := range items {
		s, ok := s.(IElConditionalExpression[T])
		if !ok {
			return errors.New("_type-assertion in ElCaseTable[T]->SetItems failed")
		} else {
			e.items = append(e.items, s)
		}
	}
	return nil
}

// CONSTRUCTOR
func NewElCaseTable[T IElTerminal]() *ElCaseTable[T] {
	elcasetable := new(ElCaseTable[T])
	// Constants
	return elcasetable
}

// BUILDER
type ElCaseTableBuilder[T IElTerminal] struct {
	elcasetable *ElCaseTable[T]
	errors      []error
}

func NewElCaseTableBuilder[T IElTerminal]() *ElCaseTableBuilder[T] {
	return &ElCaseTableBuilder[T]{
		elcasetable: NewElCaseTable[T](),
		errors:      make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Expressing generating the input value for the case table.
func (i *ElCaseTableBuilder[T]) SetTestValue(v IElValueGenerator) *ElCaseTableBuilder[T] {
	i.AddError(i.elcasetable.SetTestValue(v))
	return i
}

/*
*
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElCaseTableBuilder[T]) SetItems(v []IElDecisionBranch[T]) *ElCaseTableBuilder[T] {
	i.AddError(i.elcasetable.SetItems(v))
	return i
}

// From: ElDecisionTable
// result expression of conditional, if its condition evaluates to True.
func (i *ElCaseTableBuilder[T]) SetElse(v T) *ElCaseTableBuilder[T] {
	i.AddError(i.elcasetable.SetElse(v))
	return i
}

func (i *ElCaseTableBuilder[T]) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElCaseTableBuilder[T]) Build() *ElCaseTable[T] {
	return i.elcasetable
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElCaseTable[T]) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElCaseTable[T]) IsBoolean() bool {
	return false
}
