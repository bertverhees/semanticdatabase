package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type defining a generalised module concept. Descendants define actual
	structure and contents.
*/

type IBmmModule interface {
}

type BmmModule struct {
	// List of feature groups in this class.
	FeatureGroups	List < BMM_FEATURE_GROUP >	`yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	Features	List < BMM_FORMAL_ELEMENT >	`yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	Scope	IBmmModel	`yaml:"scope" json:"scope" xml:"scope"`
}

