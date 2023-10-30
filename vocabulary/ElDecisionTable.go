package vocabulary

/**
Meta-type for decision tables. Generic on the meta-type of the result attribute
of the branches, to allow specialised forms of if/else and case structures to be
created.
*/

// Interface definition
type IElDecisionTable[T IBmmSimpleType] interface {
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElDecisionTable[T IBmmSimpleType] struct {
	// embedded for Inheritance
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	Items []IElDecisionBranch `yaml:"items" json:"items" xml:"items"`
	// Result expression of conditional, if its condition evaluates to True.
	Else T `yaml:"else" json:"else" xml:"else"`
}

// CONSTRUCTOR
func NewElDecisionTable[T IBmmSimpleType]() *ElDecisionTable[T] {
	eldecisiontable := new(ElDecisionTable[T])
	// Constants
	return eldecisiontable
}

// BUILDER
type ElDecisionTableBuilder[T IBmmSimpleType] struct {
	eldecisiontable *ElDecisionTable[T]
}

func NewElDecisionTableBuilder[T IBmmSimpleType]() *ElDecisionTableBuilder[T] {
	return &ElDecisionTableBuilder[T]{
		eldecisiontable: NewElDecisionTable[T](),
	}
}

//BUILDER ATTRIBUTES
/**
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElDecisionTableBuilder[T]) SetItems(v []IElDecisionBranch) *ElDecisionTableBuilder[T] {
	i.eldecisiontable.Items = v
	return i
}

// Result expression of conditional, if its condition evaluates to True.
func (i *ElDecisionTableBuilder[T]) SetElse(v T) *ElDecisionTableBuilder[T] {
	i.eldecisiontable.Else = v
	return i
}

func (i *ElDecisionTableBuilder[T]) Build() *ElDecisionTable[T] {
	return i.eldecisiontable
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElDecisionTable[T]) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElDecisionTable[T]) IsBoolean() bool {
	return false
}
