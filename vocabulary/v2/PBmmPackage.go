package v2

import (
	"vocabulary"
)

/**
	Persisted form of a package as a tree structure whose nodes can contain more
	packages and/or classes.
*/

type IPBmmPackage interface {
	Merge ( other vocabulary.IPBmmPackage ) 
	CreateBmmPackageDefinition (  ) 
}

type PBmmPackage struct {
	/**
		Name of the package from schema; this name may be qualified if it is a top-level
		package within the schema, or unqualified. Persistent attribute.
	*/
	Name	string	`yaml:"name" json:"name" xml:"name"`
	// List of classes in this package. Persistent attribute.
	Classes	[]string	`yaml:"classes" json:"classes" xml:"classes"`
	// BMM_PACKAGE created by create_bmm_package_definition .
	BmmPackageDefinition	IBmmPackage	`yaml:"bmmpackagedefinition" json:"bmmpackagedefinition" xml:"bmmpackagedefinition"`
}

/**
	Merge packages and classes from other (from an included P_BMM_SCHEMA ) into this
	package.
*/
func (p *PBmmPackage) Merge ( other vocabulary.IPBmmPackage )  {
	return
}
/**
	Generate a BMM_PACKAGE_DEFINITION object and write it to bmm_package_definition
	.
*/
func (p *PBmmPackage) CreateBmmPackageDefinition (  )  {
	return
}
