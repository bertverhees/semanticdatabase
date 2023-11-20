package vocabulary

/**
Abstract parent of visibility representation. TODO: define schemes; probably
need to support C++/Java scheme as well as better type-based schemes.
*/

// Interface definition
type IBmmVisibility interface {
}

// Struct definition
type BmmVisibility struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmVisibility() *BmmVisibility {
	bmmvisibility := new(BmmVisibility)
	// Constants
	return bmmvisibility
}

// BUILDER
type BmmVisibilityBuilder struct {
	bmmvisibility *BmmVisibility
	errors        []error
}

func NewBmmVisibilityBuilder() *BmmVisibilityBuilder {
	return &BmmVisibilityBuilder{
		bmmvisibility: NewBmmVisibility(),
		errors:        make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
func (i *BmmVisibilityBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmVisibilityBuilder) Build() *BmmVisibility {
	return i.bmmvisibility
}

//FUNCTIONS
