package vocabulary

/**
Statement type representing an assignment from a value-generating source to a
writable entity, i.e. a variable reference or property. At the meta-model level,
may be understood as an initialisation of an existing meta-model instance.
*/

// Interface definition
type IBmmAssignment interface {
	IBmmSimpleStatement
	IBmmStatement
	IBmmStatementItem
}

// Struct definition
type BmmAssignment struct {
	// embedded for Inheritance
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
	// The target variable on the notional left-hand side of this assignment.
	Target IElValueGenerator `yaml:"target" json:"target" xml:"target"`
	// Source right hand side) of the assignment.
	Source IElExpression `yaml:"source" json:"source" xml:"source"`
}

// CONSTRUCTOR
func NewBmmAssignment() *BmmAssignment {
	bmmassignment := new(BmmAssignment)
	// Constants
	return bmmassignment
}

// BUILDER
type BmmAssignmentBuilder struct {
	bmmassignment *BmmAssignment
}

func NewBmmAssignmentBuilder() *BmmAssignmentBuilder {
	return &BmmAssignmentBuilder{
		bmmassignment: NewBmmAssignment(),
	}
}

// BUILDER ATTRIBUTES
// The target variable on the notional left-hand side of this assignment.
func (i *BmmAssignmentBuilder) SetTarget(v IElValueGenerator) *BmmAssignmentBuilder {
	i.bmmassignment.Target = v
	return i
}

// Source right hand side) of the assignment.
func (i *BmmAssignmentBuilder) SetSource(v IElExpression) *BmmAssignmentBuilder {
	i.bmmassignment.Source = v
	return i
}

func (i *BmmAssignmentBuilder) Build() *BmmAssignment {
	return i.bmmassignment
}

//FUNCTIONS
