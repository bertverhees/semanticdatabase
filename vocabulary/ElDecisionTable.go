package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for decision tables. Generic on the meta-type of the result attribute
	of the branches, to allow specialised forms of if/else and case structures to be
	created.
*/

// Interface definition
type IElDecisionTable interface {
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElDecisionTable struct {
	// embedded for Inheritance
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	/**
		Members of the chain, equivalent to branches in an if/then/else chain and cases
		in a case statement.
	*/
	Items	List < EL_DECISION_BRANCH >	`yaml:"items" json:"items" xml:"items"`
	// Result expression of conditional, if its condition evaluates to True.
	Else	T	`yaml:"else" json:"else" xml:"else"`
}

//CONSTRUCTOR
func NewElDecisionTable() *ElDecisionTable {
	eldecisiontable := new(ElDecisionTable)
	// Constants
	// From: ElTerminal
	// From: ElExpression
	return eldecisiontable
}
//BUILDER
type ElDecisionTableBuilder struct {
	eldecisiontable *ElDecisionTable
}

func NewElDecisionTableBuilder() *ElDecisionTableBuilder {
	 return &ElDecisionTableBuilder {
		eldecisiontable : NewElDecisionTable(),
	}
}

//BUILDER ATTRIBUTES
/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
*/
func (i *ElDecisionTableBuilder) SetItems ( v List < EL_DECISION_BRANCH > ) *ElDecisionTableBuilder{
	i.eldecisiontable.Items = v
	return i
}
// Result expression of conditional, if its condition evaluates to True.
func (i *ElDecisionTableBuilder) SetElse ( v T ) *ElDecisionTableBuilder{
	i.eldecisiontable.Else = v
	return i
}
	// //From: ElTerminal
	// //From: ElExpression

func (i *ElDecisionTableBuilder) Build() *ElDecisionTable {
	 return i.eldecisiontable
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElDecisionTable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElDecisionTable) IsBoolean (  )  bool {
	return nil
}
