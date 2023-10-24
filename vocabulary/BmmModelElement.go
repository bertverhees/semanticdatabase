package vocabulary

import (
	"vocabulary"
)

/**
	Abstract meta-type of BMM declared model elements. A declaration is a an element
	of a model within a context, which defines the scope of the element. Thus, a
	class definition and its property and routine definitions are model elements,
	but Types are not, since they are derived from model elements.
*/

// Interface definition
type IBmmModelElement interface {
	IsRootScope (  )  bool
}

// Struct definition
type BmmModelElement struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		Optional documentation of this element, as a keyed list. It is strongly
		recommended to use the following key /type combinations for the relevant
		purposes: "purpose": String "keywords": List<String> "use": String "misuse":
		String "references": String Other keys and value types may be freely added.
	*/
	Documentation	Hash < Any , String >	`yaml:"documentation" json:"documentation" xml:"documentation"`
	// Model element within which an element is declared.
	Scope	IBmmModelElement	`yaml:"scope" json:"scope" xml:"scope"`
	/**
		Optional meta-data of this element, as a keyed list. May be used to extend the
		meta-model.
	*/
	Extensions	Hash < Any , String >	`yaml:"extensions" json:"extensions" xml:"extensions"`
}

//CONSTRUCTOR
func NewBmmModelElement() *BmmModelElement {
	bmmmodelelement := new(BmmModelElement)
	// Constants
	return bmmmodelelement
}
//BUILDER
type BmmModelElementBuilder struct {
	bmmmodelelement *BmmModelElement
}

func NewBmmModelElementBuilder() *BmmModelElementBuilder {
	 return &BmmModelElementBuilder {
		bmmmodelelement : NewBmmModelElement(),
	}
}

//BUILDER ATTRIBUTES
// Name of this model element.
func (i *BmmModelElementBuilder) SetName ( v string ) *BmmModelElementBuilder{
	i.bmmmodelelement.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmModelElementBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmModelElementBuilder{
	i.bmmmodelelement.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmModelElementBuilder) SetScope ( v IBmmModelElement ) *BmmModelElementBuilder{
	i.bmmmodelelement.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmModelElementBuilder) SetExtensions ( v Hash < Any , String > ) *BmmModelElementBuilder{
	i.bmmmodelelement.Extensions = v
	return i
}

func (i *BmmModelElementBuilder) Build() *BmmModelElement {
	 return i.bmmmodelelement
}

//FUNCTIONS
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmModelElement) IsRootScope (  )  bool {
	return false
}
