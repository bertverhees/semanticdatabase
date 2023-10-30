package vocabulary

/**
Conditional structure used in condition chain expressions. Evaluated by
evaluating its condition , which is a Boolean-returning expression, and if this
returns True, the result is the evaluation result of expression .
*/

// Interface definition
type IElConditionalExpression[T IBmmSimpleType] interface {
	// From: EL_DECISION_BRANCH
}

// Struct definition
type ElConditionalExpression[T IBmmSimpleType] struct {
	// embedded for Inheritance
	ElDecisionBranch[T]
	// Constants
	// Attributes
	// Boolean expression defining the condition of this decision branch.
	Condition IElExpression `yaml:"condition" json:"condition" xml:"condition"`
}

// CONSTRUCTOR
func NewElConditionalExpression[T IBmmSimpleType]() *ElConditionalExpression[T] {
	elconditionalexpression := new(ElConditionalExpression[T])
	// Constants
	return elconditionalexpression
}

// BUILDER
type ElConditionalExpressionBuilder[T IBmmSimpleType] struct {
	elconditionalexpression *ElConditionalExpression[T]
}

func NewElConditionalExpressionBuilder[T IBmmSimpleType]() *ElConditionalExpressionBuilder[T] {
	return &ElConditionalExpressionBuilder[T]{
		elconditionalexpression: NewElConditionalExpression[T](),
	}
}

// BUILDER ATTRIBUTES
// Boolean expression defining the condition of this decision branch.
func (i *ElConditionalExpressionBuilder[T]) SetCondition(v IElExpression) *ElConditionalExpressionBuilder[T] {
	i.elconditionalexpression.Condition = v
	return i
}

// From: ElDecisionBranch
// Result expression of conditional, if its condition evaluates to True.
func (i *ElConditionalExpressionBuilder[T]) SetResult(v T) *ElConditionalExpressionBuilder[T] {
	i.elconditionalexpression.Result = v
	return i
}

func (i *ElConditionalExpressionBuilder[T]) Build() *ElConditionalExpression[T] {
	return i.elconditionalexpression
}

//FUNCTIONS
