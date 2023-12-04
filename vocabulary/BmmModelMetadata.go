package vocabulary

import "log"

/**
Core properties of BMM_MODEL , may be used in a serial representation as well,
such as P_BMM_SCHEMA .
*/

// Interface definition
type IBmmModelMetadata interface {
}

// Struct definition
type BmmModelMetadata struct {
	// Publisher of model expressed in the schema.
	rmPublisher string `yaml:"rmpublisher" json:"rmpublisher" xml:"rmpublisher"`
	// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
	rmRelease string `yaml:"rmrelease" json:"rmrelease" xml:"rmrelease"`
}

func (b *BmmModelMetadata) RmPublisher() string {
	return b.rmPublisher
}

func (b *BmmModelMetadata) SetRmPublisher(rmPublisher string) error {
	b.rmPublisher = rmPublisher
	return nil
}

func (b *BmmModelMetadata) RmRelease() string {
	return b.rmRelease
}

func (b *BmmModelMetadata) SetRmRelease(rmRelease string) error {
	b.rmRelease = rmRelease
	return nil
}

// CONSTRUCTOR
func NewBmmModelMetadata() *BmmModelMetadata {
	bmmmodelmetadata := new(BmmModelMetadata)
	return bmmmodelmetadata
}

// BUILDER
type BmmModelMetadataBuilder struct {
	bmmmodelmetadata *BmmModelMetadata
	errors           []error
}

func NewBmmModelMetadataBuilder() *BmmModelMetadataBuilder {
	log.Fatal("The class BmmModelMetadata is not yet supported")
	//return &BmmModelMetadataBuilder{
	//	bmmmodelmetadata: NewBmmModelMetadata(),
	//	errors:   make([]error, 0),
	//}
	return nil
}

// BUILDER ATTRIBUTES
// Publisher of model expressed in the schema.
func (i *BmmModelMetadataBuilder) SetRmPublisher(v string) *BmmModelMetadataBuilder {
	i.AddError(i.bmmmodelmetadata.SetRmPublisher(v))
	return i
}

// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelMetadataBuilder) SetRmRelease(v string) *BmmModelMetadataBuilder {
	i.AddError(i.bmmmodelmetadata.SetRmRelease(v))
	return i
}

func (i *BmmModelMetadataBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmModelMetadataBuilder) Build() *BmmModelMetadata {
	return i.bmmmodelmetadata
}

//FUNCTIONS
