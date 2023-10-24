package vocabulary

import (
	"vocabulary"
)

/**
	Abstraction of a package as a tree structure whose nodes can contain other
	packages and classes. The name may be qualified if it is a top-level package.
*/

// Interface definition
type IBmmPackage interface {
	RootClasses (  )  List < BMM_CLASS >
	Path (  )  string
	// From: BMM_PACKAGE_CONTAINER
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmPackage struct {
	// embedded for Inheritance
	BmmPackageContainer
	BmmModelElement
	// Constants
	// Attributes
	// Member modules in this package.
	Members	List < BMM_MODULE >	`yaml:"members" json:"members" xml:"members"`
}

//CONSTRUCTOR
func NewBmmPackage() *BmmPackage {
	bmmpackage := new(BmmPackage)
	// Constants
	return bmmpackage
}
//BUILDER
type BmmPackageBuilder struct {
	bmmpackage *BmmPackage
}

func NewBmmPackageBuilder() *BmmPackageBuilder {
	 return &BmmPackageBuilder {
		bmmpackage : NewBmmPackage(),
	}
}

//BUILDER ATTRIBUTES
// Member modules in this package.
func (i *BmmPackageBuilder) SetMembers ( v List < BMM_MODULE > ) *BmmPackageBuilder{
	i.bmmpackage.Members = v
	return i
}
// From: BmmPackageContainer
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmPackageBuilder) SetPackages ( v Hash <String, BMM_PACKAGE > ) *BmmPackageBuilder{
	i.bmmpackage.Packages = v
	return i
}
// From: BmmPackageContainer
// Model element within which a referenceable element is known.
func (i *BmmPackageBuilder) SetScope ( v IBmmPackageContainer ) *BmmPackageBuilder{
	i.bmmpackage.Scope = v
	return i
}
// From: BmmModelElement
// Name of this model element.
func (i *BmmPackageBuilder) SetName ( v string ) *BmmPackageBuilder{
	i.bmmpackage.Name = v
	return i
}
// From: BmmModelElement
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmPackageBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmPackageBuilder{
	i.bmmpackage.Documentation = v
	return i
}
// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmPackageBuilder) SetScope ( v IBmmModelElement ) *BmmPackageBuilder{
	i.bmmpackage.Scope = v
	return i
}
// From: BmmModelElement
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmPackageBuilder) SetExtensions ( v Hash < Any , String > ) *BmmPackageBuilder{
	i.bmmpackage.Extensions = v
	return i
}

func (i *BmmPackageBuilder) Build() *BmmPackage {
	 return i.bmmpackage
}

//FUNCTIONS
/**
	Obtain the set of top-level classes in this package, either from this package
	itself or by recursing into the structure until classes are obtained from child
	packages. Recurse into each child only far enough to find the first level of
	classes.
*/
func (b *BmmPackage) RootClasses (  )  List < BMM_CLASS > {
	return nil
}
// Full path of this package back to root package.
func (b *BmmPackage) Path (  )  string {
	return ""
}
// From: BMM_PACKAGE_CONTAINER
// Package at the path a_path .
func (b *BmmPackage) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
// From: BMM_PACKAGE_CONTAINER
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmPackage) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
// From: BMM_PACKAGE_CONTAINER
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackage) HasPackagePath ( a_path string )  bool {
	return false
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmPackage) IsRootScope (  )  bool {
	return false
}
