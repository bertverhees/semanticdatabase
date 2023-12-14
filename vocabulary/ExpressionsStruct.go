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

func (e *ElValueGenerator) SetIsWritable(isWritable bool) error {
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
/**
Meta-type for reference to a non-abstract type as an object. Assumed to be
accessible at run-time. Typically represented syntactically as TypeName or
{TypeName} . May be used as a value, or as the qualifier for a function or
constant access.
*/
type ElTypeRef struct {
	ElValueGenerator
	// Attributes
	// _type, directly from the name of the reference, e.g. {SOME_TYPE} .
	_type     IBmmType `yaml:"type" json:"type" xml:"type"`
	isMutable bool     `yaml:"ismutable" json:"ismutable" xml:"ismutable"`
}

func (e *ElTypeRef) Type() IBmmType {
	return e._type
}

func (e *ElTypeRef) SetType(_type IBmmType) error {
	if _type == nil {
		return errors.New("The Type attribute from Typeref may not be set to nil")
	}
	e._type = _type
	return nil
}

func (e *ElTypeRef) IsMutable() bool {
	return e.isMutable
}

func (e *ElTypeRef) SetIsMutable(isMutable bool) error {
	e.isMutable = isMutable
	return nil
}

// CONSTRUCTOR
func NewElTypeRef() *ElTypeRef {
	eltyperef := new(ElTypeRef)
	// Constants
	return eltyperef
}

func (e *ElTypeRef) EvalType() IBmmType {
	return nil
}

/* ======================= ElLiteral ==================== */
/**
Literal value of any type known in the model, including primitive types. Defined
via a BMM_LITERAL_VALUE .
*/
type ElLiteral struct {
	ElSimple
	// Attributes
	// The reference item from which the value of this node can be computed.
	value IBmmLiteralValue[IBmmType] `yaml:"value" json:"value" xml:"value"`
}

func (e *ElLiteral) Value() IBmmLiteralValue[IBmmType] {
	return e.value
}

func (e *ElLiteral) SetValue(value IBmmLiteralValue[IBmmType]) error {
	e.value = value
	return nil
}

// CONSTRUCTOR
func NewElLiteral() *ElLiteral {
	elliteral := new(ElLiteral)
	// Constants
	return elliteral
}

// FUNCTIONS
// Return value.type .
func (e *ElLiteral) EvalType() IBmmType {
	return nil
}

/* ======================= ElVariable ==================== */
// Abstract meta-type of any kind of symbolic variable.
type ElVariable struct {
	ElValueGenerator
	// Attributes
}

/* ======================= ElWritableVariable ==================== */
/**
Meta-type of writable variables, including routine locals and the special
variable 'result'.
*/
type ElWritableVariable struct {
	ElVariable
	// Attributes
	// Variable definition to which this reference refers.
	definition IBmmWritableVariable `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElWritableVariable) Definition() IBmmWritableVariable {
	return e.definition
}

func (e *ElWritableVariable) SetDefinition(definition IBmmWritableVariable) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElWritableVariable() *ElWritableVariable {
	elwritablevariable := new(ElWritableVariable)
	elwritablevariable.isWritable = true
	// Constants
	return elwritablevariable
}

/* ======================= ElReadonlyVariable ==================== */
type ElReadonlyVariable struct {
	ElVariable
	// Attributes
	// Variable definition to which this reference refers.
	definition IBmmReadonlyVariable `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElReadonlyVariable) Definition() IBmmReadonlyVariable {
	return e.definition
}

func (e *ElReadonlyVariable) SetDefinition(definition IBmmReadonlyVariable) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElReadonlyVariable() *ElReadonlyVariable {
	elreadonlyvariable := new(ElReadonlyVariable)
	elreadonlyvariable.isWritable = false
	// Constants
	return elreadonlyvariable
}

/* ======================= ElFeatureRef ==================== */
/**
A reference that is scoped by a containing entity and requires a context
qualifier if it is not the currently scoping entity.
*/
type ElFeatureRef struct {
	ElValueGenerator
	// Attributes
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	scoper IElValueGenerator `yaml:"scoper" json:"scoper" xml:"scoper"`
}

func (e *ElFeatureRef) Scoper() IElValueGenerator {
	return e.scoper
}

func (e *ElFeatureRef) SetScoper(scoper IElValueGenerator) error {
	e.scoper = scoper
	return nil
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

/* ======================= ElPropertyRef ==================== */
type ElPropertyRef struct {
	ElFeatureRef
	// Attributes
	// Property definition (within class).
	definition IBmmProperty `yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True.
}

func (e *ElPropertyRef) Definition() IBmmProperty {
	return e.definition
}

func (e *ElPropertyRef) SetDefinition(definition IBmmProperty) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElPropertyRef() *ElPropertyRef {
	elpropertyref := new(ElPropertyRef)
	elpropertyref.isWritable = true
	return elpropertyref
}

//FUNCTIONS
/**
_type definition (i.e. BMM meta-type definition object) of the constant, property
or variable, inferred by inspection of the current scoping instance. Return
definition.type .
*/
func (e *ElPropertyRef) EvalType() IBmmType {
	return nil
}

/* ======================= ElStaticRef ==================== */
// Reference to a writable property, either a constant or computed.
type ElStaticRef struct {
	ElFeatureRef
	// Attributes
	// Constant definition (within class).
	definition IBmmStatic `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElStaticRef) Definition() IBmmStatic {
	return e.definition
}

func (e *ElStaticRef) SetDefinition(definition IBmmStatic) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
func NewElStaticRef() *ElStaticRef {
	elstaticref := new(ElStaticRef)
	elstaticref.isWritable = false
	// Constants
	return elstaticref
}

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
