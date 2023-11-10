package vocabulary

/**
A statement that asserts the truth of its expression. If the expression
evaluates to False the execution generates an exception (depending on run-time
settings). May be rendered in syntax as assert condition or similar.
*/

// Interface definition
type IBmmAssertion interface {
	IBmmSimpleStatement
	Expression() (IElBooleanExpression, error)
	SetExpression(expression IElBooleanExpression) error
	Tag() (string, error)
	SetTag(tag string) error
}

// Struct definition
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

func (b *BmmAssertion) Expression() (IElBooleanExpression, error) {
	return b.expression, nil
}

func (b *BmmAssertion) SetExpression(expression IElBooleanExpression) error {
	b.expression = expression
	return nil
}

func (b *BmmAssertion) Tag() (string, error) {
	return b.tag, nil
}

func (b *BmmAssertion) SetTag(tag string) error {
	b.tag = tag
	return nil
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
func (i *BmmAssertionBuilder) SetExpression(v IElBooleanExpression) (*BmmAssertionBuilder,error) {
	e := i.bmmassertion.SetExpression(v)
	return i,e
}

/*
*
Optional tag, typically used to designate design intention of the assertion,
e.g. "Inv_all_members_valid" .
*/
func (i *BmmAssertionBuilder) SetTag(v string) (*BmmAssertionBuilder,error) {
	e := i.bmmassertion.SetTag(v)
	return i,e
}

func (i *BmmAssertionBuilder) Build() *BmmAssertion {
	return i.bmmassertion
}

//FUNCTIONS
