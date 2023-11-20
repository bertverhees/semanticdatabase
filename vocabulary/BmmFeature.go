package vocabulary

import "errors"

// A module-scoped formal element.

// Interface definition
type IBmmFeature interface {
	IBmmFormalElement
	//BMM_FEATURE
	IsSynthesisedGeneric() bool
	SetIsSynthesisedGeneric(isSynthesisedGeneric bool) error
	FeatureExtensions() []IBmmFeatureExtension
	SetFeatureExtensions(featureExtensions []IBmmFeatureExtension) error
	Group() IBmmFeatureGroup
	SetGroup(group IBmmFeatureGroup) error
}

// Struct definition
type BmmFeature struct {
	BmmFormalElement
	// Constants
	// Attributes
	/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
	*/
	isSynthesisedGeneric bool `yaml:"issynthesisedgeneric" json:"issynthesisedgeneric" xml:"issynthesisedgeneric"`
	// extensions to feature-level meta-types.
	featureExtensions []IBmmFeatureExtension `yaml:"featureextensions" json:"featureextensions" xml:"featureextensions"`
	// group containing this feature.
	group IBmmFeatureGroup `yaml:"group" json:"group" xml:"group"`
	// Model element within which an element is declared.
	scope IBmmClass `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmFeature) IsSynthesisedGeneric() bool {
	return b.isSynthesisedGeneric
}

func (b *BmmFeature) SetIsSynthesisedGeneric(isSynthesisedGeneric bool) error {
	b.isSynthesisedGeneric = isSynthesisedGeneric
	return nil
}

func (b *BmmFeature) FeatureExtensions() []IBmmFeatureExtension {
	return b.featureExtensions
}

func (b *BmmFeature) SetFeatureExtensions(featureExtensions []IBmmFeatureExtension) error {
	b.featureExtensions = featureExtensions
	return nil
}

func (b *BmmFeature) Group() IBmmFeatureGroup {
	return b.group
}

func (b *BmmFeature) SetGroup(group IBmmFeatureGroup) error {
	b.group = group
	return nil
}

func (b *BmmFeature) SetScope(v IBmmModelElement) error {
	s, ok := v.(IBmmClass)
	if !ok {
		return errors.New("_type-assertion to IBmmClass in BmmFeature->SetScope went wrong")
	} else {
		b.scope = s
		return nil
	}
}

// CONSTRUCTOR
// abstract, no constructor or builder
//FUNCTIONS
