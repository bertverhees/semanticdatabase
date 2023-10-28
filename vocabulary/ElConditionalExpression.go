package vocabulary

/**
Conditional structure used in condition chain expressions. Evaluated by
evaluating its condition , which is a Boolean-returning expression, and if this
returns True, the result is the evaluation result of expression .
*/

// Interface definition
type IElConditionalExpression interface {
	// From: EL_DECISION_BRANCH
}

// Struct definition
type ElConditionalExpression struct {
	// embedded for Inheritance
	ElDecisionBranch
	// Constants
	// Attributes
	// Boolean expression defining the condition of this decision branch.
	Condition IElExpression `yaml:"condition" json:"condition" xml:"condition"`
}

// CONSTRUCTOR
func NewElConditionalExpression() *ElConditionalExpression {
	elconditionalexpression := new(ElConditionalExpression)
	// Constants
	return elconditionalexpression
}

// BUILDER
type ElConditionalExpressionBuilder struct {
	elconditionalexpression *ElConditionalExpression
}

func NewElConditionalExpressionBuilder() *ElConditionalExpressionBuilder {
	return &ElConditionalExpressionBuilder{
		elconditionalexpression: NewElConditionalExpression(),
	}
}

// BUILDER ATTRIBUTES
// Boolean expression defining the condition of this decision branch.
func (i *ElConditionalExpressionBuilder) SetCondition(v IElExpression) *ElConditionalExpressionBuilder {
	i.elconditionalexpression.Condition = v
	return i
}

// From: ElDecisionBranch
// Result expression of conditional, if its condition evaluates to True.
func (i *ElConditionalExpressionBuilder) SetResult(v T) *ElConditionalExpressionBuilder {
	i.elconditionalexpression.Result = v
	return i
}

func (i *ElConditionalExpressionBuilder) Build() *ElConditionalExpression {
	return i.elconditionalexpression
}

//FUNCTIONS
