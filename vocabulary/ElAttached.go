package vocabulary

/**
A predicate on any object reference (including function call) that returns True
if the reference is attached, i.e. non-Void.
*/

// Interface definition
type IElAttached interface {
	IElPredicate
	//EL_ATTACHED
}

// Struct definition
type ElAttached struct {
	ElPredicate
	// Attributes
}

// CONSTRUCTOR
func NewElAttached() *ElAttached {
	elattached := new(ElAttached)
	// Constants
	return elattached
}

// BUILDER
type ElAttachedBuilder struct {
	elattached *ElAttached
	errors     []error
}

func NewElAttachedBuilder() *ElAttachedBuilder {
	return &ElAttachedBuilder{
		elattached: NewElAttached(),
		errors:     make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: ElPredicate
// The target instance of this predicate.
func (i *ElAttachedBuilder) SetOperand(v IElValueGenerator) *ElAttachedBuilder {
	i.AddError(i.elattached.SetOperand(v))
	return i
}

func (i *ElAttachedBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElAttachedBuilder) Build() *ElAttached {
	return i.elattached
}

// FUNCTIONS
