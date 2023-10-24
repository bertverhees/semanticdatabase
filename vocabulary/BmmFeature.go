package vocabulary

import (
	"vocabulary"
)

// A module-scoped formal element.

type IBmmFeature interface {
}

type BmmFeature struct {
	BmmFormalElement
	BmmModelElement
	/**
		True if this feature was synthesised due to generic substitution in an inherited
		type, or further constraining of a formal generic parameter.
	*/
	IsSynthesisedGeneric	bool	`yaml:"issynthesisedgeneric" json:"issynthesisedgeneric" xml:"issynthesisedgeneric"`
	// Extensions to feature-level meta-types.
	FeatureExtensions	List < BMM_FEATURE_EXTENSION >	`yaml:"featureextensions" json:"featureextensions" xml:"featureextensions"`
	// Group containing this feature.
	Group	IBmmFeatureGroup	`yaml:"group" json:"group" xml:"group"`
	// Model element within which an element is declared.
	Scope	IBmmClass	`yaml:"scope" json:"scope" xml:"scope"`
}

