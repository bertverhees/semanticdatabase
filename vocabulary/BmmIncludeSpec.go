package vocabulary

import (
	"vocabulary"
)

// Schema inclusion structure.

// Interface definition
type IBmmIncludeSpec interface {
}

// Struct definition
type BmmIncludeSpec struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
	Id	string	`yaml:"id" json:"id" xml:"id"`
}

//CONSTRUCTOR
func NewBmmIncludeSpec() *BmmIncludeSpec {
	bmmincludespec := new(BmmIncludeSpec)
	// Constants
	return bmmincludespec
}
//BUILDER
type BmmIncludeSpecBuilder struct {
	bmmincludespec *BmmIncludeSpec
}

func NewBmmIncludeSpecBuilder() *BmmIncludeSpecBuilder {
	 return &BmmIncludeSpecBuilder {
		bmmincludespec : NewBmmIncludeSpec(),
	}
}

//BUILDER ATTRIBUTES
// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
func (i *BmmIncludeSpecBuilder) SetId ( v string ) *BmmIncludeSpecBuilder{
	i.bmmincludespec.Id = v
	return i
}

func (i *BmmIncludeSpecBuilder) Build() *BmmIncludeSpec {
	 return i.bmmincludespec
}

//FUNCTIONS
