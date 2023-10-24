package vocabulary

import (
	"vocabulary"
)

// A BMM model component that contains packages and classes.

// Interface definition
type IBmmPackageContainer interface {
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmPackageContainer struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// Child packages; keys all in upper case for guaranteed matching.
	Packages	Hash <String, BMM_PACKAGE >	`yaml:"packages" json:"packages" xml:"packages"`
	// Model element within which a referenceable element is known.
	Scope	IBmmPackageContainer	`yaml:"scope" json:"scope" xml:"scope"`
}

//CONSTRUCTOR
func NewBmmPackageContainer() *BmmPackageContainer {
	bmmpackagecontainer := new(BmmPackageContainer)
	// Constants
	// From: BmmModelElement
	return bmmpackagecontainer
}
//BUILDER
type BmmPackageContainerBuilder struct {
	bmmpackagecontainer *BmmPackageContainer
}

func NewBmmPackageContainerBuilder() *BmmPackageContainerBuilder {
	 return &BmmPackageContainerBuilder {
		bmmpackagecontainer : NewBmmPackageContainer(),
	}
}

//BUILDER ATTRIBUTES
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmPackageContainerBuilder) SetPackages ( v Hash <String, BMM_PACKAGE > ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Packages = v
	return i
}
// Model element within which a referenceable element is known.
func (i *BmmPackageContainerBuilder) SetScope ( v IBmmPackageContainer ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Scope = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmPackageContainerBuilder) SetName ( v string ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmPackageContainerBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmPackageContainerBuilder) SetScope ( v IBmmModelElement ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmPackageContainerBuilder) SetExtensions ( v Hash < Any , String > ) *BmmPackageContainerBuilder{
	i.bmmpackagecontainer.Extensions = v
	return i
}

func (i *BmmPackageContainerBuilder) Build() *BmmPackageContainer {
	 return i.bmmpackagecontainer
}

//FUNCTIONS
// Package at the path a_path .
func (b *BmmPackageContainer) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmPackageContainer) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackageContainer) HasPackagePath ( a_path string )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmPackageContainer) IsRootScope (  )  bool {
	return nil
}
