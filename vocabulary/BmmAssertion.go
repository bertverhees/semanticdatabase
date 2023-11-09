package vocabulary

/**
A statement that asserts the truth of its expression. If the expression
evaluates to False the execution generates an exception (depending on run-time
settings). May be rendered in syntax as assert condition or similar.
*/

// Interface definition
type IBmmAssertion interface {
	IBmmSimpleStatement
}

// Struct definition
type BmmAssertion struct {
	BmmSimpleStatement
	// Attributes
	// Boolean-valued expression of the assertion.
	Expression IElBooleanExpression `yaml:"expression" json:"expression" xml:"expression"`
	/**
	Optional tag, typically used to designate design intention of the assertion,
	e.g. "Inv_all_members_valid" .
	*/
	Tag string `yaml:"tag" json:"tag" xml:"tag"`
}

// CONSTRUCTOR
func NewBmmAssertion() *BmmAssertion {
	bmmassertion := new(BmmAssertion)
	// Constants
	return bmmassertion
}

// BUILDER
type BmmAssertionBuilder struct {
	bmmassertion *BmmAssertion
}

func NewBmmAssertionBuilder() *BmmAssertionBuilder {
	return &BmmAssertionBuilder{
		bmmassertion: NewBmmAssertion(),
	}
}

// BUILDER ATTRIBUTES
// Boolean-valued expression of the assertion.
func (i *BmmAssertionBuilder) SetExpression(v IElBooleanExpression) *BmmAssertionBuilder {
	i.bmmassertion.Expression = v
	return i
}

/*
*
Optional tag, typically used to designate design intention of the assertion,
e.g. "Inv_all_members_valid" .
*/
func (i *BmmAssertionBuilder) SetTag(v string) *BmmAssertionBuilder {
	i.bmmassertion.Tag = v
	return i
}

func (i *BmmAssertionBuilder) Build() *BmmAssertion {
	return i.bmmassertion
}

//FUNCTIONS
