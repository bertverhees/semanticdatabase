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
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
