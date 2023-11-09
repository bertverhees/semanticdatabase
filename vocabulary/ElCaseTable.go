package vocabulary

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
}

// Struct definition
type ElCaseTable[T IElTerminal] struct {
	ElDecisionTable[T]
	// Attributes
	// Expressing generating the input value for the case table.
	TestValue IElValueGenerator `yaml:"testvalue" json:"testvalue" xml:"testvalue"`
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	Items []IElCase[T] `yaml:"items" json:"items" xml:"items"`
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
}

func NewElCaseTableBuilder[T IElTerminal]() *ElCaseTableBuilder[T] {
	return &ElCaseTableBuilder[T]{
		elcasetable: NewElCaseTable[T](),
	}
}

// BUILDER ATTRIBUTES
// Expressing generating the input value for the case table.
func (i *ElCaseTableBuilder[T]) SetTestValue(v IElValueGenerator) *ElCaseTableBuilder[T] {
	i.elcasetable.TestValue = v
	return i
}

/*
*
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElCaseTableBuilder[T]) SetItems(v []IElCase[T]) *ElCaseTableBuilder[T] {
	i.elcasetable.Items = v
	return i
}

// From: ElDecisionTable
// Result expression of conditional, if its condition evaluates to True.
func (i *ElCaseTableBuilder[T]) SetElse(v T) *ElCaseTableBuilder[T] {
	i.elcasetable.Else = v
	return i
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
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElCaseTable[T]) IsBoolean() bool {
	return false
}
