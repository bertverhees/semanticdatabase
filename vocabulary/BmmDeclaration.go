package vocabulary

import (
	"vocabulary"
)

// Declaration of a writable variable, associating a name with a type.

// Interface definition
type IBmmDeclaration interface {
	// From: BMM_SIMPLE_STATEMENT
	// From: BMM_STATEMENT
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmDeclaration struct {
	// embedded for Inheritance
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
	Name	string	`yaml:"name" json:"name" xml:"name"`
	Result	IElWritableVariable	`yaml:"result" json:"result" xml:"result"`
	// The declared type of the variable.
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmDeclaration() *BmmDeclaration {
	bmmdeclaration := new(BmmDeclaration)
	// Constants
	// From: BmmSimpleStatement
	// From: BmmStatement
	// From: BmmStatementItem
	return bmmdeclaration
}
//BUILDER
type BmmDeclarationBuilder struct {
	bmmdeclaration *BmmDeclaration
}

func NewBmmDeclarationBuilder() *BmmDeclarationBuilder {
	 return &BmmDeclarationBuilder {
		bmmdeclaration : NewBmmDeclaration(),
	}
}

//BUILDER ATTRIBUTES
func (i *BmmDeclarationBuilder) SetName ( v string ) *BmmDeclarationBuilder{
	i.bmmdeclaration.Name = v
	return i
}
func (i *BmmDeclarationBuilder) SetResult ( v IElWritableVariable ) *BmmDeclarationBuilder{
	i.bmmdeclaration.Result = v
	return i
}
	// The declared type of the variable.
func (i *BmmDeclarationBuilder) SetType ( v IBmmType ) *BmmDeclarationBuilder{
	i.bmmdeclaration.Type = v
	return i
}

func (i *BmmDeclarationBuilder) Build() *BmmDeclaration {
	 return i.bmmdeclaration
}

//FUNCTIONS
