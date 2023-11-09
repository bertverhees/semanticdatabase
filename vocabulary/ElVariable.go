package vocabulary

// Abstract meta-type of any kind of symbolic variable.

// Interface definition
type IElVariable interface {
	IElValueGenerator
	//EL_VARIABLE
}

// Struct definition
type ElVariable struct {
	ElValueGenerator
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
// From: EL_VALUE_GENERATOR
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElVariable) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElVariable) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElVariable) IsBoolean() bool {
	return false
}
