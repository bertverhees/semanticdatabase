package vocabulary

/**
Meta-type of read-only variables, including routine parameter and the special
variable 'Self'.
*/

// Interface definition
type IElReadonlyVariable interface {
	// From: EL_VARIABLE
	// From: EL_VALUE_GENERATOR
	Reference() string
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElReadonlyVariable struct {
	// embedded for Inheritance
	ElVariable
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Variable definition to which this reference refers.
	Definition IBmmReadonlyVariable `yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return False in all cases.
	IsWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// CONSTRUCTOR
func NewElReadonlyVariable() *ElReadonlyVariable {
	elreadonlyvariable := new(ElReadonlyVariable)
	// Constants
	return elreadonlyvariable
}

// BUILDER
type ElReadonlyVariableBuilder struct {
	elreadonlyvariable *ElReadonlyVariable
}

func NewElReadonlyVariableBuilder() *ElReadonlyVariableBuilder {
	return &ElReadonlyVariableBuilder{
		elreadonlyvariable: NewElReadonlyVariable(),
	}
}

// BUILDER ATTRIBUTES
// Variable definition to which this reference refers.
func (i *ElReadonlyVariableBuilder) SetDefinition(v IBmmReadonlyVariable) *ElReadonlyVariableBuilder {
	i.elreadonlyvariable.Definition = v
	return i
}

// Defined to return False in all cases.
func (i *ElReadonlyVariableBuilder) SetIsWritable(v bool) *ElReadonlyVariableBuilder {
	i.elreadonlyvariable.IsWritable = v
	return i
}

// From: ElValueGenerator
// Name used to represent the reference or other entity.
func (i *ElReadonlyVariableBuilder) SetName(v string) *ElReadonlyVariableBuilder {
	i.elreadonlyvariable.Name = v
	return i
}

func (i *ElReadonlyVariableBuilder) Build() *ElReadonlyVariable {
	return i.elreadonlyvariable
}

//FUNCTIONS
// From: EL_VALUE_GENERATOR
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElReadonlyVariable) Reference() string {
	return ""
}

// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElReadonlyVariable) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElReadonlyVariable) IsBoolean() bool {
	return false
}
