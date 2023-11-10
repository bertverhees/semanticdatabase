package vocabulary

/**
Statement type representing an assignment from a value-generating source to a
writable entity, i.e. a variable reference or property. At the meta-model level,
may be understood as an initialisation of an existing meta-model instance.
*/

// Interface definition
type IBmmAssignment interface {
	IBmmSimpleStatement
	Target() (IElValueGenerator, error)
	SetTarget(target IElValueGenerator) error
	Source() (IElExpression, error)
	SetSource(source IElExpression) error
}

// Struct definition
type BmmAssignment struct {
	BmmSimpleStatement
	// Attributes
	// The target variable on the notional left-hand side of this assignment.
	target IElValueGenerator `yaml:"target" json:"target" xml:"target"`
	// source right hand side) of the assignment.
	source IElExpression `yaml:"source" json:"source" xml:"source"`
}

func (b *BmmAssignment) Target() (IElValueGenerator, error) {
	return b.target, nil
}

func (b *BmmAssignment) SetTarget(target IElValueGenerator) error {
	b.target = target
	return nil
}

func (b *BmmAssignment) Source() (IElExpression, error) {
	return b.source, nil
}

func (b *BmmAssignment) SetSource(source IElExpression) error {
	b.source = source
	return nil
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
func (i *BmmAssignmentBuilder) SetTarget(v IElValueGenerator) (*BmmAssignmentBuilder, error) {
	e := i.bmmassignment.SetTarget(v)
	return i, e
}

// source right hand side) of the assignment.
func (i *BmmAssignmentBuilder) SetSource(v IElExpression) (*BmmAssignmentBuilder, error) {
	e := i.bmmassignment.SetSource(v)
	return i, e
}

func (i *BmmAssignmentBuilder) Build() *BmmAssignment {
	return i.bmmassignment
}

//FUNCTIONS
