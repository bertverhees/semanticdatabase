package vocabulary

// Meta-type representing a value-generating simple expression.

// Interface definition
type IElValueGenerator interface {
	IElSimple
	//EL_VALUE_GENERATOR
	Reference() string
}

// Struct definition
type ElValueGenerator struct {
	ElSimple
	// Attributes
	IsWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// Name used to represent the reference or other entity.
	Name string `yaml:"name" json:"name" xml:"name"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder

//FUNCTIONS
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElValueGenerator) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElValueGenerator) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElValueGenerator) IsBoolean() bool {
	return false
}
