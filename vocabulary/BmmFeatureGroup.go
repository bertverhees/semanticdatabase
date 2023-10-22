package vocabulary

import (
	"vocabulary"
)

/**
	A logical group of features, with a name and set of properties that applies to
	the group.
*/

type IBmmFeatureGroup interface {
}

type BmmFeatureGroup struct {
	// Name of this feature group; defaults to 'feature'.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		Set of properties of this group, represented as name/value pairs. These are
		understood to apply logically to all of the features contained within the group.
	*/
	Properties	Hash < String , String >	`yaml:"properties" json:"properties" xml:"properties"`
	// Set of features in this group.
	Features	List < BMM_FEATURE >	`yaml:"features" json:"features" xml:"features"`
	// Optional visibility to apply to all features in this group.
	Visibility	IBmmVisibility	`yaml:"visibility" json:"visibility" xml:"visibility"`
}

