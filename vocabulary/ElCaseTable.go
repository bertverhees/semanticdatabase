package vocabulary

/**
Compound expression consisting of a list of value-range / expression pairs, and
an else member that as a whole, represents a case statement flavour of decision
table. Evaluated by iterating through items and for each one, comparing input to
the item value_range . If the input is in the range, the evaluation result of
the table is that item’s result evaluation result. If no member of items has a
True-returning condition , the evaluation result is the result of evaluating the
else expression.
*/

// Interface definition
type IElCaseTable interface {
	// From: EL_DECISION_TABLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType() IBmmType
	IsBoolean() bool
}

// Struct definition
type ElCaseTable struct {
	// embedded for Inheritance
	ElDecisionTable
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Expressing generating the input value for the case table.
	TestValue IElValueGenerator `yaml:"testvalue" json:"testvalue" xml:"testvalue"`
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	Items []IElCase `yaml:"items" json:"items" xml:"items"`
}

// CONSTRUCTOR
func NewElCaseTable() *ElCaseTable {
	elcasetable := new(ElCaseTable)
	// Constants
	return elcasetable
}

// BUILDER
type ElCaseTableBuilder struct {
	elcasetable *ElCaseTable
}

func NewElCaseTableBuilder() *ElCaseTableBuilder {
	return &ElCaseTableBuilder{
		elcasetable: NewElCaseTable(),
	}
}

// BUILDER ATTRIBUTES
// Expressing generating the input value for the case table.
func (i *ElCaseTableBuilder) SetTestValue(v IElValueGenerator) *ElCaseTableBuilder {
	i.elcasetable.TestValue = v
	return i
}

/*
*
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElCaseTableBuilder) SetItems(v []IElCase) *ElCaseTableBuilder {
	i.elcasetable.Items = v
	return i
}

// From: ElDecisionTable
// Result expression of conditional, if its condition evaluates to True.
func (i *ElCaseTableBuilder) SetElse(v T) *ElCaseTableBuilder {
	i.elcasetable.Else = v
	return i
}

func (i *ElCaseTableBuilder) Build() *ElCaseTable {
	return i.elcasetable
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
Meta-type of expression entity used in type-checking and evaluation. Effected in
descendants.
*/
func (e *ElCaseTable) EvalType() IBmmType {
	return nil
}

// From: EL_EXPRESSION
/**
Post_result : Result = eval_type().equal(
{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElCaseTable) IsBoolean() bool {
	return false
}
