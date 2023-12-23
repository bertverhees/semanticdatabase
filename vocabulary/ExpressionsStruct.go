package vocabulary

import (
	"errors"
	"semanticdatabase/aom/constraints"
)

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
		return errors.New("ElValueGenerator -> name should not be set to empty")
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
{TypeName} . should be used as a value, or as the qualifier for a function or
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
		return errors.New("The Type attribute from Typeref should not be set to nil")
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
	eltyperef.isWritable = true
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
// A call made to a 'closed' agent, i.e. one with no remaining open arguments.
type ElAgentCall struct {
	// Attributes
	// The agent being called.
	agent IElAgent `yaml:"agent" json:"agent" xml:"agent"`
}

func (e *ElAgentCall) Agent() IElAgent {
	return e.agent
}

func (e *ElAgentCall) SetAgent(agent IElAgent) error {
	e.agent = agent
	return nil
}

/* ======================= ElFunctionCall ==================== */
/**
A call made on a closed function agent, returning a result. Equivalent to an
'application' of a function in Lambda calculus.
*/
type ElFunctionCall struct {
	ElAgentCall
	ElFeatureRef
	// Attributes
}

func (e *ElFunctionCall) SetAgent(agent IElAgent) error {
	s, ok := agent.(IElFunctionAgent)
	if !ok {
		return errors.New("_type-assertion to IElFunctionAgent in ElFunctionCall->SetAgent failed")
	} else {
		e.agent = s
		return nil
	}
}

// CONSTRUCTOR
func NewElFunctionCall() *ElFunctionCall {
	elfunctioncall := new(ElFunctionCall)
	elfunctioncall.isWritable = true
	// Constants
	return elfunctioncall
}

/* ======================= ElAgent ==================== */
/**
A delayed routine call, whose arguments should be open, partially closed or closed.
Generated by special reference to a routine, usually via a dedicated keyword,
such as 'lambda' or 'agent', or other special syntax. Instances should include
closed delayed calls like calculate_age (dob="1987-09-13", today="2019-06-03")
but also partially open calls such as format_structure (struct=?, style=3) ,
where struct is an open argument. Evaluation type (i.e. type of runtime
evaluated form) is BMM_SIGNATURE .
*/
type ElAgent struct {
	ElFeatureRef
	// Attributes
	// Closed arguments of a routine call as a tuple of objects.
	closedArgs IElTuple `yaml:"closedargs" json:"closedargs" xml:"closedargs"`
	/**
	Optional list of names of open arguments of the call. If not provided, and the
	name refers to a routine with more arguments than supplied in closed_args , the
	missing arguments are inferred from the definition .
	*/
	openArgs []string `yaml:"openargs" json:"openargs" xml:"openargs"`
	// Reference to definition of a routine for which this is an agent, if one exists.
	definition IBmmRoutine `yaml:"definition" json:"definition" xml:"definition"`
}

func (e *ElAgent) ClosedArgs() IElTuple {
	return e.closedArgs
}

func (e *ElAgent) SetClosedArgs(closedArgs IElTuple) error {
	e.closedArgs = closedArgs
	return nil
}

func (e *ElAgent) OpenArgs() []string {
	return e.openArgs
}

func (e *ElAgent) SetOpenArgs(openArgs []string) error {
	e.openArgs = openArgs
	return nil
}

func (e *ElAgent) Definition() IBmmRoutine {
	return e.definition
}

