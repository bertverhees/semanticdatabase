package vocabulary

// Definition of a symbolic operator associated with a function.

// Interface definition
type IBmmOperator interface {
}

// Struct definition
type BmmOperator struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Position of operator in syntactic representation.
	Position BmmOperatorPosition `yaml:"position" json:"position" xml:"position"`
	/**
	Set of String symbols that may be used to represent this operator in a textual
	representation of a BMM model.
	*/
	Symbols []string `yaml:"symbols" json:"symbols" xml:"symbols"`
	// Formal name of the operator, e.g. 'minus' etc.
	Name string `yaml:"name" json:"name" xml:"name"`
}

// CONSTRUCTOR
func NewBmmOperator() *BmmOperator {
	bmmoperator := new(BmmOperator)
	bmmoperator.Symbols = make([]string, 0)
	return bmmoperator
}

// BUILDER
type BmmOperatorBuilder struct {
	bmmoperator *BmmOperator
}

func NewBmmOperatorBuilder() *BmmOperatorBuilder {
	return &BmmOperatorBuilder{
		bmmoperator: NewBmmOperator(),
	}
}

// BUILDER ATTRIBUTES
// Position of operator in syntactic representation.
func (i *BmmOperatorBuilder) SetPosition(v BmmOperatorPosition) *BmmOperatorBuilder {
	i.bmmoperator.Position = v
	return i
}

/*
*
Set of String symbols that may be used to represent this operator in a textual
representation of a BMM model.
*/
func (i *BmmOperatorBuilder) SetSymbols(v []string) *BmmOperatorBuilder {
	i.bmmoperator.Symbols = v
	return i
}

// Formal name of the operator, e.g. 'minus' etc.
func (i *BmmOperatorBuilder) SetName(v string) *BmmOperatorBuilder {
	i.bmmoperator.Name = v
	return i
}

func (i *BmmOperatorBuilder) Build() *BmmOperator {
	return i.bmmoperator
}

//FUNCTIONS
