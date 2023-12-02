package v2

// Persistent form of BMM_MODEL_ELEMENT .

// Interface definition
type IPBmmModelElement interface {
	Documentation() string
	SetDocumentation(documentation string) error
}

// Struct definition
type PBmmModelElement struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Optional documentation of this element.
	documentation string `yaml:"documentation" json:"documentation" xml:"documentation"`
}

func (P *PBmmModelElement) Documentation() string {
	return P.documentation
}

func (P *PBmmModelElement) SetDocumentation(documentation string) error {
	P.documentation = documentation
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
