package vocabulary

// Interface definition
type IBmmFeatureExtension interface {
}

// Struct definition
type BmmFeatureExtension struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmFeatureExtension() *BmmFeatureExtension {
	bmmFeatureExtension := new(BmmFeatureExtension)
	// Constants
	return bmmFeatureExtension
}

// BUILDER
type BmmFeatureExtensionBuilder struct {
	bmmFeatureExtension *BmmFeatureExtension
	errors              []error
}

func NewBmmFeatureExtensionBuilder() *BmmFeatureExtensionBuilder {
	return &BmmFeatureExtensionBuilder{
		bmmFeatureExtension: NewBmmFeatureExtension(),
		errors:              make([]error, 0),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmFeatureExtensionBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmFeatureExtensionBuilder) Build() *BmmFeatureExtension {
	return i.bmmFeatureExtension
}
