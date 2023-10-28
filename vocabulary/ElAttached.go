package vocabulary

/**
A predicate on any object reference (including function call) that returns True
if the reference is attached, i.e. non-Void.
*/

// Interface definition
type IElAttached interface {
	// From: EL_PREDICATE
	EvalType() IBmmSimpleType
	IsBoolean() bool
}

// Struct definition
type ElAttached struct {
	// embedded for Inheritance
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewElAttached() *ElAttached {
	elattached := new(ElAttached)
	// Constants
	return elattached
}

// BUILDER
type ElAttachedBuilder struct {
	elattached *ElAttached
}

func NewElAttachedBuilder() *ElAttachedBuilder {
	return &ElAttachedBuilder{
		elattached: NewElAttached(),
	}
}

// BUILDER ATTRIBUTES
// From: ElPredicate
// The target instance of this predicate.
func (i *ElAttachedBuilder) SetOperand(v IElValueGenerator) *ElAttachedBuilder {
	i.elattached.Operand = v
	return i
}

func (i *ElAttachedBuilder) Build() *ElAttached {
	return i.elattached
}

// FUNCTIONS
// From: EL_PREDICATE
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElAttached) EvalType() IBmmSimpleType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElAttached) IsBoolean() bool {
	return false
}
