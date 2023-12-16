package vocabulary

import "errors"

/* ================ BmmStatementItem =========================== */
// Abstract parent of statement types representing a locally defined routine body.
type BmmStatementItem struct {
	// Attributes
}

// CONSTRUCTOR
func NewBmmStatementItem() *BmmStatementItem {
	bmmstatementitem := new(BmmStatementItem)
	// Constants
	return bmmstatementitem
}

/* ================ BmmStatementBlock =========================== */
/**
A statement 'block' corresponding to the programming language concept of the
same name. May be used to establish scope in specific languages.
*/

type BmmStatementBlock struct {
	BmmStatementItem
	// Child blocks of the current block.
	items []IBmmStatementItem `yaml:"items" json:"items" xml:"items"`
}

func (b *BmmStatementBlock) Items() []IBmmStatementItem {
	return b.items
}

func (b *BmmStatementBlock) SetItems(items []IBmmStatementItem) error {
	b.items = items
	return nil
}

func NewBmmStatementBlock() *BmmStatementBlock {
	bmmStatementBlock := new(BmmStatementBlock)
	bmmStatementBlock.items = make([]IBmmStatementItem, 0)
	return bmmStatementBlock
}

/* ================ BmmStatement =========================== */
/**
Abstract parent of 'statement' types that may be defined to implement BMM
Routines.
*/
type BmmStatement struct {
	BmmStatementItem
}

/* ================ BmmSimpleStatement =========================== */
/**
Simple statement, i.e. statement with one logical element - a single expression,
procedure call etc.
*/
type BmmSimpleStatement struct {
	BmmStatement
}

/* ================ BmmDeclaration =========================== */
// Declaration of a writable variable, associating a name with a type.
type BmmDeclaration struct {
	BmmSimpleStatement
	// Attributes
	name   string              `yaml:"name" json:"name" xml:"name"`
	result IElWritableVariable `yaml:"result" json:"result" xml:"result"`
	// The declared type of the variable.
	_type IBmmType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmDeclaration) Name() string {
	return b.name
}

func (b *BmmDeclaration) SetName(name string) error {
	if name == "" {
		return errors.New("name may not be set to empty in BmmDeclaration")
	}
	b.name = name
	return nil
}

func (b *BmmDeclaration) Result() IElWritableVariable {
	return b.result
}

func (b *BmmDeclaration) SetResult(result IElWritableVariable) error {
	if result == nil {
		return errors.New("result may not be set to null in BmmDeclaration")
	}
	b.result = result
	return nil
}

func (b *BmmDeclaration) Type() IBmmType {
	return b._type
}

func (b *BmmDeclaration) SetType(_type IBmmType) error {
	if _type == nil {
		return errors.New("type may not be set to null in BmmDeclaration")
	}
	b._type = _type
	return nil
}

// CONSTRUCTOR
func NewBmmDeclaration() *BmmDeclaration {
	bmmdeclaration := new(BmmDeclaration)
	// Constants
	return bmmdeclaration
}

/* ================ BmmAssignment =========================== */
/**
Statement type representing an assignment from a value-generating source to a
writable entity, i.e. a variable reference or property. At the meta-model level,
may be understood as an initialisation of an existing meta-model instance.
*/
type BmmAssignment struct {
	BmmSimpleStatement
	// Attributes
	// The target variable on the notional left-hand side of this assignment.
	target IElValueGenerator `yaml:"target" json:"target" xml:"target"`
	// source right hand side) of the assignment.
	source IElExpression `yaml:"source" json:"source" xml:"source"`
}

func (b *BmmAssignment) Target() IElValueGenerator {
	return b.target
}

func (b *BmmAssignment) SetTarget(target IElValueGenerator) error {
	if target == nil {
		return errors.New("type may not be set to null in BmmAssignment")
	}
	b.target = target
	return nil
}

func (b *BmmAssignment) Source() IElExpression {
	return b.source
}

func (b *BmmAssignment) SetSource(source IElExpression) error {
	if source == nil {
		return errors.New("type may not be set to null in BmmAssignment")
	}
	b.source = source
	return nil
}

// CONSTRUCTOR
func NewBmmAssignment() *BmmAssignment {
	bmmassignment := new(BmmAssignment)
	// Constants
	return bmmassignment
}

/* ================ BmmProcedureCall =========================== */
/**
A call made on a closed procedure agent. The method in BMM via which external
actions are achieved from within a program.
*/
type BmmProcedureCall struct {
	ElAgentCall
	BmmSimpleStatement
	// Attributes
	// The procedure agent being called.
	agent IElProcedureAgent `yaml:"agent" json:"agent" xml:"agent"`
}

func (e *BmmProcedureCall) SetAgent(agent IElAgent) error {
	if agent == nil {
		return errors.New("Agent in BmmProcedureCall may not be set to null")
	}
	s, ok := agent.(IElProcedureAgent)
	if !ok {
		return errors.New("_type-assertion in BmmProcedureCall->SetAgent failed")
	} else {
		e.agent = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmProcedureCall() *BmmProcedureCall {
	bmmprocedurecall := new(BmmProcedureCall)
	// Constants
	return bmmprocedurecall
}

/* ================ BmmAssertion =========================== */
type BmmAssertion struct {
	BmmSimpleStatement
	// Attributes
	// Boolean-valued expression of the assertion.
	expression IElBooleanExpression `yaml:"expression" json:"expression" xml:"expression"`
	/**
	Optional tag, typically used to designate design intention of the assertion,
	e.g. "Inv_all_members_valid" .
	*/
	tag string `yaml:"tag" json:"tag" xml:"tag"`
}

func (b *BmmAssertion) Expression() IElBooleanExpression {
	return b.expression
}

func (b *BmmAssertion) SetExpression(expression IElBooleanExpression) error {
	if expression == nil {
		return errors.New("Expression in BmmAssertion may not be set to null")
	}
	b.expression = expression
	return nil
}

func (b *BmmAssertion) Tag() string {
	return b.tag
}

func (b *BmmAssertion) SetTag(tag string) error {
	b.tag = tag
	return nil
}

// CONSTRUCTOR
func NewBmmAssertion() *BmmAssertion {
	bmmassertion := new(BmmAssertion)
	return bmmassertion
}

/* ================ BmmActionTable =========================== */
// Multi-branch conditional statement structure
type BmmActionTable struct {
	BmmStatement
	// Attributes
	/**
	A specialised decision table whose outputs can only be procedure agents. In
	execution, the matched agent will be invoked.
	*/
	decisionTable IBmmActionDecisionTable `yaml:"decisiontable" json:"decisiontable" xml:"decisiontable"`
}

func (b *BmmActionTable) DecisionTable() IBmmActionDecisionTable {
	return b.decisionTable
}

func (b *BmmActionTable) SetDecisionTable(decisionTable IBmmActionDecisionTable) error {
	if decisionTable == nil {
		return errors.New("decisionTable in BmmActionTable may not be set to null")
	}
	b.decisionTable = decisionTable
	return nil
}

// CONSTRUCTOR
func NewBmmActionTable() *BmmActionTable {
	bmmactiontable := new(BmmActionTable)
	// Constants
	return bmmactiontable
}

/* ================ BmmActionDecisionTable =========================== */
/*
Specialised form of Decision Table that allows only procedure call agents (lambdas) as the result of branches.
*/
type BmmActionDecisionTable struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmActionDecisionTable() *BmmActionDecisionTable {
	bmmactiondecisiontable := new(BmmActionDecisionTable)
	// Constants
	return bmmactiondecisiontable
}
