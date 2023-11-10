package vocabulary

import "errors"

/**
Meta-type defining a generalised module concept. Descendants define actual
structure and contents.
*/

// Interface definition
type IBmmModule interface {
	IBmmModelElement
	FeatureGroups() ([]IBmmFeatureGroup, error)
	SetFeatureGroups(featureGroups []IBmmFeatureGroup) error
	Features() ([]IBmmFormalElement, error)
	SetFeatures(features []IBmmFormalElement) error
}

// Struct definition
type BmmModule struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// List of feature groups in this class.
	featureGroups []IBmmFeatureGroup `yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	features []IBmmFormalElement `yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	scope IBmmModel `yaml:"scope" json:"scope" xml:"scope"` //redefined
}

func (b *BmmModule) FeatureGroups() ([]IBmmFeatureGroup, error) {
	return b.featureGroups, nil
}

func (b *BmmModule) SetFeatureGroups(featureGroups []IBmmFeatureGroup) error {
	b.featureGroups = featureGroups
	return nil
}

func (b *BmmModule) Features() ([]IBmmFormalElement, error) {
	return b.features, nil
}

func (b *BmmModule) SetFeatures(features []IBmmFormalElement) error {
	b.features = features
	return nil
}

func (b *BmmModule) Scope() (IBmmModelElement, error) {
	return b.scope, nil
}

func (b *BmmModule) SetScope(scope IBmmModelElement) error {
	s, ok := scope.(IBmmModel)
	if !ok {
		return errors.New("_type-assertion in BmmModule->SetScope went wrong")
	} else {
		b.scope = s
		return nil
	}
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
