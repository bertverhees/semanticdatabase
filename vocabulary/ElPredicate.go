package vocabulary

// Parent type of predicate of any object reference.

// Interface definition
type IElPredicate interface {
	// From: EL_EXPRESSION
	IsBoolean() bool
	// From: EL_TERMINAL
	//EL_SIMPLE
	//EL_PREDICATE
	EvalType() IBmmSimpleType
}

// Struct definition
type ElPredicate struct {
	// embedded for Inheritance
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// The target instance of this predicate.
	Operand IElValueGenerator `yaml:"operand" json:"operand" xml:"operand"`
}

// CONSTRUCTOR
func NewElPredicate() *ElPredicate {
	elpredicate := new(ElPredicate)
	// Constants
	return elpredicate
}

// BUILDER
type ElPredicateBuilder struct {
	elpredicate *ElPredicate
}

func NewElPredicateBuilder() *ElPredicateBuilder {
	return &ElPredicateBuilder{
		elpredicate: NewElPredicate(),
	}
}

// BUILDER ATTRIBUTES
// The target instance of this predicate.
func (i *ElPredicateBuilder) SetOperand(v IElValueGenerator) *ElPredicateBuilder {
	i.elpredicate.Operand = v
	return i
}

func (i *ElPredicateBuilder) Build() *ElPredicate {
	return i.elpredicate
}

// FUNCTIONS
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElPredicate) EvalType() IBmmSimpleType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElPredicate) IsBoolean() bool {
	return false
}
