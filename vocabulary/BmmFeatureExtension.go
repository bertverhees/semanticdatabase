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
	BmmFeatureExtension := new(BmmFeatureExtension)
	// Constants
	return BmmFeatureExtension
}

// BUILDER
type BmmFeatureExtensionBuilder struct {
	BmmFeatureExtension *BmmFeatureExtension
}

func NewBmmFeatureExtensionBuilder() *BmmFeatureExtensionBuilder {
	return &BmmFeatureExtensionBuilder{
		BmmFeatureExtension: NewBmmFeatureExtension(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmFeatureExtensionBuilder) Build() *BmmFeatureExtension {
	return i.BmmFeatureExtension
}
