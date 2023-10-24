package vocabulary

import (
	"vocabulary"
)

/**
	Core properties of BMM_MODEL , may be used in a serial representation as well,
	such as P_BMM_SCHEMA .
*/

// Interface definition
type IBmmModelMetadata interface {
}

// Struct definition
type BmmModelMetadata struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Publisher of model expressed in the schema.
	RmPublisher	string	`yaml:"rmpublisher" json:"rmpublisher" xml:"rmpublisher"`
	// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
	RmRelease	string	`yaml:"rmrelease" json:"rmrelease" xml:"rmrelease"`
}

//CONSTRUCTOR
func NewBmmModelMetadata() *BmmModelMetadata {
	bmmmodelmetadata := new(BmmModelMetadata)
	// Constants
	return bmmmodelmetadata
}
//BUILDER
type BmmModelMetadataBuilder struct {
	bmmmodelmetadata *BmmModelMetadata
}

func NewBmmModelMetadataBuilder() *BmmModelMetadataBuilder {
	 return &BmmModelMetadataBuilder {
		bmmmodelmetadata : NewBmmModelMetadata(),
	}
}

//BUILDER ATTRIBUTES
// Publisher of model expressed in the schema.
func (i *BmmModelMetadataBuilder) SetRmPublisher ( v string ) *BmmModelMetadataBuilder{
	i.bmmmodelmetadata.RmPublisher = v
	return i
}
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelMetadataBuilder) SetRmRelease ( v string ) *BmmModelMetadataBuilder{
	i.bmmmodelmetadata.RmRelease = v
	return i
}

func (i *BmmModelMetadataBuilder) Build() *BmmModelMetadata {
	 return i.bmmmodelmetadata
}

//FUNCTIONS
