package vocabulary

/**
Abstract parent for second-order constrained forms of first-order expression
meta-types.
*/

// Interface definition
type IElConstrained interface {
	IElExpression
	BaseExpression() IElExpression
	SetBaseExpression(baseExpression IElExpression) error
}

// Struct definition
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
	e.baseExpression = baseExpression
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
