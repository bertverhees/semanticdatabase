package vocabulary

/**
A predicate taking one external variable reference argument, that returns true
if the reference is resolvable, i.e. the external value is obtainable. Note
probably to be removed.
*/

// Interface definition
type IElDefined interface {
	IElPredicate
	//EL_DEFINED
}

// Struct definition
type ElDefined struct {
	ElPredicate
	// Attributes
}

// CONSTRUCTOR
func NewElDefined() *ElDefined {
	eldefined := new(ElDefined)
	// Constants
	return eldefined
}

// BUILDER
type ElDefinedBuilder struct {
	eldefined *ElDefined
	errors    []error
}

func NewElDefinedBuilder() *ElDefinedBuilder {
	return &ElDefinedBuilder{
		eldefined: NewElDefined(),
		errors:    make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: ElPredicate
// The target instance of this predicate.
func (i *ElDefinedBuilder) SetOperand(v IElValueGenerator) *ElDefinedBuilder {
	i.AddError(i.eldefined.SetOperand(v))
	return i
}

func (i *ElDefinedBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElDefinedBuilder) Build() *ElDefined {
	return i.eldefined
}

// FUNCTIONS
