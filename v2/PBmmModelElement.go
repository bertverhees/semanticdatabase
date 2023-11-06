package v2

// Persistent form of BMM_MODEL_ELEMENT .

// Interface definition
type IPBmmModelElement interface {
}

// Struct definition
type PBmmModelElement struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Optional documentation of this element.
	Documentation string `yaml:"documentation" json:"documentation" xml:"documentation"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder
