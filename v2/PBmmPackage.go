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
	IPBmmModelElement
	IPBmmPackageContainer
	Merge(other IPBmmPackage)
	CreateBmmPackageDefinition()
	Name() string
	SetName(name string) error
	Classes() []string
	SetClasses(classes []string) error
	BmmPackageDefinition() vocabulary.IBmmPackage
	SetBmmPackageDefinition(bmmPackageDefinition vocabulary.IBmmPackage) error
}

// Struct definition
type PBmmPackage struct {
	// embedded for Inheritance
	PBmmModelElement
	PBmmPackageContainer
	// Attributes
	/**
	name of the package from schema; this name may be qualified if it is a top-level
	package within the schema, or unqualified. Persistent attribute.
	*/
	name string `yaml:"name" json:"name" xml:"name"`
	// List of classes in this package. Persistent attribute.
	classes []string `yaml:"classes" json:"classes" xml:"classes"`
	// BMM_PACKAGE created by create_bmm_package_definition .
	bmmPackageDefinition vocabulary.IBmmPackage `yaml:"bmmpackagedefinition" json:"bmmpackagedefinition" xml:"bmmpackagedefinition"`
}

func (p *PBmmPackage) Name() string {
	return p.name
}

func (p *PBmmPackage) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmPackage) Classes() []string {
	return p.classes
}

func (p *PBmmPackage) SetClasses(classes []string) error {
	p.classes = classes
	return nil
}

func (p *PBmmPackage) BmmPackageDefinition() vocabulary.IBmmPackage {
	return p.bmmPackageDefinition
}

func (p *PBmmPackage) SetBmmPackageDefinition(bmmPackageDefinition vocabulary.IBmmPackage) error {
	p.bmmPackageDefinition = bmmPackageDefinition
	return nil
}

// CONSTRUCTOR
func NewPBmmPackage() *PBmmPackage {
	pbmmpackage := new(PBmmPackage)
	pbmmpackage.classes = make([]string, 0)
	return pbmmpackage
}

// BUILDER
type PBmmPackageBuilder struct {
	pbmmpackage *PBmmPackage
	errors      []error
}

func NewPBmmPackageBuilder() *PBmmPackageBuilder {
	return &PBmmPackageBuilder{
		pbmmpackage: NewPBmmPackage(),
		errors:      make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
name of the package from schema; this name may be qualified if it is a top-level
package within the schema, or unqualified. Persistent attribute.
*/
func (i *PBmmPackageBuilder) SetName(v string) *PBmmPackageBuilder {
	i.AddError(i.pbmmpackage.SetName(v))
	return i
}

// List of classes in this package. Persistent attribute.
func (i *PBmmPackageBuilder) SetClasses(v []string) *PBmmPackageBuilder {
	i.AddError(i.pbmmpackage.SetClasses(v))
	return i
}

// BMM_PACKAGE created by create_bmm_package_definition .
func (i *PBmmPackageBuilder) SetBmmPackageDefinition(v vocabulary.IBmmPackage) *PBmmPackageBuilder {
	i.AddError(i.pbmmpackage.SetBmmPackageDefinition(v))
	return i
}

// From: PBmmPackageContainer
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageBuilder) SetPackages(v map[string]IPBmmPackage) *PBmmPackageBuilder {
	i.AddError(i.pbmmpackage.SetPackages(v))
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmPackageBuilder) SetDocumentation(v string) *PBmmPackageBuilder {
	i.AddError(i.pbmmpackage.SetDocumentation(v))
	return i
}

func (i *PBmmPackageBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
