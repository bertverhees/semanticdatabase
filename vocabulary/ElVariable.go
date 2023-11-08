package vocabulary

// Abstract meta-type of any kind of symbolic variable.

// Interface definition
type IElVariable interface {
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
	// From: EL_TERMINAL
	// From: EL_SIMPLE
	// From: EL_VALUE_GENERATOR
	Reference() string
	//EL_VARIABLE
}

// Struct definition
type ElVariable struct {
	// embedded for Inheritance
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewElVariable() *ElVariable {
	elvariable := new(ElVariable)
	// Constants
	return elvariable
}

// BUILDER
type ElVariableBuilder struct {
	elvariable *ElVariable
}

func NewElVariableBuilder() *ElVariableBuilder {
	return &ElVariableBuilder{
		elvariable: NewElVariable(),
	}
}

// BUILDER ATTRIBUTES
// From: ElValueGenerator
func (i *ElVariableBuilder) SetIsWritable(v bool) *ElVariableBuilder {
	i.elvariable.IsWritable = v
	return i
}

// From: ElValueGenerator
// name used to represent the reference or other entity.
func (i *ElVariableBuilder) SetName(v string) *ElVariableBuilder {
	i.elvariable.Name = v
	return i
}

func (i *ElVariableBuilder) Build() *ElVariable {
	return i.elvariable
}

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
