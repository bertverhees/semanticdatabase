package vocabulary

// definition of a symbolic operator associated with a function.

// Interface definition
type IBmmOperator interface {
	Position() BmmOperatorPosition
	SetPosition(position BmmOperatorPosition) error
	Symbols() []string
	SetSymbols(symbols []string) error
	Name() string
	SetName(name string) error
}

// Struct definition
type BmmOperator struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// position of operator in syntactic representation.
	position BmmOperatorPosition `yaml:"position" json:"position" xml:"position"`
	/**
	Set of String symbols that may be used to represent this operator in a textual
	representation of a BMM model.
	*/
	symbols []string `yaml:"symbols" json:"symbols" xml:"symbols"`
	// Formal name of the operator, e.g. 'minus' etc.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (b *BmmOperator) Position() BmmOperatorPosition {
	return b.position
}

func (b *BmmOperator) SetPosition(position BmmOperatorPosition) error {
	b.position = position
	return nil
}

func (b *BmmOperator) Symbols() []string {
	return b.symbols
}

func (b *BmmOperator) SetSymbols(symbols []string) error {
	b.symbols = symbols
	return nil
}

func (b *BmmOperator) Name() string {
	return b.name
}

func (b *BmmOperator) SetName(name string) error {
	b.name = name
	return nil
}

// CONSTRUCTOR
func NewBmmOperator() *BmmOperator {
	bmmoperator := new(BmmOperator)
	bmmoperator.symbols = make([]string, 0)
	return bmmoperator
}

// BUILDER
type BmmOperatorBuilder struct {
	bmmoperator *BmmOperator
	errors      []error
}

func NewBmmOperatorBuilder() *BmmOperatorBuilder {
	return &BmmOperatorBuilder{
		bmmoperator: NewBmmOperator(),
		errors:      make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// position of operator in syntactic representation.
func (i *BmmOperatorBuilder) SetPosition(v BmmOperatorPosition) *BmmOperatorBuilder {
	i.AddError(i.bmmoperator.SetPosition(v))
	return i
}

/*
*
Set of String symbols that may be used to represent this operator in a textual
representation of a BMM model.
*/
func (i *BmmOperatorBuilder) SetSymbols(v []string) *BmmOperatorBuilder {
	i.AddError(i.bmmoperator.SetSymbols(v))
	return i
}

// Formal name of the operator, e.g. 'minus' etc.
func (i *BmmOperatorBuilder) SetName(v string) *BmmOperatorBuilder {
	i.AddError(i.bmmoperator.SetName(v))
	return i
}

func (i *BmmOperatorBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmOperatorBuilder) Build() *BmmOperator {
	return i.bmmoperator
}

//FUNCTIONS
