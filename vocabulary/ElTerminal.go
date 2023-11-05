package vocabulary

/**
Expression entities that are terminals (i.e. leaves) within operator expressions
or tuples.
*/

// Interface definition
type IElTerminal interface {
	// From: EL_EXPRESSION
	//EvalType() IBmmType //abstract
	IsBoolean() bool
	//EL_TERMINAL
}

// Struct definition
type ElTerminal struct {
	// embedded for Inheritance
	ElExpression
	// Constants
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElTerminal) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElTerminal) IsBoolean() bool {
	return false
}