func (e *ElAgent) SetDefinition(definition IBmmRoutine) error {
	e.definition = definition
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder

//FUNCTIONS

// Post_result_validity : result = open_arguments = Void
// True if there are no open arguments.
func (e *ElAgent) IsCallable() bool {
	return false
}

// Generated full reference name, including scoping elements.
func (e *ElAgent) Reference() string {
	return ""
}

func (e *ElAgent) EvalType() IBmmType {
	//IBmmRoutineType
	return nil
}

/* ======================= ElFunctionAgent ==================== */
// An agent whose signature is of a function, i.e. has a result type.
type ElFunctionAgent struct {
	ElAgent
	// Attributes
}

func (e *ElFunctionAgent) SetDefinition(definition IBmmRoutine) error {
	s, ok := definition.(IBmmFunction)
	if !ok {
		return errors.New("_type-assertion to IBmmFunction in ElFunctionAgent->SetDefinition failed")
	} else {
		e.definition = s
		return nil
	}
}

// CONSTRUCTOR
func NewElFunctionAgent() *ElFunctionAgent {
	elfunctionagent := new(ElFunctionAgent)
	elfunctionagent.openArgs = make([]string, 0)
	elfunctionagent.isWritable = false
	// Constants
	return elfunctionagent
}

func (e *ElFunctionAgent) EvalType() IBmmType {
	//IBmmFunctionType
	return nil
}

/* ======================= ElProcedureAgent ==================== */
// An agent whose signature is of a procedure, i.e. has no result type.
type ElProcedureAgent struct {
	ElAgent
	// Attributes
}

func (e *ElProcedureAgent) SetDefinition(definition IBmmRoutine) error {
	s, ok := definition.(IBmmProcedure)
	if !ok {
		return errors.New("_type-assertion to IBmmProcedure in ElProcedureAgent->SetDefinition failed")
	} else {
		e.definition = s
		return nil
	}
}

// CONSTRUCTOR
func NewElProcedureAgent() *ElProcedureAgent {
	elprocedureagent := new(ElProcedureAgent)
	elprocedureagent.openArgs = make([]string, 0)
	elprocedureagent.isWritable = false
	return elprocedureagent
}

func (e *ElProcedureAgent) EvalType() IBmmType {
	//IBmmProcedureType
	return nil
}

/* ======================= ElPredicate ==================== */
type ElPredicate struct {
	ElSimple
	// Attributes
	// The target instance of this predicate.
	operand IElValueGenerator `yaml:"operand" json:"operand" xml:"operand"`
}

func (e *ElPredicate) Operand() IElValueGenerator {
	return e.operand
}

func (e *ElPredicate) SetOperand(operand IElValueGenerator) error {
	e.operand = operand
	return nil
}

// Return {BMM_MODEL}.boolean_type_definition().
func (e *ElPredicate) EvalType() IBmmType {
	//IBmmSimpleType
	return nil
}

/* ======================= ElDefined ==================== */
/**
A predicate taking one external variable reference argument, that returns true
if the reference is resolvable, i.e. the external value is obtainable. Note
probably to be removed.
*/
type ElDefined struct {
	ElPredicate
	// Attributes
}

// CONSTRUCTOR
func NewElDefined() *ElDefined {
	eldefined := new(ElDefined)
	return eldefined
}

/* ======================= ElAttached ==================== */
/*
A predicate on any object reference (including function call) that returns True if the reference is attached, i.e. non-Void.
*/
type ElAttached struct {
	ElPredicate
	// Attributes
}

// CONSTRUCTOR
func NewElAttached() *ElAttached {
	elattached := new(ElAttached)
	return elattached
}

/* ======================= ElDecisionTable ==================== */
/*
Meta-type for decision tables. Generic on the meta-type of the result attribute of the branches,
to allow specialised forms of if/else and case structures to be created.
*/
type ElDecisionTable[T IElTerminal] struct {
	ElTerminal
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	items []IElDecisionBranch[T] `yaml:"items" json:"items" xml:"items"`
	// result expression of conditional, if its condition evaluates to True.
	_else T `yaml:"else" json:"else" xml:"else"`
}

func (e *ElDecisionTable[T]) Items() []IElDecisionBranch[T] {
	return e.items
}

func (e *ElDecisionTable[T]) SetItems(items []IElDecisionBranch[T]) error {
	e.items = items
	return nil
}

func (e *ElDecisionTable[T]) Else() T {
	return e._else
}

func (e *ElDecisionTable[T]) SetElse(_else T) error {
	e._else = _else
	return nil
}

/* ======================= ElDecisionBranch ==================== */
/*
Abstract parent of meta-types representing a branch of some kind of decision structure.
Defines result as being of the generic type T.
*/
type ElDecisionBranch[T IElTerminal] struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result expression of conditional, if its condition evaluates to True.
	result T `yaml:"result" json:"result" xml:"result"`
}

func (e *ElDecisionBranch[T]) Result() T {
	return e.result
}

func (e *ElDecisionBranch[T]) SetResult(result T) error {
	e.result = result
	return nil
}

/* ======================= ElConditionChain ==================== */
/**
Compound expression consisting of a chain of condition-gated expressions, and an
ungated else member that as a whole, represents an if/then/elseif/else chains.
Evaluated by iterating through items and for each one, evaluating its condition
, which if True, causes the evaluation result of the chain to be that item’s
result evaluation result. If no member of items has a True-returning condition ,
the evaluation result is the result of evaluating the else expression.
*/
type ElConditionChain[T IElTerminal] struct {
	ElDecisionTable[T]
	// Constants
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	items []IElConditionalExpression[T] `yaml:"items" json:"items" xml:"items"`
}

func (e *ElConditionChain[T]) SetItems(items []IElDecisionBranch[T]) error {
	e.items = make([]IElConditionalExpression[T], 0)
	for _, s := range items {
		s, ok := s.(IElConditionalExpression[T])
		if !ok {
			return errors.New("_type-assertion to IElConditionalExpression[T] in ElConditionChain[T]->SetItems failed")
		} else {
			e.items = append(e.items, s)
		}
	}
	return nil
}

