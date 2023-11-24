package vocabulary

/**
Abstract parent for second-order constrained forms of first-order expression
meta-types.
*/

// Interface definition
type IElConstrained interface {
	IElExpression
}

// Struct definition
type ElConstrained struct {
	ElExpression
	// Attributes
	// The base expression of this constrained form.
	BaseExpression IElExpression `yaml:"baseexpression" json:"baseexpression" xml:"baseexpression"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElConstrained) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElConstrained) IsBoolean() bool {
	return false
}
