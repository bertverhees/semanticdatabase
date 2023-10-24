package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type defining a generalised module concept. Descendants define actual
	structure and contents.
*/

// Interface definition
type IBmmModule interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmModule struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// List of feature groups in this class.
	FeatureGroups	List < BMM_FEATURE_GROUP >	`yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	Features	List < BMM_FORMAL_ELEMENT >	`yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	Scope	IBmmModel	`yaml:"scope" json:"scope" xml:"scope"`
}

//CONSTRUCTOR
func NewBmmModule() *BmmModule {
	bmmmodule := new(BmmModule)
	// Constants
	// From: BmmModelElement
	return bmmmodule
}
//BUILDER
type BmmModuleBuilder struct {
	bmmmodule *BmmModule
}

func NewBmmModuleBuilder() *BmmModuleBuilder {
	 return &BmmModuleBuilder {
		bmmmodule : NewBmmModule(),
	}
}

//BUILDER ATTRIBUTES
	// List of feature groups in this class.
func (i *BmmModuleBuilder) SetFeatureGroups ( v List < BMM_FEATURE_GROUP > ) *BmmModuleBuilder{
	i.bmmmodule.FeatureGroups = v
	return i
}
	// Features of this module.
func (i *BmmModuleBuilder) SetFeatures ( v List < BMM_FORMAL_ELEMENT > ) *BmmModuleBuilder{
	i.bmmmodule.Features = v
	return i
}
	// Model within which module is defined.
func (i *BmmModuleBuilder) SetScope ( v IBmmModel ) *BmmModuleBuilder{
	i.bmmmodule.Scope = v
	return i
}

func (i *BmmModuleBuilder) Build() *BmmModule {
	 return i.bmmmodule
}

//FUNCTIONS
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmModule) IsRootScope (  )  bool {
	return nil
}