func NewElConditionChain[T IElTerminal]() *ElConditionChain[T] {
	elconditionchain := new(ElConditionChain[T])
	elconditionchain.items = make([]IElConditionalExpression[T], 0)
	// Constants
	return elconditionchain
}

/* ======================= ElConditionalExpression ==================== */
/**
Conditional structure used in condition chain expressions. Evaluated by
evaluating its condition , which is a Boolean-returning expression, and if this
returns True, the result is the evaluation result of expression .
*/
type ElConditionalExpression[T IElTerminal] struct {
	ElDecisionBranch[T]
	// Attributes
	// Boolean expression defining the condition of this decision branch.
	condition IElExpression `yaml:"condition" json:"condition" xml:"condition"`
}

func (e *ElConditionalExpression[T]) Condition() IElExpression {
	return e.condition
}

func (e *ElConditionalExpression[T]) SetCondition(condition IElExpression) error {
	e.condition = condition
	return nil
}

// CONSTRUCTOR
func NewElConditionalExpression[T IElTerminal]() *ElConditionalExpression[T] {
	elconditionalexpression := new(ElConditionalExpression[T])
	return elconditionalexpression
}

/* ======================= ElCaseTable ==================== */
type ElCaseTable[T IElTerminal] struct {
	ElDecisionTable[T]
	// Attributes
	// Expressing generating the input value for the case table.
	testValue IElValueGenerator `yaml:"testvalue" json:"testvalue" xml:"testvalue"`
}

func (e *ElCaseTable[T]) TestValue() IElValueGenerator {
	return e.testValue
}

func (e *ElCaseTable[T]) SetTestValue(testValue IElValueGenerator) error {
	e.testValue = testValue
	return nil
}

func (e *ElCaseTable[T]) SetItems(items []IElDecisionBranch[T]) error {
	e.items = make([]IElDecisionBranch[T], 0)
	for _, s := range items {
		s, ok := s.(IElCase[T])
		if !ok {
			return errors.New("_type-assertion to IElCase[T] in ElCaseTable[T]->SetItems failed")
		} else {
			e.items = append(e.items, s)
		}
	}
	return nil
}

// CONSTRUCTOR
func NewElCaseTable[T IElTerminal]() *ElCaseTable[T] {
	elcasetable := new(ElCaseTable[T])
	elcasetable.items = make([]IElDecisionBranch[T], 0)
	return elcasetable
}

/* ======================= ElCase ==================== */
/**
One branch of a Case table, consisting of a value constraint (the match
criterion) and a result, of the generic parameter type T.
*/
type ElCase[T IElTerminal] struct {
	ElDecisionBranch[T]
	valueConstraint constraints.ICObject `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

func (e *ElCase[T]) ValueConstraint() constraints.ICObject {
	return e.valueConstraint
}

func (e *ElCase[T]) SetValueConstraint(valueConstraint constraints.ICObject) error {
	e.valueConstraint = valueConstraint
	return nil
}

// CONSTRUCTOR
func NewElCase[T IElTerminal]() *ElCase[T] {
	elcase := new(ElCase[T])
	return elcase
}

/* ======================= ElOperator ==================== */
// Abstract parent of operator types.
type ElOperator struct {
	ElExpression
	// Attributes
	/**
	True if the natural precedence of operators is overridden in the expression
	represented by this node of the expression tree. If True, parentheses should be
	introduced around the totality of the syntax expression corresponding to this
	operator node and its operands.
	*/
	precedenceOverridden bool `yaml:"precedenceoverridden" json:"precedenceoverridden" xml:"precedenceoverridden"`
	/**
	The symbol actually used in the expression, or intended to be used for
	serialisation. Must be a member of OPERATOR_DEF. symbols .
	*/
	symbol string `yaml:"symbol" json:"symbol" xml:"symbol"`
	/**
	Function call equivalent to this operator expression, inferred by matching
	operator against functions defined in interface of principal operand.
	*/
	call IElFunctionCall `yaml:"call" json:"call" xml:"call"`
}

func (e *ElOperator) PrecedenceOverridden() bool {
	return e.precedenceOverridden
}

func (e *ElOperator) SetPrecedenceOverridden(precedenceOverridden bool) error {
	e.precedenceOverridden = precedenceOverridden
	return nil
}

func (e *ElOperator) Symbol() string {
	return e.symbol
}

func (e *ElOperator) SetSymbol(symbol string) error {
	e.symbol = symbol
	return nil
}

func (e *ElOperator) Call() IElFunctionCall {
	return e.call
}

func (e *ElOperator) SetCall(call IElFunctionCall) error {
	e.call = call
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Operator definition derived from definition.operator_definition() .
func (e *ElOperator) OperatorDefinition() IBmmOperator {
	return nil
}

// Function call equivalent to this operator.
func (e *ElOperator) EquivalentCall() IElFunctionCall {
	return nil
}

/* ======================= ElUnaryOperator ==================== */
// Unary operator expression node.
type ElUnaryOperator struct {
	ElOperator
	// operand node.
	operand IElExpression `yaml:"operand" json:"operand" xml:"operand"`
}

func (e *ElUnaryOperator) Operand() IElExpression {
	return e.operand
}

func (e *ElUnaryOperator) SetOperand(operand IElExpression) error {
	e.operand = operand
	return nil
}

// CONSTRUCTOR
func NewElUnaryOperator() *ElUnaryOperator {
	elunaryoperator := new(ElUnaryOperator)
	return elunaryoperator
}

/* ======================= ElBinaryOperator ==================== */
type ElBinaryOperator struct {
	ElOperator
	// Attributes
	// Left operand node.
	leftOperand IElExpression `yaml:"leftoperand" json:"leftoperand" xml:"leftoperand"`
	// Right operand node.
	rightOperand IElExpression `yaml:"rightoperand" json:"rightoperand" xml:"rightoperand"`
}

func (e *ElBinaryOperator) LeftOperand() IElExpression {
	return e.leftOperand
}

func (e *ElBinaryOperator) SetLeftOperand(leftOperand IElExpression) error {
	e.leftOperand = leftOperand
	return nil
}

func (e *ElBinaryOperator) RightOperand() IElExpression {
	return e.rightOperand
}

func (e *ElBinaryOperator) SetRightOperand(rightOperand IElExpression) error {
	e.rightOperand = rightOperand
	return nil
}

// CONSTRUCTOR
func NewElBinaryOperator() *ElBinaryOperator {
	elbinaryoperator := new(ElBinaryOperator)
	return elbinaryoperator
}

/* ======================= ElTuple ==================== */
// Defines an array of optionally named items each of any type.
type ElTuple struct {
	ElExpression
	// Attributes
	/**
	items in the tuple, potentially with names. Typical use is to represent an
	argument list to routine call.
	*/
	items []IElTupleItem `yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	_type IBmmTupleType `yaml:"type" json:"type" xml:"type"`
}

