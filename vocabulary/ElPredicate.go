package vocabulary

// Parent type of predicate of any object reference.

// Interface definition
type IElPredicate interface {
	IElSimple
	//EL_PREDICATE
	EvalType() IBmmSimpleType
}

// Struct definition
type ElPredicate struct {
	ElSimple
	// Attributes
	// The target instance of this predicate.
	Operand IElValueGenerator `yaml:"operand" json:"operand" xml:"operand"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElPredicate) EvalType() IBmmSimpleType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElPredicate) IsBoolean() bool {
	return false
}
