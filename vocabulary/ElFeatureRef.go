package vocabulary

/**
A reference that is scoped by a containing entity and requires a context
qualifier if it is not the currently scoping entity.
*/

// Interface definition
type IElFeatureRef interface {
	IElValueGenerator
	//EL_FEATURE_REF
	Reference() string //redefined
}

// Struct definition
type ElFeatureRef struct {
	ElValueGenerator
	// Attributes
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	scoper IElValueGenerator `yaml:"scoper" json:"scoper" xml:"scoper"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
/**
Generated full reference name, consisting of scoping elements and name
concatenated using dot notation.
*/
func (e *ElFeatureRef) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElFeatureRef) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElFeatureRef) IsBoolean() bool {
	return false
}
