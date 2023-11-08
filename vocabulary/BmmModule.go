package vocabulary

/**
Meta-type defining a generalised module concept. Descendants define actual
structure and contents.
*/

// Interface definition
type IBmmModule interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement
}

// Struct definition
type BmmModule struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// List of feature groups in this class.
	FeatureGroups []IBmmFeatureGroup `yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	Features []IBmmFormalElement `yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	Scope IBmmModel `yaml:"scope" json:"scope" xml:"scope"` //redefined
}

// CONSTRUCTOR
//Abstract, no constructor, no builder

//FUNCTIONS
// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmModule) IsRootScope() bool {
	return false
}
