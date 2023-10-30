package v2

import (
	"SemanticDatabase/vocabulary"
)

/**
Persisted form of a package as a tree structure whose nodes can contain more
packages and/or classes.
*/

// Interface definition
type IPBmmPackage interface {
	Merge(other IPBmmPackage)
	CreateBmmPackageDefinition()
	// From: P_BMM_PACKAGE_CONTAINER
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmPackage struct {
	// embedded for Inheritance
	PBmmPackageContainer
	PBmmModelElement
	// Constants
	// Attributes
	/**
	Name of the package from schema; this name may be qualified if it is a top-level
	package within the schema, or unqualified. Persistent attribute.
	*/
	Name string `yaml:"name" json:"name" xml:"name"`
	// List of classes in this package. Persistent attribute.
	Classes []string `yaml:"classes" json:"classes" xml:"classes"`
	// BMM_PACKAGE created by create_bmm_package_definition .
	BmmPackageDefinition vocabulary.IBmmPackage `yaml:"bmmpackagedefinition" json:"bmmpackagedefinition" xml:"bmmpackagedefinition"`
}

// CONSTRUCTOR
func NewPBmmPackage() *PBmmPackage {
	pbmmpackage := new(PBmmPackage)
	// Constants
	return pbmmpackage
}

// BUILDER
type PBmmPackageBuilder struct {
	pbmmpackage *PBmmPackage
}

func NewPBmmPackageBuilder() *PBmmPackageBuilder {
	return &PBmmPackageBuilder{
		pbmmpackage: NewPBmmPackage(),
	}
}

//BUILDER ATTRIBUTES
/**
Name of the package from schema; this name may be qualified if it is a top-level
package within the schema, or unqualified. Persistent attribute.
*/
func (i *PBmmPackageBuilder) SetName(v string) *PBmmPackageBuilder {
	i.pbmmpackage.Name = v
	return i
}

// List of classes in this package. Persistent attribute.
func (i *PBmmPackageBuilder) SetClasses(v []string) *PBmmPackageBuilder {
	i.pbmmpackage.Classes = v
	return i
}

// BMM_PACKAGE created by create_bmm_package_definition .
func (i *PBmmPackageBuilder) SetBmmPackageDefinition(v vocabulary.IBmmPackage) *PBmmPackageBuilder {
	i.pbmmpackage.BmmPackageDefinition = v
	return i
}

// From: PBmmPackageContainer
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageBuilder) SetPackages(v map[IPBmmPackage]string) *PBmmPackageBuilder {
	i.pbmmpackage.Packages = v
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmPackageBuilder) SetDocumentation(v string) *PBmmPackageBuilder {
	i.pbmmpackage.Documentation = v
	return i
}

func (i *PBmmPackageBuilder) Build() *PBmmPackage {
	return i.pbmmpackage
}

//FUNCTIONS
/**
Merge packages and classes from other (from an included P_BMM_SCHEMA ) into this
package.
*/
func (p *PBmmPackage) Merge(other IPBmmPackage) {
	return
}

/*
*
Generate a BMM_PACKAGE_DEFINITION object and write it to bmm_package_definition
.
*/
func (p *PBmmPackage) CreateBmmPackageDefinition() {
	return
}
