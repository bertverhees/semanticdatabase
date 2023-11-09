package vocabulary

/**
Simple statement, i.e. statement with one logical element - a single expression,
procedure call etc.
*/

// Interface definition
type IBmmSimpleStatement interface {
	IBmmStatement
}

// Struct definition
type BmmSimpleStatement struct {
	BmmStatement
}

// CONSTRUCTOR
func NewBmmSimpleStatement() *BmmSimpleStatement {
	bmmsimplestatement := new(BmmSimpleStatement)
	// Constants
	return bmmsimplestatement
}

// BUILDER
type BmmSimpleStatementBuilder struct {
	bmmsimplestatement *BmmSimpleStatement
}

func NewBmmSimpleStatementBuilder() *BmmSimpleStatementBuilder {
	return &BmmSimpleStatementBuilder{
		bmmsimplestatement: NewBmmSimpleStatement(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmSimpleStatementBuilder) Build() *BmmSimpleStatement {
	return i.bmmsimplestatement
}

//FUNCTIONS
