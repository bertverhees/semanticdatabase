package vocabulary

/**
Conditional structure used in condition chain expressions. Evaluated by
evaluating its condition , which is a Boolean-returning expression, and if this
returns True, the result is the evaluation result of expression .
*/

// Interface definition
type IElConditionalExpression[T IElTerminal] interface {
	IElDecisionBranch[T]
	Condition() IElExpression
	SetCondition(condition IElExpression) error
}

// Struct definition
type ElConditionalExpression[T IElTerminal] struct {
	ElDecisionBranch[T]
	// Attributes
	// Boolean expression defining the condition of this decision branch.
	condition IElExpression `yaml:"condition" json:"condition" xml:"condition"`
}

func (e *ElConditionalExpression[T]) Condition() IElExpression {
	return e.condition
}

func (e *ElConditionalExpression[T]) SetCondition(condition IElExpression) error {
	e.condition = condition
	return nil
}

// CONSTRUCTOR
func NewElConditionalExpression[T IElTerminal]() *ElConditionalExpression[T] {
	elconditionalexpression := new(ElConditionalExpression[T])
	// Constants
	return elconditionalexpression
}

// BUILDER
type ElConditionalExpressionBuilder[T IElTerminal] struct {
	elconditionalexpression *ElConditionalExpression[T]
	errors                  []error
}

func NewElConditionalExpressionBuilder[T IElTerminal]() *ElConditionalExpressionBuilder[T] {
	return &ElConditionalExpressionBuilder[T]{
		elconditionalexpression: NewElConditionalExpression[T](),
		errors:                  make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Boolean expression defining the condition of this decision branch.
func (i *ElConditionalExpressionBuilder[T]) SetCondition(v IElExpression) *ElConditionalExpressionBuilder[T] {
	i.AddError(i.elconditionalexpression.SetCondition(v))
	return i
}

// From: ElDecisionBranch
// result expression of conditional, if its condition evaluates to True.
func (i *ElConditionalExpressionBuilder[T]) SetResult(v T) *ElConditionalExpressionBuilder[T] {
	i.AddError(i.elconditionalexpression.SetResult(v))
	return i
}

func (i *ElConditionalExpressionBuilder[T]) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElConditionalExpressionBuilder[T]) Build() *ElConditionalExpression[T] {
	return i.elconditionalexpression
}

//FUNCTIONS
