package vocabulary

// Abstract parent of all typed expression meta-types.

// Interface definition
type IElExpression interface {
	//EvalType() IBmmType //abstract
	IsBoolean() bool
}

// Struct definition
type ElExpression struct {
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElExpression) EvalType() IBmmType {
	return nil
}

/*
*
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElExpression) IsBoolean() bool {
	return false
}