func (e *ElTuple) Items() []IElTupleItem {
	return e.items
}

func (e *ElTuple) SetItems(items []IElTupleItem) error {
	e.items = items
	return nil
}

func (e *ElTuple) Type() IBmmTupleType {
	return e._type
}

func (e *ElTuple) SetType(_type IBmmTupleType) error {
	e._type = _type
	return nil
}

// CONSTRUCTOR
func NewElTuple() *ElTuple {
	eltuple := new(ElTuple)
	eltuple.items = make([]IElTupleItem, 0)
	return eltuple
}

func (e *ElTuple) EvalType() IBmmType {
	return nil
}

/* ======================= ElTupleItem ==================== */
// A single tuple item, with an optional name.
type ElTupleItem struct {
	// Attributes
	/**
	Reference to value entity. If Void, this indicates that the item in this
	position is Void, e.g. within a routine call parameter list.
	*/
	item IElExpression `yaml:"item" json:"item" xml:"item"`
	// Optional name of tuple item.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (e *ElTupleItem) Item() IElExpression {
	return e.item
}

func (e *ElTupleItem) SetItem(item IElExpression) error {
	e.item = item
	return nil
}

func (e *ElTupleItem) Name() string {
	return e.name
}

func (e *ElTupleItem) SetName(name string) error {
	e.name = name
	return nil
}

// CONSTRUCTOR
func NewElTupleItem() *ElTupleItem {
	eltupleitem := new(ElTupleItem)
	return eltupleitem
}

/* ======================= ElConstrained ==================== */
/**
Abstract parent for second-order constrained forms of first-order expression
meta-types.
*/
type ElConstrained struct {
	ElExpression
	// Attributes
	// The base expression of this constrained form.
	baseExpression IElExpression `yaml:"baseexpression" json:"baseexpression" xml:"baseexpression"`
}

func (e *ElConstrained) BaseExpression() IElExpression {
	return e.baseExpression
}

func (e *ElConstrained) SetBaseExpression(baseExpression IElExpression) error {
	if baseExpression == nil {
		return errors.New("BaseExpression should not be set to null")
	}
	e.baseExpression = baseExpression
	return nil
}

/* ======================= ElBooleanExpression ==================== */
// Boolean-returning expression.
type ElBooleanExpression struct {
	ElConstrained
}

// CONSTRUCTOR
func NewElBooleanExpression() *ElBooleanExpression {
	elbooleanexpression := new(ElBooleanExpression)
	return elbooleanexpression
}
