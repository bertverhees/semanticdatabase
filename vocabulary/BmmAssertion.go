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
	errors       []error
}

func NewBmmAssertionBuilder() *BmmAssertionBuilder {
	return &BmmAssertionBuilder{
		bmmassertion: NewBmmAssertion(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Boolean-valued expression of the assertion.
func (i *BmmAssertionBuilder) SetExpression(v IElBooleanExpression) *BmmAssertionBuilder {
	i.AddError(i.bmmassertion.SetExpression(v))
	return i
}

/*
*
Optional tag, typically used to designate design intention of the assertion,
e.g. "Inv_all_members_valid" .
*/
func (i *BmmAssertionBuilder) SetTag(v string) *BmmAssertionBuilder {
	i.AddError(i.bmmassertion.SetTag(v))
	return i
}
func (i *BmmAssertionBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmAssertionBuilder) Build() *BmmAssertion {
	return i.bmmassertion
}

//FUNCTIONS
