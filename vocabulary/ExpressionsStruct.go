package vocabulary

import "errors"

/* ======================= ElExpression ==================== */
// Abstract parent of all typed expression meta-types.
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

/* ======================= ElTerminal ==================== */
/**
expression entities that are terminals (i.e. leaves) within operator expressions
or tuples.
*/
type ElTerminal struct {
	ElExpression
	// Attributes
}

/* ======================= ElSimple ==================== */
// Simple terminal i.e. logically atomic expression element.
type ElSimple struct {
	ElTerminal
	// Attributes
}

/* ======================= ElValueGenerator ==================== */
// Meta-type representing a value-generating simple expression.
type ElValueGenerator struct {
	ElSimple
	// Attributes
	isWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// name used to represent the reference or other entity.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (e *ElValueGenerator) Name() string {
	return e.name
}

func (e *ElValueGenerator) SetName(name string) error {
	if name == "" {
		return errors.New("ElValueGenerator -> name may not be set to empty")
	}
	e.name = name
	return nil
}

func (e *ElValueGenerator) IsWritable() bool {
	return e.isWritable
}

func (e *ElValueGenerator) SetIsWritable(isWritable bool)error {
	e.isWritable = isWritable
	return nil
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

/* ======================= ElTypeRef ==================== */
/* ======================= ElLiteral ==================== */
/* ======================= ElVariable ==================== */
/* ======================= ElWritableVariable ==================== */
/* ======================= ElReadonlyVariable ==================== */
/* ======================= ElFeatureRef ==================== */
/* ======================= ElPropertyRef ==================== */
/* ======================= ElStaticRef ==================== */
/* ======================= ElAgentCall ==================== */
/* ======================= ElFunctionCall ==================== */
/* ======================= ElAgent ==================== */
/* ======================= ElFunctionAgent ==================== */
/* ======================= ElProcedureAgent ==================== */
/* ======================= ElPredicate ==================== */
/* ======================= ElDefined ==================== */
/* ======================= ElAttached ==================== */
/* ======================= ElDecisionTable ==================== */
/* ======================= ElDecisionBranch ==================== */
/* ======================= ElConditionChain ==================== */
/* ======================= ElConditionalExpression ==================== */
/* ======================= ElCaseTable ==================== */
/* ======================= ElCase ==================== */
/* ======================= ElUnaryOperator ==================== */
/* ======================= ElBinaryOperator ==================== */
/* ======================= ElTuple ==================== */
/* ======================= ElTupleItem ==================== */
/* ======================= ElConstrained ==================== */
/* ======================= ElBooleanExpression ==================== */
