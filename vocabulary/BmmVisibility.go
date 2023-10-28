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

//CONSTRUCTOR
func NewBmmVisibility() *BmmVisibility {
	bmmvisibility := new(BmmVisibility)
	// Constants
	return bmmvisibility
}
//BUILDER
type BmmVisibilityBuilder struct {
	bmmvisibility *BmmVisibility
}

func NewBmmVisibilityBuilder() *BmmVisibilityBuilder {
	 return &BmmVisibilityBuilder {
		bmmvisibility : NewBmmVisibility(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmVisibilityBuilder) Build() *BmmVisibility {
	 return i.bmmvisibility
}

//FUNCTIONS
