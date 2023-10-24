package vocabulary

import (
	"vocabulary"
)

/**
	A logical group of features, with a name and set of properties that applies to
	the group.
*/

// Interface definition
type IBmmFeatureGroup interface {
}

// Struct definition
type BmmFeatureGroup struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Name of this feature group; defaults to 'feature'.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		Set of properties of this group, represented as name/value pairs. These are
		understood to apply logically to all of the features contained within the group.
	*/
	Properties	map[string]string	`yaml:"properties" json:"properties" xml:"properties"`
	// Set of features in this group.
	Features	List < BMM_FEATURE >	`yaml:"features" json:"features" xml:"features"`
	// Optional visibility to apply to all features in this group.
	Visibility	IBmmVisibility	`yaml:"visibility" json:"visibility" xml:"visibility"`
}

//CONSTRUCTOR
func NewBmmFeatureGroup() *BmmFeatureGroup {
	bmmfeaturegroup := new(BmmFeatureGroup)
	// Constants
	return bmmfeaturegroup
}
//BUILDER
type BmmFeatureGroupBuilder struct {
	bmmfeaturegroup *BmmFeatureGroup
}

func NewBmmFeatureGroupBuilder() *BmmFeatureGroupBuilder {
	 return &BmmFeatureGroupBuilder {
		bmmfeaturegroup : NewBmmFeatureGroup(),
	}
}

//BUILDER ATTRIBUTES
// Name of this feature group; defaults to 'feature'.
func (i *BmmFeatureGroupBuilder) SetName ( v string ) *BmmFeatureGroupBuilder{
	i.bmmfeaturegroup.Name = v
	return i
}
/**
	Set of properties of this group, represented as name/value pairs. These are
	understood to apply logically to all of the features contained within the group.
*/
func (i *BmmFeatureGroupBuilder) SetProperties ( v map[string]string ) *BmmFeatureGroupBuilder{
	i.bmmfeaturegroup.Properties = v
	return i
}
// Set of features in this group.
func (i *BmmFeatureGroupBuilder) SetFeatures ( v List < BMM_FEATURE > ) *BmmFeatureGroupBuilder{
	i.bmmfeaturegroup.Features = v
	return i
}
// Optional visibility to apply to all features in this group.
func (i *BmmFeatureGroupBuilder) SetVisibility ( v IBmmVisibility ) *BmmFeatureGroupBuilder{
	i.bmmfeaturegroup.Visibility = v
	return i
}

func (i *BmmFeatureGroupBuilder) Build() *BmmFeatureGroup {
	 return i.bmmfeaturegroup
}

//FUNCTIONS
