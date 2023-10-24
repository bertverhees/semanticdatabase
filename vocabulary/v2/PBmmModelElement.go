package v2

import (
	"vocabulary"
)

// Persistent form of BMM_MODEL_ELEMENT .

// Interface definition
type IPBmmModelElement interface {
}

// Struct definition
type PBmmModelElement struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Optional documentation of this element.
	Documentation	string	`yaml:"documentation" json:"documentation" xml:"documentation"`
}

//CONSTRUCTOR
func NewPBmmModelElement() *PBmmModelElement {
	pbmmmodelelement := new(PBmmModelElement)
	// Constants
	return pbmmmodelelement
}
//BUILDER
type PBmmModelElementBuilder struct {
	pbmmmodelelement *PBmmModelElement
}

func NewPBmmModelElementBuilder() *PBmmModelElementBuilder {
	 return &PBmmModelElementBuilder {
		pbmmmodelelement : NewPBmmModelElement(),
	}
}

//BUILDER ATTRIBUTES
	// Optional documentation of this element.
func (i *PBmmModelElementBuilder) SetDocumentation ( v string ) *PBmmModelElementBuilder{
	i.pbmmmodelelement.Documentation = v
	return i
}

func (i *PBmmModelElementBuilder) Build() *PBmmModelElement {
	 return i.pbmmmodelelement
}

//FUNCTIONS
